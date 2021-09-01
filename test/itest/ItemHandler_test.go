package itest

import (
	"clean-arch-go/application/getitem"
	"clean-arch-go/domain/factory"
	"clean-arch-go/domain/gateway"
	"clean-arch-go/domain/service"
	"clean-arch-go/infra/database"
	infraFactory "clean-arch-go/infra/factory"
	"clean-arch-go/infra/gateway/memory"
	infraHttp "clean-arch-go/infra/http"
	"clean-arch-go/test"
	"database/sql"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

type ItemHandlerSuite struct {
	suite.Suite
	*require.Assertions
	ctrl              *gomock.Controller
	db                *sql.DB
	repositoryFactory factory.RepositoryFactory
	server            infraHttp.Http
	zipcodeClient     gateway.ZipcodeClient
	freightCalculator service.FreightCalculator
}

func TestItemHandler(t *testing.T) {
	ctx, err := test.ContainerDBStart("./../..")
	assert.NoError(t, err)
	suite.Run(t, new(ItemHandlerSuite))
	test.ContainerDBStop(ctx)
}

func (s *ItemHandlerSuite) SetupTest() {
	s.Assertions = require.New(s.T())
	s.ctrl = gomock.NewController(s.T())
	s.db = database.NewInstance().GetDB()
	pgDB := database.NewInstance()
	s.repositoryFactory = infraFactory.NewDatabaseRepositoryFactory(pgDB)
	s.zipcodeClient = memory.NewZipcodeClient()
	s.freightCalculator = service.NewFreightCalculator()
	s.server = infraHttp.NewGorillaMux()
	infraHttp.NewRouteConfig(s.server, s.repositoryFactory, s.zipcodeClient, s.freightCalculator).Build()
}

func (s ItemHandlerSuite) TearDownTest() {
	s.ctrl.Finish()
}

func (s ItemHandlerSuite) TestGetItems_noFilter() {
	err := test.DatasetTest(s.db, "./../..", "clean.sql", "item_repository/data.sql")
	s.NoError(err)
	req, _ := http.NewRequest("GET", "/api/items", nil)
	response := test.ExecuteRequest(req, s.server.(infraHttp.GorillaMux).Router)
	s.Equal(http.StatusOK, response.Code)
	var items []getitem.OutputGetItem
	err = json.Unmarshal(response.Body.Bytes(), &items)
	s.NoError(err)
	s.Len(items, 3)
}

func (s ItemHandlerSuite) TestGetItems_categoryFilter_notFound() {
	err := test.DatasetTest(s.db, "./../..", "clean.sql", "item_repository/data.sql")
	s.NoError(err)
	req, _ := http.NewRequest("GET", "/api/items?category=foo", nil)
	response := test.ExecuteRequest(req, s.server.(infraHttp.GorillaMux).Router)
	s.Equal(http.StatusNotFound, response.Code)
}

func (s ItemHandlerSuite) TestGetItems_categoryFilter_found() {
	err := test.DatasetTest(s.db, "./../..", "clean.sql", "item_repository/data.sql")
	s.NoError(err)
	req, _ := http.NewRequest("GET", "/api/items?category=Acessórios", nil)
	response := test.ExecuteRequest(req, s.server.(infraHttp.GorillaMux).Router)
	s.Equal(http.StatusOK, response.Code)
	var items []getitem.OutputGetItem
	err = json.Unmarshal(response.Body.Bytes(), &items)
	s.NoError(err)
	s.Len(items, 1)
	s.Equal("Acessórios", items[0].Category)
}
