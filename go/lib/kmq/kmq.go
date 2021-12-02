package kmq

import (
	"context"
	"log"

	kubemq "github.com/kubemq-io/kubemq-go"
)

type Kmq struct {
	client *kubemq.QueuesClient
}

func New(ctx context.Context, host, clientID string) (*Kmq, error) {
	queuesClient, err := kubemq.NewQueuesStreamClient(ctx,
		kubemq.WithAddress(host, 50000),
		kubemq.WithClientId(clientID),
		kubemq.WithTransportType(kubemq.TransportTypeGRPC))
	if err != nil {
		return nil, err
	}

	return &Kmq{
		client: queuesClient,
	}, nil

}

func (k *Kmq) Close() func() {
	return func() {
		err := k.client.Close()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (k *Kmq) Publish(ctx context.Context, messageID, channel string, message []byte) error {

	_, err := k.client.Send(ctx, kubemq.NewQueueMessage().
		SetId(messageID).
		SetChannel(channel).
		SetBody([]byte(message)))
	if err != nil {
		return err
	}

	return nil
}
