package sfapi

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Client struct {
	Target       string
	Port         int
	RequestCount int64
	Credentials  Credentials
	Timeout      int
	Version      string
	ApiUrl       string
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

type APIResult struct {
	Id     int                    `json:"id"`
	Result map[string]interface{} `json:"result,omitempty"`
	Error  struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Name    string `json:"name"`
	} `json:"error,omitempty"`
}

var (
	endpoint string
	creds    Credentials
	ver      string
	prt      int
	timeout  int
)

func init() {
	log.SetLevel(log.DebugLevel)
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
	rand.Seed(time.Now().UTC().UnixNano())

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
		return nil, err
	}

	// decode the response into a map
	r := bytes.NewReader(body)
	var result map[string]interface{}
	decErr := json.NewDecoder(r).Decode(&result)
	if decErr != nil {
		log.Error(decErr)
	}
	// log request/response as pretty json
	//var prettyJson bytes.Buffer
	//_ = json.Indent(&prettyJson, body, "", "  ")
	//log.WithField("", prettyJson.String()).Debug("request:", c.RequestCount, " method:", method, " params:", params)

	// check for errors
	errresp := APIError{}
	json.Unmarshal([]byte(body), &errresp)
	if errresp.Error.Code != 0 {
		err = errors.New("Received error response from API request")
		return nil, err
	}

	// return successful response
	return result, nil
}
