package main

import (
	"encoding/json"

	webpush "github.com/kushaltirumala/webpush-go"
)

const (
	subscription    = `{"subscription":{"endpoint":"https://fcm.googleapis.com/fcm/send/cmoXViADyTQ:APA91bFAI3R9cTbYtdgL-FHS1Af7lu53E1TSQ6MD4nUKwv6CukcRAXEgfLuoQv3evBIyzaPM8X3cfo2WyEB5OpbnBb6pocu2BZtHLONuxtDf2lozCgg7Tv7u142wsLW3MYRICEu4am5n","expirationTime":null,"keys":{"p256dh":"BJCpYiqJ2foeUfbRfr0B4bfeXcJYwhuK19PM7lGQuW3TkNDORz9vdkpAoQCrE0snWXq-YP3qPJdAd8100w7NiOA","auth":"txyIaJGBXcz9Bj4z9WtreQ"}}}`
	vapidPublicKey  = "BG3BBbmoXKl89s1bX5_iK2KPlxZkiOaYVpC9dzO7Wjm7kSw7FFtI9YmoeIcO8Bd-ktm3Am1qKD6isrEbB-C1aDg"
	vapidPrivateKey = "lNs1xvkU9rzDmwm1b5OxgbfIRQD5t9ZbplSWZsvQi4I"
)

func main() {
	// Decode subscription
	s := &webpush.Subscription{}
	json.Unmarshal([]byte(subscription), s)

	// Send Notification
	resp, err := webpush.SendNotification([]byte("Test"), s, &webpush.Options{
		Subscriber:      "example@example.com", // Do not include "mailto:"
		VAPIDPublicKey:  vapidPublicKey,
		VAPIDPrivateKey: vapidPrivateKey,
		TTL:             30,
	})
	if err != nil {
		// TODO: Handle error
	}
	defer resp.Body.Close()
}
