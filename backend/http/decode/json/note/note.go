package note

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi"

	"github.com/luanngominh/mnotes/backend/endpoints/note"
)

//CreateNoteDecode decode http request
func CreateNoteDecode(_ context.Context, r *http.Request) (interface{}, error) {
	var req note.CreateNoteRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

//GetNoteDecode decode http request
func GetNoteDecode(_ context.Context, r *http.Request) (interface{}, error) {
	getQueryEncode := chi.URLParam(r, "get_query")

	getQueryDecode, err := base64.StdEncoding.DecodeString(getQueryEncode)
	if err != nil {
		return nil, ErrGetQuery
	}

	queryParser := strings.Split(string(getQueryDecode), "&")
	if len(queryParser) != 2 {
		return nil, ErrGetQuery
	}

	con, err := strconv.Atoi(strings.Replace(queryParser[0], "con=", "", 1))

	if err != nil {
		return nil, ErrGetQuery
	}

	limit, err := strconv.Atoi(strings.Replace(queryParser[1], "limit=", "", 1))
	if err != nil {
		return nil, ErrGetQuery
	}

	return note.GetNoteRequest{Continue: con, Limit: limit}, err
}

//DeleteNoteDecode decode http request
func DeleteNoteDecode(_ context.Context, r *http.Request) (interface{}, error) {
	var req note.DeleteNoteRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}
