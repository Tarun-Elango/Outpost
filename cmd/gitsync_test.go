package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestKeyInSSHAgentParsesLoadedIdentity(t *testing.T) {
	keyPath := "/tmp/id_ed25519"
	fingerprint := "SHA256:abc123"

	orig := execCommand
	t.Cleanup(func() { execCommand = orig })
	execCommand = func(name string, args ...string) *exec.Cmd {
		if len(args) == 2 && args[0] == "-lf" && strings.HasSuffix(name, "ssh-keygen") {
			return fakeCommand(t, 0, fmt.Sprintf("256 %s user@example.com (ED25519)\n", fingerprint))
		}
		if len(args) == 1 && args[0] == "-l" && strings.HasSuffix(name, "ssh-add") {
			return fakeCommand(t, 0, fmt.Sprintf("256 %s user@example.com (ED25519)\n", fingerprint))
		}
		return orig(name, args...)
	}

	loaded, err := keyInSSHAgent(keyPath)
	if err != nil {
		t.Fatalf("keyInSSHAgent: %v", err)
	}
	if !loaded {
		t.Fatal("expected key to be reported as loaded")
	}
}

func TestKeyInSSHAgentEmptyAgent(t *testing.T) {
	orig := execCommand
	t.Cleanup(func() { execCommand = orig })
	execCommand = func(name string, args ...string) *exec.Cmd {
		if len(args) == 2 && args[0] == "-lf" && strings.HasSuffix(name, "ssh-keygen") {
			return fakeCommand(t, 0, "256 SHA256:abc123 user@example.com (ED25519)\n")
		}
		if len(args) == 1 && args[0] == "-l" && strings.HasSuffix(name, "ssh-add") {
			return fakeCommand(t, 1, "The agent has no identities.\n")
		}
		return orig(name, args...)
	}

	loaded, err := keyInSSHAgent("/tmp/id_ed25519")
	if err != nil {
		t.Fatalf("keyInSSHAgent: %v", err)
	}
	if loaded {
		t.Fatal("expected key not to be loaded")
	}
}

func fakeCommand(t *testing.T, code int, output string) *exec.Cmd {
	t.Helper()
	script := filepath.Join(t.TempDir(), "fake-cmd.sh")
	content := "#!/bin/sh\n"
	content += "cat <<'EOF'\n" + output
	if !strings.HasSuffix(output, "\n") {
		content += "\n"
	}
	content += "EOF\n"
	content += fmt.Sprintf("exit %d\n", code)
	if err := os.WriteFile(script, []byte(content), 0700); err != nil {
		t.Fatalf("write fake ssh-add: %v", err)
	}
	return exec.Command(script)
}
