package database

import (
	"clean-arch-go/domain/entity"
	"clean-arch-go/domain/repository"
	infraDB "clean-arch-go/infra/database"
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
)

type OrderRepositoryDatabase struct {
	db infraDB.Database
}

func NewOrderRepositoryDatabase(db infraDB.Database) repository.OrderRepository {
	return OrderRepositoryDatabase{db: db}
}

func (r OrderRepositoryDatabase) Save(order *entity.Order) error {
	insert := "insert into ccca.order (coupon_code, code, cpf, issue_date, shipping_cost, serial) values ($1, $2, $3, $4, $5, $6) returning id"
	result, err := r.db.Exec(context.Background(), insert, order.Coupon.Code, order.Code.Value, order.Document.Value(), order.IssueDate, order.ShippingCost, order.Sequence)
	if err != nil {
		return fmt.Errorf("could not save order: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("could not retrieve rows affected")
	}
	if rowsAffected < 1 {
		return fmt.Errorf("no rows affected when save order")
	}
	return nil
}

func (r OrderRepositoryDatabase) Get(code string) (*entity.Order, error) {
	order := &entity.Order{Coupon: entity.Coupon{}}
	var document string
	var couponCode sql.NullString
	var orderCode string
	err := r.db.One(context.Background(), "select id, coupon_code, code, cpf, issue_date, shipping_cost, serial from ccca.order where code = $1", code).
		Scan(&order.ID, &couponCode, &orderCode, &document, &order.IssueDate, &order.ShippingCost, &order.Sequence)
	if err != nil {
		return nil, fmt.Errorf("could not scan order: %w", err)
	}
	order.Code = entity.OrderCode{Value: orderCode}
	if couponCode.Valid {
		order.Coupon.Code = couponCode.String
	}
	order.Document = entity.NewCPF(document)
	if !order.Document.Valid() {
		return nil, fmt.Errorf("invalid cpf: %w", err)
	}
	rows, err := r.db.Many(context.Background(), "select id_item, price, quantity from ccca.order_item where id_order = $1", order.ID)
	for rows.Next() {
		var itemID uuid.UUID
		var price float64
		var quantity float64
		err = rows.Scan(&itemID, &price, &quantity)
		if err != nil {
			return nil, fmt.Errorf("could not scan order item: %w", err)
		}
		order.AddItem(itemID, price, quantity)
	}
	return order, nil
}

func (r OrderRepositoryDatabase) Count() int {
	var count int
	_ = r.db.One(context.Background(), "select count(*)::int as count from ccca.order").Scan(&count)
	return count
}
