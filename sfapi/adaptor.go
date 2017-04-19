package sfapi

import (
	log "github.com/Sirupsen/logrus"
	"github.com/solidfire/solidfire-sdk-golang/sftypes"
)

func (c *Client) AddAccountAdaptor(request sftypes.AddAccountRequest) (sftypes.AddAccountResult, error) {
	params := map[string]interface{}{"username": request.Username}

	if request.InitiatorSecret.Secret == sftypes.CHAPSecretAutoGenerate() {
		params["initiatorSecret"] = nil
	} else if len(request.InitiatorSecret.Secret) > 0 {
		params["initiatorSecret"] = request.InitiatorSecret
	}

	if request.TargetSecret.Secret == sftypes.CHAPSecretAutoGenerate() {
		params["targetSecret"] = nil
	} else if len(request.TargetSecret.Secret) > 0 {
		params["targetSecret"] = request.TargetSecret
	}

	if request.Attributes != nil {
		params["attributes"] = request.Attributes
	}

	response, err := c.SendRequest("AddAccount", params)
	if err != nil {
		log.Errorf("Err: %v", err)
	}
	var result sftypes.AddAccountResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
	}
	return result, nil
}

func (c *Client) CreateScheduleAdaptor(request sftypes.CreateScheduleRequest) (sftypes.CreateScheduleResult, error) {
	log.Fatal("Not implemented")
	return sftypes.CreateScheduleResult{}, nil
}

func (c *Client) GetDriveStatsAdaptor(request sftypes.GetDriveStatsRequest) (sftypes.GetDriveStatsResult, error) {
	log.Fatal("Not implemented")
	return sftypes.GetDriveStatsResult{}, nil
}

func (c *Client) GetNodeStatsAdaptor(request sftypes.GetNodeStatsRequest) (sftypes.GetNodeStatsResult, error) {
	log.Fatal("Not implemented")
	return sftypes.GetNodeStatsResult{}, nil
}

func (c *Client) GetScheduleAdaptor(request sftypes.GetScheduleRequest) (sftypes.GetScheduleResult, error) {
	log.Fatal("Not implemented")
	return sftypes.GetScheduleResult{}, nil
}

func (c *Client) InvokeSFApiAdaptor(request sftypes.InvokeSFApiRequest) (interface{}, error) {
	log.Fatal("Not implemented")
	return nil, nil
}

func (c *Client) ListSchedulesAdaptor() (sftypes.ListSchedulesResult, error) {
	log.Fatal("Not implemented")
	return sftypes.ListSchedulesResult{}, nil
}

func (c *Client) ModifyScheduleAdaptor(request sftypes.ModifyScheduleRequest) (sftypes.ModifyScheduleResult, error) {
	log.Fatal("Not implemented")
	return sftypes.ModifyScheduleResult{}, nil
}

func (c *Client) ModifyAccountAdaptor(request sftypes.ModifyAccountRequest) (sftypes.ModifyAccountResult, error) {
	params := map[string]interface{}{
		"accountID": request.AccountID}

	if len(request.Username) > 0 {
		params["username"] = request.Username
	}
	if len(request.Username) > 0 {
		params["status"] = request.Status
	}
	if request.InitiatorSecret.Secret == sftypes.CHAPSecretAutoGenerate() {
		params["initiatorSecret"] = nil
	} else if len(request.InitiatorSecret.Secret) > 0 {
		params["initiatorSecret"] = request.InitiatorSecret
	}
	if request.TargetSecret.Secret == sftypes.CHAPSecretAutoGenerate() {
		params["targetSecret"] = nil
	} else if len(request.TargetSecret.Secret) > 0 {
		params["targetSecret"] = request.TargetSecret
	}
	if request.Attributes != nil {
		params["attributes"] = request.Attributes
	}
	response, err := c.SendRequest("ModifyAccount", params)
	if err != nil {
		log.Errorf("Err: %v", err)
	}
	var result sftypes.ModifyAccountResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
	}
	return result, nil
}
