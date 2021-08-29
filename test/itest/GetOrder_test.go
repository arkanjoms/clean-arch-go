package itest

import (
	"clean-arch-go/application"
	infraDatabase "clean-arch-go/infra/database"
	infraFactory "clean-arch-go/infra/factory"
	"clean-arch-go/test"
	"database/sql"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type GetOrderSuite struct {
	suite.Suite
	*require.Assertions
	ctrl *gomock.Controller
	db   *sql.DB
	uc   *application.GetOrder
}

func TestGetOrder(t *testing.T) {
	ctx, err := test.ContainerDBStart("./../..")
	assert.NoError(t, err)
	suite.Run(t, new(GetOrderSuite))
	test.ContainerDBStop(ctx)
}

func (s *GetOrderSuite) SetupTest() {
	s.Assertions = require.New(s.T())
	s.ctrl = gomock.NewController(s.T())
	database := infraDatabase.PGDatabase{}
	pgDB := database.GetTestInstance()
	repoFactory := infraFactory.NewDatabaseRepositoryFactory(pgDB)
	s.db = pgDB.GetDB()
	s.uc = application.NewGetOrder(repoFactory)
}

func (s GetOrderSuite) TearDownTest() {
	s.ctrl.Finish()
}

func (s GetOrderSuite) TestExecute() {
	err := test.DatasetTest(s.db, "./../..", "clean.sql", "place_order/data_with_order.sql")
	s.NoError(err)
	output, err := s.uc.Execute("202100000009")
	s.NoError(err)
	s.NotNil(output)
	s.Len(output.OrderItens, 1)
}
