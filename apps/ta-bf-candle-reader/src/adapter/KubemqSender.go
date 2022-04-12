package adapter

import (
	"context"
	"log"

	"github.com/kubemq-io/kubemq-go"
	"github.com/kubemq-io/kubemq-go/pkg/uuid"
)

type KubemqConnection struct {
	client *kubemq.EventsClient
	ctx    context.Context
}

func KubemqNewConnection(context context.Context) (conn *KubemqConnection) {
	client, err := kubemq.NewEventsClient(context,
		kubemq.WithAddress("ta-kubemq", 50000),
		kubemq.WithClientId("go-sdk-cookbook-pubsub-events-load-balance"),
		kubemq.WithTransportType(kubemq.TransportTypeGRPC))
	if err != nil {
		log.Fatal(err)
	}
	return &KubemqConnection{client: client, ctx: context}
}

func (conn KubemqConnection) Send(obj string) {
	err := conn.client.Send(conn.ctx, kubemq.NewEvent().SetId(uuid.New()).SetChannel("bf-candle").SetBody([]byte(obj)))
	if err != nil {
		log.Printf("error sending event %s, error: %s", obj, err)
	}
}

func (conn KubemqConnection) Ping() {
	// info, err := conn.client.Ping(conn.ctx)
	// if err != nil {
	// 	log.Fatalf("kubemq client error %s", err)
	// } else {
	// 	log.Printf("kubemq ping %v", info)
	// }
}

func (conn KubemqConnection) Close() {
	conn.client.Close()
	log.Println("closing Kubemq connection")
}
