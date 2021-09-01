package database

import (
	"clean-arch-go/domain/repository"
	infraDatabase "clean-arch-go/infra/database"
	"clean-arch-go/test"
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
	pgDB := infraDatabase.NewInstance()
	s.repo = NewItemRepositoryDatabase(pgDB)
	s.db = pgDB.GetDB()
}

func (s ItemRepositoryDatabaseSuite) TearDownTest() {
	s.ctrl.Finish()
}

func (s ItemRepositoryDatabaseSuite) TestGetById_ok() {
	err := test.DatasetTest(s.db, "./../../..", "clean.sql", "item_repository/data.sql")
	s.NoError(err)
	id := uuid.MustParse("36ed8660-feaa-4add-94c5-441792e8a0c2")
	item, err := s.repo.GetById(id)
	s.NoError(err)
	s.NotEmpty(item)
	s.Equal(item.ID, id)
	s.Equal("Cabo", item.Description)
	s.Equal("Acessórios", item.Category)
}
