package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/lokichoggio/gateway/internal/types"

	"github.com/tal-tech/go-zero/core/metric"
	"github.com/tal-tech/go-zero/core/trace/tracespec"
)

const (
	SuccessCode = 0
)

// A WithBodyResponseWriter is a helper to delay sealing a http.ResponseWriter on writing body.
type WithBodyResponseWriter struct {
	Writer http.ResponseWriter
	Body   []byte
}

// Header returns the http header.
func (w *WithBodyResponseWriter) Header() http.Header {
	return w.Writer.Header()
}

// Write writes bytes into w.
func (w *WithBodyResponseWriter) Write(bytes []byte) (int, error) {
	w.Body = bytes
	return w.Writer.Write(bytes)
}

// WriteHeader writes code into w, and not sealing the writer.
func (w *WithBodyResponseWriter) WriteHeader(code int) {
	w.Writer.WriteHeader(code)
}

const serverNamespace = "http_server"

var metricServiceCodeTotal = metric.NewCounterVec(&metric.CounterVecOpts{
	Namespace: serverNamespace,
	Subsystem: "requests",
	Name:      "service_code_total",
	Help:      "http server requests status code count.",
	Labels:    []string{"path", "code", "trace"},
})

func traceIdFromContext(ctx context.Context) string {
	t, ok := ctx.Value(tracespec.TracingKey).(tracespec.Trace)
	if !ok {
		return ""
	}

	return t.TraceId()
}

func ServiceMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bw := &WithBodyResponseWriter{Writer: w}

		defer func() {
			resp := types.BaseResp{}
			_ = json.Unmarshal(bw.Body, &resp)

			metricServiceCodeTotal.Inc(r.URL.Path, strconv.FormatInt(resp.ErrCode, 10), traceIdFromContext(r.Context()))
		}()

		next.ServeHTTP(bw, r)
	}
}
