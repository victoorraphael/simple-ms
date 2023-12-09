package rest

import (
	"context"
	"encoding/json"
)

type Provider interface {
	Get(ctx context.Context, url string, data RequestData) (resp Response, err error)
	Post(ctx context.Context, url string, data RequestData) (resp Response, err error)
	Put(ctx context.Context, url string, data RequestData) (resp Response, err error)
	Patch(ctx context.Context, url string, data RequestData) (resp Response, err error)
	Delete(ctx context.Context, url string, data RequestData) (resp Response, err error)
}

type RequestData struct {
	Body    any
	Headers map[string]string
}

type Response struct {
	Body       []byte
	Header     map[string]string
	StatusCode int
}

func (r Response) Scan(dest any) error {
	return json.Unmarshal(r.Body, dest)
}
