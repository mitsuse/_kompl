package kompl

import (
	"encoding/json"
	"net/http"
)

func (s *Server) getCandidates(writer http.ResponseWriter, requst *http.Request) {
	header := writer.Header()
	header.Set("Content-Type", "application/json")

	contextSeqJson := requst.FormValue("context")
	prefix := requst.FormValue("prefix")

	var contextSeq []string
	if err := json.Unmarshal([]byte(contextSeqJson), &contextSeq); err != nil {
		// TODO: Return error response.
		return
	}

	candSeq := s.Predictor().Predict(contextSeq, prefix, 10)

	encoder := json.NewEncoder(writer)
	encoder.Encode(candSeq)
}
