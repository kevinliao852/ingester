package main

import (
	"fmt"
	v1 "inx/github.com/mukappalambda/my-trader/proto/message/v1"
	"inx/pkg/client"
	"inx/pkg/ingester"
	"inx/pkg/inputs"
	"inx/pkg/metadatastore"
	"inx/pkg/store"
	"log"
)

// set up the ingestion server
//

func main() {
	// run ingestion server
	gc := client.GrpcClient()
	v1MessageServiceClient := v1.NewMessageServiceClient(gc)
	fmt.Println("starting ingestion server", v1MessageServiceClient)

	dbUrl := "postgres://postgres:password@localhost:5432/demo?sslmode=disable"
	metadataStore := metadatastore.NewStore(dbUrl)
	err := metadataStore.CreateDB()
	if err != nil {
		log.Fatal(err)
	}
	ap := store.NewAppendLogger("ingester.log")
	ingest := ingester.NewIngester(ingester.IngesterConfig{
		Logger: ap,
	},
		v1MessageServiceClient,
	)
	go ingest.Start()

	// set up the http server
	inputs.SetUpIngester(ingest)

}
