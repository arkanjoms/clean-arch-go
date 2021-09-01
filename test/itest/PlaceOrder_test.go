package itest

import (
	"clean-arch-go/application/placeorder"
	"clean-arch-go/domain/service"
	"clean-arch-go/infra/database"
	"clean-arch-go/infra/factory"
	"clean-arch-go/infra/gateway/memory"
	"clean-arch-go/test"
	"database/sql"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type PlaceOrderSuite struct {
	suite.Suite
	*require.Assertions
	ctrl *gomock.Controller
	db   *sql.DB
	uc   placeorder.PlaceOrder
}

func TestNewPlaceOrder(t *testing.T) {
	ctx, err := test.ContainerDBStart("./../..")
	assert.NoError(t, err)
	suite.Run(t, new(PlaceOrderSuite))
	test.ContainerDBStop(ctx)
}

func (s *PlaceOrderSuite) SetupTest() {
	s.Assertions = require.New(s.T())
	s.ctrl = gomock.NewController(s.T())
	zipcodeClient := memory.NewZipcodeClient()
	freightCalculator := service.NewFreightCalculator()
	pgDB := database.NewInstance()
	repositoryFactory := factory.NewDatabaseRepositoryFactory(pgDB)
	s.db = database.NewInstance().GetDB()
	s.uc = placeorder.NewPlaceOrder(zipcodeClient, freightCalculator, repositoryFactory)
}

func (s PlaceOrderSuite) TearDownTest() {
	s.ctrl.Finish()
}

func (s PlaceOrderSuite) TestUseCoupon_Valid_20Percent() {
	err := test.DatasetTest(s.db, "./../..", "clean.sql", "place_order/data.sql")
	s.NoError(err)
	input := placeorder.InputPlaceOrder{
		Document: "05272720784",
		Items: []placeorder.InputPlaceOrderItem{
			{ItemID: uuid.MustParse("5549d46f-20d3-4d48-9cbe-80acc2b5cbb9"), Quantity: 2},
			{ItemID: uuid.MustParse("cf3dfb32-f654-42b6-be0b-d698eae8a146"), Quantity: 1},
			{ItemID: uuid.MustParse("36ed8660-feaa-4add-94c5-441792e8a0c2"), Quantity: 3},
		},
		CouponCode:     "VALE20",
		ZipcodeOrigin:  "A",
		ZipcodeDestiny: "B",
	}
	output, err := s.uc.Execute(input)
	s.NoError(err)
	s.NotNil(output)
	s.NotEmpty(output)
	s.Equal(5982.0, output.Total)
	s.Equal(310.0, output.ShippingCost)
}

func (s PlaceOrderSuite) TestUseCoupon_InvalidCoupon() {
	err := test.DatasetTest(s.db, "./../..", "clean.sql", "place_order/data.sql")
	s.NoError(err)
	input := placeorder.InputPlaceOrder{
		Document: "05272720784",
		Items: []placeorder.InputPlaceOrderItem{
			{ItemID: uuid.MustParse("5549d46f-20d3-4d48-9cbe-80acc2b5cbb9"), Quantity: 2},
			{ItemID: uuid.MustParse("cf3dfb32-f654-42b6-be0b-d698eae8a146"), Quantity: 1},
			{ItemID: uuid.MustParse("36ed8660-feaa-4add-94c5-441792e8a0c2"), Quantity: 3},
		},
		CouponCode:     "AMAZING_COUPON",
		ZipcodeOrigin:  "A",
		ZipcodeDestiny: "B",
	}
	output, err := s.uc.Execute(input)
	s.NoError(err)
	s.NotNil(output)
	s.NotEmpty(output)
	s.Equal(7400.0, output.Total)
	s.Equal(310.0, output.ShippingCost)
}

func (s PlaceOrderSuite) TestUseCoupon_ExpiredCoupon() {
	err := test.DatasetTest(s.db, "./../..", "clean.sql", "place_order/data.sql")
	s.NoError(err)
	input := placeorder.InputPlaceOrder{
		Document: "05272720784",
		Items: []placeorder.InputPlaceOrderItem{
			{ItemID: uuid.MustParse("5549d46f-20d3-4d48-9cbe-80acc2b5cbb9"), Quantity: 2},
			{ItemID: uuid.MustParse("cf3dfb32-f654-42b6-be0b-d698eae8a146"), Quantity: 1},
			{ItemID: uuid.MustParse("36ed8660-feaa-4add-94c5-441792e8a0c2"), Quantity: 3},
		},
		CouponCode:     "VALE20_EXPIRED",
		ZipcodeOrigin:  "A",
		ZipcodeDestiny: "B",
	}
	output, err := s.uc.Execute(input)
	s.NoError(err)
	s.NotNil(output)
	s.NotEmpty(output)
	s.Equal(7400.0, output.Total)
	s.Equal(310.0, output.ShippingCost)
}

func (s PlaceOrderSuite) TestUseCoupon_CalcOrderCode() {
	err := test.DatasetTest(s.db, "./../..", "clean.sql", "place_order/data.sql")
	s.NoError(err)
	location, _ := time.LoadLocation("America/Sao_Paulo")
	input := placeorder.InputPlaceOrder{
		Document:  "05272720784",
		IssueDate: time.Date(2021, time.Month(8), 20, 0, 0, 0, 0, location),
		Items: []placeorder.InputPlaceOrderItem{
			{ItemID: uuid.MustParse("5549d46f-20d3-4d48-9cbe-80acc2b5cbb9"), Quantity: 2},
			{ItemID: uuid.MustParse("cf3dfb32-f654-42b6-be0b-d698eae8a146"), Quantity: 1},
			{ItemID: uuid.MustParse("36ed8660-feaa-4add-94c5-441792e8a0c2"), Quantity: 3},
		},
		CouponCode:     "AMAZING_COUPON",
		ZipcodeOrigin:  "A",
		ZipcodeDestiny: "B",
	}
	output, err := s.uc.Execute(input)
	s.NoError(err)
	s.NotNil(output)
	s.NotEmpty(output)
	s.Equal(7400.0, output.Total)
	s.Equal(310.0, output.ShippingCost)
	s.Equal("202100000001", output.Code)
}
