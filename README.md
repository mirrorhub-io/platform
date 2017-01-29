# Mirrorhub.io Platform

[![Go Report Card](https://goreportcard.com/badge/github.com/mirrorhub-io/platform)](https://goreportcard.com/report/github.com/mirrorhub-io/platform)
[![Docker Repository on Quay](https://quay.io/repository/mirrorhub/platform/status "Docker Repository on Quay")](https://quay.io/repository/mirrorhub/platform)

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
