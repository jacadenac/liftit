package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, htmlBody)
	})
	http.ListenAndServe(":5000", nil)
}

const htmlBody = `
<!DOCTYPE html>
<html>
  <head>
    <style>
      html, body {
        height: 100%;
        margin: 0;
        padding: 0;
      }
      #map {
        height: 100%;
      }
    </style>
  </head>
  <body>
    <h1>Prueba Backend Lifttit</h1>
  </body>
</html>
`
