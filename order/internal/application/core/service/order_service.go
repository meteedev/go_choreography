package service

import (
	"context"
	"encoding/json"
	"log"

	"github.com/meteedev/go_choreography/constant"
	"github.com/meteedev/go_choreography/order/internal/adapter/db"
	"github.com/meteedev/go_choreography/order/internal/application/core/domain"
	"github.com/meteedev/go_choreography/order/internal/ports"
	"github.com/meteedev/go_choreography/pkg/event"
	"github.com/meteedev/go_choreography/pkg/messenger"
)

type OrderService struct {
	OrderRepo      ports.OrderRepositoryPort
	MessageService messenger.MessengerService
}

func NewOrderService(m messenger.MessengerService, r ports.OrderRepositoryPort) *OrderService {
	return &OrderService{
		OrderRepo:      r,
		MessageService: m,
	}
}

func (o OrderService) CreateOrder(ctx context.Context, order event.OrderCreateEvent) (*domain.Order, error) {
	log.Println("in service call CreateOrder")

	order.Status = domain.OrderStatusProcessing
	orderInsert := ConvertOrderEventToParams(order)

	var orderDb db.Order
	var orderItemDbs []db.OrderItem

	err := o.OrderRepo.ExecTx(ctx, func(q *db.Queries) error {

		var e error
		orderDb, e = q.InsertOrder(ctx, orderInsert)
		if e != nil {
			log.Println(e.Error())
			return e
		}

		orderItemsInsert := ConvertOrderEventItemsToParams(order.OrderItems, orderDb.ID)

		for _, orderItem := range orderItemsInsert {
			orderItemDb, e := q.InsertOrderItems(ctx, orderItem)
			if e != nil {
				log.Println(e.Error())
				return e
			}
			orderItemDbs = append(orderItemDbs, orderItemDb)
		}

		return nil

	})

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	// assign uuid
	order.ID = orderDb.ID

	msg, err := json.Marshal(order)

	if err != nil {
		return nil, nil
	}

	//o.MessageService.Messenger.Publish(ctx, msg, constant.OrderQueue,false)
	o.MessageService.Messenger.Publish(ctx, msg, constant.OrderQueue, false)
	return nil, nil
}

func (o OrderService) UpdateOrder(ctx context.Context, event event.OrderUpdateEvent) (*domain.Order, error) {
	log.Println("in service call UpdateOrder ", event)
	return nil, nil
}
