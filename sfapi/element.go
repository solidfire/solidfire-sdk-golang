package sfapi

import (
	log "github.com/Sirupsen/logrus"
	"github.com/solidfire/solidfire-sdk-golang/sftypes"
	"github.com/ottemo/mapstructure"
)

func (c *Client) AddAccount(request sftypes.AddAccountRequest) (*sftypes.AddAccountResult, error) {
	// Adaptor
	result, err := c.AddAccountAdaptor(request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil
}

func (c *Client) AddClusterAdmin(request sftypes.AddClusterAdminRequest) (*sftypes.AddClusterAdminResult, error) {
	response, err := c.SendRequest("AddClusterAdmin", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.AddClusterAdminResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) AddDrives(request sftypes.AddDrivesRequest) (*sftypes.AddDrivesResult, error) {
	response, err := c.SendRequest("AddDrives", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.AddDrivesResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) AddInitiatorsToVolumeAccessGroup(request sftypes.AddInitiatorsToVolumeAccessGroupRequest) (*sftypes.ModifyVolumeAccessGroupResult, error) {
	response, err := c.SendRequest("AddInitiatorsToVolumeAccessGroup", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.ModifyVolumeAccessGroupResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) AddLdapClusterAdmin(request sftypes.AddLdapClusterAdminRequest) (*sftypes.AddLdapClusterAdminResult, error) {
	response, err := c.SendRequest("AddLdapClusterAdmin", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.AddLdapClusterAdminResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) AddNodes(request sftypes.AddNodesRequest) (*sftypes.AddNodesResult, error) {
	response, err := c.SendRequest("AddNodes", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.AddNodesResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) AddVirtualNetwork(request sftypes.AddVirtualNetworkRequest) (*sftypes.AddVirtualNetworkResult, error) {
	response, err := c.SendRequest("AddVirtualNetwork", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.AddVirtualNetworkResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) AddVolumesToVolumeAccessGroup(request sftypes.AddVolumesToVolumeAccessGroupRequest) (*sftypes.ModifyVolumeAccessGroupResult, error) {
	response, err := c.SendRequest("AddVolumesToVolumeAccessGroup", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.ModifyVolumeAccessGroupResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) CancelClone(request sftypes.CancelCloneRequest) error {
	response, err := c.SendRequest("CancelClone", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return err
	}
	if response != nil {
		log.Debugf("CancelClone: ", response)
	}
	return err
}

func (c *Client) CancelGroupClone(request sftypes.CancelGroupCloneRequest) error {
	response, err := c.SendRequest("CancelGroupClone", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return err
	}
	if response != nil {
		log.Debugf("CancelGroupClone: ", response)
	}
	return err
}

func (c *Client) ClearClusterFaults(request sftypes.ClearClusterFaultsRequest) error {
	response, err := c.SendRequest("ClearClusterFaults", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return err
	}
	if response != nil {
		log.Debugf("ClearClusterFaults: ", response)
	}
	return err
}

func (c *Client) CloneMultipleVolumes(request sftypes.CloneMultipleVolumesRequest) (*sftypes.CloneMultipleVolumesResult, error) {
	response, err := c.SendRequest("CloneMultipleVolumes", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.CloneMultipleVolumesResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) CloneVolume(request sftypes.CloneVolumeRequest) (*sftypes.CloneVolumeResult, error) {
	response, err := c.SendRequest("CloneVolume", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.CloneVolumeResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) CompleteClusterPairing(request sftypes.CompleteClusterPairingRequest) (*sftypes.CompleteClusterPairingResult, error) {
	response, err := c.SendRequest("CompleteClusterPairing", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.CompleteClusterPairingResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) CompleteVolumePairing(request sftypes.CompleteVolumePairingRequest) error {
	response, err := c.SendRequest("CompleteVolumePairing", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return err
	}
	if response != nil {
		log.Debugf("CompleteVolumePairing: ", response)
	}
	return err
}

func (c *Client) CopyVolume(request sftypes.CopyVolumeRequest) (*sftypes.CopyVolumeResult, error) {
	response, err := c.SendRequest("CopyVolume", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.CopyVolumeResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) CreateBackupTarget(request sftypes.CreateBackupTargetRequest) (*sftypes.CreateBackupTargetResult, error) {
	response, err := c.SendRequest("CreateBackupTarget", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.CreateBackupTargetResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) CreateCluster(request sftypes.CreateClusterRequest) error {
	response, err := c.SendRequest("CreateCluster", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return err
	}
	if response != nil {
		log.Debugf("CreateCluster: ", response)
	}
	return err
}

func (c *Client) CreateGroupSnapshot(request sftypes.CreateGroupSnapshotRequest) (*sftypes.CreateGroupSnapshotResult, error) {
	response, err := c.SendRequest("CreateGroupSnapshot", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.CreateGroupSnapshotResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) CreateInitiators(request sftypes.CreateInitiatorsRequest) (*sftypes.CreateInitiatorsResult, error) {
	response, err := c.SendRequest("CreateInitiators", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.CreateInitiatorsResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) CreateSchedule(request sftypes.CreateScheduleRequest) (*sftypes.CreateScheduleResult, error) {
	// Adaptor
	result, err := c.CreateScheduleAdaptor(request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil
}

func (c *Client) CreateSnapshot(request sftypes.CreateSnapshotRequest) (*sftypes.CreateSnapshotResult, error) {
	response, err := c.SendRequest("CreateSnapshot", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.CreateSnapshotResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) CreateStorageContainer(request sftypes.CreateStorageContainerRequest) (*sftypes.CreateStorageContainerResult, error) {
	response, err := c.SendRequest("CreateStorageContainer", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.CreateStorageContainerResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) CreateSupportBundle(request sftypes.CreateSupportBundleRequest) (*sftypes.CreateSupportBundleResult, error) {
	response, err := c.SendRequest("CreateSupportBundle", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.CreateSupportBundleResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) CreateVolume(request sftypes.CreateVolumeRequest) (*sftypes.CreateVolumeResult, error) {
	response, err := c.SendRequest("CreateVolume", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.CreateVolumeResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) CreateVolumeAccessGroup(request sftypes.CreateVolumeAccessGroupRequest) (*sftypes.CreateVolumeAccessGroupResult, error) {
	response, err := c.SendRequest("CreateVolumeAccessGroup", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.CreateVolumeAccessGroupResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) DeleteAllSupportBundles() (*sftypes.DeleteAllSupportBundlesResult, error) {
	response, err := c.SendRequest("DeleteAllSupportBundles", nil)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.DeleteAllSupportBundlesResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) DeleteGroupSnapshot(request sftypes.DeleteGroupSnapshotRequest) error {
	response, err := c.SendRequest("DeleteGroupSnapshot", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return err
	}
	if response != nil {
		log.Debugf("DeleteGroupSnapshot: ", response)
	}
	return err
}

func (c *Client) DeleteInitiators(request sftypes.DeleteInitiatorsRequest) error {
	response, err := c.SendRequest("DeleteInitiators", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return err
	}
	if response != nil {
		log.Debugf("DeleteInitiators: ", response)
	}
	return err
}

func (c *Client) DeleteSnapshot(request sftypes.DeleteSnapshotRequest) error {
	response, err := c.SendRequest("DeleteSnapshot", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return err
	}
	if response != nil {
		log.Debugf("DeleteSnapshot: ", response)
	}
	return err
}

func (c *Client) DeleteStorageContainers(request sftypes.DeleteStorageContainersRequest) error {
	response, err := c.SendRequest("DeleteStorageContainers", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return err
	}
	if response != nil {
		log.Debugf("DeleteStorageContainers: ", response)
	}
	return err
}

func (c *Client) DeleteVolume(request sftypes.DeleteVolumeRequest) (*sftypes.DeleteVolumeResult, error) {
	response, err := c.SendRequest("DeleteVolume", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.DeleteVolumeResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) DeleteVolumeAccessGroup(request sftypes.DeleteVolumeAccessGroupRequest) error {
	response, err := c.SendRequest("DeleteVolumeAccessGroup", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return err
	}
	if response != nil {
		log.Debugf("DeleteVolumeAccessGroup: ", response)
	}
	return err
}

func (c *Client) DeleteVolumes(request sftypes.DeleteVolumesRequest) (*sftypes.DeleteVolumesResult, error) {
	response, err := c.SendRequest("DeleteVolumes", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.DeleteVolumesResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) DisableEncryptionAtRest() error {
	response, err := c.SendRequest("DisableEncryptionAtRest", nil)
	if err != nil {
		log.Errorf("Err: %v", err)
		return err
	}
	if response != nil {
		log.Debugf("DisableEncryptionAtRest: ", response)
	}
	return err
}

func (c *Client) DisableLdapAuthentication() error {
	response, err := c.SendRequest("DisableLdapAuthentication", nil)
	if err != nil {
		log.Errorf("Err: %v", err)
		return err
	}
	if response != nil {
		log.Debugf("DisableLdapAuthentication: ", response)
	}
	return err
}

func (c *Client) DisableSnmp() error {
	response, err := c.SendRequest("DisableSnmp", nil)
	if err != nil {
		log.Errorf("Err: %v", err)
		return err
	}
	if response != nil {
		log.Debugf("DisableSnmp: ", response)
	}
	return err
}

func (c *Client) EnableEncryptionAtRest() error {
	response, err := c.SendRequest("EnableEncryptionAtRest", nil)
	if err != nil {
		log.Errorf("Err: %v", err)
		return err
	}
	if response != nil {
		log.Debugf("EnableEncryptionAtRest: ", response)
	}
	return err
}

func (c *Client) EnableFeature(request sftypes.EnableFeatureRequest) error {
	response, err := c.SendRequest("EnableFeature", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return err
	}
	if response != nil {
		log.Debugf("EnableFeature: ", response)
	}
	return err
}

func (c *Client) EnableLdapAuthentication(request sftypes.EnableLdapAuthenticationRequest) error {
	response, err := c.SendRequest("EnableLdapAuthentication", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return err
	}
	if response != nil {
		log.Debugf("EnableLdapAuthentication: ", response)
	}
	return err
}

func (c *Client) EnableSnmp(request sftypes.EnableSnmpRequest) error {
	response, err := c.SendRequest("EnableSnmp", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return err
	}
	if response != nil {
		log.Debugf("EnableSnmp: ", response)
	}
	return err
}

func (c *Client) GetAPI() (*sftypes.GetAPIResult, error) {
	response, err := c.SendRequest("GetAPI", nil)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.GetAPIResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) GetAccountByID(request sftypes.GetAccountByIDRequest) (*sftypes.GetAccountResult, error) {
	response, err := c.SendRequest("GetAccountByID", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.GetAccountResult
	decoder, err := mapstructure.NewDecoder(sftypes.GetAccountResultDecodeHookConfig(&result))
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) GetAccountByName(request sftypes.GetAccountByNameRequest) (*sftypes.GetAccountResult, error) {
	response, err := c.SendRequest("GetAccountByName", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.GetAccountResult
	decoder, err := mapstructure.NewDecoder(sftypes.GetAccountResultDecodeHookConfig(&result))
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) GetAccountEfficiency(request sftypes.GetAccountEfficiencyRequest) (*sftypes.GetEfficiencyResult, error) {
	response, err := c.SendRequest("GetAccountEfficiency", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.GetEfficiencyResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) GetAsyncResult(request sftypes.GetAsyncResultRequest) (interface{}, error) {
	response, err := c.SendRequest("GetAsyncResult", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result interface{}
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) GetBackupTarget(request sftypes.GetBackupTargetRequest) (*sftypes.GetBackupTargetResult, error) {
	response, err := c.SendRequest("GetBackupTarget", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.GetBackupTargetResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) GetBootstrapConfig() (*sftypes.GetBootstrapConfigResult, error) {
	response, err := c.SendRequest("GetBootstrapConfig", nil)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.GetBootstrapConfigResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) GetClusterCapacity() (*sftypes.GetClusterCapacityResult, error) {
	response, err := c.SendRequest("GetClusterCapacity", nil)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.GetClusterCapacityResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) GetClusterConfig() (*sftypes.GetClusterConfigResult, error) {
	response, err := c.SendRequest("GetClusterConfig", nil)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.GetClusterConfigResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) GetClusterFullThreshold() (*sftypes.GetClusterFullThresholdResult, error) {
	response, err := c.SendRequest("GetClusterFullThreshold", nil)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.GetClusterFullThresholdResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) GetClusterHardwareInfo(request sftypes.GetClusterHardwareInfoRequest) (*sftypes.GetClusterHardwareInfoResult, error) {
	response, err := c.SendRequest("GetClusterHardwareInfo", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.GetClusterHardwareInfoResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) GetClusterInfo() (*sftypes.GetClusterInfoResult, error) {
	response, err := c.SendRequest("GetClusterInfo", nil)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.GetClusterInfoResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) GetClusterMasterNodeID() (*sftypes.GetClusterMasterNodeIDResult, error) {
	response, err := c.SendRequest("GetClusterMasterNodeID", nil)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.GetClusterMasterNodeIDResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) GetClusterState(request sftypes.GetClusterStateRequest) (*sftypes.GetClusterStateResult, error) {
	response, err := c.SendRequest("GetClusterState", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.GetClusterStateResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) GetClusterStats() (*sftypes.GetClusterStatsResult, error) {
	response, err := c.SendRequest("GetClusterStats", nil)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.GetClusterStatsResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) GetClusterVersionInfo() (*sftypes.GetClusterVersionInfoResult, error) {
	response, err := c.SendRequest("GetClusterVersionInfo", nil)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.GetClusterVersionInfoResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) GetCompleteStats() error {
	response, err := c.SendRequest("GetCompleteStats", nil)
	if err != nil {
		log.Errorf("Err: %v", err)
		return err
	}
	if response != nil {
		log.Debugf("GetCompleteStats: ", response)
	}
	return err
}

func (c *Client) GetConfig() (*sftypes.GetConfigResult, error) {
	response, err := c.SendRequest("GetConfig", nil)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.GetConfigResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) GetCurrentClusterAdmin() (*sftypes.GetCurrentClusterAdminResult, error) {
	response, err := c.SendRequest("GetCurrentClusterAdmin", nil)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.GetCurrentClusterAdminResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) GetDefaultQoS() (*sftypes.VolumeQOS, error) {
	response, err := c.SendRequest("GetDefaultQoS", nil)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.VolumeQOS
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) GetDriveConfig() (*sftypes.GetDriveConfigResult, error) {
	response, err := c.SendRequest("GetDriveConfig", nil)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.GetDriveConfigResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) GetDriveHardwareInfo(request sftypes.GetDriveHardwareInfoRequest) (*sftypes.GetDriveHardwareInfoResult, error) {
	response, err := c.SendRequest("GetDriveHardwareInfo", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.GetDriveHardwareInfoResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) GetDriveStats(request sftypes.GetDriveStatsRequest) (*sftypes.GetDriveStatsResult, error) {
	// Adaptor
	result, err := c.GetDriveStatsAdaptor(request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetFeatureStatus(request sftypes.GetFeatureStatusRequest) (*sftypes.GetFeatureStatusResult, error) {
	response, err := c.SendRequest("GetFeatureStatus", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.GetFeatureStatusResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) GetHardwareConfig() (*sftypes.GetHardwareConfigResult, error) {
	response, err := c.SendRequest("GetHardwareConfig", nil)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.GetHardwareConfigResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) GetHardwareInfo() (*sftypes.GetHardwareInfoResult, error) {
	response, err := c.SendRequest("GetHardwareInfo", nil)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.GetHardwareInfoResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) GetIpmiConfig(request sftypes.GetIpmiConfigRequest) (*sftypes.GetIpmiConfigResult, error) {
	response, err := c.SendRequest("GetIpmiConfig", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.GetIpmiConfigResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) GetIpmiInfo() (*sftypes.GetIpmiInfoResult, error) {
	response, err := c.SendRequest("GetIpmiInfo", nil)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.GetIpmiInfoResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) GetLdapConfiguration() (*sftypes.GetLdapConfigurationResult, error) {
	response, err := c.SendRequest("GetLdapConfiguration", nil)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.GetLdapConfigurationResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) GetLimits() (*sftypes.GetLimitsResult, error) {
	response, err := c.SendRequest("GetLimits", nil)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.GetLimitsResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) GetLoginSessionInfo() (*sftypes.GetLoginSessionInfoResult, error) {
	response, err := c.SendRequest("GetLoginSessionInfo", nil)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.GetLoginSessionInfoResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) GetNetworkConfig() (*sftypes.GetNetworkConfigResult, error) {
	response, err := c.SendRequest("GetNetworkConfig", nil)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.GetNetworkConfigResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) GetNodeHardwareInfo(request sftypes.GetNodeHardwareInfoRequest) (*sftypes.GetNodeHardwareInfoResult, error) {
	response, err := c.SendRequest("GetNodeHardwareInfo", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.GetNodeHardwareInfoResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) GetNodeStats(request sftypes.GetNodeStatsRequest) (*sftypes.GetNodeStatsResult, error) {
	// Adaptor
	result, err := c.GetNodeStatsAdaptor(request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetNtpInfo() (*sftypes.GetNtpInfoResult, error) {
	response, err := c.SendRequest("GetNtpInfo", nil)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.GetNtpInfoResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) GetNvramInfo(request sftypes.GetNvramInfoRequest) (*sftypes.GetNvramInfoResult, error) {
	response, err := c.SendRequest("GetNvramInfo", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.GetNvramInfoResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) GetOrigin() (*sftypes.GetOriginResult, error) {
	response, err := c.SendRequest("GetOrigin", nil)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.GetOriginResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) GetPendingOperation() (*sftypes.GetPendingOperationResult, error) {
	response, err := c.SendRequest("GetPendingOperation", nil)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.GetPendingOperationResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) GetRawStats() error {
	response, err := c.SendRequest("GetRawStats", nil)
	if err != nil {
		log.Errorf("Err: %v", err)
		return err
	}
	if response != nil {
		log.Debugf("GetRawStats: ", response)
	}
	return err
}

func (c *Client) GetRemoteLoggingHosts() (*sftypes.GetRemoteLoggingHostsResult, error) {
	response, err := c.SendRequest("GetRemoteLoggingHosts", nil)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.GetRemoteLoggingHostsResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) GetSchedule(request sftypes.GetScheduleRequest) (*sftypes.GetScheduleResult, error) {
	// Adaptor
	result, err := c.GetScheduleAdaptor(request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetSnmpACL() (*sftypes.GetSnmpACLResult, error) {
	response, err := c.SendRequest("GetSnmpACL", nil)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.GetSnmpACLResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) GetSnmpInfo() (*sftypes.GetSnmpInfoResult, error) {
	response, err := c.SendRequest("GetSnmpInfo", nil)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.GetSnmpInfoResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) GetSnmpState() (*sftypes.GetSnmpStateResult, error) {
	response, err := c.SendRequest("GetSnmpState", nil)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.GetSnmpStateResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) GetSnmpTrapInfo() (*sftypes.GetSnmpTrapInfoResult, error) {
	response, err := c.SendRequest("GetSnmpTrapInfo", nil)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.GetSnmpTrapInfoResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) GetStorageContainerEfficiency(request sftypes.GetStorageContainerEfficiencyRequest) (*sftypes.GetStorageContainerEfficiencyResult, error) {
	response, err := c.SendRequest("GetStorageContainerEfficiency", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.GetStorageContainerEfficiencyResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) GetSystemStatus() (*sftypes.GetSystemStatusResult, error) {
	response, err := c.SendRequest("GetSystemStatus", nil)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.GetSystemStatusResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) GetVirtualVolumeCount() (*sftypes.GetVirtualVolumeCountResult, error) {
	response, err := c.SendRequest("GetVirtualVolumeCount", nil)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.GetVirtualVolumeCountResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) GetVolumeAccessGroupEfficiency(request sftypes.GetVolumeAccessGroupEfficiencyRequest) (*sftypes.GetEfficiencyResult, error) {
	response, err := c.SendRequest("GetVolumeAccessGroupEfficiency", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.GetEfficiencyResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) GetVolumeAccessGroupLunAssignments(request sftypes.GetVolumeAccessGroupLunAssignmentsRequest) (*sftypes.GetVolumeAccessGroupLunAssignmentsResult, error) {
	response, err := c.SendRequest("GetVolumeAccessGroupLunAssignments", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.GetVolumeAccessGroupLunAssignmentsResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) GetVolumeCount() (*sftypes.GetVolumeCountResult, error) {
	response, err := c.SendRequest("GetVolumeCount", nil)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.GetVolumeCountResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) GetVolumeEfficiency(request sftypes.GetVolumeEfficiencyRequest) (*sftypes.GetVolumeEfficiencyResult, error) {
	response, err := c.SendRequest("GetVolumeEfficiency", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.GetVolumeEfficiencyResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) GetVolumeStats(request sftypes.GetVolumeStatsRequest) (*sftypes.GetVolumeStatsResult, error) {
	response, err := c.SendRequest("GetVolumeStats", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.GetVolumeStatsResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) InvokeSFApi(request sftypes.InvokeSFApiRequest) (interface{}, error) {
	// Adaptor
	result, err := c.InvokeSFApiAdaptor(request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil
}

func (c *Client) ListAccounts(request sftypes.ListAccountsRequest) (*sftypes.ListAccountsResult, error) {
	response, err := c.SendRequest("ListAccounts", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.ListAccountsResult
	decoder, err := mapstructure.NewDecoder(sftypes.ListAccountsResultDecodeHookConfig(&result))
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) ListActiveNodes() (*sftypes.ListActiveNodesResult, error) {
	response, err := c.SendRequest("ListActiveNodes", nil)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.ListActiveNodesResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) ListActivePairedVolumes(request sftypes.ListActivePairedVolumesRequest) (*sftypes.ListActivePairedVolumesResult, error) {
	response, err := c.SendRequest("ListActivePairedVolumes", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.ListActivePairedVolumesResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) ListActiveVolumes(request sftypes.ListActiveVolumesRequest) (*sftypes.ListActiveVolumesResult, error) {
	response, err := c.SendRequest("ListActiveVolumes", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.ListActiveVolumesResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) ListAllNodes() (*sftypes.ListAllNodesResult, error) {
	response, err := c.SendRequest("ListAllNodes", nil)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.ListAllNodesResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) ListAsyncResults(request sftypes.ListAsyncResultsRequest) (*sftypes.ListAsyncResultsResult, error) {
	response, err := c.SendRequest("ListAsyncResults", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.ListAsyncResultsResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) ListBackupTargets() (*sftypes.ListBackupTargetsResult, error) {
	response, err := c.SendRequest("ListBackupTargets", nil)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.ListBackupTargetsResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) ListBulkVolumeJobs() (*sftypes.ListBulkVolumeJobsResult, error) {
	response, err := c.SendRequest("ListBulkVolumeJobs", nil)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.ListBulkVolumeJobsResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) ListClusterAdmins() (*sftypes.ListClusterAdminsResult, error) {
	response, err := c.SendRequest("ListClusterAdmins", nil)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.ListClusterAdminsResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) ListClusterFaults(request sftypes.ListClusterFaultsRequest) (*sftypes.ListClusterFaultsResult, error) {
	response, err := c.SendRequest("ListClusterFaults", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.ListClusterFaultsResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) ListClusterPairs() (*sftypes.ListClusterPairsResult, error) {
	response, err := c.SendRequest("ListClusterPairs", nil)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.ListClusterPairsResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) ListDeletedVolumes(request sftypes.ListDeletedVolumesRequest) (*sftypes.ListDeletedVolumesResult, error) {
	response, err := c.SendRequest("ListDeletedVolumes", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.ListDeletedVolumesResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) ListDriveHardware(request sftypes.ListDriveHardwareRequest) (*sftypes.ListDriveHardwareResult, error) {
	response, err := c.SendRequest("ListDriveHardware", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.ListDriveHardwareResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) ListDriveStats(request sftypes.ListDriveStatsRequest) (*sftypes.ListDriveStatsResult, error) {
	response, err := c.SendRequest("ListDriveStats", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.ListDriveStatsResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) ListDrives() (*sftypes.ListDrivesResult, error) {
	response, err := c.SendRequest("ListDrives", nil)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.ListDrivesResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) ListEvents(request sftypes.ListEventsRequest) (*sftypes.ListEventsResult, error) {
	response, err := c.SendRequest("ListEvents", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.ListEventsResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) ListFibreChannelPortInfo() (*sftypes.ListFibreChannelPortInfoResult, error) {
	response, err := c.SendRequest("ListFibreChannelPortInfo", nil)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.ListFibreChannelPortInfoResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) ListFibreChannelSessions() (*sftypes.ListFibreChannelSessionsResult, error) {
	response, err := c.SendRequest("ListFibreChannelSessions", nil)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.ListFibreChannelSessionsResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) ListGroupSnapshots(request sftypes.ListGroupSnapshotsRequest) (*sftypes.ListGroupSnapshotsResult, error) {
	response, err := c.SendRequest("ListGroupSnapshots", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.ListGroupSnapshotsResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) ListISCSISessions() (*sftypes.ListISCSISessionsResult, error) {
	response, err := c.SendRequest("ListISCSISessions", nil)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.ListISCSISessionsResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) ListInitiators(request sftypes.ListInitiatorsRequest) (*sftypes.ListInitiatorsResult, error) {
	response, err := c.SendRequest("ListInitiators", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.ListInitiatorsResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) ListNetworkInterfaces() (*sftypes.ListNetworkInterfacesResult, error) {
	response, err := c.SendRequest("ListNetworkInterfaces", nil)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.ListNetworkInterfacesResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) ListNodeFibreChannelPortInfo() (*sftypes.ListNodeFibreChannelPortInfoResult, error) {
	response, err := c.SendRequest("ListNodeFibreChannelPortInfo", nil)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.ListNodeFibreChannelPortInfoResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) ListNodeStats() (*sftypes.ListNodeStatsResult, error) {
	response, err := c.SendRequest("ListNodeStats", nil)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.ListNodeStatsResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) ListPendingActiveNodes() (*sftypes.ListPendingActiveNodesResult, error) {
	response, err := c.SendRequest("ListPendingActiveNodes", nil)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.ListPendingActiveNodesResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) ListPendingNodes() (*sftypes.ListPendingNodesResult, error) {
	response, err := c.SendRequest("ListPendingNodes", nil)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.ListPendingNodesResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) ListProtocolEndpoints(request sftypes.ListProtocolEndpointsRequest) (*sftypes.ListProtocolEndpointsResult, error) {
	response, err := c.SendRequest("ListProtocolEndpoints", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.ListProtocolEndpointsResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) ListSchedules() (*sftypes.ListSchedulesResult, error) {
	// Adaptor
	result, err := c.ListSchedulesAdaptor()
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil
}

func (c *Client) ListServices() (*sftypes.ListServicesResult, error) {
	response, err := c.SendRequest("ListServices", nil)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.ListServicesResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) ListSnapshots(request sftypes.ListSnapshotsRequest) (*sftypes.ListSnapshotsResult, error) {
	response, err := c.SendRequest("ListSnapshots", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.ListSnapshotsResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) ListStorageContainers(request sftypes.ListStorageContainersRequest) (*sftypes.ListStorageContainersResult, error) {
	response, err := c.SendRequest("ListStorageContainers", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.ListStorageContainersResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) ListSyncJobs() (*sftypes.ListSyncJobsResult, error) {
	response, err := c.SendRequest("ListSyncJobs", nil)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.ListSyncJobsResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) ListTests() (*sftypes.ListTestsResult, error) {
	response, err := c.SendRequest("ListTests", nil)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.ListTestsResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) ListUtilities() (*sftypes.ListUtilitiesResult, error) {
	response, err := c.SendRequest("ListUtilities", nil)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.ListUtilitiesResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) ListVirtualNetworks(request sftypes.ListVirtualNetworksRequest) (*sftypes.ListVirtualNetworksResult, error) {
	response, err := c.SendRequest("ListVirtualNetworks", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.ListVirtualNetworksResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) ListVirtualVolumeBindings(request sftypes.ListVirtualVolumeBindingsRequest) (*sftypes.ListVirtualVolumeBindingsResult, error) {
	response, err := c.SendRequest("ListVirtualVolumeBindings", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.ListVirtualVolumeBindingsResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) ListVirtualVolumeHosts(request sftypes.ListVirtualVolumeHostsRequest) (*sftypes.ListVirtualVolumeHostsResult, error) {
	response, err := c.SendRequest("ListVirtualVolumeHosts", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.ListVirtualVolumeHostsResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) ListVirtualVolumeTasks(request sftypes.ListVirtualVolumeTasksRequest) (*sftypes.ListVirtualVolumeTasksResult, error) {
	response, err := c.SendRequest("ListVirtualVolumeTasks", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.ListVirtualVolumeTasksResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) ListVirtualVolumes(request sftypes.ListVirtualVolumesRequest) (*sftypes.ListVirtualVolumesResult, error) {
	response, err := c.SendRequest("ListVirtualVolumes", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.ListVirtualVolumesResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) ListVolumeAccessGroups(request sftypes.ListVolumeAccessGroupsRequest) (*sftypes.ListVolumeAccessGroupsResult, error) {
	response, err := c.SendRequest("ListVolumeAccessGroups", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.ListVolumeAccessGroupsResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) ListVolumeStats(request sftypes.ListVolumeStatsRequest) (*sftypes.ListVolumeStatsResult, error) {
	response, err := c.SendRequest("ListVolumeStats", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.ListVolumeStatsResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) ListVolumeStatsByAccount(request sftypes.ListVolumeStatsByAccountRequest) (*sftypes.ListVolumeStatsByAccountResult, error) {
	response, err := c.SendRequest("ListVolumeStatsByAccount", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.ListVolumeStatsByAccountResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) ListVolumeStatsByVirtualVolume(request sftypes.ListVolumeStatsByVirtualVolumeRequest) (*sftypes.ListVolumeStatsByVirtualVolumeResult, error) {
	response, err := c.SendRequest("ListVolumeStatsByVirtualVolume", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.ListVolumeStatsByVirtualVolumeResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) ListVolumeStatsByVolume(request sftypes.ListVolumeStatsByVolumeRequest) (*sftypes.ListVolumeStatsByVolumeResult, error) {
	response, err := c.SendRequest("ListVolumeStatsByVolume", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.ListVolumeStatsByVolumeResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) ListVolumeStatsByVolumeAccessGroup(request sftypes.ListVolumeStatsByVolumeAccessGroupRequest) (*sftypes.ListVolumeStatsByVolumeAccessGroupResult, error) {
	response, err := c.SendRequest("ListVolumeStatsByVolumeAccessGroup", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.ListVolumeStatsByVolumeAccessGroupResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) ListVolumes(request sftypes.ListVolumesRequest) (*sftypes.ListVolumesResult, error) {
	response, err := c.SendRequest("ListVolumes", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.ListVolumesResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) ListVolumesForAccount(request sftypes.ListVolumesForAccountRequest) (*sftypes.ListVolumesForAccountResult, error) {
	response, err := c.SendRequest("ListVolumesForAccount", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.ListVolumesForAccountResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) ModifyAccount(request sftypes.ModifyAccountRequest) (*sftypes.ModifyAccountResult, error) {
	// Adaptor
	result, err := c.ModifyAccountAdaptor(request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil
}

func (c *Client) ModifyBackupTarget(request sftypes.ModifyBackupTargetRequest) error {
	response, err := c.SendRequest("ModifyBackupTarget", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return err
	}
	if response != nil {
		log.Debugf("ModifyBackupTarget: ", response)
	}
	return err
}

func (c *Client) ModifyClusterAdmin(request sftypes.ModifyClusterAdminRequest) error {
	response, err := c.SendRequest("ModifyClusterAdmin", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return err
	}
	if response != nil {
		log.Debugf("ModifyClusterAdmin: ", response)
	}
	return err
}

func (c *Client) ModifyClusterFullThreshold(request sftypes.ModifyClusterFullThresholdRequest) (*sftypes.ModifyClusterFullThresholdResult, error) {
	response, err := c.SendRequest("ModifyClusterFullThreshold", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.ModifyClusterFullThresholdResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) ModifyGroupSnapshot(request sftypes.ModifyGroupSnapshotRequest) (*sftypes.ModifyGroupSnapshotResult, error) {
	response, err := c.SendRequest("ModifyGroupSnapshot", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.ModifyGroupSnapshotResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) ModifyInitiators(request sftypes.ModifyInitiatorsRequest) (*sftypes.ModifyInitiatorsResult, error) {
	response, err := c.SendRequest("ModifyInitiators", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.ModifyInitiatorsResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) ModifySchedule(request sftypes.ModifyScheduleRequest) (*sftypes.ModifyScheduleResult, error) {
	// Adaptor
	result, err := c.ModifyScheduleAdaptor(request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil
}

func (c *Client) ModifySnapshot(request sftypes.ModifySnapshotRequest) (*sftypes.ModifySnapshotResult, error) {
	response, err := c.SendRequest("ModifySnapshot", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.ModifySnapshotResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) ModifyStorageContainer(request sftypes.ModifyStorageContainerRequest) (*sftypes.ModifyStorageContainerResult, error) {
	response, err := c.SendRequest("ModifyStorageContainer", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.ModifyStorageContainerResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) ModifyVirtualNetwork(request sftypes.ModifyVirtualNetworkRequest) (*sftypes.AddVirtualNetworkResult, error) {
	response, err := c.SendRequest("ModifyVirtualNetwork", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.AddVirtualNetworkResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) ModifyVolume(request sftypes.ModifyVolumeRequest) (*sftypes.ModifyVolumeResult, error) {
	response, err := c.SendRequest("ModifyVolume", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.ModifyVolumeResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) ModifyVolumeAccessGroup(request sftypes.ModifyVolumeAccessGroupRequest) (*sftypes.ModifyVolumeAccessGroupResult, error) {
	response, err := c.SendRequest("ModifyVolumeAccessGroup", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.ModifyVolumeAccessGroupResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) ModifyVolumeAccessGroupLunAssignments(request sftypes.ModifyVolumeAccessGroupLunAssignmentsRequest) (*sftypes.ModifyVolumeAccessGroupLunAssignmentsResult, error) {
	response, err := c.SendRequest("ModifyVolumeAccessGroupLunAssignments", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.ModifyVolumeAccessGroupLunAssignmentsResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) ModifyVolumePair(request sftypes.ModifyVolumePairRequest) error {
	response, err := c.SendRequest("ModifyVolumePair", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return err
	}
	if response != nil {
		log.Debugf("ModifyVolumePair: ", response)
	}
	return err
}

func (c *Client) ModifyVolumes(request sftypes.ModifyVolumesRequest) (*sftypes.ModifyVolumesResult, error) {
	response, err := c.SendRequest("ModifyVolumes", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.ModifyVolumesResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) PurgeDeletedVolume(request sftypes.PurgeDeletedVolumeRequest) error {
	response, err := c.SendRequest("PurgeDeletedVolume", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return err
	}
	if response != nil {
		log.Debugf("PurgeDeletedVolume: ", response)
	}
	return err
}

func (c *Client) PurgeDeletedVolumes(request sftypes.PurgeDeletedVolumesRequest) error {
	response, err := c.SendRequest("PurgeDeletedVolumes", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return err
	}
	if response != nil {
		log.Debugf("PurgeDeletedVolumes: ", response)
	}
	return err
}

func (c *Client) RemoveAccount(request sftypes.RemoveAccountRequest) error {
	response, err := c.SendRequest("RemoveAccount", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return err
	}
	if response != nil {
		log.Debugf("RemoveAccount: ", response)
	}
	return err
}

func (c *Client) RemoveBackupTarget(request sftypes.RemoveBackupTargetRequest) error {
	response, err := c.SendRequest("RemoveBackupTarget", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return err
	}
	if response != nil {
		log.Debugf("RemoveBackupTarget: ", response)
	}
	return err
}

func (c *Client) RemoveClusterAdmin(request sftypes.RemoveClusterAdminRequest) error {
	response, err := c.SendRequest("RemoveClusterAdmin", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return err
	}
	if response != nil {
		log.Debugf("RemoveClusterAdmin: ", response)
	}
	return err
}

func (c *Client) RemoveClusterPair(request sftypes.RemoveClusterPairRequest) error {
	response, err := c.SendRequest("RemoveClusterPair", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return err
	}
	if response != nil {
		log.Debugf("RemoveClusterPair: ", response)
	}
	return err
}

func (c *Client) RemoveDrives(request sftypes.RemoveDrivesRequest) (*sftypes.AsyncHandleResult, error) {
	response, err := c.SendRequest("RemoveDrives", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.AsyncHandleResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) RemoveInitiatorsFromVolumeAccessGroup(request sftypes.RemoveInitiatorsFromVolumeAccessGroupRequest) (*sftypes.ModifyVolumeAccessGroupResult, error) {
	response, err := c.SendRequest("RemoveInitiatorsFromVolumeAccessGroup", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.ModifyVolumeAccessGroupResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) RemoveNodes(request sftypes.RemoveNodesRequest) error {
	response, err := c.SendRequest("RemoveNodes", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return err
	}
	if response != nil {
		log.Debugf("RemoveNodes: ", response)
	}
	return err
}

func (c *Client) RemoveVirtualNetwork(request sftypes.RemoveVirtualNetworkRequest) error {
	response, err := c.SendRequest("RemoveVirtualNetwork", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return err
	}
	if response != nil {
		log.Debugf("RemoveVirtualNetwork: ", response)
	}
	return err
}

func (c *Client) RemoveVolumePair(request sftypes.RemoveVolumePairRequest) error {
	response, err := c.SendRequest("RemoveVolumePair", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return err
	}
	if response != nil {
		log.Debugf("RemoveVolumePair: ", response)
	}
	return err
}

func (c *Client) RemoveVolumesFromVolumeAccessGroup(request sftypes.RemoveVolumesFromVolumeAccessGroupRequest) (*sftypes.ModifyVolumeAccessGroupResult, error) {
	response, err := c.SendRequest("RemoveVolumesFromVolumeAccessGroup", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.ModifyVolumeAccessGroupResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) ResetDrives(request sftypes.ResetDrivesRequest) (*sftypes.ResetDrivesResult, error) {
	response, err := c.SendRequest("ResetDrives", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.ResetDrivesResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) ResetNode(request sftypes.ResetNodeRequest) (*sftypes.ResetNodeResult, error) {
	response, err := c.SendRequest("ResetNode", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.ResetNodeResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) RestartNetworking(request sftypes.RestartNetworkingRequest) (interface{}, error) {
	response, err := c.SendRequest("RestartNetworking", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result interface{}
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) RestartServices(request sftypes.RestartServicesRequest) (interface{}, error) {
	response, err := c.SendRequest("RestartServices", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result interface{}
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) RestoreDeletedVolume(request sftypes.RestoreDeletedVolumeRequest) error {
	response, err := c.SendRequest("RestoreDeletedVolume", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return err
	}
	if response != nil {
		log.Debugf("RestoreDeletedVolume: ", response)
	}
	return err
}

func (c *Client) RollbackToGroupSnapshot(request sftypes.RollbackToGroupSnapshotRequest) (*sftypes.RollbackToGroupSnapshotResult, error) {
	response, err := c.SendRequest("RollbackToGroupSnapshot", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.RollbackToGroupSnapshotResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) RollbackToSnapshot(request sftypes.RollbackToSnapshotRequest) (*sftypes.RollbackToSnapshotResult, error) {
	response, err := c.SendRequest("RollbackToSnapshot", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.RollbackToSnapshotResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) SecureEraseDrives(request sftypes.SecureEraseDrivesRequest) (*sftypes.AsyncHandleResult, error) {
	response, err := c.SendRequest("SecureEraseDrives", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.AsyncHandleResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) SetClusterConfig(request sftypes.SetClusterConfigRequest) (*sftypes.SetClusterConfigResult, error) {
	response, err := c.SendRequest("SetClusterConfig", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.SetClusterConfigResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) SetConfig(request sftypes.SetConfigRequest) (*sftypes.SetConfigResult, error) {
	response, err := c.SendRequest("SetConfig", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.SetConfigResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) SetDefaultQoS(request sftypes.SetDefaultQoSRequest) (*sftypes.SetDefaultQoSResult, error) {
	response, err := c.SendRequest("SetDefaultQoS", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.SetDefaultQoSResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) SetLoginSessionInfo(request sftypes.SetLoginSessionInfoRequest) error {
	response, err := c.SendRequest("SetLoginSessionInfo", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return err
	}
	if response != nil {
		log.Debugf("SetLoginSessionInfo: ", response)
	}
	return err
}

func (c *Client) SetNetworkConfig(request sftypes.SetNetworkConfigRequest) (*sftypes.SetNetworkConfigResult, error) {
	response, err := c.SendRequest("SetNetworkConfig", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.SetNetworkConfigResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) SetNtpInfo(request sftypes.SetNtpInfoRequest) error {
	response, err := c.SendRequest("SetNtpInfo", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return err
	}
	if response != nil {
		log.Debugf("SetNtpInfo: ", response)
	}
	return err
}

func (c *Client) SetRemoteLoggingHosts(request sftypes.SetRemoteLoggingHostsRequest) error {
	response, err := c.SendRequest("SetRemoteLoggingHosts", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return err
	}
	if response != nil {
		log.Debugf("SetRemoteLoggingHosts: ", response)
	}
	return err
}

func (c *Client) SetSnmpACL(request sftypes.SetSnmpACLRequest) error {
	response, err := c.SendRequest("SetSnmpACL", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return err
	}
	if response != nil {
		log.Debugf("SetSnmpACL: ", response)
	}
	return err
}

func (c *Client) SetSnmpInfo(request sftypes.SetSnmpInfoRequest) error {
	response, err := c.SendRequest("SetSnmpInfo", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return err
	}
	if response != nil {
		log.Debugf("SetSnmpInfo: ", response)
	}
	return err
}

func (c *Client) SetSnmpTrapInfo(request sftypes.SetSnmpTrapInfoRequest) error {
	response, err := c.SendRequest("SetSnmpTrapInfo", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return err
	}
	if response != nil {
		log.Debugf("SetSnmpTrapInfo: ", response)
	}
	return err
}

func (c *Client) Shutdown(request sftypes.ShutdownRequest) (*sftypes.ShutdownResult, error) {
	response, err := c.SendRequest("Shutdown", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.ShutdownResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) SnmpSendTestTraps() (*sftypes.SnmpSendTestTrapsResult, error) {
	response, err := c.SendRequest("SnmpSendTestTraps", nil)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.SnmpSendTestTrapsResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) StartBulkVolumeRead(request sftypes.StartBulkVolumeReadRequest) (*sftypes.StartBulkVolumeReadResult, error) {
	response, err := c.SendRequest("StartBulkVolumeRead", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.StartBulkVolumeReadResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) StartBulkVolumeWrite(request sftypes.StartBulkVolumeWriteRequest) (*sftypes.StartBulkVolumeWriteResult, error) {
	response, err := c.SendRequest("StartBulkVolumeWrite", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.StartBulkVolumeWriteResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) StartClusterPairing() (*sftypes.StartClusterPairingResult, error) {
	response, err := c.SendRequest("StartClusterPairing", nil)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.StartClusterPairingResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) StartVolumePairing(request sftypes.StartVolumePairingRequest) (*sftypes.StartVolumePairingResult, error) {
	response, err := c.SendRequest("StartVolumePairing", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.StartVolumePairingResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) TestConnectEnsemble(request sftypes.TestConnectEnsembleRequest) (*sftypes.TestConnectEnsembleResult, error) {
	response, err := c.SendRequest("TestConnectEnsemble", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.TestConnectEnsembleResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) TestConnectMvip(request sftypes.TestConnectMvipRequest) (*sftypes.TestConnectMvipResult, error) {
	response, err := c.SendRequest("TestConnectMvip", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.TestConnectMvipResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) TestConnectSvip(request sftypes.TestConnectSvipRequest) (*sftypes.TestConnectSvipResult, error) {
	response, err := c.SendRequest("TestConnectSvip", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.TestConnectSvipResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) TestDrives(request sftypes.TestDrivesRequest) (*sftypes.TestDrivesResult, error) {
	response, err := c.SendRequest("TestDrives", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.TestDrivesResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) TestLdapAuthentication(request sftypes.TestLdapAuthenticationRequest) (*sftypes.TestLdapAuthenticationResult, error) {
	response, err := c.SendRequest("TestLdapAuthentication", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.TestLdapAuthenticationResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) TestPing(request sftypes.TestPingRequest) (*sftypes.TestPingResult, error) {
	response, err := c.SendRequest("TestPing", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.TestPingResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}

func (c *Client) UpdateBulkVolumeStatus(request sftypes.UpdateBulkVolumeStatusRequest) (*sftypes.UpdateBulkVolumeStatusResult, error) {
	response, err := c.SendRequest("UpdateBulkVolumeStatus", request)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	var result sftypes.UpdateBulkVolumeStatusResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}
	return &result, nil

}
