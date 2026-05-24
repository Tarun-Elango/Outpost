package cmd

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"devbox-cli/internal/api"
)

// snapshotItem represents a snapshot as returned by the API.
type snapshotItem struct {
	AmiID    string `json:"amiId"`
	Name     string `json:"name"`
	State    string `json:"state"`
	BoxAwsID string `json:"boxAwsId"`
}

// Snapshots dispatches snapshot sub-commands.
//
//	devbox snapshots              → list all user snapshots
//	devbox snapshots ls <boxId>   → list snapshots for a specific box
//	devbox snapshots delete <amiId> → delete a snapshot
func Snapshots(args []string) {
	if TestMode {
		fmt.Println("[test] snapshots: done")
		return
	}

	if len(args) == 0 {
		snapshotsList()
		return
	}

	sub := args[0]
	subArgs := args[1:]

	switch sub {
	case "ls":
		if len(subArgs) < 1 {
			fmt.Fprintln(os.Stderr, "usage: devbox snapshots ls <boxId>")
			os.Exit(1)
		}
		snapshotsListByBox(subArgs[0])
	case "delete":
		if len(subArgs) < 1 {
			fmt.Fprintln(os.Stderr, "usage: devbox snapshots delete <amiId>")
			os.Exit(1)
		}
		snapshotsDelete(subArgs[0])
	default:
		fmt.Fprintf(os.Stderr, "snapshots: unknown sub-command %q\n", sub)
		fmt.Fprintln(os.Stderr, "usage: devbox snapshots [ls <boxId> | delete <amiId>]")
		os.Exit(1)
	}
}

func snapshotsList() {
	client, err := api.NewDefault()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	resp, err := client.Get("/v1/snapshots")
	if err != nil {
		fmt.Fprintf(os.Stderr, "snapshots failed: %v\n", err)
		os.Exit(1)
	}
	if err := api.CheckStatus(resp); err != nil {
		fmt.Fprintf(os.Stderr, "snapshots failed: %v\n", err)
		os.Exit(1)
	}

	var items []snapshotItem
	if err := api.DecodeJSON(resp, &items); err != nil {
		fmt.Fprintf(os.Stderr, "snapshots failed: %v\n", err)
		os.Exit(1)
	}

	if len(items) == 0 {
		fmt.Println("No snapshots found.")
		return
	}

	printSnapshotTable(items)
}

func snapshotsListByBox(boxID string) {
	client, err := api.NewDefault()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	resp, err := client.Get("/v1/boxes/" + boxID + "/snapshots")
	if err != nil {
		fmt.Fprintf(os.Stderr, "snapshots failed: %v\n", err)
		os.Exit(1)
	}
	if err := api.CheckStatus(resp); err != nil {
		fmt.Fprintf(os.Stderr, "snapshots failed: %v\n", err)
		os.Exit(1)
	}

	var items []snapshotItem
	if err := api.DecodeJSON(resp, &items); err != nil {
		fmt.Fprintf(os.Stderr, "snapshots failed: %v\n", err)
		os.Exit(1)
	}

	if len(items) == 0 {
		fmt.Printf("No snapshots found for box %s.\n", boxID)
		return
	}

	printSnapshotTable(items)
}

func snapshotsDelete(amiID string) {
	client, err := api.NewDefault()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	resp, err := client.Delete("/v1/snapshots/" + amiID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "snapshot delete failed: %v\n", err)
		os.Exit(1)
	}

	if resp.StatusCode == http.StatusNotFound {
		resp.Body.Close()
		fmt.Fprintf(os.Stderr, "snapshot %s not found\n", amiID)
		os.Exit(1)
	}

	if err := api.CheckStatus(resp); err != nil {
		fmt.Fprintf(os.Stderr, "snapshot delete failed: %v\n", err)
		os.Exit(1)
	}
	resp.Body.Close()

	fmt.Printf("Snapshot %s deleted.\n", amiID)
}

func printSnapshotTable(items []snapshotItem) {
	fmt.Printf("%-24s  %-20s  %-12s  %s\n", "AMI ID", "NAME", "STATE", "BOX ID")
	fmt.Println(strings.Repeat("-", 90))
	for _, s := range items {
		fmt.Printf("%-24s  %-20s  %-12s  %s\n", s.AmiID, s.Name, s.State, s.BoxAwsID)
	}
}
