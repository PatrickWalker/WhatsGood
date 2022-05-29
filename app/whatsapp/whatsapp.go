package whatsapp

import (
	"fmt"
	"net/http"
	"time"
)

//create the client with the key for config
type WhatsAppClient struct {
	PhoneID    string
	BusinessID string
	Token      string
	Template   string
	HttpClient *http.Client
}

func NewClient(phoneID, businessID, token, template string) *WhatsAppClient {
	return &WhatsAppClient{
		PhoneID:    phoneID,
		BusinessID: businessID,
		Token:      token,
		Template:   template,
		HttpClient: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}

// Send is used to send the message with the template with recipients and vars
func (wc *WhatsAppClient) Send(recipients []string, vars map[string]string) {
	fmt.Printf("Token %s vars %v \n", wc.Token, vars)
}

//send the message using a template, variables and the numbers?
