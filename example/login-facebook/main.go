package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"go.uber.org/zap"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	config := oauth2.Config{
		ClientID:     "",
		ClientSecret: "",
		RedirectURL:  "http://localhost:3000/facebook/redirect",
		Endpoint:     facebook.Endpoint,
		Scopes: []string{
			"email",
		},
	}

	// Redirect user to Facebook for login.
	url := config.AuthCodeURL("state", oauth2.AccessTypeOffline)
	http.HandleFunc("/facebook", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, url, http.StatusFound)
	})

	// Handle Facebook callback
	http.HandleFunc("/facebook/redirect", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		token, err := config.Exchange(context.Background(), code)
		if err != nil {
			logger.Error("Error exchanging token", zap.Error(err))
			http.Error(w, "Failed to exchange token", http.StatusInternalServerError)
			return
		}

		// Store token and possibly the user ID in the session.
		fmt.Fprintf(w, "Token: %v\n", token)
	})

	log.Fatal(http.ListenAndServe(":3000", nil))
}
