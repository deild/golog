# golog

[![GitHub Release](https://img.shields.io/github/release/deild/golog.svg)](https://github.com/deild/golog/releases/latest)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/deild/golog)
[![License](http://img.shields.io/:license-apache-blue.svg)](http://www.apache.org/licenses/LICENSE-2.0.html)

[![Build Status](https://travis-ci.org/deild/golog.svg)](https://travis-ci.org/deild/golog)
[![Coverage](http://gocover.io/_badge/github.com/deild/golog?0)](http://gocover.io/github.com/deild/golog)
[![Go Report Card](https://goreportcard.com/badge/github.com/deild/golog)](https://goreportcard.com/report/github.com/deild/golog)

`golog` is an easy and lightweight CLI tool to time track your tasks. The goal is to enable to track concurrent from small to big tasks.

![demo](http://i.imgur.com/o2F0JbW.gif?1)

## Overview

I work in a very fast paced company, and I'm always receiving requests, plus *a lot* of small requests and I've struggled to find a tool that fit my needs. We do use other tools to track the time spent on a task, but sometimes it gets so overwhelming that it's just not worth to create a bunch of small tasks and track them **but you do want to track them**. If you have your terminal always opened like me, `golog` is perfect for this environments, you can log multiple tasks at the same time without going to your browser/proj management tool to improve productiveness.

## Installation

Make sure you have a working Go environment (go 1.1+ is *required*). [See the install instructions](http://golang.org/doc/install.html).

To get `golog`, run:

```sh
go get github.com/deild/golog
```

To install it in your path so that `golog`can be easily used:

```sh
cd $GOPATH/src/github.com/deild/golog
GOBIN="/usr/local/bin" go install
```

### Enabling autocomplete

Copy `autocomplete/bash_autocomplete` into `/etc/bash_completion.d/golog`.
Don't forget to source the file to make it active in the current shell.

```sh
sudo cp autocomplete/bash_autocomplete /etc/bash_completion.d/golog
source /etc/bash_completion.d/golog
```

Alternatively, you can just source `autocomplete/bash_autocomplete` in your bash configuration with `$PROG` set to golog.

```sh
PROG=golog source "$GOPATH/src/github.com/deild/golog/autocomplete/bash_autocomplete"
```

If using `zsh` use `zsh_autocomplete`

```sh
PROG=golog source "$GOPATH/src/github.com/deild/golog/autocomplete/zsh_autocomplete"
```

## Getting Started

The **start** command will start tracking time for a given taskname. **Note that a taskname cannot have white spaces**, they serve as identifiers.

```sh
> golog start {taskname}
```

To stop tracking use the **stop** command, if you want to **resume** a stopped task just golog start {taskname} again.

```sh
> golog stop {taskname}
```

With the **list** command you can see all your tasks and see which of them are active.

```sh
> golog list
0h:1m:10s    create-readme (running)
0h:0m:44s    do-some-task
```

If you only want to check the status of one task, use the **status** command.

```sh
> golog status create-readme
0h:3m:55s    create-readme (running)
```

The lifetime of the info I usually need is very short (actually is just a workday), in the next day it's unlikely that i'll need previous info. This is one case where **clear** command is handy.

```sh
> golog clear
All tasks were deleted.
```

## Contribution Guidelines

@TODO
If you have any questions feel free to link @deild to the issue in question and we can review it together.
