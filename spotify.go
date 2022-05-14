package spotify

import (
    "fmt"
    "net/http"
    "log"
    "strings"
)

func Apicall(a string) response string {
	
	// response, err := http.Get(fmt.Sprintf("https://api.spotify.com/v1/recommendations?limit=%s&market=SE&seed_genres=%v", songs, genres))
	response := a
	return response
}