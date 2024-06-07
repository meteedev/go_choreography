package constant

const (

	//publish by order
	//consume by inventory
	Order_created = "order_created"

	//publish by inventory , payment
	//consume by order
	Order_update = "order_update"

	//publish by inventory
	//consume by payment
	//consume by order
	Inventory_reserved = "inventory_reserved" // payment consume

	//publish by inventory
	//consume by order
	Inventory_failed = "inventory_failed" //order consume

	//publish by payment
	//consume by inventory
	Inventory_compensate = "inventory_compensate" //inventory consume

	//publish by payment
	//consume by order
	Payment_processed = "payment_completed" //order_consume

	//publish by payment
	//fan out to inventory_reserved , order_update
	Payment_failed = "payment_failed"
)
