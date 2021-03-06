package logger

import (
	// "fmt"
	check "gopkg.in/check.v1"
	"testing"
	"strings"
	logger "github.com/tiny-libs/logger-go"
	capturer "github.com/kami-zh/go-capturer"
)

var _ = check.Suite(new(Suite))

func Test(t *testing.T) {
	check.TestingT(t)
}

type Suite struct{}

func (s *Suite) TestLog(ch *check.C) {
	out := capturer.CaptureStdout(func() {
		logger.Log("foo", "boo", "acme")
	})
	parts := strings.Fields(out)

	ch.Check(parts[2], check.Equals, "logger_test.go:22")
	ch.Check(parts[3], check.Equals, "\x1b[90m[LOG]\x1b[0m")
	ch.Check(strings.Join(parts[4:], " "), check.Equals, "foo boo acme")
}
func (s *Suite) TestInfo(ch *check.C) {
	out := capturer.CaptureStdout(func() {
		logger.Info("foo", "boo", "acme")
	})
	parts := strings.Fields(out)

	ch.Check(parts[2], check.Equals, "logger_test.go:32")
	ch.Check(parts[3], check.Equals, "\x1b[36m[INFO]\x1b[0m")
	ch.Check(strings.Join(parts[4:], " "), check.Equals, "foo boo acme")
}
func (s *Suite) TestWarning(ch *check.C) {
	out := capturer.CaptureStdout(func() {
		logger.Warning("foo", "boo", "acme")
	})
	parts := strings.Fields(out)

	ch.Check(parts[2], check.Equals, "logger_test.go:42")
	ch.Check(parts[3], check.Equals, "\x1b[93m[WARN]\x1b[0m")
	ch.Check(strings.Join(parts[4:], " "), check.Equals, "foo boo acme")
}
func (s *Suite) TestError(ch *check.C) {
	out := capturer.CaptureStdout(func() {
		logger.Error("foo", "boo", "acme")
	})
	parts := strings.Fields(out)

	ch.Check(parts[2], check.Equals, "logger_test.go:52")
	ch.Check(parts[3], check.Equals, "\x1b[31m[ERROR]\x1b[0m")
	ch.Check(strings.Join(parts[4:], " "), check.Equals, "foo boo acme")
}
func (s *Suite) TestDump(ch *check.C) {
	// out := capturer.CaptureStdout(func() {
	logger.Dump("foo", "boo", "acme")
	// })
	// parts := strings.Fields(out)

	// fmt.Println(parts)
	// ch.Check(parts[2], check.Equals, "logger_test.go:52")
	// ch.Check(parts[3], check.Equals, "\x1b[31m[ERROR]\x1b[0m")
	// ch.Check(strings.Join(parts[4:], " "), check.Equals, "foo boo acme")
}