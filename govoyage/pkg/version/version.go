package version

import (
	"bytes"
	"fmt"
	"os"
	"runtime"
	"strings"
	"text/template"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

// Build information
var (
	Version   string
	Revision  string
	Branch    string
	BuildUser string
	BuildDate string
	AppName   = "govoyage"
	Uptime    = time.Now()
	GoVersion = runtime.Version()
	Platform  = runtime.GOOS + "/" + runtime.GOARCH
)

var versionInfoTmpl = `
{{.program}}, version {{.version}} (branch: {{.branch}}, revision: {{.revision}})
  build user:       {{.buildUser}}
  build date:       {{.buildDate}}
  go version:       {{.goVersion}}
  platform:         {{.platform}}
`

// Runtime returns server runtime information
func Runtime() map[string]any {
	return map[string]any{
		"app":    AppName,
		"pid":    os.Getpid(),
		"build":  Info(),
		"uptime": Uptime,
	}
}

// Info returns version and branch information
func Info() map[string]string {
	return map[string]string{
		"version":   Version,
		"branch":    Branch,
		"buildUser": BuildUser,
		"goVersion": GoVersion,
	}
}

// String returns version information string.
func String() string {
	m := map[string]string{
		"program":   AppName,
		"version":   Version,
		"revision":  Revision,
		"branch":    Branch,
		"buildUser": BuildUser,
		"buildDate": BuildDate,
		"goVersion": GoVersion,
		"platform":  Platform,
	}
	t := template.Must(template.New("version").Parse(versionInfoTmpl))

	var buf bytes.Buffer
	if err := t.ExecuteTemplate(&buf, "version", m); err != nil {
		panic(err)
	}
	return strings.TrimSpace(buf.String())
}

// NewCollector exports metrics about program build info
func NewCollector(program string) prometheus.Collector {
	return prometheus.NewGaugeFunc(
		prometheus.GaugeOpts{
			Namespace: program,
			Name:      "build_info",
			Help:      fmt.Sprintf("%s build info with platform and goversion", program),
			ConstLabels: prometheus.Labels{
				"job":       program,
				"branch":    Branch,
				"version":   Version,
				"revision":  Revision,
				"platform":  Platform,
				"goversion": GoVersion,
				"builduser": BuildUser,
			},
		},
		func() float64 { return 1 },
	)
}
