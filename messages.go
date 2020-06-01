package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/Saur4ig/mescon"
)

func logInfo(port int, apps []Application) error {
	var sb strings.Builder
	sb.WriteString("Server started on:\n")
	sb.WriteString(fmt.Sprintf("port - %d\n", port))
	for i, val := range apps {
		sb.WriteString(fmt.Sprintf("app - %d; endpoint - %s; token - %s; key - %s\n", i, val.endpoint, val.token, val.secretParamKey))
	}
	message, err := mescon.GenAny(sb.String())
	if err != nil {
		return err
	}
	log.Println(message)
	return nil
}
