package dto

import (
	"encoding/json"
	"fmt"
)

const (
	AddSubscription    = "AddSubscription"
	DeleteSubscription = "DeleteSubscription"
)

type addressMessage struct {
	Subscriptions map[string][]string `json:"subscriptions"`
	Operation     string              `json:"operation"`
}

func NewAddressMessage() *addressMessage {
	return &addressMessage{}
}

func (a *addressMessage) ToBytes() []byte {
	bytes, _ := json.Marshal(a)
	fmt.Println(bytes)

	newAddressMessageEvent := &addressMessage{}
	json.Unmarshal(bytes, newAddressMessageEvent)
	fmt.Println(newAddressMessageEvent)

	return bytes
}
