package unattended_install_client

import (
    "bytes"
    "github.com/stretchr/testify/assert"
    "io/ioutil"
    "net/http"
    "testing"
)

type httpClientMock struct {
    StatusCode   int
    ResponseBody string
    Requests     []httpRequestLogEntry
}

type httpRequestLogEntry struct {
    Path string
    Body string
}

func (m *httpClientMock) Do(req *http.Request) (*http.Response, error) {
    // create a new reader with that JSON
    b := ioutil.NopCloser(bytes.NewReader([]byte(m.ResponseBody)))
    resp := http.Response{
        Status:           "",
        StatusCode:       m.StatusCode,
        Proto:            "",
        ProtoMajor:       0,
        ProtoMinor:       0,
        Header:           nil,
        Body:             b,
        ContentLength:    0,
        TransferEncoding: nil,
        Close:            false,
        Uncompressed:     false,
        Trailer:          nil,
        Request:          nil,
        TLS:              nil,
    }

    buffer := new(bytes.Buffer)
    _, _ = buffer.ReadFrom(req.Body)

    m.Requests = append(m.Requests, httpRequestLogEntry{
        req.URL.Path,
        buffer.String(),
    })

    return &resp, nil
}

func TestReadOsRequestCorrectPath(t *testing.T) {
    mockClient := httpClientMock{StatusCode: 200}

    client := NewClient("1234")
    client.client = &mockClient
    _, _ = client.ReadOperationSystems()

    assert.EqualValues(t, 1, len(mockClient.Requests))
    assert.EqualValues(t, "/uii/oslist", mockClient.Requests[0].Path)
}

func TestReadOsReturnsCorrectError(t *testing.T) {
    mockClient := httpClientMock{StatusCode: 400,
        ResponseBody: "{\n\t\"errors\": [\n\t\t\"embedded error\"\n\t],\n\t\"timestamp\": \"2022-01-16T18:12:52+0000\",\n\t\"statuscode\": 400,\n\t\"instance\": \"/url\"\n}"}

    client := NewClient("1234")
    client.client = &mockClient
    _, err := client.ReadOperationSystems()

    assert.Equal(t, "embedded error", err.Error())
}

func TestReadOsReturnsCorrectErrorOnNonStandardResult(t *testing.T) {
    mockClient := httpClientMock{StatusCode: 400,
        ResponseBody: "this is no json"}

    client := NewClient("1234")
    client.client = &mockClient
    _, err := client.ReadOperationSystems()

    assert.Equal(t, "request was not successful and returned status 400 and body \"this is no json\"", err.Error())
}
