package main

import (
	"log"

	"github.com/mattn/go-xmpp"
)

func main() {
	options := xmpp.Options{
		Host:     "example.com",
		User:     "user@example.com",
		Password: "password",
		NoTLS:    true,
		Debug:    true,
	}

	client, err := options.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	// Set initial presence status to "online"
	if err := client.Send(xmpp.NewPresence(xmpp.AvailableType)); err != nil {
		log.Fatal(err)
	}

	// Subscribe to presence updates from other users
	if err := client.Send(xmpp.NewPresence(xmpp.SubscribeType)); err != nil {
		log.Fatal(err)
	}

	for {
		stanza, err := client.Recv()
		if err != nil {
			log.Fatal(err)
		}

		switch v := stanza.(type) {
		case xmpp.Presence:
			// Update the presence store with the latest presence information
			updatePresenceStore(v.From, v.Type)
		}
	}
}

func updatePresenceStore(user string, presenceType string) {
	// Implement this function to update the presence store with the latest presence information for the given user
	// For example, you might use a database or in-memory data structure to store this information
}

//This code creates a new XMPP client and connects to the server at example.com. It then sets its own initial presence
//status to "online" and subscribes to presence updates from other users. Finally, it enters a loop to receive incoming
//stanzas (XML messages) from the server, and updates the presence store with the latest presence information whenever
//it receives a presence stanza.
//
//This is just a simple example, and there are many additional features and considerations you might want to add to a
//chat system that uses XMPP. For example, you might want to handle authentication, error handling, messaging, and more.
