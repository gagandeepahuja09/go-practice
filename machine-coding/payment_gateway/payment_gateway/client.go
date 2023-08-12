package payment_gateway

import (
	"errors"
	"log"

	"github.com/google/uuid"
)

const (
	ErrClientNotFound = "client not found"
)

type Client struct {
	id       string
	name     string
	isActive bool
}

func AddClient(name string) *Client {
	return &Client{
		id:       uuid.NewString(),
		name:     name,
		isActive: true,
	}
}

func RemoveClient(c *Client) {
	c.isActive = false
}

func HasClient(c *Client) bool {
	return c != nil && c.isActive
}

func AddSupportForMethods(c *Client, pmList []PaymentMethod) error {
	if !HasClient(c) {
		return errors.New(ErrClientNotFound)
	}
	methodsSupported := pg.clientMethodsSupported[c]
	if methodsSupported == nil {
		pg.clientMethodsSupported[c] = pmList
	} else {
		pg.clientMethodsSupported[c] = append(methodsSupported, pmList...)
	}
	return nil
}

func ListSupportedMethods(c *Client) error {
	if !HasClient(c) {
		return errors.New(ErrClientNotFound)
	}
	log.Printf("METHODS_SUPPORTED: %+v\n", pg.clientMethodsSupported[c])
	return nil
}
