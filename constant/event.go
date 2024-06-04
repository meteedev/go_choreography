package constant

const (
	OrderUpdateQueue   = "order_update"
	OrderQueue         = "order_queue"
	OrderFailedPayment = "order_fail_payment"
	OrderFail          = "order_failed"
	OrderCompensate    = "order_compensate"

	InventoryQueue       = "inventory_queue"
	InventoryEvent       = "inventory"
	InventoryFailedQueue = "inventory_failed"
	InventoryCompensate  = "inventory_compensate"

	PaymentQueue         = "payment_queue"
	PaymentBindingFailed = "payment_failed"
)
