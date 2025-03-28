package slogtest_test

import (
	"testing"

	"github.com/khulnasoft/clog"
	"github.com/khulnasoft/clog/slogtest"
)

func TestSlogTest(t *testing.T) {
	ctx := slogtest.Context(t)

	clog.FromContext(ctx).With("foo", "bar").Infof("hello world")
	clog.FromContext(ctx).With("bar", "baz").Infof("me again")
	clog.FromContext(ctx).With("baz", true).Infof("okay last one")

	clog.FromContext(ctx).Debug("hello debug")
	clog.FromContext(ctx).Info("hello info")
	clog.FromContext(ctx).Warn("hello warn")
	clog.FromContext(ctx).Error("hello error")
}
