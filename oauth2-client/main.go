package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const clientID = "7995a42830508914de35"
const clientSecret = "fd1e487f24f3945c6a8f0209467fc8f6d12c3762"

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	httpClient := http.Client{}

	http.HandleFunc("/oauth/redirect", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			fmt.Fprintf(os.Stdout, "could not parse query: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		code := r.FormValue("code")

		reqURL := fmt.Sprintf(
			"https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s",
			clientID, clientSecret, code,
		)
		req, err := http.NewRequest(http.MethodPost, reqURL, nil)
		if err != nil {
			fmt.Fprintf(os.Stdout, "could not create HTTP request : %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		req.Header.Set("accept", "application/json")

		res, err := httpClient.Do(req)
		if err != nil {
			fmt.Fprintf(os.Stdout, "could not send HTTP request : %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer res.Body.Close()

		var t OauthAccessResponse
		if err := json.NewDecoder(res.Body).Decode(&t); err != nil {
			fmt.Fprintf(os.Stdout, "could not parse JSON response : %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.Header().Set("Location", "/welcome.html?access_token="+t.AccessToken)
		w.WriteHeader(http.StatusFound)

	})

	fmt.Println("Running server on localhost:8080")
	http.ListenAndServe(":8080", nil)
}

type OauthAccessResponse struct {
	AccessToken string `json:"access_token"`
}
