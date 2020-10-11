package parser

import (
	"fmt"
	"strings"

	"github.com/epsxy/gommitizen/pkg/git"
)

//GenerateMarkdownChangelog generates a markdown-formatted string containing
//the changelog for the list of commits provided.
func GenerateMarkdownChangelog(commits []git.Commit) string {
	breakingChanges := []string{}
	features := []string{}
	fixes := []string{}
	chores := []string{}
	others := []string{}

	for _, commit := range commits {
		if commit.IsMergeCommit {
			continue
		}
		if commit.IsBreakingChange {
			breakingChanges = append(breakingChanges, parseMarkdownCommit(commit))
			continue
		}
		if commit.Type == "feat" {
			features = append(features, parseMarkdownCommit(commit))
			continue
		}
		if commit.Type == "fix" {
			fixes = append(fixes, parseMarkdownCommit(commit))
			continue
		}
		if commit.Type == "chore" {
			chores = append(chores, parseMarkdownCommit(commit))
			continue
		}
		others = append(others, parseMarkdownCommit(commit))
	}

	title := "Changelog\n\n"
	var markdown *string = &title
	parseSection(markdown, "Breaking Changes", breakingChanges)
	parseSection(markdown, "Features", features)
	parseSection(markdown, "Fixes", fixes)
	parseSection(markdown, "Chores", chores)
	parseSection(markdown, "Others", others)
	return *markdown
}

func parseMarkdownCommit(c git.Commit) string {
	fmt.Println(c)
	return fmt.Sprintf("- %s : %s", c.Hash, c.Msg)
}

func parseSection(markdown *string, title string, commits []string) {
	if len(commits) == 0 {
		return
	}
	*markdown = *markdown + fmt.Sprintf("## %s\n\n", title) + strings.Join(commits, "\n") + "\n\n"
}
