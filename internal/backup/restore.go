package backup

import (
	"database/sql"
	"encoding/json"
	"os"
	"path/filepath"
)

// RestoreConfigIfNeeded copies config.json from the latest backup when the live
// file is missing or unreadable. Best-effort: errors are ignored.
func RestoreConfigIfNeeded() {
	path, err := configPath()
	if err != nil {
		return
	}
	if isConfigValid(path) {
		return
	}
	restoreFile(path, configFile)
}

// RestoreDBIfNeeded copies devbox.db from the latest backup when the live file
// is missing or unusable. Best-effort: errors are ignored.
func RestoreDBIfNeeded() {
	path, err := dbPath()
	if err != nil {
		return
	}
	if isDBValid(path) {
		return
	}
	restoreFile(path, dbFile)
}

func configPath() (string, error) {
	_, configPath, err := devboxPaths()
	return configPath, err
}

func dbPath() (string, error) {
	_, dbPath, err := devboxPaths()
	return dbPath, err
}

func isConfigValid(path string) bool {
	data, err := os.ReadFile(path)
	if err != nil {
		return false
	}
	var cfg map[string]json.RawMessage
	return json.Unmarshal(data, &cfg) == nil
}

func isDBValid(path string) bool {
	if !fileExists(path) {
		return false
	}
	conn, err := sql.Open("sqlite", path)
	if err != nil {
		return false
	}
	defer func() { _ = conn.Close() }()
	conn.SetMaxOpenConns(1)
	return conn.Ping() == nil
}

func latestBackupDir() (string, bool) {
	dir, err := backupDir()
	if err != nil {
		return "", false
	}
	latest, ok := latestBackupTime(dir)
	if !ok {
		return "", false
	}
	return filepath.Join(dir, latest.Format(timestampLayout)), true
}

func restoreFile(destPath, name string) {
	backupRoot, ok := latestBackupDir()
	if !ok {
		return
	}
	src := filepath.Join(backupRoot, name)
	if !fileExists(src) {
		return
	}
	if err := os.MkdirAll(filepath.Dir(destPath), 0700); err != nil {
		return
	}
	if err := copyFile(src, destPath); err != nil {
		return
	}
	_ = os.Chmod(destPath, 0600)
}
