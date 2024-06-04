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

	r.DeclareQueue(constant.OrderQueue, false)
	r.DeclareQueue(constant.OrderUpdateQueue, false)
	r.DeclareQueue(constant.OrderCompensate, false)
	r.DeclareQueue(constant.InventoryCompensate, false)
	r.DeclareQueue(constant.InventoryQueue, false)
	r.DeclareQueue(constant.PaymentQueue, false)

	r.DeclareExchange(constant.OrderFail, "fanout", false)
	r.BindQueue(constant.OrderCompensate, constant.OrderFail, "")
	r.BindQueue(constant.InventoryCompensate, constant.OrderFail, "")

}
