package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getIP(w http.ResponseWriter, r *http.Request) {
	ip := r.RemoteAddr

	cleanip := strings.Split(ip, ":")
	fmt.Println(cleanip[0])

	io.WriteString(w, cleanip[0])
}

func getIndex(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "aui.laker.tech - see docs")
}

func main() {
	http.HandleFunc("/", getIndex)
	http.HandleFunc("/ip", getIP)

	err := http.ListenAndServe(":3333", nil)
	println(err)
}
