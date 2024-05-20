package ingester

import (
	"context"
	"fmt"
	v1 "inx/github.com/mukappalambda/my-trader/proto/message/v1"
	"time"

	"google.golang.org/genproto/googleapis/type/datetime"

	"github.com/google/uuid"
)

type Ingester struct {
	queue      chan string
	logger     IngesterLogger
	grpcClient v1.MessageServiceClient
}

type IngesterLogger interface {
	LogMessage(message string)
	RetrieveLogs(cursor *int)
}

type IngesterConfig struct {
	Logger IngesterLogger
}

func NewIngester(ic IngesterConfig, client v1.MessageServiceClient,
) *Ingester {
	return &Ingester{
		queue:      make(chan string),
		logger:     ic.Logger,
		grpcClient: client,
	}

}

func (i *Ingester) Ingest(data string) {
	i.queue <- data
}

type IngestData struct {
	Data      string
	Id        string
	Timestamp int64
}

func (i *Ingester) Start() {
	go func() {
		for {
			data := <-i.queue
			unixNano := time.Now().UnixNano()

			// uuid for id
			id := uuid.New().String()

			ingestData := IngestData{
				Data:      data,
				Id:        id,
				Timestamp: unixNano,
			}

			if i.logger != nil {
				i.logger.LogMessage(fmt.Sprintf("%v", ingestData))
			}

			t := time.Unix(0, unixNano)

			ctx := context.Background()

			fmt.Println("putting message", i.grpcClient)

			r, err := i.grpcClient.PutMessage(ctx, &v1.MessageRequest{
				Topic:   "test-topic",
				Message: data,
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
				fmt.Println(err)
			}

			fmt.Println(r)
		}
	}()
}

func (i *Ingester) Stop() {
	close(i.queue)
}
