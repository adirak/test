// Package p contains an HTTP Cloud Function.
package p

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"

	"example.com/cloudfunction/loop"
)

// HelloWorld prints the JSON encoded "message" field in the body
// of the request or "Hello, World!" if there isn't one.
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	var d struct {
		Message string `json:"message"`
	}

	word := loop.GenWord()

	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		fmt.Fprint(w, word)
		return
	}
	if d.Message == "" {
		fmt.Fprint(w, word)
		return
	}
	fmt.Fprint(w, html.EscapeString(d.Message))
}
