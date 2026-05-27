package cmd

import (
	"fmt"
	"strings"
	"os"

	"devbox-cli/internal/api"

)
type Template struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	StartupScript string `json:"startupScript"`
}

func Templates(args []string) {
	fmt.Println("Fetching Templates")
	if TestMode {
		fmt.Println("[test] templates: done")
		return
	}

	client, err := api.NewDefault()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	resp, err := client.Get("/v1/boxes/templates")
	if err != nil {
		api.FailBox("templates", err)
	}
	
	if err := api.CheckStatus(resp); err != nil {
		api.FailBox("templates", err)
	}

	// build a table of templates
	var templates []Template
	if err := api.DecodeJSON(resp, &templates); err != nil { // decode the response body into the templates slice
		api.FailBox("templates", err)
	}

	fmt.Printf("%-24s  %-20s  %-10s\n", "ID", "NAME", "DESCRIPTION")
	fmt.Println(strings.Repeat("-", 100))
	for _, t := range templates {
		fmt.Printf("%-24s  %-20s  %-10s\n", t.ID, t.Name, t.Description)
	}
}