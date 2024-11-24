package entities

import (
	"github.com/google/uuid"
)

type (
	OrderPaymentStatus string
)

func (s OrderPaymentStatus) String() string {
	return string(s)
}

const (
	ORDER_PAYMENT_PENDING          OrderPaymentStatus = "PENDING"
	ORDER_PAYMENT_PAID             OrderPaymentStatus = "PAID"
	ORDER_PAYMENT_REJECTED         OrderPaymentStatus = "REJECTED"
	ORDER_PAYMENT_AWAITING_PAYMENT OrderPaymentStatus = "AWAITING_PAYMENT"
)

type Payment struct {
	id            string
	orderID       string
	amount				float64
	paymentStatus OrderPaymentStatus
}

func CreatePayment(orderID *string, amount float64) *Payment {
	return RestorePayment(
		uuid.NewString(),
		orderID,
		amount,
		ORDER_PAYMENT_PENDING,
	)
}

func RestorePayment(
	id string,
	orderID *string,
	amount float64,
	paymentStatus OrderPaymentStatus,
) *Payment {
	return &Payment{
		id:            id,
		orderID:       *orderID,
		amount:				 amount,		
		paymentStatus: paymentStatus,
	}
}

func (p *Payment) GetID() string {
	return p.id
}

func (p *Payment) GetOrderID() string {
	return p.orderID
}

func (p *Payment) GetAmount() float64 {
	return p.amount
}

func (p *Payment) GetPaymentStatus() OrderPaymentStatus {
	return p.paymentStatus
}

func (p *Payment) AwaitingPayment() {
	p.paymentStatus = ORDER_PAYMENT_AWAITING_PAYMENT
}

func (p *Payment) PaymentReceived() {
	p.paymentStatus = ORDER_PAYMENT_PAID
}

func (p *Payment) PaymentRejected() {
	p.paymentStatus = ORDER_PAYMENT_REJECTED
}


// type Order struct {
// 	id                string
// 	customerId        *string
// 	items             []*valueobjects.OrderItem
// 	paymentStatus     OrderPaymentStatus
// 	preparationStatus OrderPreparationStatus
// }
//
// func CreateOpenOrder(customerId *string) *Order {
// 	return RestoreOrder(
// 		uuid.NewString(),
// 		customerId,
// 		make([]*valueobjects.OrderItem, 0),
// 		ORDER_PAYMENT_PENDING,
// 		ORDER_PREPARATION_AWAITING,
// 	)
// }
//
// func RestoreOrder(
// 	id string,
// 	customerId *string,
// 	items []*valueobjects.OrderItem,
// 	paymentStatus OrderPaymentStatus,
// 	preparationStatus OrderPreparationStatus,
// ) *Order {
// 	return &Order{
// 		id:                id,
// 		customerId:        customerId,
// 		items:             items,
// 		paymentStatus:     paymentStatus,
// 		preparationStatus: preparationStatus,
// 	}
// }
//
// func (o *Order) GetId() string {
// 	return o.id
// }
//
// func (o *Order) GetCustomerId() *string {
// 	return o.customerId
// }
//
// func (o *Order) GetItems() []*valueobjects.OrderItem {
// 	return o.items
// }
//
// func (o *Order) GetPaymentStatus() OrderPaymentStatus {
// 	return o.paymentStatus
// }
//
// func (o *Order) AwaitingPayment() {
// 	o.paymentStatus = ORDER_PAYMENT_AWAITING_PAYMENT
// }
//
// func (o *Order) PaymentReceived() {
// 	o.paymentStatus = ORDER_PAYMENT_PAID
// 	o.SetPreparationStatus(ORDER_PREPARATION_RECEIVED)
// }
//
// func (o *Order) PaymentRejected() {
// 	o.paymentStatus = ORDER_PAYMENT_REJECTED
// 	o.SetPreparationStatus(ORDER_PREPARATION_CANCELED)
// }
//
// func (o *Order) GetPreparationStatus() OrderPreparationStatus {
// 	return o.preparationStatus
// }
//
// func (o *Order) SetPreparationStatus(status OrderPreparationStatus) error {
// 	if !IsValidOrderPreparationStatus(status) {
// 		return errors.ErrInvalidPreparationStatus
// 	}
// 	o.preparationStatus = status
// 	return nil
// }
//
// func (o *Order) GetTotal() float64 {
// 	var total float64
// 	for _, item := range o.items {
// 		total = total + item.GetTotal()
// 	}
// 	return total
// }
//
// func (o *Order) FindOrderItem(productName string) *valueobjects.OrderItem {
// 	for _, item := range o.items {
// 		if item.GetProductName() == productName {
// 			return item
// 		}
// 	}
// 	return nil
// }
//
// func (o *Order) AddItem(product *Product, quantity int) {
// 	amount := product.GetPrice()
// 	productName := product.GetName()
// 	item := o.FindOrderItem(productName)
// 	if item == nil {
// 		item = valueobjects.NewOrderItem(amount, 0, productName)
// 		o.items = append(o.items, item)
// 	}
// 	quantity = item.GetQuantity() + quantity
// 	item.SetQuatity(quantity)
// }
//
