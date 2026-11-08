package git

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGetGitDir_WitGitDirectory(t *testing.T) {
	dir := t.TempDir()

	if err := os.Mkdir(filepath.Join(dir, ".git"), 0755); err != nil {
		t.Fatalf("Test directory could not be created: %v", err)
	}

	got, err := GetGitDir(dir)

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if !got {
		t.Error("GetGitDir() = false, expected true")
	}

}

func TestGetGitDir_WithoutGitDirectory(t *testing.T) {

	dir := t.TempDir()

	got, err := GetGitDir(dir)

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if got {
		t.Error("GetGitDir() = true, expected false")
	}

}

func TestGitGetDir_WithoutExistingDirectory(t *testing.T) {
	got, err := GetGitDir("")

	if err == nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if got {
		t.Error("GetGitDir() = true, expected")
	}
}
