package middleware

import (
	"log/slog"
	"net/http"
	"net/http/pprof"
)

// Profiler is a middleware for pprof
type Profiler struct {
	Endpoint string
}

// NewProfiler creates a new profiler instance
func NewProfiler(prefixs ...string) *Profiler {
	prefix := "/debug/pprof"
	if len(prefixs) > 0 {
		prefix = prefixs[0]
	}
	slog.Info("server profiler start", "path", prefix)
	return &Profiler{Endpoint: prefix}
}

// Register injects routers on server mux
func (p *Profiler) Register(mux *http.ServeMux) {
	mux.Handle(p.Endpoint+"/", http.HandlerFunc(pprof.Index))
	mux.Handle(p.Endpoint+"/trace", http.HandlerFunc(pprof.Trace))
	mux.Handle(p.Endpoint+"/symbol", http.HandlerFunc(pprof.Symbol))
	mux.Handle(p.Endpoint+"/cmdline", http.HandlerFunc(pprof.Cmdline))
	mux.Handle(p.Endpoint+"/profile", http.HandlerFunc(pprof.Profile))
	mux.Handle(p.Endpoint+"/heap", http.HandlerFunc(pprof.Handler("heap").ServeHTTP))
	mux.Handle(p.Endpoint+"/block", http.HandlerFunc(pprof.Handler("block").ServeHTTP))
	mux.Handle(p.Endpoint+"/mutex", http.HandlerFunc(pprof.Handler("mutex").ServeHTTP))
	mux.Handle(p.Endpoint+"/allocs", http.HandlerFunc(pprof.Handler("allocs").ServeHTTP))
	mux.Handle(p.Endpoint+"/goroutine", http.HandlerFunc(pprof.Handler("goroutine").ServeHTTP))
	mux.Handle(p.Endpoint+"/threadcreate", http.HandlerFunc(pprof.Handler("threadcreate").ServeHTTP))
}
