package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func readFile(f string) string {
	file, err := os.Open(f)
	if err != nil {
		// handle the error here
		return ""
	}
	defer file.Close()
	// get the file size
	stat, err := file.Stat()
	if err != nil {
		return ""
	}
	// read the file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)

	if err != nil {
		return ""
	}
	//returns the content of the file
	return string(bs)
}

func index(res http.ResponseWriter, req *http.Request) {
	pageData := readFile("index.html")
	// sets the content type
	res.Header().Set(
		"Content-Type",
		"text/html; charset=utf-8",
	)
	//rendes the sting as an html page.
	io.WriteString(
		res,
		pageData,
	)
}

func main() {

	port := ":9000"
	http.HandleFunc("/", index)
	//out put on command line to indicate the server is running.
	fmt.Println("Server running on port", port)
	http.ListenAndServe(port, nil)

}
