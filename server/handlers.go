package server

import (
	"encoding/json"
	"net/http"
	"strings"
)

/*
This request handler returns the cmpletion canidates for the given context and prefix.
*/
func (s *Server) getCandidates(writer http.ResponseWriter, requst *http.Request) {
	header := writer.Header()
	header.Set("Content-Type", "application/json")

	var state *State
	encodedState := requst.FormValue("state")
	decoder := json.NewDecoder(strings.NewReader(encodedState))

	if err := decoder.Decode(&state); err != nil {
		// TODO: Return error response.
		return
	}

	order := s.Predictor().Order()
	context := make([]string, 0, order)
	tokenSeq := s.Tokenizer().Tokenize(state.Context)

	emptySize := -1 * (len(tokenSeq) - order + 1)
	for i := 0; i < emptySize; i++ {
		context = append(context, "")
	}

	tokenIndex := len(tokenSeq) - order + len(context) + 1
	for i := tokenIndex; len(context) < order-1; i++ {
		context = append(context, tokenSeq[i])
	}

	candSeq := s.Predictor().Predict(context, state.Prefix, state.K)

	encoder := json.NewEncoder(writer)
	encoder.Encode(candSeq)
}

func (s *Server) getDescription(writer http.ResponseWriter, requst *http.Request) {
	header := writer.Header()
	header.Set("Content-Type", "application/json")

	descMap := make(map[string]interface{})
	descMap["order"] = s.Predictor().Order()

	encoder := json.NewEncoder(writer)
	encoder.Encode(descMap)
}

type State struct {
	Context string `json:"context"`
	Prefix  string `json:"prefix"`
	K       int    `json:"k"`
}
