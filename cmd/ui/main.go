package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	zoStorage "github.com/ddddddO/zo/storage"
)

var (
	bucketName string
	gcs        storage
)

func init() {
	bucketName = os.Getenv("BUCKET_NAME")
	if bucketName == "" {
		log.Fatal("Not setup BUCKET_NAME")
	}

	// TODO: cloud runのこのサービス専用サービスアカウント作ってそれでgcsにアクセスするようにする。
	gcs = zoStorage.NewGCS(bucketName)
}

type storage interface {
	GetURLs() (urls []string, err error)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Listening on port %s", port)

	http.HandleFunc("/", indexHandler(bucketName, gcs))

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func indexHandler(bucketName string, gcs storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		urls, err := gcs.GetURLs()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "failed to GetURLs")
			return
		}

		const link = `<a href="%s">%s</a><br>`
		links := ""
		for _, u := range urls {
			links += fmt.Sprintf(link, u, u)
		}
		indexHTML := fmt.Sprintf(htmlTemplate, links)

		fmt.Fprint(w, indexHTML)
	}
}

const htmlTemplate = `
<!DOCTYPE html>
<html>
<head>
    <meta name="viewport" content="width=device-width,initial-scale=1">
    <title>zo</title>
</head>
<body>
    <div id="main">
        %s
    </div>
</body>
</html>
`
