# Mirrorhub.io Platform

[![Go Report Card](https://goreportcard.com/badge/github.com/mirrorhub-io/platform)](https://goreportcard.com/report/github.com/mirrorhub-io/platform)
[![Docker Repository on Quay](https://quay.io/repository/mirrorhub/platform/status "Docker Repository on Quay")](https://quay.io/repository/mirrorhub/platform)
[![codebeat badge](https://codebeat.co/badges/605c5d37-6f31-44c9-afeb-a0833251b930)](https://codebeat.co/projects/github-com-mirrorhub-io-platform)
[![GoDoc](https://godoc.org/github.com/mirrorhub-io/platform?status.svg)](https://godoc.org/github.com/mirrorhub-io/platform)

**Coming soon**

## API

Authorize (HTTP-Header) from Client

```
Grpc-Metadata-ClientToken: <Token>
```

Authorize (HTTP-Header) from Frontend

```
# Only first time ...
Grpc-Metadata-ContactEmail: <Email>
Grpc-Metadata-ContactPassword: <Pass>

# Session Token
Grpc-Metadata-ContactToken: <Token>
```

Current routes.

```
GET /v1/mirrors
POST /v1/mirrors
```
