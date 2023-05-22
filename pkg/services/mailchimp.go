package services

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/dnnybanh/goxen_backend/pkg/models"
)

type Mailchimp struct {
	APIKey  string
	ListID  string
	BaseURL string
}

type Subscriber struct {
	EmailAddress string `json:"email_address"`
	Status       string `json:"status"`
	FirstName    string `json:"merge_fields,omitempty"`
	LastName     string `json:"merge_fields,omitempty"`
}

func (m *Mailchimp) SubscribeUser(user models.User) error {
	// Create a new subscriber
	subscriber := &Subscriber{
		EmailAddress: user.Email,
		Status:       "subscribed",
		FirstName:    user.FirstName,
		LastName:     user.LastName,
	}

	// Encode the subscriber as JSON
	body, err := json.Marshal(subscriber)
	if err != nil {
		return err
	}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", m.BaseURL+"/lists/"+m.ListID+"/members", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	// Set the headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "apikey "+m.APIKey)

	// Send the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// TODO: Check the HTTP response and return an error if it's not a 2xx status code

	return nil
}
