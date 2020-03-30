package handler 

import (
	"fmt"
	"net/http"
)

var name string

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello " + name)
}

func init() {
	name = "allen"
}


