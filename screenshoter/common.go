package screenshoter

import (
	"runtime"
	"time"

	"github.com/ddddddO/zo/executor"
)

// ref: https://github.com/migueldemoura/myazo/blob/1c191360eb00ac30f5039828b26933b112a6f542/client/src/myazo.py#L37-L51
var cmds = map[string]string{
	"windows": "snippingtool",
	"darwin":  "screencapture",
	"linux":   "gnome-screenshot",
}

func New(options ...optFunc) executor.Screenshoter {
	option := &option{}
	for _, opt := range options {
		opt(option)
	}

	switch runtime.GOOS {
	case "windows":
		shoter := newShoterWindows(cmds[runtime.GOOS])
		shoter.withTimeout(option.timeout)
		return shoter
	case "darwin":
		// TODO:
		panic("not yet supported!")
	case "linux":
		// TODO:
		panic("not yet supported!")
	default:
		panic("not supported!")
	}
}

type option struct {
	timeout time.Duration
}

type optFunc func(*option)

func WithTimeout(t time.Duration) optFunc {
	return func(o *option) {
		o.timeout = t
	}
}
