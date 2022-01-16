package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"strings"
	"text/template"

	"github.com/bjornm82/go-swagger-nifi/client"
	"github.com/bjornm82/go-swagger-nifi/client/flow"
	"github.com/bjornm82/go-swagger-nifi/client/process_groups"
	"github.com/bjornm82/go-swagger-nifi/models"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

func main() {

	c := New("localhost", 8080, false)

	v, err := c.findRootFlowID()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(v)

	v2, err := c.createProcessorGroupOnRoot("hell-now")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(v2.Payload.ID)

	v3, err := c.findRootProcessorGroupByName("hell-now")
	if err != nil {
		log.Fatalln(err)
	}
	log.Fatalln(v3.Component.Name)

	type Inventory struct {
		Material string
		Count    uint
	}
	sweaters := Inventory{"wool", 17}
	// dat, err := ioutil.ReadFile("templates/s3-list-pg.xml")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	tmpl, err := template.New("s3-list-pg.xml").ParseFiles("templates/s3-list-pg.xml")
	if err != nil {
		panic(err)
	}
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, sweaters)
	if err != nil {
		panic(err)
	}

	transport := httptransport.New("localhost:8080", "nifi-api", []string{"http"})
	cl := client.New(transport, strfmt.Default)

	p := flow.NewGetFlowParams()
	p.ID = "root"

	d, err := cl.Flow.GetFlow(p, nil)
	if err != nil {
		log.Println("error")
		log.Fatalln(err)
	}

	log.Println(d.Payload.ProcessGroupFlow.ID)

	pgtemp := process_groups.NewUploadTemplateParams()
	pgtemp.SetID(d.Payload.ProcessGroupFlow.ID)
	re := strings.NewReader(buf.String())
	pgtemp.SetTemplate(runtime.NamedReader("s3-list-pg", re))
	temp, err := cl.ProcessGroups.UploadTemplate(pgtemp, nil)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(temp.Error())
	log.Println(temp.Payload)
	// TODO: temp has empty payload, xml response not supported according to: https://github.com/go-swagger/go-swagger/issues/2245

	itempPar := process_groups.NewInstantiateTemplateParams()
	itempPar.SetID(temp.Payload.Template.ID)

	itemp, err := cl.ProcessGroups.InstantiateTemplate(itempPar, nil)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(itemp.Payload.Flow.ProcessGroups)

	pgcl := cl.ProcessGroups

	pg := process_groups.NewCreateProcessGroupParams()
	model := models.ProcessGroupEntity{}

	value := `{"revision":{"version":0},"component":{"name":"some-name-here","position":{"x":369,"y":278}}}`
	err = model.UnmarshalBinary([]byte(value))
	if err != nil {
		log.Fatalln(err)
	}
	pg.SetBody(&model)
	pg.SetID(d.Payload.ProcessGroupFlow.ID)

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

type cl struct {
	*client.NiFiRest
}

func New(host string, port int, usessl bool) *cl {
	var scheme = "http"
	if usessl {
		scheme = "https"
	}

	// tr := &http.Transport{
	// 	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	// }

	httptransport.BasicAuth("admin", "nonono")

	transport := httptransport.New(fmt.Sprintf("%s:%d", host, port), "nifi-api", []string{scheme})

	return &cl{
		client.New(transport, strfmt.Default),
	}
}

func (c *cl) createProcessorGroupOnRoot(name string) (*process_groups.CreateProcessGroupCreated, error) {
	id, err := c.findRootFlowID()
	if err != nil {
		return nil, err
	}
	pg := process_groups.NewCreateProcessGroupParams()
	pge := models.ProcessGroupEntity{}
	var version0 = int64(0)
	rev := models.RevisionDTO{
		Version: &version0,
	}
	pge.Revision = &rev
	pgdto := models.ProcessGroupDTO{}
	pgdto.Name = name
	pge.Component = &pgdto

	pg.SetBody(&pge)
	pg.SetID(id)

	pgcl := c.ProcessGroups

	return pgcl.CreateProcessGroup(pg, nil)
}

func (c *cl) findRootFlowID() (string, error) {
	p := flow.NewGetFlowParams()
	p.ID = "root"

	d, err := c.Flow.GetFlow(p, nil)
	if err != nil {
		return "", err
	}
	return d.Payload.ProcessGroupFlow.ID, nil
}

func (c *cl) findRootProcessorGroupByName(name string) (*models.ProcessGroupEntity, error) {
	if name == "" {
		return nil, errors.New("no group found without name")
	}
	p := flow.NewGetFlowParams()
	p.ID = "root"

	d, err := c.Flow.GetFlow(p, nil)
	if err != nil {
		return nil, err
	}

	pgf := d.Payload.ProcessGroupFlow
	pg := pgf.Flow.ProcessGroups
	for _, v := range pg {
		if v.Component.Name == name {
			return v, nil
		}
	}

	return nil, fmt.Errorf("no group found with name: %s", name)
}

func (c *cl) createProcessorGroupControllerService() {

}

func (c *cl) uploadTemplate(sweaters interface{}) {
	tmpl, err := template.New("s3-list-pg.xml").ParseFiles("templates/s3-list-pg.xml")
	if err != nil {
		panic(err)
	}
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, sweaters)
	if err != nil {
		panic(err)
	}
}
