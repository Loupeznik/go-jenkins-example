package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
	"time"
)

func TestMain(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	cases := []struct {
		Name           string
		Args           []string
		ExpectedExit   int
		ExpectedOutput string
	}{
		{"30 days", []string{"30"}, 0, time.Now().AddDate(0, 0, -30).Local().Format(time.RFC1123)[:16]},
		{"30 days without argument", []string{}, 0, time.Now().AddDate(0, 0, -30).Local().Format(time.RFC1123)[:16]},
		{"60 days", []string{"60"}, 0, time.Now().AddDate(0, 0, -60).Local().Format(time.RFC1123)[:16]},
		{"invalid argument", []string{"test"}, 1, "Argument `test` is not a number"},
	}

	for _, tc := range cases {
		os.Args = append([]string{tc.Name}, tc.Args...)

		var output bytes.Buffer

		exitCode := renderResult(&output)

		if tc.ExpectedExit != exitCode {
			t.Errorf("Wrong exit code for args: %v, expected: %v, got: %v",
				tc.Args, tc.ExpectedExit, exitCode)
		}

		if !strings.Contains(output.String(), tc.ExpectedOutput) {
			t.Errorf("Wrong output for args: %v, expected %v, got: %v",
				tc.Args, tc.ExpectedOutput, output.String())
		}
	}
}
