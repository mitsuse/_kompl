package server

import (
	"encoding/json"
	"net/http"
	"strings"
)

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

	var context []string
	if len(state.Context.Tokens) == 0 {
		context = make([]string, 0, state.K)
		tokenSeq := s.Tokenizer().Tokenize(state.Context.Chars)

		emptySize := -1 * (len(tokenSeq) - state.K)
		for i := 0; i < emptySize; i++ {
			context = append(context, "")
		}

		tokenIndex := len(tokenSeq) + state.K - len(context) - 1
		for i := tokenIndex; len(context) < state.K; i++ {
			context = append(context, tokenSeq[i])
		}
	} else {
		context = state.Context.Tokens
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
	Context *Context `json:"context"`
	Prefix  string   `json:"prefix"`
	K       int      `json:"k"`
}

type Context struct {
	Tokens []string `json:"tokens"`
	Chars  string   `json:"chars"`
}
