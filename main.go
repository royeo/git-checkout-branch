package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/manifoldco/promptui"
)

var (
	flags      *flag.FlagSet
	listSize   int
	listAll    bool
	listRemote bool
)

const usage = `The git-checkout-branch command is a wrapper for the git command that
supports for switching branches interactively.

Usage:
  git checkout-branch [flags]

Flags:
  -r    List the remote-tracking branches
  -a    List both remote-tracking branches and local branches
  -n    Set the number of branches displayed in the list, defaults to 10
`

const help = `Error: unknown command "%s" for "git checkout-branch"
Run 'git checkout-branch help' for usage.
`

func initFlags() {
	flags = flag.NewFlagSet("", flag.ExitOnError)
	flags.Usage = func() {
		fmt.Fprint(stdout(), usage)
	}
	flags.IntVar(&listSize, "n", 10, "")
	flags.BoolVar(&listRemote, "r", false, "")
	flags.BoolVar(&listAll, "a", false, "")
}

func main() {
	initFlags()
	flags.Parse(os.Args[1:])

	if len(os.Args) > 1 {
		cmd := os.Args[1]
		if cmd == "help" {
			flags.Usage()
			return
		}
		if !isFlag(cmd) {
			fmt.Fprintf(stderr(), help, cmd)
			return
		}
	}

	var branches []*Branch
	switch {
	case listAll:
		branches = AllBranches()
	case listRemote:
		branches = RemoteBranches()
	default:
		branches = LocalBranches()
	}

	if len(branches) == 0 {
		return
	}

	branch := selectBranch(branches, listSize)
	if branch != nil {
		CheckoutBranch(branch)
	}
}

func selectBranch(branches []*Branch, size int) *Branch {
	iconSelect := promptui.Styler(promptui.FGGreen)("*")

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   iconSelect + " {{ .Name | green }}",
		Inactive: "  {{ .Name }}",
		Selected: promptui.IconGood + " {{ .Name }}",
	}

	searcher := func(input string, index int) bool {
		b := branches[index]
		name := strings.Replace(strings.ToLower(b.Name), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)
		return strings.Contains(name, input)
	}

	label := strconv.Itoa(len(branches)) + " Branches:"

	prompt := promptui.Select{
		Label:     label,
		Items:     branches,
		Templates: templates,
		Size:      size,
		Searcher:  searcher,
	}

	i, _, err := prompt.Run()
	if err != nil {
		return nil
	}
	return branches[i]
}
