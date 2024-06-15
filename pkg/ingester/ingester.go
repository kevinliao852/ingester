package ingester

import (
	"fmt"
	v1 "inx/github.com/mukappalambda/my-trader/proto/message/v1"
	"time"

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

// TODO startWithContext
// TODO add a cxt
// TODO graceful shutdown

func (i *Ingester) Start() {

	for {
		data := <-i.queue
		// metadata
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

	}
}

func (i *Ingester) Stop() {
	close(i.queue)
}
