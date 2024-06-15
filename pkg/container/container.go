package container

import (
	"context"
	"fmt"
	v1 "inx/github.com/mukappalambda/my-trader/proto/message/v1"
	"time"

	"google.golang.org/genproto/googleapis/type/datetime"
)

var _ Publisher = (*MessageServiceClientContainer)(nil)

type MessageServiceClientContainer struct {
	client v1.MessageServiceClient
	topic  string
}

func NewMessageServiceClientContainer() *MessageServiceClientContainer {
	return &MessageServiceClientContainer{}
}

func (c *MessageServiceClientContainer) Publish(message string) {
	unixNano := time.Now().UnixNano()
	t := time.Unix(0, unixNano)
	ctx := context.Background()
	fmt.Println("putting message", c.client)
	r, err := c.client.PutMessage(ctx, &v1.MessageRequest{
		Topic:   c.topic,
		Message: message,
		CreatedAt: &datetime.DateTime{
			Year:    int32(t.Year()),
			Month:   int32(t.Month()),
			Day:     int32(t.Day()),
			Hours:   int32(t.Hour()),
			Minutes: int32(t.Minute()),
			Seconds: int32(t.Second()),
		},
	})

	if err != nil {
		fmt.Println(err.Error())
	} else {

		fmt.Println(r.String())
	}
}
