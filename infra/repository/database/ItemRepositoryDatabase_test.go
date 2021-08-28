package database

import (
	"clean-arch-go/domain/repository"
	infraDB "clean-arch-go/infra/database"
	"clean-arch-go/ops/test"
	"database/sql"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ItemRepositoryDatabaseSuite struct {
	suite.Suite
	*require.Assertions
	ctrl *gomock.Controller
	db   *sql.DB
	repo repository.ItemRepository
}

func TestItemRepositoryDatabase(t *testing.T) {
	ctx, err := test.ContainerDBStart("./../../..")
	assert.NoError(t, err)
	suite.Run(t, new(ItemRepositoryDatabaseSuite))
	test.ContainerDBStop(ctx)
}

func (s *ItemRepositoryDatabaseSuite) SetupTest() {
	s.Assertions = require.New(s.T())
	s.ctrl = gomock.NewController(s.T())
	database := infraDB.PGDatabase{}
	pgDB := database.GetTestInstance()
	s.repo = NewItemRepositoryDatabase(pgDB)
	s.db = pgDB.GetDB()
}

func (s ItemRepositoryDatabaseSuite) TearDownTest() {
	s.ctrl.Finish()
}

func (s ItemRepositoryDatabaseSuite) TestGetById_ok() {
	err := test.DatasetTest(s.db, "./../../..", "clean.sql", "item_repository/data.sql")
	s.NoError(err)
	id := uuid.MustParse("c2e06611-8d13-4d2e-8406-ff878ae61ded")
	item, err := s.repo.GetById(id)
	s.NoError(err)
	s.NotEmpty(item)
	s.Equal(item.ID, id)
	s.Equal(item.Description, "Palheta")
}
