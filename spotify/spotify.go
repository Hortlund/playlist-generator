package spotify

import (
    "fmt"
	"net/http"
)

func Apicall(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love spotify")
}
