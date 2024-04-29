# pulumi-gophers

This repo allows you to deploy your apps with Pulumi and the Docker Pulumi provider.

Please read the article in order to understand everything about this code :-).

## Pre-requisites

* Install pulumi CLI:

```
brew install pulumi/tap/pulumi
```

* Pulumi account and access token (only if you wants to use Pulumi Cloud instead of using a local state)

## Creation of the Pulumi Go app

Init/log in the project:

```bash
$ pulumi login --local
```

Deploy:

```bash
$ pulumi up
```

## Cleanup

```bash
$ pulumi destroy
```
