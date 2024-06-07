package queue

import (
	"context"
	"log"

	"github.com/meteedev/go_choreography/config"
	"github.com/meteedev/go_choreography/constant"
	"github.com/meteedev/go_choreography/pkg/messenger"
)

func SetUpQueue(ctx context.Context) {

	r, err := messenger.NewRabbitMQ(ctx, config.GetAmqpURL())

	if err != nil {
		log.Fatalf("Failed to initialize RabbitMQ Publisher: %v", err)
	}

	r.DeclareQueue(constant.Order_created, false)
	r.DeclareQueue(constant.Order_update, false)
	r.DeclareQueue(constant.Inventory_reserved, false)
	r.DeclareQueue(constant.Inventory_failed, false)
	r.DeclareQueue(constant.Inventory_compensate, false)
	r.DeclareQueue(constant.Payment_processed, false)

	// fan out  payment failed to  inv compensate and order update  
	r.DeclareExchange(constant.Payment_failed, "fanout", false)
	r.BindQueue(constant.Inventory_compensate, constant.Payment_failed, "")
	r.BindQueue(constant.Order_update, constant.Payment_failed, "")

}
