# pulumi-gophers

This repo allows you to deploy your apps with Pulumi and the Docker Pulumi provider.

Please read the article in order to understand everything about this code :-).

## Pre-requisites

* Install pulumi CLI:

```
brew install pulumi/tap/pulumi
```

* OVHcloud API credentials.

* Pulumi account and access token

## Creation of the Pulumi Go app

Init the project:

```
$ pulumi new go --force
```

Get the dependencies:

```
$ go get github.com/pulumi/pulumi-docker/sdk/v3@v3.6.1
$ go get github.com/pulumi/pulumi/sdk/v3@v3.44.2
```

Edit the `main.go` file.

Define ports:

```bash
$ pulumi config set gophersAPIPort 8080
$ pulumi config set gophersAPIWatcherPort 8000
```

Ask Go to download dependencies and update `go.mod` file:

```bash
$ go mod tidy
```

Deploy:

```bash
$ pulumi up
```

## Cleanup

```bash
$ pulumi destroy
```

## Usage
