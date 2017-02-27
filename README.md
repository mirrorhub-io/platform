# Mirrorhub.io Platform

[![Go Report Card](https://goreportcard.com/badge/github.com/mirrorhub-io/platform)](https://goreportcard.com/report/github.com/mirrorhub-io/platform)
[![Docker Repository on Quay](https://quay.io/repository/mirrorhub/platform/status "Docker Repository on Quay")](https://quay.io/repository/mirrorhub/platform)
[![codebeat badge](https://codebeat.co/badges/605c5d37-6f31-44c9-afeb-a0833251b930)](https://codebeat.co/projects/github-com-mirrorhub-io-platform)
[![GoDoc](https://godoc.org/github.com/mirrorhub-io/platform?status.svg)](https://godoc.org/github.com/mirrorhub-io/platform)

Our goal is to provide MirrorAsAService for everybody. Everybody knows it, you're running an application stack and you don't care about traffic. It is not as much to provide anybody a mirroring functionality but it wouldn't hurt to do.

We are trying to provide you as soon as possible a full stack solution which allows you to stop or hold mirroring at any time you want, with configurable traffic and storage limits. Possibly you want to support any open source project with up to 100GB storage but your are only able to provide 30% of them. 

**Coming soon**

## CLI

### Installation

Download latest prebuilt binaries package. Select the correct binary and just run it.

```
Mirrorhub root command.

Usage:
  mirrorhub [command]

Available Commands:
  api          Start mirrorhub api
  autocomplete Generate shell autocompletion script for Mirrorhub
  client       Mirrorhub API-Client
  gateway      Start mirrorhub rest-gateway

Flags:
      --config string   config file (default is $HOME/.mirrorhub.yaml)

Use "mirrorhub [command] --help" for more information about a command.
```

A config could looks like the following.

```yaml
Email: huimoo@example.org
Password: supersecurepassword
API:
  base: localhost:9000
```

## API

**Checkout protocol buffers under controllers/proto/api.proto**

Authorize (HTTP-Header) from Client

```
Grpc-Metadata-ClientToken: <Token>
```

Authorize (HTTP-Header) from Frontend

```
Grpc-Metadata-ContactToken: <Token>
```
