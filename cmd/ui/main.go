package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	zoStorage "github.com/ddddddO/zo/storage"
)

var (
	gcs storage
)

func init() {
	bucketName := os.Getenv("BUCKET_NAME")
	if bucketName == "" {
		log.Fatal("Not setup BUCKET_NAME")
	}

	// TODO: cloud runのこのサービス専用サービスアカウント作ってそれでgcsにアクセスするようにする。
	var err error
	gcs, err = zoStorage.NewGCS(bucketName)
	if err != nil {
		log.Fatal(err)
	}
}

type storage interface {
	GetAttrs() (attrs [][4]string, err error)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Listening on port %s", port)

	http.HandleFunc("/", indexHandler(gcs))

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func indexHandler(gcs storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		// name/url/owner/created
		attrs, err := gcs.GetAttrs()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "failed to GetAttrs")
			return
		}

		table := buildTable(attrs)
		indexHTML := fmt.Sprintf(htmlTemplate, table)
		fmt.Fprint(w, indexHTML)
	}
}

func buildTable(attrs [][4]string) string {
	table := ""
	const tmpl = `<tr><th scope="row">%d</th><td><a href="%s">%s</a></td><td>%s</td><td>%s</td></tr>`
	for i, attr := range attrs {
		no := i + 1
		name := attr[0]
		url := attr[1]
		user := attr[2]
		created := attr[3]
		table += fmt.Sprintf(tmpl, no, url, name, user, created)
	}
	return table
}

const htmlTemplate = `
<!DOCTYPE html>
<html>
<head>
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.2.1/css/bootstrap.min.css" integrity="sha384-GJzZqFGwb1QTTN6wy59ffF1BuGJpLSa9DkKMp0DgiMDm4iYMj70gZWKYbI706tWS" crossorigin="anonymous">
    <meta name="viewport" content="width=device-width,initial-scale=1">
    <title>zo</title>
</head>
<body>
    <div id="main">
        <table class="table table-striped table-dark">
            <thead>
              <tr>
                <th scope="col">#</th>
                <th scope="col">Name</th>
                <th scope="col">User</th>
                <th scope="col">Created</th>
              </tr>
            </thead>
            <tbody>
                %s
            </tbody>
          </table>
    </div>
</body>
</html>
`
