package format

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"net/http"
	"zskparker.com/foundation/base/pb"
)

func EncodeHTTPGenericResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if f, ok := response.(endpoint.Failer); ok && f.Failed() != nil {
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func Metadata() kithttp.RequestFunc {
	return func(ctx context.Context, request *http.Request) context.Context {
		meta := &fs_base.Metadata{}
		meta.Device = request.Header.Get("X-User-Device")
		meta.ClientId = request.Header.Get("X-Client-Id")
		meta.Ip = request.Header.Get("X-Real-IP")
		meta.Token = request.Header.Get("Authorization")
		meta.UserAgent = request.Header.Get("User-Agent")
		meta.Api = request.URL.Path
		return context.WithValue(ctx, "meta", meta)
	}
}
