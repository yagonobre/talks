package main

import (
	"encoding/json"
	"log"
	"net/http"

	auth "k8s.io/client-go/pkg/apis/authentication/v1beta1"
)

const (
	token string = "mytoken"
	user  string = "foo"
	group string = "bar"
)

func authenticationError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"apiVersion": "authentication.k8s.io/v1beta1",
		"kind":       "TokenReview",
		"status": auth.TokenReviewStatus{
			Authenticated: false,
		},
	})
}

func authenticationSucess(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	trs := auth.TokenReviewStatus{
		Authenticated: true,
		User: auth.UserInfo{
			Username: user,
			UID:      group,
		},
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"apiVersion": "authentication.k8s.io/v1beta1",
		"kind":       "TokenReview",
		"status":     trs,
	})
}

func authentication(w http.ResponseWriter, r *http.Request) {
	var tr auth.TokenReview
	err := json.NewDecoder(r.Body).Decode(&tr)
	if err != nil {
		return
	}

	if tr.Spec.Token != token {
		authenticationError(w, r)
		return
	}
	authenticationSucess(w, r)
}

func main() {
	http.HandleFunc("/authenticate", authentication)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
