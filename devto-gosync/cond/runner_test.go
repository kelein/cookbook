package cond

import "testing"

func Test_startRunner(t *testing.T) { startRunner() }

func BenchmarkStartRunner(b *testing.B) {
	for b.Loop() {
		startRunner()
	}
	b.ReportAllocs()
}
