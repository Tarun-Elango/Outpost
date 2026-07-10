package service

import (
	"testing"

	localDb "outpost-cli/service/localDb"
)

func TestUniqueImportName(t *testing.T) {
	t.Setenv("HOME", t.TempDir())
	db, err := localDb.Open()
	if err != nil {
		t.Fatalf("open db: %v", err)
	}
	t.Cleanup(func() { _ = db.Close() })

	if err := db.InsertInstance("box-1", "i-11111111111111111", "alpha", LocalUserID, "running", "t3.micro", "us-east-1", "aws"); err != nil {
		t.Fatalf("insert: %v", err)
	}

	got, err := uniqueImportName(db, LocalUserID, "alpha", "i-22222222222222222", false)
	if err != nil || got != "alpha-2" {
		t.Fatalf("collision: got %q err %v, want alpha-2", got, err)
	}

	got, err = uniqueImportName(db, LocalUserID, "i-33333333333333333", "i-33333333333333333", false)
	if err != nil || got != "imported-33333333" {
		t.Fatalf("id-shaped: got %q err %v, want imported-33333333", got, err)
	}
}
