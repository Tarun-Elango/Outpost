package localDb

import (
	"strings"
	"testing"
)

func TestValidateTemplateNameAvailableForRenameAllowsCurrentName(t *testing.T) {
	db := newTestDB(t)

	if err := db.InsertTemplate("tmpl-1", LocalUserID, "alpha", ""); err != nil {
		t.Fatalf("insert template: %v", err)
	}

	if err := db.ValidateTemplateNameAvailableForRename("alpha", LocalUserID, "alpha"); err != nil {
		t.Fatalf("validate current name for rename: %v", err)
	}
}

func TestValidateTemplateNameAvailableForRenameRejectsAnotherTemplateName(t *testing.T) {
	db := newTestDB(t)

	if err := db.InsertTemplate("tmpl-1", LocalUserID, "alpha", ""); err != nil {
		t.Fatalf("insert first template: %v", err)
	}
	if err := db.InsertTemplate("tmpl-2", LocalUserID, "beta", ""); err != nil {
		t.Fatalf("insert second template: %v", err)
	}

	err := db.ValidateTemplateNameAvailableForRename("beta", LocalUserID, "alpha")
	if err == nil {
		t.Fatal("expected duplicate name error")
	}
	if !strings.Contains(err.Error(), "template name already exists: beta") {
		t.Fatalf("unexpected duplicate name error: %v", err)
	}
}

func TestUpdateTemplateNamePersistsTrimmedName(t *testing.T) {
	db := newTestDB(t)

	if err := db.InsertTemplate("tmpl-1", LocalUserID, "alpha", "echo hello"); err != nil {
		t.Fatalf("insert template: %v", err)
	}

	if err := db.UpdateTemplateName("alpha", LocalUserID, " beta "); err != nil {
		t.Fatalf("update template name: %v", err)
	}

	record, err := db.GetTemplateByNameAndUserID("beta", LocalUserID)
	if err != nil {
		t.Fatalf("get renamed template: %v", err)
	}
	if record.Name != "beta" {
		t.Fatalf("expected name beta, got %q", record.Name)
	}
	if record.StartupScript.String != "echo hello" {
		t.Fatalf("expected startup script preserved, got %q", record.StartupScript.String)
	}
}
