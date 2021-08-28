package database

import (
	"clean-arch-go/domain/repository"
	infraDB "clean-arch-go/infra/database"
	"clean-arch-go/ops/test"
	"database/sql"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type CouponRepositoryDatabaseSuite struct {
	suite.Suite
	*require.Assertions
	ctrl *gomock.Controller
	db   *sql.DB
	repo repository.CouponRepository
}

func TestCouponRepositoryDatabase(t *testing.T) {
	ctx, err := test.ContainerDBStart("./../../..")
	assert.NoError(t, err)
	suite.Run(t, new(CouponRepositoryDatabaseSuite))
	test.ContainerDBStop(ctx)
}

func (s *CouponRepositoryDatabaseSuite) SetupTest() {
	s.Assertions = require.New(s.T())
	s.ctrl = gomock.NewController(s.T())
	database := infraDB.PGDatabase{}
	pgDB := database.GetTestInstance()
	s.repo = NewCouponRepositoryDatabase(pgDB)
	s.db = pgDB.GetDB()
}

func (s CouponRepositoryDatabaseSuite) TearDownTest() {
	s.ctrl.Finish()
}

func (s CouponRepositoryDatabaseSuite) TestFindByCode() {
	err := test.DatasetTest(s.db, "./../../..", "clean.sql", "coupon/data.sql")
	s.NoError(err)
	coupon, err := s.repo.FindByCode("VALE20")
	s.NoError(err)
	s.NotEmpty(coupon)
	s.Equal("VALE20", coupon.Code)
	s.Equal(20.0, coupon.Percentage)
	s.True(coupon.Valid())
}

func (s CouponRepositoryDatabaseSuite) TestFindByCode_expired() {
	err := test.DatasetTest(s.db, "./../../..", "clean.sql", "coupon/data.sql")
	s.NoError(err)
	coupon, err := s.repo.FindByCode("VALE20_EXPIRED")
	s.NoError(err)
	s.NotEmpty(coupon)
	s.Equal("VALE20_EXPIRED", coupon.Code)
	s.Equal(20.0, coupon.Percentage)
	s.False(coupon.Valid())
}

func (s CouponRepositoryDatabaseSuite) TestFindByCode_notFound() {
	err := test.DatasetTest(s.db, "./../../..", "clean.sql", "coupon/data.sql")
	s.NoError(err)
	coupon, err := s.repo.FindByCode("AMAZING_COUPON")
	s.NoError(err)
	s.Empty(coupon)
}
