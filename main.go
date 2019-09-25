package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var (
	listSize   int
	listAll    bool
	listRemote bool
	isDelete   bool
	hideHelp   bool
)

var rootCmd = &cobra.Command{
	Use:   "git checkout-branch",
	Short: "Checkout git branches more efficiently.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			cmd.Help()
			os.Exit(1)
		}

		var branches []*Branch
		switch {
		case listAll:
			branches = allBranches()
		case listRemote:
			branches = remoteBranches()
		default:
			branches = localBranches()
		}

		if len(branches) == 0 {
			return
		}

		branch := selectBranch(branches, listSize, hideHelp)
		if branch == nil {
			return
		}
		if isDelete {
			deleteBranch(branch)
			return
		}
		checkoutBranch(branch)
	},
}

func init() {
	rootCmd.Flags().BoolVarP(&listRemote, "remotes", "r", false, "")
	rootCmd.Flags().BoolVarP(&listAll, "all", "a", false, "")
	rootCmd.Flags().IntVarP(&listSize, "number", "n", 10, "")
	rootCmd.Flags().BoolVarP(&isDelete, "delete", "d", false, "")
	rootCmd.Flags().BoolVarP(&hideHelp, "hide-help", "", false, "")

	rootCmd.SetUsageFunc(func(*cobra.Command) error {
		usage := `Usage:
  git checkout-branch [flags]

Flags:
  -a, --all          List both remote-tracking branches and local branches
  -r, --remotes      List the remote-tracking branches
  -n, --number       Set the number of branches displayed in the list (default 10)
  -d, --delete       Delete a branch
      --hide-help    Hide the help information`
		fmt.Println(usage)
		os.Exit(1)
		return nil
	})
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func selectBranch(branches []*Branch, size int, hideHelp bool) *Branch {
	iconSelect := promptui.Styler(promptui.FGGreen)("*")

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   iconSelect + " {{ .Name | green }}",
		Inactive: "  {{ .Name }}",
	}

	searcher := func(input string, index int) bool {
		b := branches[index]
		name := strings.Replace(strings.ToLower(b.Name), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)
		return strings.Contains(name, input)
	}

	label := strconv.Itoa(len(branches)) + " Branches:"

	prompt := promptui.Select{
		Label:        label,
		Items:        branches,
		Templates:    templates,
		Size:         size,
		Searcher:     searcher,
		HideHelp:     hideHelp,
		HideSelected: true,
	}

	i, _, err := prompt.Run()
	if err != nil {
		return nil
	}
	return branches[i]
}
