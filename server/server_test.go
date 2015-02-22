package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/mitsuse/kompl/predictor"
)

func createReader() io.Reader {
	byteSeq, _ := Asset("test/wiki.txt")

	return bytes.NewReader(byteSeq)
}

func TestNew(t *testing.T) {
	order := 3

	p, err := predictor.Build(order, createReader())
	if err != nil {
		template := "Failed to build a predictor: %s"
		t.Errorf(template, err.Error())
		return
	}

	port := "8080"
	server := New(port, p)

	if server.Port() != port {
		template := "The server should use the port %s, but will use %s."
		t.Errorf(template, port, server.Port())
		return
	}

	if server.Predictor() != p {
		template := "The server should use *Predictor@%p, but will use %p."
		t.Errorf(template, p, server.Predictor())
		return
	}
}

func TestServerGetDecription(t *testing.T) {
	order := 3

	p, err := predictor.Build(order, createReader())
	if err != nil {
		template := "Failed to build a predictor: %s"
		t.Errorf(template, err.Error())
		return
	}

	port := "8080"
	server := New(port, p)

	testServer := httptest.NewServer(http.HandlerFunc(server.getDescription))
	defer testServer.Close()

	response, err := http.Get(testServer.URL)
	if err != nil {
		template := "Failed to get: %s"
		t.Errorf(template, err.Error())
		return
	}

	if response.StatusCode != 200 {
		template := "Failed to get: status = %03d"
		t.Errorf(template, response.StatusCode)
		return
	}

	jsonObj := make(map[string]int)

	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&jsonObj); err != nil {
		template := "Failed to read the body: %s"
		t.Errorf(template, err.Error())
		return
	}

	predictorOrder, hasKey := jsonObj["order"]
	if !hasKey {
		t.Error("The response of \"getDescription\" should have the key \"order\".")
		return
	}

	if predictorOrder != order {
		template := "The order of \"getDescription\"'s response should be %d, but is %d."
		t.Errorf(template, order, predictorOrder)
		return
	}
}

func TestServerGetCandidates(t *testing.T) {
	order := 3

	p, err := predictor.Build(order, createReader())
	if err != nil {
		template := "Failed to build a predictor: %s"
		t.Errorf(template, err.Error())
		return
	}

	port := "8080"
	server := New(port, p)

	testServer := httptest.NewServer(http.HandlerFunc(server.getCandidates))
	defer testServer.Close()

	rawParams := "context=[\"also\", \"commonly\"]&prefix=ref"
	params, err := url.ParseQuery(rawParams)
	if err != nil {
		template := "Failed to parse parameters: %s"
		t.Errorf(template, rawParams)
		return
	}

	url := fmt.Sprintf("%s?%s", testServer.URL, params.Encode())
	expectedSeq := []string{"referred"}

	response, err := http.Get(url)
	if err != nil {
		template := "Failed to get: %s"
		t.Errorf(template, err.Error())
		return
	}

	if response.StatusCode != 200 {
		template := "Failed to get: status = %03d"
		t.Errorf(template, response.StatusCode)
		return
	}

	candidateSeq := []string{}

	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&candidateSeq); err != nil {
		template := "Failed to read the body: %s"
		t.Errorf(template, err.Error())
		return
	}

	if len(expectedSeq) != len(candidateSeq) {
		template := "The size of candidates should be %d, but is %d."
		t.Errorf(template, len(expectedSeq), len(candidateSeq))
		return
	}

	for i := range expectedSeq {
		if expectedSeq[i] != candidateSeq[i] {
			template := "The canidate should be %s, but is %s."
			t.Errorf(template, expectedSeq[i], candidateSeq[i])
			return
		}
	}
}
