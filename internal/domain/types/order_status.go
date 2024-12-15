package types

type OrderStatus string

const (
	OrderStatusNew        OrderStatus = "new"
	OrderStatusReadyToPay OrderStatus = "ready_to_pay"
	OrderStatusPayed      OrderStatus = "payed"
)

func (os OrderStatus) String() string {
	return string(os)
}
