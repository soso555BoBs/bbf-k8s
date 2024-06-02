package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	f, err := os.Open("/etc/config/myconfig.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	defer f.Close()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, f)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, buf.String())
	})

	// port 取得
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}

	log.Println("Starting server on port " + port)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
