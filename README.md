# git checkout-branch

[![CircleCI](https://circleci.com/gh/royeo/git-checkout-branch.svg?style=shield)](https://circleci.com/gh/royeo/git-checkout-branch)
[![Go Report Card](https://goreportcard.com/badge/royeo/git-checkout-branch)](https://goreportcard.com/report/royeo/git-checkout-branch)

Switch git branches interactively.

The `git checkout-branch` command is a custom git command to improve the efficiency of switching branches.

## Overview

![](https://raw.githubusercontent.com/royeo/static/master/gif/git-checkout-branch.gif)

Instructions:
- Use the arrow keys to navigate: `↓` `↑` `→` `←`
- You can also move up and down using `j` and `k`
- Use `/` to toggle search

## Installation

Use `go get` to install the `git checkout-branch` command:

```sh
go get github.com/royeo/git-checkout-branch
```

> Make sure the `$GOPATH/bin` folder is in your `PATH`.

If you are using the GO1.11 module, use the following command to install:

```sh
GO111MODULE=off go get github.com/royeo/git-checkout-branch
```

It is recommended to set up an alias for `checkout-branch`, such as `cb`.

```sh
git config --global alias.cb checkout-branch
```

## Usage

Use `git checkout-branch help` for help information.

```
Switch git branch interactively.

Usage:
  git checkout-branch [flags]

Flags:
  -r    List the remote-tracking branches
  -a    List both remote-tracking branches and local branches
  -n    Set the number of branches displayed in the list, defaults to 10
```

## License

MIT Copyright (c) 2019 Royeo
