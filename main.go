package main

import (
	"log"

	"github.com/bjornm82/swagger-nifi/client"
	"github.com/bjornm82/swagger-nifi/client/flow"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

func main() {

	transport := httptransport.New("localhost:8080", "nifi-api", []string{"http"})
	// transport.Producers["application/json"] = runtime.JSONProducer()
	// transport.Consumers["application/json"] = runtime.JSONConsumer()
	cl := client.New(transport, strfmt.Default)

	p := flow.NewGetFlowParams()
	p.ID = "root"

	d, err := cl.Flow.GetFlow(p, nil)
	if err != nil {
		log.Println("error")
		log.Fatalln(err)
	}

	b, err := d.Payload.MarshalBinary()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(b))

	log.Println(*d.Payload.Permissions.CanRead)
	log.Println(*d.Payload.Permissions.CanWrite)
	log.Println(d.Payload.ProcessGroupFlow.Flow.ProcessGroups[0].ID)
}
