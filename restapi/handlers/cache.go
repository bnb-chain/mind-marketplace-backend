package handlers

import (
	"bytes"
	"encoding/gob"
	"net/http"
	"strings"

	"github.com/bnb-chain/mind-marketplace-backend/util"
	"github.com/coocood/freecache"
)

const (
	defaultCacheExpireSeconds = 5
)

type Response struct {
	Body   []byte
	Header http.Header
}

func (r Response) Bytes() []byte {
	var b bytes.Buffer
	enc := gob.NewEncoder(&b)
	_ = enc.Encode(&r)

	return b.Bytes()
}

func toResponse(b []byte) Response {
	var r Response
	dec := gob.NewDecoder(bytes.NewReader(b))
	_ = dec.Decode(&r)

	return r
}

func getRouteCacheSeconds(path string) int {
	switch path {
	default:
		return defaultCacheExpireSeconds
	}
}

func handleCache(handler http.Handler, config *util.ServerConfig) http.Handler {
	cacheSize := 100 * 1024 * 1024
	cache := freecache.NewCache(cacheSize)

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Only cache GET request
		if !config.APIConfig.EnableCache || r.Method != http.MethodGet {
			handler.ServeHTTP(w, r)
			return
		}

		key := r.URL.String()
		data, err := cache.Get([]byte(key))
		if err == nil {
			response := toResponse(data)
			for k, v := range response.Header {
				w.Header().Set(k, strings.Join(v, ","))
			}
			_, err = w.Write(response.Body)
			if err != nil {
				util.Logger.Errorf("response cache write error, err=%s", err.Error())
			}
			return
		}

		rw := newResponseWriter(w)
		handler.ServeHTTP(rw, r)

		if rw.statusCode == http.StatusOK {
			response := Response{
				Body:   rw.body,
				Header: rw.header,
			}
			data = response.Bytes()
			expireSeconds := getRouteCacheSeconds(r.URL.Path)
			if err := cache.Set([]byte(key), data, expireSeconds); err != nil {
				util.Logger.Errorf("response cache set error, err=%s", err.Error())
			}
		}
	})
	return h
}
