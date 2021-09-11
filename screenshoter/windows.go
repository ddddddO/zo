package screenshoter

import (
	"context"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/pkg/errors"
)

type shoterWindows struct {
	cmd     string
	timeout time.Duration
}

func newShoterWindows(cmd string) *shoterWindows {
	if _, err := exec.LookPath(cmd); err != nil {
		panic(err)
	}

	return &shoterWindows{
		cmd: cmd,
	}
}

func (s *shoterWindows) withTimeout(t time.Duration) {
	s.timeout = t
}

func (s *shoterWindows) Capture() (string, error) {
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

	// TODO: どうすれば保存されたスクショのパスを取得できるか。一旦固定
	// 案１：ここで、標準入力受け取る処理をして、ユーザーからスクショのパスをもらう
	dir := `C:\Users\lbfde\OneDrive\画像\スクリーンショット`
	name := "キャプチャ.PNG"
	return filepath.Join(dir, name), nil
}
