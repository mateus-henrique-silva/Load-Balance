package rest

import (
	"encoding/json"
	"net/http"
)

func RequestBody(w http.ResponseWriter, r *http.Request, v interface{}) error {
	jsonDec := json.NewDecoder(r.Body)
	if err := jsonDec.Decode(v); err != nil {
		SendError(w, err)
		return err
	}
	Send(w, v)
	return nil

}
