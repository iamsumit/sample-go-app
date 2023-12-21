package slogger_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/iamsumit/sample-go-app/pkg/slogger"
)

func TestSLogger(t *testing.T) {
	tests := []struct {
		Name string
		F    func(buf *bytes.Buffer)
		Want string
	}{
		{
			Name: "TextLogger",
			F: func(buf *bytes.Buffer) {
				buf.Reset()
				l := slogger.New(buf, true).TextLogger()
				l.Info("text-formatted-log", "key", "value")
			},
			Want: "msg=text-formatted-log key=value",
		},
		{
			Name: "WithTintLogger",
			F: func(buf *bytes.Buffer) {
				buf.Reset()
				l := slogger.New(buf, true).TintLogger()
				l.Info("tint-formatted-log", "key", "value")
			},
			Want: "tint-formatted-log",
		},
		{
			Name: "WithJSONLogger",
			F: func(buf *bytes.Buffer) {
				buf.Reset()
				l := slogger.New(buf, true).JSONLogger()
				l.Info("json-formatted-log", "key", "value")
			},
			Want: "\"msg\":\"json-formatted-log\",\"key\":\"value\"",
		},
		{
			Name: "WithTextGroupLogger",
			F: func(buf *bytes.Buffer) {
				buf.Reset()
				l := slogger.New(buf, true).TextLogger()
				l = slogger.WithGroup(l)
				l.Info("group-log", "key", "value")
			},
			Want: "program_info.go_version=go1",
		},
	}

	var buf bytes.Buffer
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			tt.F(&buf)
			if result := buf.String(); !strings.Contains(result, tt.Want) {
				t.Errorf("Expected log message '%s' to contain '%s'", result, tt.Want)
			}
		})
	}
}
