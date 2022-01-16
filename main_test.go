package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelf(t *testing.T) {
	var scheme = "https"

	var host = "localhost"
	var port = 8443
	var path = "nifi-api"

	skipTlsClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	url := fmt.Sprintf("%s://%s:%d/%s/access/token", scheme, host, port, path)

	rd := strings.NewReader("username=admin&password=ctsBtRBKHRAx69EqUghvvgEvjnaLjFEB")

	req, err := http.NewRequest("POST", url, rd)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	resp, err := skipTlsClient.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(resp.Header)
	assert.Equal(t, "err", string(b))
	log.Println(resp.StatusCode)
}

// func TestMainClient(t *testing.T) {

// 	var scheme = "https"

// 	var host = "localhost"
// 	var port = 8443
// 	var path = "/nifi-api/"

// 	skipTlsClient := &http.Client{
// 		Transport: &http.Transport{
// 			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
// 		},
// 		Timeout: 30,
// 	}

// 	tr := httptransport.NewWithClient(fmt.Sprintf("%s:%d", host, port), path, []string{scheme}, skipTlsClient)

// 	tr.Debug = true

// 	c := client.New(tr, nil)

// 	var user = "admin"
// 	var pass = "ctsBtRBKHRAx69EqUghvvgEvjnaLjFEB"

// 	p := access.CreateAccessTokenParams{
// 		Username: &user,
// 		Password: &pass,
// 		Context:  context.Background(),
// 	}

// 	re, err := c.Access.CreateAccessToken(&p)
// 	if err != nil {
// 		log.Println(re)
// 		log.Fatalln(err)
// 	}

// 	log.Println(re.Payload)

// }
