# Go choreography
- Order module
- Inventory module
- Payment module
- Messaging bus by rabbitMQ

## How to run


- start docker-compose for init rabbitmq 

- set environment variable 
    AMQP_URL="amqp://guest:guest@xx.xx.xx.161:5672/"
    DATABASE_URL="host=yourip port=5433 user=user password=password dbname=order sslmode=disable"
    (dbname change to  "order" , "inventory" , "payment" )

- ORDER
    - run  queue.go  in path \order\cmd\queue  for init queue
    - run  echo.go   in path \order\cmd\server for start http server
    - run  consumer.go   in path \order\cmd\consumer\ for order consumer event

- INVENTORY 
    - run  consumer.go   in path \inventory\cnd\consumer\ for inventory consumer event

- PAYMENT 
    - run  consumer.go   in path \payment\cnd\consumer\ for inventory consumer event

- example json for post  /orders path
```json
{
  "customer_id": 987654321,
  "status": "success",
  "order_items": [
    {
      "product_code": "PROD001",
      "unit_price": 10,
      "quantity": 2
    },
    {
      "product_code": "PROD002",
      "unit_price": 100,
      "quantity": 1
    }
  ],
  "created_at": 1672531199
}
```