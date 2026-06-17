package service

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func sshConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".ssh", "config"), nil // should be ~/.ssh/config
}

func readSSHConfig() (string, error) {
	path, err := sshConfigPath() // get the path to the ssh config file
	if err != nil {
		return "", err
	}
	data, err := os.ReadFile(path) // read the ssh config file
	if err != nil {
		if os.IsNotExist(err) {
			return "", nil
		}
		return "", err
	}
	return string(data), nil
}

func writeSSHConfig(content string) error {
	path, err := sshConfigPath()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(path), 0700); err != nil {
		return err
	}
	return os.WriteFile(path, []byte(content), 0600) // write the ssh config file
}

type hostBlock struct {
	start int
	end   int
	hosts []string
}

func parseHostBlocks(content string) []hostBlock {
	lines := strings.Split(content, "\n")
	var blocks []hostBlock
	var current *hostBlock

	for i, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(strings.ToLower(trimmed), "host ") {
			if current != nil {
				current.end = i
				blocks = append(blocks, *current)
			}
			parts := strings.Fields(trimmed)
			hosts := []string{}
			if len(parts) > 1 {
				hosts = parts[1:]
			}
			current = &hostBlock{start: i, end: len(lines), hosts: hosts}
		}
	}
	if current != nil {
		blocks = append(blocks, *current)
	}
	return blocks
}

func hostInBlock(block hostBlock, name string) bool {
	for _, h := range block.hosts {
		if h == name {
			return true
		}
	}
	return false
}

func devboxHostName(name string) string {
	return "devbox-" + name
}

func findBlockByHost(content, name string) (hostBlock, bool) {
	// for each block in the ssh config file, check if the host is in the block
	for _, block := range parseHostBlocks(content) {
		if hostInBlock(block, name) {
			return block, true
		}
	}
	return hostBlock{}, false
}

func formatHostBlock(name, ipAddress string) string {
	return fmt.Sprintf(
		"Host %s\n    HostName %s\n    User ec2-user\n    IdentityFile ~/.ssh/id_ed25519\n    StrictHostKeyChecking accept-new\n",
		name, ipAddress,
	)
}

// add new host in .ssh/config ( name, ip address)
func AddHost(name, ipAddress string) error {
	host := devboxHostName(name)
	content, err := readSSHConfig()
	if err != nil {
		return err
	}
	if _, found := findBlockByHost(content, host); found { // check if the host already exists
		return fmt.Errorf("host %q already exists", host)
	}
	if content != "" && !strings.HasSuffix(content, "\n") { // if empty or not a new line, add a new line
		content += "\n"
	}
	content += formatHostBlock(host, ipAddress)
	return writeSSHConfig(content)
}

// update ip address in .ssh/config (name, ip address), for the given host with name, update the ip address
func UpdateHost(name, ipAddress string) error {
	host := devboxHostName(name)
	content, err := readSSHConfig() // get content of ssh config file
	if err != nil {
		return err
	}
	block, found := findBlockByHost(content, host)
	if !found {
		return fmt.Errorf("host %q does not exist", host)
	}

	lines := strings.Split(content, "\n")      // the entire ssh config file as a list of lines
	blockLines := lines[block.start:block.end] // get the lines in the block

	hostnameUpdated := false
	for i, line := range blockLines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(strings.ToLower(trimmed), "hostname ") {
			indent := line[:len(line)-len(strings.TrimLeft(line, " \t"))] // get the indent of the line
			blockLines[i] = indent + "HostName " + ipAddress              // update the block line i with the new ip address
			hostnameUpdated = true
			break
		}
	}
	if !hostnameUpdated {
		insert := []string{"    HostName " + ipAddress}
		blockLines = append(blockLines[:1], append(insert, blockLines[1:]...)...)
	}

	newLines := append(lines[:block.start], append(blockLines, lines[block.end:]...)...) //[ everything before the host block ] + [ the updated block ] + [ everything after the host block ]
	return writeSSHConfig(strings.Join(newLines, "\n"))
}

// syncSSHHostIP updates the HostName for an existing entry, or adds one if missing.
func syncSSHHostIP(name, ipAddress string) error {
	if ipAddress == "" {
		return nil
	}
	if err := UpdateHost(name, ipAddress); err != nil {
		if !strings.Contains(err.Error(), "does not exist") {
			return err
		}
		return AddHost(name, ipAddress)
	}
	return nil
}

// delete host (name)
func DeleteHost(name string) error {
	host := devboxHostName(name)
	content, err := readSSHConfig()
	if err != nil {
		return err
	}
	block, found := findBlockByHost(content, host)
	if !found {
		return nil // if the host does not exist, do nothing
	}

	lines := strings.Split(content, "\n")                         // the entire ssh config file as a list of lines
	newLines := append(lines[:block.start], lines[block.end:]...) // [ everything before the host block ] + [ everything after the host block ]
	return writeSSHConfig(strings.Join(newLines, "\n"))
}
