package cheatsheet

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/p1ck4rd/famulus/pkg/parser"
)

func TestGenerate(t *testing.T) {
	res := &parser.Result{
		Commands: []parser.Command{
			{Name: "Test Command", Shortcuts: []parser.Shortcut{{"Q"}}},
			{Name: "Other test Command", Shortcuts: []parser.Shortcut{{"Q", "W"}}},
			{Name: "Command with multiple shortcuts", Shortcuts: []parser.Shortcut{{"Q"}, {"W"}}},
			{Name: "Command with nonexistent button", Shortcuts: []parser.Shortcut{{"TEST"}}},
			{Name: "Command without shortcuts"},
			{Name: "Command with special keys", Shortcuts: []parser.Shortcut{{"."}, {"/"}}},
		},
		Tools: []parser.Tool{
			{Name: "Test Tool", Shortcut: parser.Shortcut{"Q"}},
			{Name: "Other test Tool", Shortcut: parser.Shortcut{"Q", "W"}},
			{Name: "Tool with nonexistent button", Shortcut: parser.Shortcut{"TEST"}},
			{Name: "Tool without shortcuts"},
			{Name: "Tool with special keys", Shortcut: parser.Shortcut{"."}},
		},
	}

	refFile := filepath.FromSlash(
		"../../test/testdata/cheatsheet/output.pdf",
	)
	testFile := filepath.FromSlash(
		"../../test/testdata/cheatsheet/test_output.pdf",
	)

	defer os.Remove(testFile)

	var err error
	if err = Generate(res, testFile); err != nil {
		t.Error(err)
	}

	refStat, err := os.Stat(refFile)
	if err != nil {
		t.Error(err)
	}

	testStat, err := os.Stat(testFile)
	if err != nil {
		t.Error(err)
	}

	refSize, testSize := refStat.Size(), testStat.Size()
	if refSize != testSize {
		t.Errorf("File sizes are not equal: %v != %v", refSize, testSize)
	}
}
