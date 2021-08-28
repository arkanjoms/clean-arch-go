package application

import (
	infraDB "clean-arch-go/infra/database"
	"clean-arch-go/infra/factory"
	"clean-arch-go/ops/test"
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
	uc   *GetOrder
}

func TestGetOrder(t *testing.T) {
	ctx, err := test.ContainerDBStart("./..")
	assert.NoError(t, err)
	suite.Run(t, new(GetOrderSuite))
	test.ContainerDBStop(ctx)
}

func (s *GetOrderSuite) SetupTest() {
	s.Assertions = require.New(s.T())
	s.ctrl = gomock.NewController(s.T())
	database := infraDB.PGDatabase{}
	pgDB := database.GetTestInstance()
	repoFactory := factory.NewDatabaseRepositoryFactory(pgDB)
	s.db = pgDB.GetDB()
	s.uc = NewGetOrder(repoFactory)
}

func (s GetOrderSuite) TearDownTest() {
	s.ctrl.Finish()
}

func (s GetOrderSuite) TestExecute() {
	err := test.DatasetTest(s.db, "./..", "clean.sql", "place_order/data_with_order.sql")
	s.NoError(err)
	output, err := s.uc.Execute("202100000009")
	s.NoError(err)
	s.NotNil(output)
	s.Len(output.OrderItens, 1)
}
