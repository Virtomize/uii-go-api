[![Donate](https://img.shields.io/badge/Donate-PayPal-green.svg)](https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=VBXHBYFU44T5W&source=url)
[![GoDoc](https://img.shields.io/badge/godoc-reference-green.svg)](https://godoc.org/github.com/Virtomize/uii-go-api)
[![Go Report Card](https://goreportcard.com/badge/github.com/Virtomize/uii-go-api)](https://goreportcard.com/report/github.com/virtomize/uii-go-api)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/virtomize/uii-go-api/blob/master/LICENSE)
[![Website](https://img.shields.io/badge/%40-Virtomize-%231e828c)](https://virtomize.com)
[![Twitter](https://badgen.net/badge/icon/twitter?icon=twitter&label)](https://twitter.com/virtomize)
[![LinkedIn](https://img.shields.io/badge/linkedIn-%20-blue.svg?style=social&logo=linkedin)](https://www.linkedin.com/company/virtomize/)

# UUI Golang API
This repository contains the go client implementation of [**Virtomize Unattended Install Images API**](https://uii.virtomize.com/).

## Add to project

To use this client, first add it to your Go project by running `go get` from your command line

``` bash 
go get github.com/virtomize/uii-go-api
```

## Build an ISO
Building an ISO requires two steps.
1. Create a client object by using the `NewClient` function. 
This requires the API token created in the UI (see below). 

    ```go
    c, err := client.NewClient("myTokenHere")
    ```

2. Building the ISO by using `Build` on the client object. 
    ```go
     c, err := uiiclient.NewClient("myTokenHere")
     if (err != nil){
         // handle error
     }
   
     err = c.Build("/tmp/my-iso.iso", uiiclient.BuildArgs{
       Distribution: "debian",
       Version:      "11",
       Hostname:     "my-debian",
       Networks: []uiiclient.NetworkArgs{
         DHCP:       true,
         NoInternet: false,
       },
     }, uiiclient.BuildOpts{})
     if err != nil {
       // handle error
     }
   ```
   
   `Build` requires 3 parameters: 
   - A path to the output file
   - A set of required parameters represented by the `BuildArgs` struct. 
    This struct should be fully initialized with valid values.
   - A set of optional parameters represented by the `BuildOpts` struct.
   All members of this struct can be left uninitialized and the API will pick sensible default values.

## Register an API token
Before using this client library, you need to register an API token under your account.
For this, login into the [website](https://uii.virtomize.com/) and go to the "API-Token" tab.

![API-Token tab](https://github.com/Virtomize/uii-go-api/blob/master/doc/api-token.png "API-Token tab")

There, create a new API token by clicking "Create" and give it a name. Then click save.
![Create a token](https://github.com/Virtomize/uii-go-api/blob/master/doc/api-token-create.png "Create a token")

A new token will be created.
Copy this token, as it will be hidden, once you log out.
There, create a new API token by clicking "Create" and give it a name. Then click save.
![Save token](https://github.com/Virtomize/uii-go-api/blob/master/doc/api-token-created.png "Save token")

# Contribution

Thank you for participating to this project.
Please see our [Contribution Guidlines](https://github.com/virtomize/uii-go-api/blob/master/CONTRIBUTING.md) for more information.

## Pre-Commit

This repo uses [pre-commit hooks](https://pre-commit.com/). Please install pre-commit and do `pre-commit install`

## Conventional Commits

Format commit messaged according to [Conventional Commits standard](https://www.conventionalcommits.org/en/v1.0.0/).

## Semantic Versioning

Whenever you need to version something make use of [Semantic Versioning](https://semver.org).
