package helpers

import (
	"net/http"

	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

func ConvertFasthttpToHttpRequest(fasthttp *fasthttp.RequestCtx) *http.Request {
	r := &http.Request{}
	fasthttpadaptor.ConvertRequest(fasthttp, r, true)

	return r
}
