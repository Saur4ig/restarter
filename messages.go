package main

import (
	"fmt"
	"log"

	"github.com/Saur4ig/mescon"
)

func logInfo(port int, endpoint, token, tokenQueryKey string) error {
	m := fmt.Sprintf("Server started on:\nport - %d\nendpoint - %s\ntoken - %s\nkey - %s", port, endpoint, token, tokenQueryKey)
	message, err := mescon.GenMultiLineMessage(50, m, "")
	if err != nil {
		return err
	}
	log.Println(message)
	return nil
}
