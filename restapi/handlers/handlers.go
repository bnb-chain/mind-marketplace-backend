package handlers

import (
	"context"
	"net/http"
	"runtime/debug"

	"github.com/go-openapi/errors"
	"github.com/google/uuid"
	"github.com/rs/cors"

	"github.com/bnb-chain/greenfield-data-marketplace-backend/service"
	"github.com/bnb-chain/greenfield-data-marketplace-backend/util"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
	body       []byte
	header     http.Header
}

func newResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{w, http.StatusOK, []byte{}, http.Header{}}
}

func (rw *responseWriter) Write(body []byte) (int, error) {
	rw.body = body
	return rw.ResponseWriter.Write(body)
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.header = rw.ResponseWriter.Header()
	rw.ResponseWriter.WriteHeader(code)
}

// SetupHandler enable CORS, handler metrics
func SetupHandler(handler http.Handler, app string, config *util.ServerConfig) http.Handler {

	cacheHandler := handleCache(handler, config)

	contextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		ctx = context.WithValue(ctx, "env", config.Env)           //nolint:staticcheck
		ctx = context.WithValue(ctx, "traceID", uuid.NewString()) //nolint:staticcheck
		r = r.WithContext(ctx)

		cacheHandler.ServeHTTP(w, r)
	})

	panicHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				util.Logger.Errorf("panic, err=%v, stack=%v", err, string(debug.Stack()))
				e := err.(error)
				errors.ServeError(w, r, e)
				return
			}
		}()
		contextHandler.ServeHTTP(w, r)
	})

	h := handleMetrics(panicHandler, app)

	handleCORS := cors.AllowAll().Handler
	return handleCORS(h)
}

func Error(err error) (int64, string) {
	switch e := err.(type) {
	case service.Err:
		return e.Code, e.Message
	case nil:
		return service.NoErr.Code, service.NoErr.Message
	default:
		return service.InternalErr.Code, err.Error()
	}
}
