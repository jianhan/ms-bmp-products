package db

import "cloud.google.com/go/firestore"

type Base struct {
	client *firestore.Client
	path   string
}

func (b *Base) Close() error {
	return b.client.Close()
}
