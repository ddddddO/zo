package screenshoter

import (
	"context"
	"fmt"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/pkg/errors"
)

type shoterWindows struct {
	cmd     string
	timeout time.Duration
}

func newShoterWindows(cmd string) (*shoterWindows, error) {
	if _, err := exec.LookPath(cmd); err != nil {
		return nil, errors.WithStack(err)
	}

	return &shoterWindows{
		cmd: cmd,
	}, nil
}

func (s *shoterWindows) withTimeout(t time.Duration) {
	s.timeout = t
}

func (s *shoterWindows) Capture() (string, error) {
	fmt.Println("Input screenshoted file path:")

	ctx := context.Background()
	if s.timeout != 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, s.timeout)
		defer cancel()
	}

	cmd := exec.CommandContext(ctx, s.cmd)
	if err := cmd.Run(); err != nil {
		return "", errors.WithStack(err)
	}

	// NOTE: ユーザーからスクショのパスをもらう
	var p string
	fmt.Scanln(&p)
	path, err := filepath.Abs(p) // FIXME:
	if err != nil {
		return "", errors.WithStack(err)
	}

	return path, nil
}
