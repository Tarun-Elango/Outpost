package cmd

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"

	"devbox-cli/internal/api"
)

const (
	devboxReadyPath    = "/var/lib/devbox/ready"
	devboxReadyMessage = "the user data script is completed"
)

// defaultKeyPath returns the path to the user's default SSH private key,
// trying id_ed25519 then id_rsa under ~/.ssh. Returns "" if none found.
func defaultKeyPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	for _, name := range []string{"id_ed25519", "id_rsa"} {
		p := filepath.Join(home, ".ssh", name)
		if _, err := os.Stat(p); err == nil {
			return p
		}
	}
	return ""
}

// SshStatusResponse is returned by GET /v2/boxes/{id}/ssh-status.
type SshStatusResponse struct {
	Ready    bool `json:"ready"`
	Instance *Box `json:"instance"`
}

func sshBaseArgs(identity, portArg string) []string {
	argv := []string{
		"-p", portArg,
		"-o", "ConnectTimeout=15",
		"-o", "StrictHostKeyChecking=accept-new", // TODO: StrictHostKeyChecking=yes plus managing known_hosts
	}
	if identity != "" {
		argv = append([]string{"-i", identity}, argv...)
	}
	return argv
}

// checkDevboxReady runs one SSH probe for the user-data ready marker.
func checkDevboxReady(sshBin, identity, user, host, portArg string) bool {
	target := fmt.Sprintf("%s@%s", user, host)
	argv := append([]string{sshBin}, sshBaseArgs(identity, portArg)...)
	argv = append(argv,
		"-o", "BatchMode=yes",
		target,
		"cat", devboxReadyPath,
	)
	out, err := exec.Command(argv[0], argv[1:]...).CombinedOutput()
	return err == nil && strings.TrimSpace(string(out)) == devboxReadyMessage
}

// SSH checks EC2 health and the devbox ready marker, then execs ssh.
func SSH(args []string) {
	if TestMode {
		fmt.Println("[test] ssh: done")
		return
	}
	fs := flag.NewFlagSet("ssh", flag.ExitOnError)
	user := fs.String("u", "ec2-user", "SSH username")
	port := fs.Int("p", 22, "SSH port")
	identity := fs.String("i", defaultKeyPath(), "path to SSH private key")
	fs.Usage = func() {
		fmt.Fprintln(os.Stderr, "usage: devbox ssh -v -i [-u user] [-p port] [-i identity] <id> [-- ssh-args...]")
		fs.PrintDefaults()
	}

	// Split args on "--" to allow passing raw flags to ssh.
	var extra []string
	for i, a := range args {
		if a == "--" {
			extra = args[i+1:]
			args = args[:i]
			break
		}
	}

	if err := fs.Parse(args); err != nil {
		os.Exit(1)
	}
	if fs.NArg() < 1 {
		fs.Usage()
		os.Exit(1)
	}
	id := fs.Arg(0)

	client, err := api.NewDefault()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	resp, err := client.Get("/v2/boxes/" + id + "/ssh-status") // check the two conditions from server
	if err != nil {
		fmt.Fprintf(os.Stderr, "ssh: %v\n", err)
		os.Exit(1)
	}
	if err := api.CheckStatus(resp); err != nil {
		fmt.Fprintf(os.Stderr, "ssh: %v\n", err)
		os.Exit(1)
	}

	var status SshStatusResponse // check response type
	if err := api.DecodeJSON(resp, &status); err != nil {
		fmt.Fprintf(os.Stderr, "ssh: %v\n", err)
		os.Exit(1)
	}

	if !status.Ready {
		fmt.Fprintln(os.Stderr, "ssh: box is not ready yet (EC2 status checks still pending)")
		os.Exit(1)
	}
	if status.Instance == nil {
		fmt.Fprintln(os.Stderr, "ssh: server reported ready but returned no instance details, try the command again in a few minutes.")
		os.Exit(1)
	}

	b := *status.Instance
	if b.PublicIP == "" {
		fmt.Fprintln(os.Stderr, "ssh: box has no IP address (is it running?)")
		os.Exit(1)
	}
	if b.Status != "running" {
		fmt.Fprintf(os.Stderr, "ssh: box is %s, not running\n", b.Status)
		os.Exit(1)
	}

	sshBin, err := exec.LookPath("ssh")
	if err != nil {
		fmt.Fprintln(os.Stderr, "ssh: ssh binary not found in PATH")
		os.Exit(1)
	}

	target := fmt.Sprintf("%s@%s", *user, b.PublicIP)
	portArg := fmt.Sprintf("%d", *port)

	if !checkDevboxReady(sshBin, *identity, *user, b.PublicIP, portArg) {
		fmt.Fprintln(os.Stderr, "ssh: devbox is not ready yet — try again in a minute")
		os.Exit(1)
	}

	fmt.Fprintf(os.Stderr, "Connecting to %s (box %s)...\n", target, id)

	argv := append([]string{sshBin}, sshBaseArgs(*identity, portArg)...) // create ssh command
	argv = append(argv, target)
	argv = append(argv, extra...)

	if err := syscall.Exec(sshBin, argv, os.Environ()); err != nil {
		fmt.Fprintf(os.Stderr, "ssh: exec failed: %v\n", err)
		os.Exit(1)
	}
}
