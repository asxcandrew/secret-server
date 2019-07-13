package middleware

import (
	"context"
	stdhttp "net/http"
	"regexp"

	"github.com/go-kit/kit/transport/http"
)

const (
	AcceptHeaderJSONContextKey = "AcceptHeaderJSON"
	AcceptHeaderXMLContextKey  = "AcceptHeaderXML"
)

func HTTPToContext() http.RequestFunc {
	return func(ctx context.Context, r *stdhttp.Request) context.Context {
		header := r.Header.Get("Accept")

		if ok, _ := regexp.MatchString("json", header); ok {
			return context.WithValue(ctx, AcceptHeaderJSONContextKey, true)
		}

		if ok, _ := regexp.MatchString("xml", header); ok {
			return context.WithValue(ctx, AcceptHeaderXMLContextKey, true)
		}

		return ctx
	}
}
