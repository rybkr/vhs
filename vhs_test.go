package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestDuplicateFrame(t *testing.T) {
	dir := t.TempDir()

	cursor := []byte("cursor-image")
	text := []byte("text-image")
	writeFrame := func(format string, frame int, data []byte) {
		t.Helper()
		if err := os.WriteFile(filepath.Join(dir, fmt.Sprintf(format, frame)), data, 0o600); err != nil {
			t.Fatal(err)
		}
	}
	writeFrame(cursorFrameFormat, 1, cursor)
	writeFrame(textFrameFormat, 1, text)

	if err := duplicateFrame(dir, 1, 2); err != nil {
		t.Fatalf("duplicateFrame: %v", err)
	}

	for _, tc := range []struct {
		format string
		want   []byte
	}{
		{cursorFrameFormat, cursor},
		{textFrameFormat, text},
	} {
		got, err := os.ReadFile(filepath.Join(dir, fmt.Sprintf(tc.format, 2)))
		if err != nil {
			t.Fatalf("duplicated frame missing: %v", err)
		}
		if string(got) != string(tc.want) {
			t.Errorf("duplicated %s = %q, want %q", tc.format, got, tc.want)
		}
	}

	if err := duplicateFrame(dir, 3, 4); err == nil {
		t.Error("expected error duplicating a missing frame")
	}
}
