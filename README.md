Learn more about [Virtomize](virtomize.com) 
or follow us on social media 
[Twitter](http://twitter.com/virtomize) 
[LinkedIn](http://www.linkedin.com/company/virtomize)

# UUI Golang API
This repository contains an implementation of a client for the [**Virtomize Unattended Install Images API**](https://uii.virtomize.com/).

## Add to project

To use this client, first add it to your Go project by running `go get` from your command line

``` bash 
go get github.com/virtomize/uii-go-api
```

## Build an ISO
Building an ISO requires two steps.
1. Create a client object by using the `NewClient` function. 
This requires the API token created in the UI (see below). 

```
    c := client.NewClient("myTokenHere")
```

2. Building the ISO by using `Build` on the client object. 
    ```
   err := c.Build(l.Output, 
        client.BuildArgs{
            Distribution: "debian",
            Version:      "10",
            Hostname:     "host",
            Networks: []client.NetworkArgs{            
                DHCP:       true,
                NoInternet: false,
            },
       },         
       client.BuildOpts{})
   ```
   
   `Build` requires 3 parameters: 
   - A path to the output file
   - A set of required parameters represented by the `BuildArgs` struct. 
    This struct should be fully initialized with valid values.
   - A set of optional parameters represented by the `BuildOpts` struct.
   All members of this struct can be left uninitialized and the API will pick sensible default values.

## Register an API token
Before using this client library, you need to register an API token under your account.
For this, login into the [website](virtomize.com) and go to the "API-Token" tab.

![API-Token tab](https://github.com/Virtomize/uii_go_api/blob/60f79a50fc429f630eba553aaf057e6daa12ef97/doc/api-token.png "API-Token tab")

There, create a new API token by clicking "Create" and give it a name. Then click save.
![Create a token](https://github.com/Virtomize/uii_go_api/blob/60f79a50fc429f630eba553aaf057e6daa12ef97/doc/api-token-create.png "Create a token")

A new token will be created.
Copy this token, as it will be hidden, once you log out.
There, create a new API token by clicking "Create" and give it a name. Then click save.
![Save token](https://github.com/Virtomize/uii_go_api/blob/60f79a50fc429f630eba553aaf057e6daa12ef97/doc/api-token-created.png "Save token")

