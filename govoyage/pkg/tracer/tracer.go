package tracer

import (
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

// Setup sets up the tracing system
func Setup() error {
	cfg := config.Configuration{
		Sampler: &config.SamplerConfig{},
		Reporter: &config.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: time.Second * 1,
		},
	}
	tracer, closer, err := cfg.New(
		"Govoyage", config.Logger(jaeger.StdLogger),
	)
	opentracing.SetGlobalTracer(tracer)
	defer closer.Close()
	return err
}
