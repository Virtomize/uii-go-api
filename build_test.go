package unattended_install_client

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestBuildRequestCorrectPath(t *testing.T) {
    mockClient := httpClientMock{StatusCode: 400}

    client := NewClient("1234")
    client.client = &mockClient
    _ = client.Build("foo", BuildArgs{
        Distribution: "UnitTest",
        Version:      "1",
        Hostname:     "Horst",
        Networks:     []NetworkArgs{{DHCP: true}},
    }, BuildOpts{})

    assert.EqualValues(t, 1, len(mockClient.Requests))
    assert.EqualValues(t, "/uii/images", mockClient.Requests[0].Path)
}

func TestBuildSendsValidJson(t *testing.T) {
    mockClient := httpClientMock{StatusCode: 400}

    client := NewClient("1234")
    client.client = &mockClient
    _ = client.Build("foo", BuildArgs{
        Distribution: "UnitTest",
        Version:      "1",
        Hostname:     "Horst",
        Networks:     []NetworkArgs{{DHCP: true}},
    }, BuildOpts{})

    expectedJson :=
        "{\"arch\":\"\",\"dist\":\"UnitTest\",\"hostname\":\"Horst\",\"keyboard\":\"\",\"locale\":\"\",\"networks\":[{\"dhcp\":true}],\"packages\":null,\"password\":\"\",\"sshkeys\":null,\"sshpasswordauth\":false,\"timezone\":\"\",\"version\":\"1\"}"

    assert.EqualValues(t, 1, len(mockClient.Requests))
    assert.EqualValues(t, expectedJson, mockClient.Requests[0].Body)
}

func TestBuildReturnsCorrectError(t *testing.T) {
    mockClient := httpClientMock{StatusCode: 400,
        ResponseBody: "{\n\t\"errors\": [\n\t\t\"none of your networks has a valid internet connection\"\n\t],\n\t\"timestamp\": \"2022-01-16T18:12:52+0000\",\n\t\"statuscode\": 400,\n\t\"instance\": \"/images\"\n}"}

    client := NewClient("1234")
    client.client = &mockClient
    err := client.Build("foo", BuildArgs{
        Distribution: "UnitTest",
        Version:      "1",
        Hostname:     "Horst",
        Networks:     []NetworkArgs{NetworkArgs{DHCP: true}},
    }, BuildOpts{})

    assert.Equal(t, "none of your networks has a valid internet connection", err.Error())
}

func TestBuildReturnsCorrectErrorOnNonStandardResult(t *testing.T) {
    mockClient := httpClientMock{StatusCode: 400,
        ResponseBody: "this is no json"}

    client := NewClient("1234")
    client.client = &mockClient
    err := client.Build("foo", BuildArgs{
        Distribution: "UnitTest",
        Version:      "1",
        Hostname:     "Horst",
        Networks:     []NetworkArgs{NetworkArgs{DHCP: true}},
    }, BuildOpts{})

    assert.Equal(t, "request was not successful and returned status 400 and body \"this is no json\"", err.Error())
}
