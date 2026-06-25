package cmd

import (
	"fmt"
	"os"
	"strings"

	"devbox-cli/internal/api"
	"devbox-cli/service"
)

const templateNewUsageLine = "usage: devbox template new <name> [command string]"

// Template dispatches template sub-commands.
//
//	devbox template new <name> [command string] → create a template
//	devbox template delete <id>                  → delete a template
//	devbox template rename <id> <new-name>       → rename a template
func Template(args []string) {
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "usage: devbox template <new|delete|rename> ...")
		os.Exit(1)
	}

	sub := args[0]
	subArgs := args[1:]

	switch sub {
	case "new":
		TemplateNew(subArgs)
	case "delete":
		TemplateDelete(subArgs)
	case "rename":
		TemplateRename(subArgs)
	default:
		fmt.Fprintf(os.Stderr, "error: unknown subcommand %q (expected %q, %q, or %q)\n", sub, "new", "delete", "rename")
		os.Exit(1)
	}
}

// ParseTemplateNewArgs parses arguments after the "template new" subcommand.
// Expected shape: <name> [command parts...]
func ParseTemplateNewArgs(args []string) (name, startupScript string, err error) {
	if len(args) == 0 {
		return "", "", fmt.Errorf("template name is required")
	}

	name = strings.TrimSpace(args[0])
	if name == "" {
		return "", "", fmt.Errorf("template name is required")
	}
	if strings.HasPrefix(name, "--") {
		return "", "", fmt.Errorf("template name cannot be a flag")
	}

	// remaining args is the startup script
	for _, arg := range args[1:] {
		if strings.HasPrefix(arg, "--") {
			return "", "", fmt.Errorf("unknown flag %q", arg)
		}
	}

	if len(args) > 1 {
		startupScript = strings.Join(args[1:], " ") // join the args with a space
		startupScript = strings.TrimSpace(startupScript)
	}
	return name, startupScript, nil
}

func templateNewUsage() string {
	return templateNewUsageLine
}

// TemplateNew creates a user-owned startup template.
// Usage: devbox template new <name> [command string]

// this is to create a new template
func TemplateNew(args []string) {
	name, startupScript, err := ParseTemplateNewArgs(args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		fmt.Fprintln(os.Stderr, templateNewUsage())
		os.Exit(1)
	}

	if TestMode {
		fmt.Printf("[test] template new: name=%q", name)
		if startupScript != "" {
			fmt.Printf(" startupScript=%q", startupScript)
		}
		fmt.Println()
		return
	}

	mode, err := service.EnsureLocalModeAndGetCurrMode()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	var created service.Template
	if mode == "local" {
		rt := mustOpenRuntime()
		defer func() { _ = rt.Close() }()
		tmpl, err := rt.CreateTemplate(name, startupScript, service.LocalUserID)
		if err != nil {
			api.FailBox("template new", err)
		}
		created = *tmpl
	} else {
		client, err := api.NewDefault()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}

		body := map[string]string{"name": name}
		if startupScript != "" {
			body["startupScript"] = startupScript
		}

		resp, err := client.Post("/v1/boxes/templates", body)
		if err != nil {
			api.FailBox("template new", err)
		}
		if err := api.CheckStatus(resp); err != nil {
			api.FailBox("template new", err)
		}
		if err := api.DecodeJSON(resp, &created); err != nil {
			api.FailBox("template new", err)
		}
	}

	fmt.Printf("Template created.\n")
	fmt.Printf("  Name: %s\n", created.Name)
	if created.Description != "" {
		fmt.Printf("  Description: %s\n", created.Description)
	}
	if startupScript != "" {
		fmt.Printf("\n  Use: devbox create --template %s <box-name>\n", created.Name)
	} else {
		fmt.Printf("\n  Add a startup command later or use as-is with:\n")
		fmt.Printf("  devbox create --template %s <box-name>\n", created.Name)
	}
}

const templateDeleteUsageLine = "usage: devbox template delete <id>"

// TemplateDelete deletes a user-owned startup template.
// Usage: devbox template delete <id>
func TemplateDelete(args []string) {
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "error: template id is required")
		fmt.Fprintln(os.Stderr, templateDeleteUsageLine)
		os.Exit(1)
	}

	id := strings.TrimSpace(args[0])
	if id == "" {
		fmt.Fprintln(os.Stderr, "error: template id is required")
		fmt.Fprintln(os.Stderr, templateDeleteUsageLine)
		os.Exit(1)
	}
	if strings.HasPrefix(id, "--") {
		fmt.Fprintf(os.Stderr, "error: unknown flag %q\n", id)
		os.Exit(1)
	}
	if len(args) > 1 {
		fmt.Fprintln(os.Stderr, templateDeleteUsageLine)
		os.Exit(1)
	}

	if TestMode {
		fmt.Printf("[test] template delete: id=%q\n", id)
		return
	}

	mode, err := service.EnsureLocalModeAndGetCurrMode()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	if mode != "local" {
		fmt.Fprintf(os.Stderr, "error: template delete is only supported in local mode\n")
		os.Exit(1)
	}

	rt := mustOpenRuntime()
	defer func() { _ = rt.Close() }()
	if err := rt.DeleteTemplate(id, service.LocalUserID); err != nil {
		if strings.Contains(err.Error(), "not found") {
			fmt.Fprintf(os.Stderr, "template %s not found\n", id)
		} else {
			fmt.Fprintf(os.Stderr, "template delete failed: %v\n", err)
		}
		os.Exit(1)
	}

	fmt.Printf("Template %s deleted.\n", id)
}

const templateRenameUsageLine = "usage: devbox template rename <id> <new-name>"

// TemplateRename updates a user-owned template name.
// Usage: devbox template rename <id> <new-name>
func TemplateRename(args []string) {
	if len(args) < 2 {
		fmt.Fprintln(os.Stderr, "error: template id and new name are required")
		fmt.Fprintln(os.Stderr, templateRenameUsageLine)
		os.Exit(1)
	}

	id := strings.TrimSpace(args[0])
	newName := strings.TrimSpace(args[1])
	if id == "" {
		fmt.Fprintln(os.Stderr, "error: template id is required")
		fmt.Fprintln(os.Stderr, templateRenameUsageLine)
		os.Exit(1)
	}
	if newName == "" {
		fmt.Fprintln(os.Stderr, "error: new template name is required")
		fmt.Fprintln(os.Stderr, templateRenameUsageLine)
		os.Exit(1)
	}
	if strings.HasPrefix(id, "--") || strings.HasPrefix(newName, "--") {
		fmt.Fprintf(os.Stderr, "error: unknown flag\n")
		os.Exit(1)
	}
	if len(args) > 2 {
		fmt.Fprintln(os.Stderr, templateRenameUsageLine)
		os.Exit(1)
	}

	if TestMode {
		fmt.Printf("[test] template rename: id=%q newName=%q\n", id, newName)
		return
	}

	mode, err := service.EnsureLocalModeAndGetCurrMode()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	if mode != "local" {
		fmt.Fprintf(os.Stderr, "error: template rename is only supported in local mode\n")
		os.Exit(1)
	}

	rt := mustOpenRuntime()
	defer func() { _ = rt.Close() }()
	renamed, err := rt.RenameTemplate(id, newName, service.LocalUserID)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			fmt.Fprintf(os.Stderr, "template %s not found\n", id)
		} else if strings.Contains(err.Error(), "already exists") {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
		} else {
			fmt.Fprintf(os.Stderr, "template rename failed: %v\n", err)
		}
		os.Exit(1)
	}

	fmt.Printf("Template %s renamed to %s.\n", id, renamed.Name)
}
