package test

import (
	"testing"

	log "github.com/Sirupsen/logrus"
	"github.com/solidfire/solidfire-sdk-golang/sfapi"
	"github.com/stretchr/testify/assert"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

// This constructs a Client with specified or default values
func TestSfApi_NewFromOpts(t *testing.T) {
	c, err := sfapi.NewFromOpts("10.117.61.42", "admin", "admin", "", 0, 0)

	assert.True(t, c.Target == "10.117.61.42")
	assert.True(t, c.Credentials.Username == "admin")
	assert.True(t, c.Credentials.Password == "admin")
	assert.True(t, c.Port == 443)
	assert.True(t, c.Version == "6.0")
	assert.True(t, c.Timeout == 40)
	assert.Nil(t, err)
}

// This is similar to the factory methods in the other SDKs that will auto-discover the current version and name of the target cluster
func TestSfApi_Create(t *testing.T) {
	c, err := sfapi.Create("10.117.61.42", "admin", "admin", "", 0, 0)
	assert.True(t, c.Target == "10.117.61.42")
	assert.True(t, c.Credentials.Username == "admin")
	assert.True(t, c.Credentials.Password == "admin")
	assert.True(t, c.Port == 443)
	assert.True(t, c.Version == "9.0")
	assert.True(t, c.Timeout == 40)
	assert.True(t, c.Name == "AHcluster")
	assert.Nil(t, err)
}

// This tests running a simple API call
func TestSfApi_SendRequest_GetAPI(t *testing.T) {
	c, err := sfapi.NewFromOpts("10.117.61.42", "admin", "admin", "", 0, 0)
	assert.Nil(t, err)
	result, e := c.GetAPI()
	assert.Nil(t, e)
	assert.True(t, result.CurrentVersion == 9.0)
	assert.True(t, len(result.SupportedVersions) > 0)
}

// This tests running an API call that returns an interface in the result and then decodes the a value in the interface.
func TestSfApi_SendRequest_ListTests(t *testing.T) {
	c, err := sfapi.NewFromOpts("10.117.61.42", "admin", "admin", "", 442, 0)
	assert.Nil(t, err)
	result, e := c.ListTests()
	assert.Nil(t, e)
	var tests []string
	decoder, err := sfapi.GetDecoder(&tests)
	err = decoder.Decode(result.Tests)
	if err != nil {
		log.Error("Err: %v", err)
	}
	assert.True(t, len(tests) > 0)
}
