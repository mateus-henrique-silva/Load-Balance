package rest

import (
	"io"
	"net/http"
	"time"

	"go.mod/entity"
)

func ReadRequest(w http.ResponseWriter, body io.Reader) float64 {
	v, err := io.ReadAll(body)
	if err != nil {
		SendError(w, err)
	}
	size := float64(len(v) / 1024)

	Send(w, entity.Size{
		SizeRequest: size,
		Time:        time.Now(),
	})
	return size
}
