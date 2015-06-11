package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		_, header, _ := r.FormFile("file")
		file, _ := header.Open()
		path := fmt.Sprintf("files/%s", header.Filename)
		buf, _ := ioutil.ReadAll(file)
		ioutil.WriteFile(path, buf, 0644)
		http.Redirect(w, r, "/"+path, 301)
	} else {
		http.Redirect(w, r, "/", 301)
	}
}
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<html>
    <head>
        <title> Dank Upload Page</title>
    </head>
    <body>
        <form action="/upload" method="post" enctype="multipart/form-data">
            <input type="file" id="file" name="file">
            <input type="submit" name="submit" value="submit">
        </form>
    </body>
</html>`)
}
func main() {
	staticServer := http.StripPrefix("/files/", http.FileServer(http.Dir("files/")))
	http.HandleFunc("/", index)
	http.HandleFunc("/upload", upload)
	http.Handle("/files/", staticServer)
	panic(http.ListenAndServe(":8080", nil))
}
