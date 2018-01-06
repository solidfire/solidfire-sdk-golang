package sfapi

import (
	"errors"

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
	log.Error("Not implemented")
	return sftypes.CreateScheduleResult{}, errors.New("function CreateScheduleAdaptor is Not Implemented")
}

func (c *Client) GetDriveStatsAdaptor(request sftypes.GetDriveStatsRequest) (sftypes.GetDriveStatsResult, error) {
	log.Error("Not implemented")
	return sftypes.GetDriveStatsResult{}, errors.New("function GetDriveStatsAdaptor is Not Implemented")
}

func (c *Client) GetNodeStatsAdaptor(request sftypes.GetNodeStatsRequest) (sftypes.GetNodeStatsResult, error) {
	log.Error("Not implemented")
	return sftypes.GetNodeStatsResult{}, errors.New("function GetNodeStatsAdaptor is Not Implemented")
}

func (c *Client) GetScheduleAdaptor(request sftypes.GetScheduleRequest) (sftypes.GetScheduleResult, error) {
	log.Error("Not implemented")
	return sftypes.GetScheduleResult{}, errors.New("function GetScheduleAdaptor is Not Implemented")
}

func (c *Client) InvokeSFApiAdaptor(request sftypes.InvokeSFApiRequest) (interface{}, error) {
	log.Error("Not implemented")
	return nil, errors.New("function InvokeSFApiAdaptor is Not Implemented")
}

func (c *Client) ListSchedulesAdaptor() (sftypes.ListSchedulesResult, error) {
	log.Error("Not implemented")
	return sftypes.ListSchedulesResult{}, errors.New("function ListSchedulesAdaptor is Not Implemented")
}

func (c *Client) ModifyScheduleAdaptor(request sftypes.ModifyScheduleRequest) (sftypes.ModifyScheduleResult, error) {
	log.Error("Not implemented")
	return sftypes.ModifyScheduleResult{}, errors.New("function GetScheduleAdaptor is Not Implemented")
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
