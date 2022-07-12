/*
Package uiiclient implementing Virtomize UII REST API

Simple example:

	c,err := uiiclient.NewClient("myTokenHere")
    if err != nil {
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

*/
package uiiclient
