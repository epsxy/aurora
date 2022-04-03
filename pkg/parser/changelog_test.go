package parser

import (
	"testing"

	"github.com/epsxy/aurora/pkg/git"
)

// TestParseMarkdownCommit verifies we correctly parse a Commit to markdown syntax
func TestParseMarkdownCommit(t *testing.T) {
	commit := git.Commit{
		Hash:  "eed3af1",
		Scope: "client",
		Type:  "feat",
		Msg:   "add cookies",
	}
	wanted := "- eed3af1 : add cookies"

	res := parseMarkdownCommit(commit)
	if res != wanted {
		t.Fatalf("test failed, wanted: %s, got: %s", res, wanted)
	}
}
