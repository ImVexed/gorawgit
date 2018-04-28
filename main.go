package main

import (
	"fmt"
	"io"
	"mime"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc(`/`, serveFile)
	if err := http.ListenAndServe(`:80`, nil); err != nil {
		panic(err)
	}
}

func serveFile(w http.ResponseWriter, r *http.Request) {
	fileExt := strings.Split(r.URL.Path, `.`)

	if fileExt == nil {
		w.WriteHeader(400)
		fmt.Println(`Failed to parse URL path: `, r.URL.Path)
		return
	}
	fileMime := mime.TypeByExtension(`.` + fileExt[len(fileExt)-1])

	w.Header().Set(`Content-Type`, fileMime)

	fileData, err := http.Get(`https://raw.githubusercontent.com` + r.URL.Path)

	if err != nil {
		w.WriteHeader(400)
		fmt.Println(err)
		return
	}

	defer fileData.Body.Close()

	_, err = io.Copy(w, fileData.Body)

	if err != nil {
		w.WriteHeader(500)
		fmt.Println(err)
		return
	}
	fmt.Println(`Serving `, r.URL.Path, ` to `, r.RemoteAddr)
}
