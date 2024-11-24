package repositories

import (
	// "encoding/json"
	"tech-challenge-fase-1/internal/core/entities"
	valueobjects "tech-challenge-fase-1/internal/core/value_objects"
	"tech-challenge-fase-1/internal/infra/database"
)

type PaymentRepositoryDB struct {
	conn database.ConnectionDB
}

func NewOrderRepositoryDB(conn database.ConnectionDB) *PaymentRepositoryDB {
	return &PaymentRepositoryDB{conn: conn}
}

func (r *PaymentRepositoryDB) Insert(order *entities.Payment) error {
	// sql := `
	// INSERT INTO orders(id, customer_id, items, payment_status, preparation_status)
	// VALUES ($1, $2, $3, $4, $5)
	// `
	// return r.conn.Exec(
	// 	sql,
	// 	order.GetId(),
	// 	order.GetCustomerId(),
	// 	newOrderItemHelperList(order.GetItems()),
	// 	order.GetPaymentStatus().String(),
	// 	order.GetPreparationStatus().String(),
	// )
	return nil
}

func (r *PaymentRepositoryDB) FindPaymentByOrderID(orderId string) (*entities.Payment, error) {
	// sql := `
	// SELECT
	// 	id,
	// 	customer_id,
	// 	items,
	// 	payment_status,
	// 	preparation_status
	// FROM orders 
	// WHERE id = $1`
	// row := r.conn.QueryRow(sql, orderId)
	// return r.toEntity(row)
	return nil, nil
}

func (r *PaymentRepositoryDB) Update(order *entities.Payment) error {
	// sql := `
	// UPDATE orders 
	// SET 
	// 	customer_id = $1,
	// 	items = $2, 
	// 	payment_status = $3,
	// 	preparation_status = $4
	// WHERE id = $5;`
	// return r.conn.Exec(
	// 	sql,
	// 	order.GetCustomerId(),
	// 	newOrderItemHelperList(order.GetItems()),
	// 	order.GetPaymentStatus().String(),
	// 	order.GetPreparationStatus().String(),
	// 	order.GetId(),
	// )
	return nil
}

func (r *PaymentRepositoryDB) toEntity(row database.RowDB) (*entities.Payment, error) {
	// var id string
	// var customerId *string
	// var items []*orderItemHelper
	// var paymentStatus entities.OrderPaymentStatus
	// var preparationStatus entities.OrderPreparationStatus
	// err := row.Scan(&id, &customerId, &items, &paymentStatus, &preparationStatus)
	// if err != nil {
	// 	if err.Error() == ErrNotFound {
	// 		return nil, ErrOrderNotFound
	// 	}
	// 	return nil, err
	// }
	// return entities.RestoreOrder(
	// 	id,
	// 	customerId,
	// 	orderItemsFromHelper(items),
	// 	paymentStatus,
	// 	preparationStatus,
	// ), nil
	return nil, nil
}

// type orderItemHelper struct {
// 	Amount      float64  `json:"amount,omitempty"`
// 	Quantity    int    `json:"quantity,omitempty"`
// 	ProductName string `json:"product_name,omitempty"`
// }
//
// func orderItemsFromHelper(orderItemsHelper []*orderItemHelper) []*valueobjects.OrderItem {
// 	orderItems := make([]*valueobjects.OrderItem, 0)
// 	for _, item := range orderItemsHelper {
// 		orderItems = append(
// 			orderItems,
// 			valueobjects.NewOrderItem(item.Amount, item.Quantity, item.ProductName),
// 		)
// 	}
// 	return orderItems
// }
//
// func newOrderItemHelperList(orderItems []*valueobjects.OrderItem) []*orderItemHelper {
// 	orderItemsHelper := make([]*orderItemHelper, 0)
// 	for _, item := range orderItems {
// 		orderItemsHelper = append(
// 			orderItemsHelper,
// 			&orderItemHelper{
// 				Amount: item.GetAmount(),
// 				Quantity: item.GetQuantity(),
// 				ProductName: item.GetProductName(),
// 			},
// 		)
// 	}
// 	return orderItemsHelper
// }
