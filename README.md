# Mirrorhub.io Platform

[![Go Report Card](https://goreportcard.com/badge/github.com/mirrorhub-io/platform)](https://goreportcard.com/report/github.com/mirrorhub-io/platform)
[![Docker Repository on Quay](https://quay.io/repository/mirrorhub/platform/status "Docker Repository on Quay")](https://quay.io/repository/mirrorhub/platform)
[![codebeat badge](https://codebeat.co/badges/605c5d37-6f31-44c9-afeb-a0833251b930)](https://codebeat.co/projects/github-com-mirrorhub-io-platform)
[![GoDoc](https://godoc.org/github.com/mirrorhub-io/platform?status.svg)](https://godoc.org/github.com/mirrorhub-io/platform)

Our goal is to provide MirrorAsAService for everybody. Everybody knows it, you're running an application stack and you don't care about traffic. It is not as much to provide anybody a mirroring functionality but it wouldn't hurt to do.

We are trying to provide you as soon as possible a full stack solution which allows you to stop or hold mirroring at any time you want, with configurable traffic and storage limits. Possibly you want to support any open source project with up to 100GB storage but your are only able to provide 30% of them. 

**Coming soon**

## API

Authorize (HTTP-Header) from Client

```
Grpc-Metadata-ClientToken: <Token>
```

Authorize (HTTP-Header) from Frontend

```
Grpc-Metadata-ContactToken: <Token>
```

Current routes.

```
GET /v1/mirrors
POST /v1/mirrors
GET /v1/contacts/self   -> Get contact info (Auth!)
PUT /v1/contacts/self   -> Update contact info, new token will be set (Auth!)
POST /v1/contacts       -> Create new contact {name: <display_name>, email: ..., password: ...}
POST /v1/contacts/auth  -> Login contact {email: <email>, password: <pass>}
```
