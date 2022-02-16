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

### Step 1 : Configuration file
```yaml
# API Version also used to protect against API or Schema changes
# Actually, the available API versions are: 1
apiVersion: 1
# List of your proxies
# You can add as many proxies as you want with logic:
# 1 proxy per final domain
proxies:
- # namespace is the domain name used for following entries
  # This can be a subdomain, a domain or a full domain name
  # subdomain.example.org, example.org or example.org/subdomain
  namespace: atomys.codes
  # All entries of this namespace will be proxied to the following address
  # Key are the name and the entrypoint/path of your proxied packages
  # Value is the current URL of your package. The Destination URL must
  # end with a valid protocol.
  # Allowed protocol are: "bzr", "fossil", "git", "hg", "svn".
  entries:
    # Redirect go-get import to atomys.codes/go-proxy
    go-proxy: https://github.com/42Atomys/go-proxy.git
    # Redirect go-get import to atomys.codes/dns-updater
    dns-updater: https://github.com/42Atomys/dns-updater.git
    # Redirect go-get import to atomys.codes/subpath/gw2api-go
    subpath/gw2api-go: https://gitlab.com/Atomys/gw2api-go.git
```

### Step 2: Launch it 🚀

> **TIPS**: When you create your routing, configure it to only take into account the query params `go-get=1` to follow the GoLang directive (https://pkg.go.dev/cmd/go#hdr-Remote_import_paths)

### With Kubernetes

If you want to use kubernetes, for production or personnal use, refere to example/kubernetes:

https://github.com/42Atomys/go-proxy/tree/main/examples/kubernetes


### With Docker image

You can use the docker image [atomys/go-proxy](https://hub.docker.com/r/atomys/go-proxy) in a very simplistic way

```sh
# Basic launch instruction using the default configuration path
docker run -it --rm -p 8080:8080 -v ${PWD}/myconfig.yaml:/config/goproxy.yaml atomys/go-proxy:latest
# Use custom configuration file
docker run -it --rm -p 8080:8080 -v ${PWD}/myconfig.yaml:/myconfig.yaml atomys/go-proxy:latest serve --config /myconfig.yaml
```

### With pre-builded binary

```sh
./goproxy serve --config config.yaml -p 8080
```

## To-Do

TO-Do is moving on Project Section: https://github.com/42Atomys/go-proxy/projects?type=beta

# Contribution

All pull requests and issues on GitHub will welcome.

All contributions are welcome :)

## Thanks
