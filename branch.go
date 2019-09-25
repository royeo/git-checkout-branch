package main

import (
	"strings"
)

// Branch describes a Git branch.
type Branch struct {
	Name string
}

// currentBranch returns the current branch.
func currentBranch() *Branch {
	name := strings.TrimSpace(cmdOutput("git", "rev-parse", "--abbrev-ref", "HEAD"))
	return &Branch{Name: name}
}

// localBranches returns the local branches.
func localBranches() []*Branch {
	return splitBranch(cmdOutput("git", "branch"))
}

// allBranches returns both remote-tracking branches and local branches.
func allBranches() []*Branch {
	return splitBranch(cmdOutput("git", "branch", "-a"))
}

// remoteBranches returns the remote-tracking branches.
func remoteBranches() []*Branch {
	return splitBranch(cmdOutput("git", "branch", "-r"))
}

func splitBranch(output string) []*Branch {
	o := strings.Replace(output, "*", "", -1)
	names := strings.Split(o, "\n")
	var branches []*Branch
	for _, name := range names {
		if len(name) == 0 {
			continue
		}
		name = strings.TrimSpace(name)
		branches = append(branches, &Branch{Name: name})
	}
	return branches
}

// checkoutBranch switch to the selected branch.
func checkoutBranch(b *Branch) {
	cmdRun("git", "checkout", extractBranch(b.Name))
}

func deleteBranch(b *Branch) {
	cmdRun("git", "branch", "-d", extractBranch(b.Name))
}

func extractBranch(name string) string {
	if strings.Contains(name, "->") {
		s := strings.Split(name, "->")
		return strings.TrimSpace(s[0])
	}
	return name
}
