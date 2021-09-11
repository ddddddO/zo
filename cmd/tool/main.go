package main

import (
	"log"

	"github.com/ddddddO/zo/executor"
	"github.com/ddddddO/zo/screenshoter"
	"github.com/ddddddO/zo/storage"
)

func main() {
	const (
		gcsBucket = "testddddddo"
		// TODO: ユーザー毎に取得するようにする
		adcpath = `\\wsl$\Debian\home\ochi\.config\gcloud\legacy_credentials\lbfdeatq@gmail.com\adc.json`
	)
	sh := screenshoter.New()
	st := storage.NewGCS(gcsBucket, adcpath)
	ex := executor.New(sh, st)

	if err := ex.Execute(); err != nil {
		log.Fatalf("%+v", err)
	}
}
