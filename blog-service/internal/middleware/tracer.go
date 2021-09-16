package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/go-programming-tour-book/blog-service/global"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

func Tracing() func(c *gin.Context) {
	return func(c *gin.Context) {
		var newCtx context.Context
		var span opentracing.Span
		// var traceID string
		// var spanID string
		// var spanContext = span.Context()
		// 这里类型断言会出现invalid memory address or nil pointer dereference wtf!!!
		// switch spanContext.(type) {
		// case jaeger.SpanContext:
		// 	jaegerContext := spanContext.(jaeger.SpanContext)
		// 	traceID = jaegerContext.TraceID().String()
		// 	spanID = jaegerContext.SpanID().String()
		// }
		spanCtx, err := opentracing.GlobalTracer().Extract(
			opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(c.Request.Header),
		)
		if err != nil {
			span, newCtx = opentracing.StartSpanFromContextWithTracer(
				c.Request.Context(),
				global.Tracer,
				c.Request.URL.Path,
			)
		} else {
			span, newCtx = opentracing.StartSpanFromContextWithTracer(
				c.Request.Context(),
				global.Tracer,
				c.Request.URL.Path,
				opentracing.ChildOf(spanCtx),
				opentracing.Tag{Key: string(ext.Component), Value: "HTTP"},
			)
		}
		defer span.Finish()

		// c.Set("X-Trace-ID", traceID)
		// c.Set("X-Span-ID", spanID)
		c.Request = c.Request.WithContext(newCtx)
		c.Next()
	}
}
