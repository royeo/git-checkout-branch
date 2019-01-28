package main

import (
	"strings"
)

// Branch describes a Git branch.
type Branch struct {
	Name string
}

// CurrentBranch returns the current branch.
func CurrentBranch() *Branch {
	name := strings.TrimSpace(cmdOutput("git", "rev-parse", "--abbrev-ref", "HEAD"))
	return &Branch{Name: name}
}

// LocalBranches returns the local branches.
func LocalBranches() []*Branch {
	return splitBranch(cmdOutput("git", "branch"))
}

// AllBranches returns both remote-tracking branches and local branches.
func AllBranches() []*Branch {
	return splitBranch(cmdOutput("git", "branch", "-a"))
}

// RemoteBranches returns the remote-tracking branches.
func RemoteBranches() []*Branch {
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

// CheckoutBranch switch to the selected branch.
func CheckoutBranch(b *Branch) {
	cmdRun("git", "checkout", extractBranch(b.Name))
}

func extractBranch(name string) string {
	if strings.Contains(name, "->") {
		s := strings.Split(name, "->")
		return strings.TrimSpace(s[0])
	}
	return name
}
