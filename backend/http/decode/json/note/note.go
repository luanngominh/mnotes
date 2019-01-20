package note

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/luanngominh/mnotes/backend/endpoints/note"
)

//CreateNoteDecode decode http request
func CreateNoteDecode(_ context.Context, r *http.Request) (interface{}, error) {
	var req note.CreateNoteRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}
