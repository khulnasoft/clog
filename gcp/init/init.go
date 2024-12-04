package init

import (
	"log/slog"
	"os"

	"github.com/khulnasoft/clog"
	"github.com/khulnasoft/clog/gcp"
)

// Set up structured logging at Info+ level.
func init() {
	level := slog.LevelInfo
	if e, ok := os.LookupEnv("LOG_LEVEL"); ok {
		if err := level.UnmarshalText([]byte(e)); err != nil {
			clog.Fatalf("slog: invalid log level: %v", err)
		}
	}
	slog.SetDefault(slog.New(gcp.NewHandler(level)))
}
