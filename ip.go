package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
)

type Host struct {
	Ip    string   `json:"ip"`
	Names []string `json:"hosts"`
}

var (
	flport = flag.String("port", "8080", "Listen port")
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	remote := strings.Split(r.Host, ":")
	names, err := net.LookupAddr(remote[0])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	rHost := Host{remote[0], names}
	rJson, err := json.Marshal(rHost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("%s %v", rHost.Ip, rHost.Names)
	fmt.Fprintf(w, "%s", rJson)
}

func main() {
	flag.Parse()
	http.HandleFunc("/", mainHandler)
	http.ListenAndServe(":"+*flport, nil)
}
