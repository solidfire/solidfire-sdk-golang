package sfapi

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type Client struct {
	Target       string
	Port         int
	RequestCount int64
	Credentials  Credentials
	Timeout      int
	Version      string
	ApiUrl       string
	Name         string
}

type Credentials struct {
	Username string
	Password string
}

type APIError struct {
	Id    int `json:"id"`
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Name    string `json:"name"`
	} `json:"error"`
}

var (
	endpoint string
	creds    Credentials
	ver      string
	prt      int
	timeout  int
)

func init() {

}

func Create(target string, username string, password string, version string, port int, timeoutSecs int) (c *Client, err error) {
	client, err := NewFromOpts(target, username, password, ver, port, timeoutSecs)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}

	// set to current version
	getApiResult, err := client.GetAPI()
	if err != nil {
		log.Errorf("Error retrieving Cluster version: %v", err)
		return nil, err
	}
	client.Version = strconv.FormatFloat(getApiResult.CurrentVersion, 'f', 1, 64)

	if client.Port == 443 {
		getClusterInfoResult, err := client.GetClusterInfo()
		if err != nil {
			log.Errorf("Error retrieving Cluster name: %v", err)
			return nil, err
		}
		client.Name = getClusterInfoResult.ClusterInfo.Name
	} else if client.Port == 442 {
		getConfigResult, err := client.GetConfig()
		if err != nil {
			log.Errorf("Error retrieving Node name: %v", err)
			return nil, err
		}
		client.Name = getConfigResult.Config.Cluster.Name
	}
	return client, err
}

func NewFromOpts(target string, username string, password string, version string, port int, timeoutSecs int) (c *Client, err error) {
	endpoint = target
	creds = Credentials{Username: username, Password: password}
	ver = version
	prt = port
	timeout = timeoutSecs
	return New()
}

func New() (c *Client, err error) {
	if endpoint == "" {
		log.Error("Target is not set, unable to issue requests")
		err = errors.New("Unable to issue json-rpc requests without specifying Target")
		return nil, err
	}

	if creds.Username == "" || creds.Password == "" {
		log.Error("Credentials are not set, unable to issue requests")
		err = errors.New("Unable to issue json-rpc requests without specifying Credentials")
		return nil, err
	}

	if prt == 0 {
		prt = 443
	}

	if timeout == 0 {
		timeout = 40
	}

	if ver == "" {
		ver = "6.0"
	}

	SFClient := &Client{
		Target:      endpoint,
		Credentials: creds,
		Version:     ver,
		Port:        prt,
		Timeout:     timeout}
	SFClient.ApiUrl = "https://" + endpoint + ":" + strconv.Itoa(prt) + "/json-rpc/" + ver
	return SFClient, nil
}

func (c *Client) SendRequest(method string, params interface{}) (response map[string]interface{}, err error) {
	// increment the request counter
	c.RequestCount++

	// create the request json
	data, err := json.Marshal(map[string]interface{}{
		"method": method,
		"id":     c.RequestCount,
		"params": params,
	})

	// create the http client with proper settings
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true},
		},
	}

	// make the request
	req, err := http.NewRequest("POST", c.ApiUrl, strings.NewReader(string(data)))
	req.SetBasicAuth(c.Credentials.Username, c.Credentials.Password)
	log.Debugf("Request: %+v", string(data))
	resp, err := client.Do(req)
	if err != nil {
		log.Errorf("Error encountered posting request: %v", err)
		return nil, err
	}

	// read the response into a byte array
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	log.Debugf("Response: %+v", string(body))
	if err != nil {
		log.Errorf("Error reading response body: %v", err)
		return nil, err
	}

	// decode the response into a map
	r := bytes.NewReader(body)
	var result map[string]interface{}
	err = json.NewDecoder(r).Decode(&result)
	if err != nil {
		log.Errorf("Error decoding response into map: %v", err)
		return nil, err
	}

	// check for errors
	errresp := APIError{}
	json.Unmarshal([]byte(body), &errresp)
	if errresp.Error.Code != 0 {
		err = errors.New("Received error response from API request: " + *&errresp.Error.Message)
		return nil, err
	}

	// return successful response
	return result, nil
}
