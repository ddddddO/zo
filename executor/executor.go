package executor

import (
	"fmt"

	"github.com/pkg/errors"
)

type Screenshoter interface {
	Capture() (filepath string, err error)
}

type storage interface {
	Upload(filepath string) (hashedFileName string, err error)
	GetURL(hashedFileName string) (url string, err error)
	Close() error
}

type executor struct {
	shoter  Screenshoter
	storage storage
}

func New(shoter Screenshoter, storage storage) *executor {
	return &executor{
		shoter:  shoter,
		storage: storage,
	}
}

func (e *executor) Execute() error {
	defer e.storage.Close()

	filepath, err := e.shoter.Capture()
	if err != nil {
		return errors.WithStack(err)
	}

	hashedFileName, err := e.storage.Upload(filepath)
	if err != nil {
		return errors.WithStack(err)
	}

	url, err := e.storage.GetURL(hashedFileName)
	if err != nil {
		return errors.WithStack(err)
	}

	// TODO: 各OSでクリップボードにコピーできればいいと思う
	fmt.Println(url)

	return nil
}
