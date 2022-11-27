package tests

import (
	"bytes"
	"io"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := &bytes.Buffer{}
	type args struct {
		writer io.Writer
		name   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"A", args{buffer, "Torres"}, "Hello, Torres"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Greet(tt.args.writer, tt.args.name)
			got := buffer.String()
			if got != tt.want {
				t.Errorf("Greet() got = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestCountdown(t *testing.T) {
	type args struct {
		seconds int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"A", args{3}, `3
2
1
Go!`},
		{"B", args{0}, `Go!`},
		{"C", args{-3}, `Go!`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := &bytes.Buffer{}
			Countdown(buf, tt.args.seconds)
			got := buf.String()
			if got != tt.want {
				t.Errorf("Countdown() got = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestCountdownWithSleeper(t *testing.T) {
	type args struct {
		seconds int
	}
	tests := []struct {
		name       string
		args       args
		wantWriter string
	}{
		{"A", args{3}, `3
2
1
Go!`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := &bytes.Buffer{}
			sleeper := &SpySleeper{}
			CountdownWithSleeper(writer, sleeper)
			if gotWriter := writer.String(); gotWriter != tt.wantWriter {
				t.Errorf("CountdownWithSleeper() = %v, want %v", gotWriter, tt.wantWriter)
			}

			if sleeper.Calls != 4 {
				t.Errorf("CountdownWithSleeper() sleeper calls = %v, want 4", sleeper.Calls)
			}
		})
	}
}
