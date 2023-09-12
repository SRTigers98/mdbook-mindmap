package main

import (
	_ "embed"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"text/template"
)

var (
	//go:embed changelog.tmpl.md
	changelogTmpl []byte
	tmpl          *template.Template
)

func init() {
	tmpl = template.Must(template.New("changelog").Parse(string(changelogTmpl)))
}

func main() {
	if len(os.Args) < 2 {
		panic("tag name is missing")
	}
	tag := os.Args[1]

	lastTag := getLastTag()
	history := getCommitHistory(lastTag)
	feats, fixes := groupCommits(history)

	changelog, err := os.Create(fmt.Sprintf("%s.md", tag))
	if err != nil {
		panic(err)
	}

	tmpl.Execute(changelog, map[string]any{
		"tag":   tag,
		"pre":   strings.Contains(tag, "-"),
		"feats": feats,
		"fixes": fixes,
	})
}

func getLastTag() string {
	lastTagCmd := exec.Command("git", "describe", "--tags", "--abbrev=0")

	if lastTag, err := lastTagCmd.Output(); err != nil {
		return ""
	} else {
		return string(lastTag)
	}
}

func getCommitHistory(lastTag string) []string {
	var gitLogCmd *exec.Cmd
	if len(lastTag) == 0 {
		gitLogCmd = exec.Command("git", "--no-pager", "log", "--format=%s")
	} else {
		historyRange := fmt.Sprintf("%s..HEAD", lastTag)
		gitLogCmd = exec.Command("git", "--no-pager", "log", historyRange, "--format=%s")
	}

	if log, err := gitLogCmd.Output(); err != nil {
		panic(err)
	} else {
		return strings.Split(string(log), "\n")
	}
}

func groupCommits(history []string) ([]string, []string) {
	var feats []string
	var fixes []string

	for _, commit := range history {
		if strings.HasPrefix(commit, "feat:") {
			feats = append(feats, commit)
		}

		if strings.HasPrefix(commit, "fix:") {
			fixes = append(fixes, commit)
		}
	}

	return feats, fixes
}
