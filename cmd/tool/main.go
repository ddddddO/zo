package main

import (
	"log"

	"github.com/ddddddO/zo/screenshoter"
	"github.com/ddddddO/zo/storage"
	"github.com/ddddddO/zo/tool"
)

func main() {
	const (
		gcsBucket = "testddddddo"
		// TODO: ユーザー毎に取得するようにする
		adcpath = `\\wsl$\Debian\home\ochi\.config\gcloud\legacy_credentials\lbfdeatq@gmail.com\adc.json`
	)
	sh, err := screenshoter.New()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	st, err := storage.NewGCSWithCredential(gcsBucket, adcpath)
	if err != nil {
		log.Fatalf("%+v", err)
	}
	ex := tool.NewExecutor(sh, st)

	if err := ex.Execute(); err != nil {
		log.Fatalf("%+v", err)
	}
}
