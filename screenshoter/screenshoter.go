package screenshoter

import (
	"runtime"
	"time"

	"github.com/pkg/errors"
)

type screenshoter interface {
	Capture() (filepath string, err error)
}

// ref: https://github.com/migueldemoura/myazo/blob/1c191360eb00ac30f5039828b26933b112a6f542/client/src/myazo.py#L37-L51
var cmds = map[string]string{
	"windows": "snippingtool",
	"darwin":  "screencapture",
	"linux":   "gnome-screenshot",
}

var ErrNoImpl = errors.New("implementation is not exsist")

func New(options ...optFunc) (screenshoter, error) {
	option := &option{}
	for _, opt := range options {
		opt(option)
	}

	switch runtime.GOOS {
	case "windows":
		shoter, err := newShoterWindows(cmds[runtime.GOOS])
		if err != nil {
			return nil, errors.WithStack(err)
		}
		shoter.withTimeout(option.timeout)
		return shoter, nil
	case "darwin":
		// TODO:
		return nil, ErrNoImpl
	case "linux":
		// TODO:
		return nil, ErrNoImpl
	default:
		return nil, ErrNoImpl
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
