package inputs

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Ingester interface {
	Ingest(data string)
}

func SetUpIngester(ingester Ingester) {
	ingestHandler := func(w http.ResponseWriter, r *http.Request) {
		data := r.Body
		defer r.Body.Close()

		dataBytes, err := ioutil.ReadAll(data)

		if err != nil {
			panic(err)
		}

		stringData := string(dataBytes)
		ingester.Ingest(stringData)
	}

	receiverHandler := func(w http.ResponseWriter, r *http.Request) {
	}

	http.HandleFunc("/ingest", ingestHandler)
	http.HandleFunc("/receiver", receiverHandler)
	fmt.Println("http server started on port 8080")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err)
	}

}
