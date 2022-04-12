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

func (conn KubemqConnection) Subscribe() {
	channel := "bf-candle"
	conn.client.Subscribe(conn.ctx, &kubemq.EventsSubscription{
		Channel:  channel,
		Group:    "ta-candle-store",
		ClientId: uuid.New(),
	}, func(msg *kubemq.Event, err error) {
		if err != nil {
			log.Fatal(err)
		} else {
			log.Printf("Receiver A - Event Received:\nEventID: %s\nChannel: %s\nMetadata: %s\nBody: %s\n", msg.Id, msg.Channel, msg.Metadata, msg.Body)
		}
	})
}

func (conn KubemqConnection) Close() {
	conn.client.Close()
	log.Println("closing Kubemq connection")
}
