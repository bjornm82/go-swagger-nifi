package main

import (
	"log"

	"github.com/bjornm82/go-swagger-nifi/client"
	"github.com/bjornm82/go-swagger-nifi/client/flow"
	"github.com/bjornm82/go-swagger-nifi/client/process_groups"
	"github.com/bjornm82/go-swagger-nifi/models"
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

	log.Println(d.Payload.ProcessGroupFlow.ID)

	pgcl := cl.ProcessGroups

	pg := process_groups.NewCreateProcessGroupParams()
	model := models.ProcessGroupEntity{}

	value := `{"revision":{"version":0},"component":{"name":"some-name-here","position":{"x":369,"y":278}}}`
	err = model.UnmarshalBinary([]byte(value))
	if err != nil {
		log.Fatalln(err)
	}
	pg.SetBody(&model)
	pg.ID = d.Payload.ProcessGroupFlow.ID

	gr, err := pgcl.CreateProcessGroup(pg, nil)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(gr.Payload.ID)

	b, err := d.Payload.MarshalBinary()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(b))

	log.Println(d.Payload.ProcessGroupFlow.Flow.ProcessGroups[0].ID)
}
