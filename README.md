[![Coverage Status](https://coveralls.io/repos/github/xplorfin/go-template/badge.svg?branch=master)](https://coveralls.io/github/xplorfin/go-template?branch=master)
[![Renovate enabled](https://img.shields.io/badge/renovate-enabled-brightgreen.svg)](https://app.renovatebot.com/dashboard#github/xplorfin/go-template)
[![Build status](https://github.com/xplorfin/go-template/workflows/test/badge.svg)](https://github.com/xplorfin/go-template/actions?query=workflow%3Atest)
[![Build status](https://github.com/xplorfin/go-template/workflows/goreleaser/badge.svg)](https://github.com/xplorfin/go-template/actions?query=workflow%3Agoreleaser)
[![](https://godoc.org/github.com/xplorfin/go-template?status.svg)](https://godoc.org/github.com/xplorfin/go-template)
[![Go Report Card](https://goreportcard.com/badge/github.com/xplorfin/go-template)](https://goreportcard.com/report/github.com/xplorfin/go-template)

# What is this?

This repository serves as a template for our open source projects. It uses [urfave/cli](https://github.com/urfave/cli) as for creating a cli. The release process follows semver, with an optional test on every build and automerges passing dependency upgrades. You can optionally use [ci skip] to skip a release

# For [Entropy](http://entropy.rocks/) Organization Members

While we're working to open source as much of our code as possible, this will take time so if you need to create a private repository [please do so here](https://github.com/xplorfin/go-template-private#what-is-this).  


# Using this repository:

1. Create a new repository from this template on github
1. Find and replace: `go-template` with `your-repo-name` across the entire project
1. [Mark the appropriate package as public](https://docs.github.com/en/packages/guides/configuring-access-control-and-visibility-for-container-images#configuring-visibility-of-container-images-for-an-organization)
1. *Optional*: If you won't be writing test (you should!) you can delete `.github/workflows/test.yml`. Tests are enabled by default, so you'll need to switch the `if` line with `if: false` or delete the [test](.github/workflows/test.yml) file 

# Library only releases

If your repo does not have a binary, you can remove the 

# Directory structure:

1. [`.github`](.github): contains configs for [making sure users don't create issues](.github/ISSUE_TEMPLATE), [workflows](.github/workflows/go.yml) for releasing and a disabled worfklow for [testing](.github/workflows/test.yml) as well as our [renovatebot config](.github/renovate.json)
1. [`cmd`](cmd): contains start command and [cli interface](https://github.com/urfave/cli)
1. [`docker`](docker): contains docker files for deploying
1. [`internal`](internal): for packages you don't want exported

_Note_: we follow the standard [go project layout](https://github.com/golang-standards/project-layout). Modules required by other packages should go in the `pkg` library
