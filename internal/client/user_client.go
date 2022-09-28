package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type RefreshTokenRequest struct {
	ClientID     string `json:"client_id,omitempty"`
	ClientSecret string `json:"client_secret,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	GrantType    string `json:"grant_type,omitempty"`
}

const (
	POST              = "POST"
	client_id         = "dataprovider.gsuite.client_id"
	client_secret     = "dataprovider.gsuite.client_secret"
	grant_type        = "refresh_token"
	refresh_token_url = "dataprovider.gsuite.refresh_token_url"
)

func main() {
	payload, err := json.Marshal(RefreshTokenRequest{
		ClientID:     client_id,
		ClientSecret: client_secret,
		RefreshToken: refresh_token_url,
		GrantType:    grant_type,
	})
	if err != nil {
		log.Println(err)
	}

	rt := RefreshTokenRequest{}
	fmt.Println(json.Unmarshal(payload, &rt))
	fmt.Println(strings.NewReader(string(payload)))
}

// func getUserInput() *user.Model {
// 	var name string
// 	var subs string
// 	var id string

// 	id = uuid.New().String()
// 	fmt.Println("Enter Name")
// 	fmt.Scanf("%s", name)
// 	fmt.Println("Enter Subs")
// 	fmt.Scanf("%s", subs)

// 	return createUser(id, name, subs)
// }

// func createUser(id string, name string, subs string) *user.Model {
// 	return &user.Model{ID: id,
// 		Name:         name,
// 		Subscription: subs}
// }
