package output

import (
	"os"
	"strings"
	"testing"
)

func TestSetOutput_File(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "github-output-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	t.Setenv("GITHUB_OUTPUT", tmpFile.Name())

	if err := SetOutput("key", "value"); err != nil {
		t.Fatalf("SetOutput() error = %v", err)
	}

	data, err := os.ReadFile(tmpFile.Name())
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(string(data), "key=value") {
		t.Errorf("output file contains %q, want key=value", string(data))
	}
}

func TestSetOutput_Multiline(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "github-output-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	t.Setenv("GITHUB_OUTPUT", tmpFile.Name())

	if err := SetOutput("body", "line1\nline2"); err != nil {
		t.Fatalf("SetOutput() error = %v", err)
	}

	data, err := os.ReadFile(tmpFile.Name())
	if err != nil {
		t.Fatal(err)
	}

	content := string(data)
	if !strings.Contains(content, "body<<EOF") {
		t.Errorf("output file contains %q, want heredoc format", content)
	}
}
