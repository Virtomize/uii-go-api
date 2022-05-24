package unattended_install_client

import (
    "bytes"
    "crypto/tls"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
)

type VirtomizeCloudApiClient struct {
    token  string
    client httpClient
}

// Build builds a ISO for a given operating system configuration.
func (c VirtomizeCloudApiClient) Build(filePath string, requiredArguments BuildArgs, optionalArguments BuildOpts) error {
    uri := constructURLForEndpoint("/images")
    jsonString, err := mergeToJsonByteSlice(requiredArguments, optionalArguments)
    if err != nil {
        return err
    }

    req, err := http.NewRequest("POST", uri, bytes.NewBuffer(jsonString))
    if err != nil {
        return err
    }
    req.Header.Add("Content-Type", "application/json")
    req.Header.Add("Authorization", "Bearer "+c.token)
    addSharedHeaders(&req.Header)

    resp, err := c.client.Do(req)
    defer resp.Body.Close()
    if err != nil {
        return err
    }

    // check the StatusCode and parse your response
    b, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return err
    }

    if resp.StatusCode != http.StatusOK {
        returnedError := parseError(b)
        if returnedError != "" {
            return fmt.Errorf(returnedError)
        }

        return fmt.Errorf("request was not successful and returned status %d and body \"%s\"", resp.StatusCode, string(b))
    }

    // return type is os.File
    var file *os.File
    file, err = os.Create(filePath)
    if err != nil {
        return err
    }

    _, err = file.Write(b)
    if err != nil {
        _ = os.Remove(filePath)
        return err
    }

    return file.Close()
}

// ReadOperationSystems reads all available operating systems.
func (c VirtomizeCloudApiClient) ReadOperationSystems() ([]Os, error) {
    uri := constructURLForEndpoint("/oslist")
    stringBody, err := c.requestStringResult(http.MethodGet, uri, []byte{})
    if err != nil {
        return []Os{}, err
    }

    var jsonResponse struct {
        Embedded []Os `json:"_embedded"`
    }

    err = json.Unmarshal([]byte(stringBody), &jsonResponse)
    if err != nil {
        return []Os{}, err
    }

    return jsonResponse.Embedded, nil
}

// ReadPackages reads all available packages for a given OS.
func (c VirtomizeCloudApiClient) ReadPackages(requiredArguments PackageArgs, optionalArguments PackageOpts) ([]string, error) {
    jsonPayload, err := mergeToJsonByteSlice(requiredArguments, optionalArguments)
    if err != nil {
        return []string{}, err
    }

    uri := constructURLForEndpoint("/packages")
    stringBody, err := c.requestStringResult(http.MethodPost, uri, jsonPayload)
    if err != nil {
        return []string{}, err
    }

    var jsonResponse struct {
        Embedded PackageListResponse `json:"_embedded"`
    }

    err = json.Unmarshal([]byte(stringBody), &jsonResponse)
    if err != nil {
        return []string{}, err
    }

    return jsonResponse.Embedded.Packages, nil
}

// NewClient creates a new cloud api client.
func NewClient(apiToken string) VirtomizeCloudApiClient {
    var tranport = http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
    tranport.TLSClientConfig.InsecureSkipVerify = false
    return VirtomizeCloudApiClient{
        token:  apiToken,
        client: &http.Client{Transport: &tranport},
    }
}

func constructURLForEndpoint(path string) string {
    return fmt.Sprintf("https://api.virtomize.com/uii%s", path)
}
