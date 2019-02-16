# git-checkout-branch

[![CircleCI](https://circleci.com/gh/royeo/git-checkout-branch.svg?style=shield)](https://circleci.com/gh/royeo/git-checkout-branch)
[![Go Report Card](https://goreportcard.com/badge/royeo/git-checkout-branch)](https://goreportcard.com/report/royeo/git-checkout-branch)

The `git-checkout-branch` tool is a command-line tool for improving the efficiency of switching git branches.

## Features

- Switch git branch interactively
- Search git branch dynamically

## Overview

![](https://raw.githubusercontent.com/royeo/static/master/gif/git-checkout-branch.gif)

Instructions:
- Use the arrow keys to navigate: `↓` `↑` `→` `←`
- You can also move up and down using `j` and `k`
- Use `/` to toggle search
- Press `ctrl + c` to quit

## Installation

You can install the `git-checkout-branch` binary from GitHub Releases.

```sh
curl -sSL https://github.com/royeo/git-checkout-branch/releases/download/v0.2.0/git-checkout-branch-`uname -s`-`uname -m` -o /usr/local/bin/git-checkout-branch && chmod +x /usr/local/bin/git-checkout-branch
```

You can also use `go get` to install the `git-checkout-branch` binary:

```sh
go get -u github.com/royeo/git-checkout-branch
```

> make sure the `$GOPATH/bin` folder is in your `PATH`.

It is recommended to set up an alias for `checkout-branch`, such as `cb`.

```sh
git config --global alias.cb checkout-branch
```

## Usage

Use `git checkout-branch help` for help information.

```
Checkout git branches more efficiently.

Usage:
  git checkout-branch [flags]

Flags:
  -a, --all          List both remote-tracking branches and local branches
  -r, --remotes      List the remote-tracking branches
  -n, --number       Set the number of branches displayed in the list (default 10)
      --hide-help    Hide the help information
```

## License

MIT Copyright (c) 2019 Royeo
