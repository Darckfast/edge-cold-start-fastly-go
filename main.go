package main

import (
	"compute-starter-kit-go/types"
	"context"
	"time"

	"github.com/fastly/compute-sdk-go/fsthttp"
	"github.com/mailru/easyjson"
)

// The entry point for your application.
//
// Use this function to define your main request handling logic. It could be
// used to route based on the request properties (such as method or path), send
// the request to a backend, make completely new requests, and/or generate
// synthetic responses.

func main() {
	fsthttp.ServeFunc(func(ctx context.Context, w fsthttp.ResponseWriter, r *fsthttp.Request) {
		easyjson.MarshalToWriter(types.Payload{
			Time: time.Now().UnixMilli(),
		}, w)
	})
}
