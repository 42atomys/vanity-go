# Simple Go Proxy

[![Release 🎉](https://github.com/42Atomys/go-proxy/actions/workflows/release.yaml/badge.svg)](https://github.com/42Atomys/go-proxy/actions/workflows/release.yaml)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/42atomys/go-proxy?label=last%20release)
![GitHub contributors](https://img.shields.io/github/contributors/42Atomys/go-proxy?color=blueviolet)
![GitHub Repo stars](https://img.shields.io/github/stars/42atomys/go-proxy?color=blueviolet)
[![Docker Pull](https://img.shields.io/docker/pulls/atomys/go-proxy)](https://hub.docker.com/r/atomys/go-proxy)
[![Docker Pull](https://img.shields.io/docker/image-size/atomys/go-proxy)](https://hub.docker.com/r/atomys/go-proxy)

Simple go application that allows you to share your code with a custom domain name instead of github or gitlab links or other git protocols.

Say goodbye to `github.com/42Atomys/go-proxy` imports. Say hello to `atomys.codes/go-proxy` imports 🎉

## Motivation

At the beginning to clarify my code especially with gitlab and subfolders (ex: `gitlab.com/subgroup-a/subgroup-b/subgroup-n/project`) importing files was not sexy. And a problem occurred when the folder architecture changed!

Error 404 in all directions, in all repositories. I created a classic `index.html` file but having to connect to a server to do vim (sorry emacs) is very annoying.

With [42Stellar](https://github.com/42Stellar) project, I told myself that I didn't want to redo everything, so I put this repo online

## Usage

TODO

## To-Do

- [ ] app core
  - [ ] code
  - [ ] tests
  - [ ] release
- [ ] app configuration
  - [ ] code
  - [ ] tests
  - [ ] release
- [ ] allow multiples domains
  - [ ] code
  - [ ] tests
  - [ ] release
- [ ] allow multiples redirect per domain
  - [ ] code
  - [ ] tests
  - [ ] release
- [ ] create a ctl to simplify configuration manipulation
  - [ ] code
  - [ ] tests
  - [ ] release

# Contribution

All pull requests and issues on GitHub will welcome.

All contributions are welcome :)

## Thanks
