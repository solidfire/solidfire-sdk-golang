package test

import (
	"github.com/solidfire/solidfire-sdk-golang/sfapi"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func init() {
	//log.SetOutput(os.Stdout)
}

func TestNewClient(t *testing.T) {
	c, err := sfapi.NewFromOpts("10.117.61.42", "admin", "admin", "", 0, 0)

	assert.True(t, c.Target == "10.117.61.42")
	assert.True(t, c.Credentials.Username == "admin")
	assert.True(t, c.Credentials.Password == "admin")
	assert.True(t, c.Port == 443)
	assert.True(t, c.Version == "6.0")
	assert.True(t, c.Timeout == 40)
	assert.Nil(t, err)
}

func TestSfApi_SendRequest_GetAPI(t *testing.T) {
	c, err := sfapi.NewFromOpts("10.117.61.42", "admin", "admin", "", 0, 0)
	assert.Nil(t, err)
	result, e := c.GetAPI()
	assert.Nil(t, e)
	assert.True(t, result.CurrentVersion == 9.0)
	assert.True(t, len(result.SupportedVersions) > 0)
}

func TestSfApi_SendRequest_ListTests(t *testing.T) {
	c, err := sfapi.NewFromOpts("10.117.61.42", "admin", "admin", "", 442, 0)
	assert.Nil(t, err)
	result, e := c.ListTests()
	assert.Nil(t, e)
	var tests []string
	decoder, err := sfapi.GetDecoder(&tests)
	err = decoder.Decode(result.Tests)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	assert.True(t, len(tests) > 0)
}
