package adapter

import (
	"context"
	"log"

	"github.com/kubemq-io/kubemq-go"
	"github.com/kubemq-io/kubemq-go/pkg/uuid"
)

type KubemqConnection struct {
	client   *kubemq.EventsClient
	ctx      context.Context
	clientid string
}

func KubemqNewConnection(context context.Context) (conn *KubemqConnection) {
	clientid := uuid.New()
	client, err := kubemq.NewEventsClient(context,
		kubemq.WithAddress("ta-kubemq", 50000),
		kubemq.WithClientId(clientid),
		kubemq.WithTransportType(kubemq.TransportTypeGRPC))
	if err != nil {
		log.Fatal(err)
	}

	return &KubemqConnection{client: client, ctx: context, clientid: clientid}
}

func (conn KubemqConnection) Subscribe(onEvent func(msg *kubemq.Event, err error)) {
	channel := "bf-candle"
	conn.client.Subscribe(conn.ctx, &kubemq.EventsSubscription{
		Channel:  channel,
		Group:    "ta-candle-store",
		ClientId: uuid.New(),
	}, onEvent)
}

func (conn KubemqConnection) Close() {
	conn.client.Close()
	log.Println("closing Kubemq connection")
}
