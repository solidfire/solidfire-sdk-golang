package sfapi

import (
	"github.com/solidfire/solidfire-sdk-golang/sftypes"
	"log"
)

func (c *Client) AddAccount(request sftypes.AddAccountRequest) (sftypes.AddAccountResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("AddAccount", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.AddAccountResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) AddClusterAdmin(request sftypes.AddClusterAdminRequest) (sftypes.AddClusterAdminResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("AddClusterAdmin", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.AddClusterAdminResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) AddDrives(request sftypes.AddDrivesRequest) (sftypes.AddDrivesResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("AddDrives", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.AddDrivesResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) AddInitiatorsToVolumeAccessGroup(request sftypes.AddInitiatorsToVolumeAccessGroupRequest) (sftypes.ModifyVolumeAccessGroupResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("AddInitiatorsToVolumeAccessGroup", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.ModifyVolumeAccessGroupResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) AddLdapClusterAdmin(request sftypes.AddLdapClusterAdminRequest) (sftypes.AddLdapClusterAdminResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("AddLdapClusterAdmin", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.AddLdapClusterAdminResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) AddNodes(request sftypes.AddNodesRequest) (sftypes.AddNodesResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("AddNodes", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.AddNodesResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) AddVirtualNetwork(request sftypes.AddVirtualNetworkRequest) (sftypes.AddVirtualNetworkResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("AddVirtualNetwork", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.AddVirtualNetworkResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) AddVolumesToVolumeAccessGroup(request sftypes.AddVolumesToVolumeAccessGroupRequest) (sftypes.ModifyVolumeAccessGroupResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("AddVolumesToVolumeAccessGroup", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.ModifyVolumeAccessGroupResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) CancelClone(request sftypes.CancelCloneRequest) (sftypes.CancelCloneResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("CancelClone", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.CancelCloneResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) CancelGroupClone(request sftypes.CancelGroupCloneRequest) (sftypes.CancelGroupCloneResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("CancelGroupClone", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.CancelGroupCloneResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) ClearClusterFaults(request sftypes.ClearClusterFaultsRequest) (sftypes.ClearClusterFaultsResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("ClearClusterFaults", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.ClearClusterFaultsResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) CloneMultipleVolumes(request sftypes.CloneMultipleVolumesRequest) (sftypes.CloneMultipleVolumesResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("CloneMultipleVolumes", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.CloneMultipleVolumesResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) CloneVolume(request sftypes.CloneVolumeRequest) (sftypes.CloneVolumeResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("CloneVolume", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.CloneVolumeResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) CompleteClusterPairing(request sftypes.CompleteClusterPairingRequest) (sftypes.CompleteClusterPairingResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("CompleteClusterPairing", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.CompleteClusterPairingResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) CompleteVolumePairing(request sftypes.CompleteVolumePairingRequest) (sftypes.CompleteVolumePairingResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("CompleteVolumePairing", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.CompleteVolumePairingResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) CopyVolume(request sftypes.CopyVolumeRequest) (sftypes.CopyVolumeResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("CopyVolume", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.CopyVolumeResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) CreateBackupTarget(request sftypes.CreateBackupTargetRequest) (sftypes.CreateBackupTargetResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("CreateBackupTarget", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.CreateBackupTargetResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) CreateCluster(request sftypes.CreateClusterRequest) (sftypes.CreateClusterResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("CreateCluster", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.CreateClusterResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) CreateGroupSnapshot(request sftypes.CreateGroupSnapshotRequest) (sftypes.CreateGroupSnapshotResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("CreateGroupSnapshot", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.CreateGroupSnapshotResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) CreateInitiators(request sftypes.CreateInitiatorsRequest) (sftypes.CreateInitiatorsResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("CreateInitiators", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.CreateInitiatorsResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) CreateSchedule(request sftypes.CreateScheduleRequest) (sftypes.CreateScheduleResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("CreateSchedule", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.CreateScheduleResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) CreateSnapshot(request sftypes.CreateSnapshotRequest) (sftypes.CreateSnapshotResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("CreateSnapshot", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.CreateSnapshotResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) CreateStorageContainer(request sftypes.CreateStorageContainerRequest) (sftypes.CreateStorageContainerResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("CreateStorageContainer", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.CreateStorageContainerResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) CreateSupportBundle(request sftypes.CreateSupportBundleRequest) (sftypes.CreateSupportBundleResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("CreateSupportBundle", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.CreateSupportBundleResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) CreateVolume(request sftypes.CreateVolumeRequest) (sftypes.CreateVolumeResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("CreateVolume", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.CreateVolumeResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) CreateVolumeAccessGroup(request sftypes.CreateVolumeAccessGroupRequest) (sftypes.CreateVolumeAccessGroupResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("CreateVolumeAccessGroup", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.CreateVolumeAccessGroupResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) DeleteAllSupportBundles() (sftypes.DeleteAllSupportBundlesResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("DeleteAllSupportBundles", nil)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.DeleteAllSupportBundlesResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) DeleteGroupSnapshot(request sftypes.DeleteGroupSnapshotRequest) (sftypes.DeleteGroupSnapshotResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("DeleteGroupSnapshot", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.DeleteGroupSnapshotResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) DeleteInitiators(request sftypes.DeleteInitiatorsRequest) (sftypes.DeleteInitiatorsResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("DeleteInitiators", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.DeleteInitiatorsResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) DeleteSnapshot(request sftypes.DeleteSnapshotRequest) (sftypes.DeleteSnapshotResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("DeleteSnapshot", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.DeleteSnapshotResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) DeleteStorageContainers(request sftypes.DeleteStorageContainersRequest) (sftypes.DeleteStorageContainerResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("DeleteStorageContainers", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.DeleteStorageContainerResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) DeleteVolume(request sftypes.DeleteVolumeRequest) (sftypes.DeleteVolumeResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("DeleteVolume", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.DeleteVolumeResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) DeleteVolumeAccessGroup(request sftypes.DeleteVolumeAccessGroupRequest) (sftypes.DeleteVolumeAccessGroupResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("DeleteVolumeAccessGroup", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.DeleteVolumeAccessGroupResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) DeleteVolumes(request sftypes.DeleteVolumesRequest) (sftypes.DeleteVolumesResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("DeleteVolumes", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.DeleteVolumesResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) DisableEncryptionAtRest() (sftypes.DisableEncryptionAtRestResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("DisableEncryptionAtRest", nil)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.DisableEncryptionAtRestResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) DisableLdapAuthentication() (sftypes.DisableLdapAuthenticationResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("DisableLdapAuthentication", nil)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.DisableLdapAuthenticationResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) DisableSnmp() (sftypes.DisableSnmpResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("DisableSnmp", nil)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.DisableSnmpResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) EnableEncryptionAtRest() (sftypes.EnableEncryptionAtRestResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("EnableEncryptionAtRest", nil)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.EnableEncryptionAtRestResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) EnableFeature(request sftypes.EnableFeatureRequest) (sftypes.EnableFeatureResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("EnableFeature", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.EnableFeatureResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) EnableLdapAuthentication(request sftypes.EnableLdapAuthenticationRequest) (sftypes.EnableLdapAuthenticationResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("EnableLdapAuthentication", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.EnableLdapAuthenticationResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) EnableSnmp(request sftypes.EnableSnmpRequest) (sftypes.EnableSnmpResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("EnableSnmp", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.EnableSnmpResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) GetAPI() (sftypes.GetAPIResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("GetAPI", nil)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.GetAPIResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) GetAccountByID(request sftypes.GetAccountByIDRequest) (sftypes.GetAccountResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("GetAccountByID", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.GetAccountResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) GetAccountByName(request sftypes.GetAccountByNameRequest) (sftypes.GetAccountResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("GetAccountByName", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.GetAccountResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) GetAccountEfficiency(request sftypes.GetAccountEfficiencyRequest) (sftypes.GetEfficiencyResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("GetAccountEfficiency", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.GetEfficiencyResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) GetAsyncResult(request sftypes.GetAsyncResultRequest) (interface{}, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("GetAsyncResult", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result interface{}
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) GetBackupTarget(request sftypes.GetBackupTargetRequest) (sftypes.GetBackupTargetResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("GetBackupTarget", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.GetBackupTargetResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) GetBootstrapConfig() (sftypes.GetBootstrapConfigResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("GetBootstrapConfig", nil)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.GetBootstrapConfigResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) GetClusterCapacity() (sftypes.GetClusterCapacityResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("GetClusterCapacity", nil)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.GetClusterCapacityResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) GetClusterConfig() (sftypes.GetClusterConfigResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("GetClusterConfig", nil)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.GetClusterConfigResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) GetClusterFullThreshold() (sftypes.GetClusterFullThresholdResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("GetClusterFullThreshold", nil)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.GetClusterFullThresholdResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) GetClusterHardwareInfo(request sftypes.GetClusterHardwareInfoRequest) (sftypes.GetClusterHardwareInfoResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("GetClusterHardwareInfo", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.GetClusterHardwareInfoResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) GetClusterInfo() (sftypes.GetClusterInfoResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("GetClusterInfo", nil)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.GetClusterInfoResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) GetClusterMasterNodeID() (sftypes.GetClusterMasterNodeIDResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("GetClusterMasterNodeID", nil)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.GetClusterMasterNodeIDResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) GetClusterState(request sftypes.GetClusterStateRequest) (sftypes.GetClusterStateResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("GetClusterState", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.GetClusterStateResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) GetClusterStats() (sftypes.GetClusterStatsResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("GetClusterStats", nil)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.GetClusterStatsResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) GetClusterVersionInfo() (sftypes.GetClusterVersionInfoResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("GetClusterVersionInfo", nil)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.GetClusterVersionInfoResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) GetCompleteStats() (interface{}, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("GetCompleteStats", nil)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result interface{}
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) GetConfig() (sftypes.GetConfigResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("GetConfig", nil)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.GetConfigResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) GetCurrentClusterAdmin() (sftypes.GetCurrentClusterAdminResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("GetCurrentClusterAdmin", nil)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.GetCurrentClusterAdminResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) GetDefaultQoS() (sftypes.VolumeQOS, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("GetDefaultQoS", nil)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.VolumeQOS
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) GetDriveConfig() (sftypes.GetDriveConfigResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("GetDriveConfig", nil)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.GetDriveConfigResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) GetDriveHardwareInfo(request sftypes.GetDriveHardwareInfoRequest) (sftypes.GetDriveHardwareInfoResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("GetDriveHardwareInfo", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.GetDriveHardwareInfoResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) GetDriveStats(request sftypes.GetDriveStatsRequest) (sftypes.GetDriveStatsResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("GetDriveStats", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.GetDriveStatsResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) GetFeatureStatus(request sftypes.GetFeatureStatusRequest) (sftypes.GetFeatureStatusResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("GetFeatureStatus", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.GetFeatureStatusResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) GetHardwareConfig() (sftypes.GetHardwareConfigResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("GetHardwareConfig", nil)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.GetHardwareConfigResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) GetHardwareInfo() (sftypes.GetHardwareInfoResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("GetHardwareInfo", nil)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.GetHardwareInfoResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) GetIpmiConfig(request sftypes.GetIpmiConfigRequest) (sftypes.GetIpmiConfigResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("GetIpmiConfig", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.GetIpmiConfigResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) GetIpmiInfo(request sftypes.GetIpmiInfoRequest) (sftypes.GetIpmiInfoResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("GetIpmiInfo", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.GetIpmiInfoResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) GetLdapConfiguration() (sftypes.GetLdapConfigurationResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("GetLdapConfiguration", nil)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.GetLdapConfigurationResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) GetLimits() (sftypes.GetLimitsResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("GetLimits", nil)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.GetLimitsResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) GetLoginSessionInfo() (sftypes.GetLoginSessionInfoResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("GetLoginSessionInfo", nil)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.GetLoginSessionInfoResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) GetNetworkConfig() (sftypes.GetNetworkConfigResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("GetNetworkConfig", nil)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.GetNetworkConfigResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) GetNodeHardwareInfo(request sftypes.GetNodeHardwareInfoRequest) (sftypes.GetNodeHardwareInfoResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("GetNodeHardwareInfo", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.GetNodeHardwareInfoResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) GetNodeStats(request sftypes.GetNodeStatsRequest) (sftypes.GetNodeStatsResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("GetNodeStats", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.GetNodeStatsResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) GetNtpInfo() (sftypes.GetNtpInfoResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("GetNtpInfo", nil)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.GetNtpInfoResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) GetNvramInfo() (sftypes.GetNvramInfoResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("GetNvramInfo", nil)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.GetNvramInfoResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) GetOrigin(request sftypes.GetOriginRequest) (sftypes.GetOriginResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("GetOrigin", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.GetOriginResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) GetPendingOperation() (sftypes.GetPendingOperationResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("GetPendingOperation", nil)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.GetPendingOperationResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) GetRawStats() (interface{}, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("GetRawStats", nil)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result interface{}
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) GetRemoteLoggingHosts() (sftypes.GetRemoteLoggingHostsResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("GetRemoteLoggingHosts", nil)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.GetRemoteLoggingHostsResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) GetSchedule(request sftypes.GetScheduleRequest) (sftypes.GetScheduleResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("GetSchedule", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.GetScheduleResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) GetSnmpACL() (sftypes.GetSnmpACLResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("GetSnmpACL", nil)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.GetSnmpACLResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) GetSnmpInfo() (sftypes.GetSnmpInfoResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("GetSnmpInfo", nil)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.GetSnmpInfoResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) GetSnmpState() (sftypes.GetSnmpStateResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("GetSnmpState", nil)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.GetSnmpStateResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) GetSnmpTrapInfo(request sftypes.GetSnmpTrapInfoRequest) (sftypes.GetSnmpTrapInfoResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("GetSnmpTrapInfo", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.GetSnmpTrapInfoResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) GetStorageContainerEfficiency(request sftypes.GetStorageContainerEfficiencyRequest) (sftypes.GetStorageContainerEfficiencyResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("GetStorageContainerEfficiency", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.GetStorageContainerEfficiencyResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) GetSystemStatus() (sftypes.GetSystemStatusResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("GetSystemStatus", nil)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.GetSystemStatusResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) GetVirtualVolumeCount() (sftypes.GetVirtualVolumeCountResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("GetVirtualVolumeCount", nil)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.GetVirtualVolumeCountResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) GetVolumeAccessGroupEfficiency(request sftypes.GetVolumeAccessGroupEfficiencyRequest) (sftypes.GetEfficiencyResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("GetVolumeAccessGroupEfficiency", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.GetEfficiencyResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) GetVolumeAccessGroupLunAssignments(request sftypes.GetVolumeAccessGroupLunAssignmentsRequest) (sftypes.GetVolumeAccessGroupLunAssignmentsResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("GetVolumeAccessGroupLunAssignments", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.GetVolumeAccessGroupLunAssignmentsResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) GetVolumeCount() (sftypes.GetVolumeCountResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("GetVolumeCount", nil)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.GetVolumeCountResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) GetVolumeEfficiency(request sftypes.GetVolumeEfficiencyRequest) (sftypes.GetVolumeEfficiencyResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("GetVolumeEfficiency", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.GetVolumeEfficiencyResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) GetVolumeStats(request sftypes.GetVolumeStatsRequest) (sftypes.GetVolumeStatsResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("GetVolumeStats", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.GetVolumeStatsResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) InvokeSFApi(request sftypes.InvokeSFApiRequest) (interface{}, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("InvokeSFApi", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result interface{}
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) ListAccounts(request sftypes.ListAccountsRequest) (sftypes.ListAccountsResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("ListAccounts", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.ListAccountsResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) ListActiveNodes() (sftypes.ListActiveNodesResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("ListActiveNodes", nil)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.ListActiveNodesResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) ListActivePairedVolumes(request sftypes.ListActivePairedVolumesRequest) (sftypes.ListActivePairedVolumesResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("ListActivePairedVolumes", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.ListActivePairedVolumesResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) ListActiveVolumes(request sftypes.ListActiveVolumesRequest) (sftypes.ListActiveVolumesResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("ListActiveVolumes", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.ListActiveVolumesResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) ListAllNodes() (sftypes.ListAllNodesResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("ListAllNodes", nil)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.ListAllNodesResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) ListAsyncResults(request sftypes.ListAsyncResultsRequest) (sftypes.ListAsyncResultsResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("ListAsyncResults", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.ListAsyncResultsResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) ListBackupTargets() (sftypes.ListBackupTargetsResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("ListBackupTargets", nil)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.ListBackupTargetsResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) ListBulkVolumeJobs() (sftypes.ListBulkVolumeJobsResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("ListBulkVolumeJobs", nil)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.ListBulkVolumeJobsResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) ListClusterAdmins(request sftypes.ListClusterAdminsRequest) (sftypes.ListClusterAdminsResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("ListClusterAdmins", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.ListClusterAdminsResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) ListClusterFaults(request sftypes.ListClusterFaultsRequest) (sftypes.ListClusterFaultsResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("ListClusterFaults", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.ListClusterFaultsResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) ListClusterPairs() (sftypes.ListClusterPairsResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("ListClusterPairs", nil)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.ListClusterPairsResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) ListDeletedVolumes() (sftypes.ListDeletedVolumesResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("ListDeletedVolumes", nil)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.ListDeletedVolumesResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) ListDriveHardware(request sftypes.ListDriveHardwareRequest) (sftypes.ListDriveHardwareResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("ListDriveHardware", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.ListDriveHardwareResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) ListDriveStats(request sftypes.ListDriveStatsRequest) (sftypes.ListDriveStatsResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("ListDriveStats", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.ListDriveStatsResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) ListDrives() (sftypes.ListDrivesResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("ListDrives", nil)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.ListDrivesResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) ListEvents(request sftypes.ListEventsRequest) (sftypes.ListEventsResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("ListEvents", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.ListEventsResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) ListFibreChannelPortInfo() (sftypes.ListFibreChannelPortInfoResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("ListFibreChannelPortInfo", nil)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.ListFibreChannelPortInfoResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) ListFibreChannelSessions() (sftypes.ListFibreChannelSessionsResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("ListFibreChannelSessions", nil)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.ListFibreChannelSessionsResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) ListGroupSnapshots(request sftypes.ListGroupSnapshotsRequest) (sftypes.ListGroupSnapshotsResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("ListGroupSnapshots", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.ListGroupSnapshotsResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) ListISCSISessions() (sftypes.ListISCSISessionsResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("ListISCSISessions", nil)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.ListISCSISessionsResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) ListInitiators(request sftypes.ListInitiatorsRequest) (sftypes.ListInitiatorsResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("ListInitiators", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.ListInitiatorsResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) ListNetworkInterfaces() (sftypes.ListNetworkInterfacesResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("ListNetworkInterfaces", nil)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.ListNetworkInterfacesResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) ListNodeFibreChannelPortInfo() (sftypes.ListNodeFibreChannelPortInfoResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("ListNodeFibreChannelPortInfo", nil)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.ListNodeFibreChannelPortInfoResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) ListNodeStats() (sftypes.ListNodeStatsResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("ListNodeStats", nil)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.ListNodeStatsResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) ListPendingActiveNodes() (sftypes.ListPendingActiveNodesResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("ListPendingActiveNodes", nil)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.ListPendingActiveNodesResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) ListPendingNodes() (sftypes.ListPendingNodesResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("ListPendingNodes", nil)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.ListPendingNodesResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) ListProtocolEndpoints(request sftypes.ListProtocolEndpointsRequest) (sftypes.ListProtocolEndpointsResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("ListProtocolEndpoints", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.ListProtocolEndpointsResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) ListSchedules() (sftypes.ListSchedulesResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("ListSchedules", nil)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.ListSchedulesResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) ListServices() (sftypes.ListServicesResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("ListServices", nil)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.ListServicesResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) ListSnapshots(request sftypes.ListSnapshotsRequest) (sftypes.ListSnapshotsResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("ListSnapshots", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.ListSnapshotsResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) ListStorageContainers(request sftypes.ListStorageContainersRequest) (sftypes.ListStorageContainersResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("ListStorageContainers", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.ListStorageContainersResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) ListSyncJobs() (sftypes.ListSyncJobsResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("ListSyncJobs", nil)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.ListSyncJobsResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) ListTests() (sftypes.ListTestsResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("ListTests", nil)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.ListTestsResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) ListUtilities() (sftypes.ListUtilitiesResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("ListUtilities", nil)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.ListUtilitiesResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) ListVirtualNetworks(request sftypes.ListVirtualNetworksRequest) (sftypes.ListVirtualNetworksResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("ListVirtualNetworks", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.ListVirtualNetworksResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) ListVirtualVolumeBindings(request sftypes.ListVirtualVolumeBindingsRequest) (sftypes.ListVirtualVolumeBindingsResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("ListVirtualVolumeBindings", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.ListVirtualVolumeBindingsResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) ListVirtualVolumeHosts(request sftypes.ListVirtualVolumeHostsRequest) (sftypes.ListVirtualVolumeHostsResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("ListVirtualVolumeHosts", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.ListVirtualVolumeHostsResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) ListVirtualVolumeTasks(request sftypes.ListVirtualVolumeTasksRequest) (sftypes.ListVirtualVolumeTasksResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("ListVirtualVolumeTasks", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.ListVirtualVolumeTasksResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) ListVirtualVolumes(request sftypes.ListVirtualVolumesRequest) (sftypes.ListVirtualVolumesResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("ListVirtualVolumes", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.ListVirtualVolumesResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) ListVolumeAccessGroups(request sftypes.ListVolumeAccessGroupsRequest) (sftypes.ListVolumeAccessGroupsResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("ListVolumeAccessGroups", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.ListVolumeAccessGroupsResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) ListVolumeStats(request sftypes.ListVolumeStatsRequest) (sftypes.ListVolumeStatsResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("ListVolumeStats", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.ListVolumeStatsResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) ListVolumeStatsByAccount() (sftypes.ListVolumeStatsByAccountResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("ListVolumeStatsByAccount", nil)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.ListVolumeStatsByAccountResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) ListVolumeStatsByVirtualVolume(request sftypes.ListVolumeStatsByVirtualVolumeRequest) (sftypes.ListVolumeStatsByVirtualVolumeResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("ListVolumeStatsByVirtualVolume", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.ListVolumeStatsByVirtualVolumeResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) ListVolumeStatsByVolume() (sftypes.ListVolumeStatsByVolumeResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("ListVolumeStatsByVolume", nil)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.ListVolumeStatsByVolumeResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) ListVolumeStatsByVolumeAccessGroup(request sftypes.ListVolumeStatsByVolumeAccessGroupRequest) (sftypes.ListVolumeStatsByVolumeAccessGroupResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("ListVolumeStatsByVolumeAccessGroup", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.ListVolumeStatsByVolumeAccessGroupResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) ListVolumes(request sftypes.ListVolumesRequest) (sftypes.ListVolumesResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("ListVolumes", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.ListVolumesResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) ListVolumesForAccount(request sftypes.ListVolumesForAccountRequest) (sftypes.ListVolumesForAccountResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("ListVolumesForAccount", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.ListVolumesForAccountResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) ModifyAccount(request sftypes.ModifyAccountRequest) (sftypes.ModifyAccountResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("ModifyAccount", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.ModifyAccountResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) ModifyBackupTarget(request sftypes.ModifyBackupTargetRequest) (sftypes.ModifyBackupTargetResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("ModifyBackupTarget", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.ModifyBackupTargetResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) ModifyClusterAdmin(request sftypes.ModifyClusterAdminRequest) (sftypes.ModifyClusterAdminResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("ModifyClusterAdmin", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.ModifyClusterAdminResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) ModifyClusterFullThreshold(request sftypes.ModifyClusterFullThresholdRequest) (sftypes.ModifyClusterFullThresholdResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("ModifyClusterFullThreshold", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.ModifyClusterFullThresholdResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) ModifyGroupSnapshot(request sftypes.ModifyGroupSnapshotRequest) (sftypes.ModifyGroupSnapshotResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("ModifyGroupSnapshot", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.ModifyGroupSnapshotResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) ModifyInitiators(request sftypes.ModifyInitiatorsRequest) (sftypes.ModifyInitiatorsResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("ModifyInitiators", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.ModifyInitiatorsResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) ModifySchedule(request sftypes.ModifyScheduleRequest) (sftypes.ModifyScheduleResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("ModifySchedule", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.ModifyScheduleResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) ModifySnapshot(request sftypes.ModifySnapshotRequest) (sftypes.ModifySnapshotResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("ModifySnapshot", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.ModifySnapshotResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) ModifyStorageContainer(request sftypes.ModifyStorageContainerRequest) (sftypes.ModifyStorageContainerResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("ModifyStorageContainer", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.ModifyStorageContainerResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) ModifyVirtualNetwork(request sftypes.ModifyVirtualNetworkRequest) (sftypes.AddVirtualNetworkResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("ModifyVirtualNetwork", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.AddVirtualNetworkResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) ModifyVolume(request sftypes.ModifyVolumeRequest) (sftypes.ModifyVolumeResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("ModifyVolume", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.ModifyVolumeResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) ModifyVolumeAccessGroup(request sftypes.ModifyVolumeAccessGroupRequest) (sftypes.ModifyVolumeAccessGroupResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("ModifyVolumeAccessGroup", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.ModifyVolumeAccessGroupResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) ModifyVolumeAccessGroupLunAssignments(request sftypes.ModifyVolumeAccessGroupLunAssignmentsRequest) (sftypes.ModifyVolumeAccessGroupLunAssignmentsResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("ModifyVolumeAccessGroupLunAssignments", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.ModifyVolumeAccessGroupLunAssignmentsResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) ModifyVolumePair(request sftypes.ModifyVolumePairRequest) (sftypes.ModifyVolumePairResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("ModifyVolumePair", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.ModifyVolumePairResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) ModifyVolumes(request sftypes.ModifyVolumesRequest) (sftypes.ModifyVolumesResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("ModifyVolumes", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.ModifyVolumesResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) PurgeDeletedVolume(request sftypes.PurgeDeletedVolumeRequest) (sftypes.PurgeDeletedVolumeResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("PurgeDeletedVolume", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.PurgeDeletedVolumeResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) PurgeDeletedVolumes(request sftypes.PurgeDeletedVolumesRequest) (sftypes.PurgeDeletedVolumesResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("PurgeDeletedVolumes", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.PurgeDeletedVolumesResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) RemoveAccount(request sftypes.RemoveAccountRequest) (sftypes.RemoveAccountResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("RemoveAccount", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.RemoveAccountResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) RemoveBackupTarget(request sftypes.RemoveBackupTargetRequest) (sftypes.RemoveBackupTargetResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("RemoveBackupTarget", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.RemoveBackupTargetResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) RemoveClusterAdmin(request sftypes.RemoveClusterAdminRequest) (sftypes.RemoveClusterAdminResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("RemoveClusterAdmin", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.RemoveClusterAdminResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) RemoveClusterPair(request sftypes.RemoveClusterPairRequest) (sftypes.RemoveClusterPairResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("RemoveClusterPair", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.RemoveClusterPairResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) RemoveDrives(request sftypes.RemoveDrivesRequest) (sftypes.AsyncHandleResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("RemoveDrives", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.AsyncHandleResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) RemoveInitiatorsFromVolumeAccessGroup(request sftypes.RemoveInitiatorsFromVolumeAccessGroupRequest) (sftypes.ModifyVolumeAccessGroupResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("RemoveInitiatorsFromVolumeAccessGroup", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.ModifyVolumeAccessGroupResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) RemoveNodes(request sftypes.RemoveNodesRequest) (sftypes.RemoveNodesResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("RemoveNodes", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.RemoveNodesResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) RemoveVirtualNetwork(request sftypes.RemoveVirtualNetworkRequest) (sftypes.RemoveVirtualNetworkResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("RemoveVirtualNetwork", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.RemoveVirtualNetworkResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) RemoveVolumePair(request sftypes.RemoveVolumePairRequest) (sftypes.RemoveVolumePairResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("RemoveVolumePair", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.RemoveVolumePairResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) RemoveVolumesFromVolumeAccessGroup(request sftypes.RemoveVolumesFromVolumeAccessGroupRequest) (sftypes.ModifyVolumeAccessGroupResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("RemoveVolumesFromVolumeAccessGroup", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.ModifyVolumeAccessGroupResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) ResetDrives(request sftypes.ResetDrivesRequest) (sftypes.ResetDrivesResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("ResetDrives", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.ResetDrivesResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) ResetNode(request sftypes.ResetNodeRequest) (sftypes.ResetNodeResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("ResetNode", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.ResetNodeResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) RestartNetworking(request sftypes.RestartNetworkingRequest) (interface{}, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("RestartNetworking", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result interface{}
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) RestartServices(request sftypes.RestartServicesRequest) (interface{}, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("RestartServices", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result interface{}
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) RestoreDeletedVolume(request sftypes.RestoreDeletedVolumeRequest) (sftypes.RestoreDeletedVolumeResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("RestoreDeletedVolume", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.RestoreDeletedVolumeResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) RollbackToGroupSnapshot(request sftypes.RollbackToGroupSnapshotRequest) (sftypes.RollbackToGroupSnapshotResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("RollbackToGroupSnapshot", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.RollbackToGroupSnapshotResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) RollbackToSnapshot(request sftypes.RollbackToSnapshotRequest) (sftypes.RollbackToSnapshotResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("RollbackToSnapshot", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.RollbackToSnapshotResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) SecureEraseDrives(request sftypes.SecureEraseDrivesRequest) (sftypes.AsyncHandleResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("SecureEraseDrives", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.AsyncHandleResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) SetClusterConfig(request sftypes.SetClusterConfigRequest) (sftypes.SetClusterConfigResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("SetClusterConfig", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.SetClusterConfigResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) SetConfig(request sftypes.SetConfigRequest) (sftypes.SetConfigResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("SetConfig", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.SetConfigResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) SetDefaultQoS(request sftypes.SetDefaultQoSRequest) (sftypes.SetDefaultQoSResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("SetDefaultQoS", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.SetDefaultQoSResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) SetLoginSessionInfo(request sftypes.SetLoginSessionInfoRequest) (sftypes.SetLoginSessionInfoResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("SetLoginSessionInfo", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.SetLoginSessionInfoResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) SetNetworkConfig(request sftypes.SetNetworkConfigRequest) (sftypes.SetNetworkConfigResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("SetNetworkConfig", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.SetNetworkConfigResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) SetNtpInfo(request sftypes.SetNtpInfoRequest) (sftypes.SetNtpInfoResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("SetNtpInfo", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.SetNtpInfoResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) SetRemoteLoggingHosts(request sftypes.SetRemoteLoggingHostsRequest) (sftypes.SetRemoteLoggingHostsResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("SetRemoteLoggingHosts", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.SetRemoteLoggingHostsResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) SetSnmpACL(request sftypes.SetSnmpACLRequest) (sftypes.SetSnmpACLResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("SetSnmpACL", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.SetSnmpACLResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) SetSnmpInfo(request sftypes.SetSnmpInfoRequest) (sftypes.SetSnmpInfoResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("SetSnmpInfo", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.SetSnmpInfoResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) SetSnmpTrapInfo(request sftypes.SetSnmpTrapInfoRequest) (sftypes.SetSnmpTrapInfoResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("SetSnmpTrapInfo", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.SetSnmpTrapInfoResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) Shutdown(request sftypes.ShutdownRequest) (sftypes.ShutdownResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("Shutdown", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.ShutdownResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) SnmpSendTestTraps() (sftypes.SnmpSendTestTrapsResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("SnmpSendTestTraps", nil)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.SnmpSendTestTrapsResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) StartBulkVolumeRead(request sftypes.StartBulkVolumeReadRequest) (sftypes.StartBulkVolumeReadResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("StartBulkVolumeRead", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.StartBulkVolumeReadResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) StartBulkVolumeWrite(request sftypes.StartBulkVolumeWriteRequest) (sftypes.StartBulkVolumeWriteResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("StartBulkVolumeWrite", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.StartBulkVolumeWriteResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) StartClusterPairing() (sftypes.StartClusterPairingResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("StartClusterPairing", nil)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.StartClusterPairingResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) StartVolumePairing(request sftypes.StartVolumePairingRequest) (sftypes.StartVolumePairingResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("StartVolumePairing", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.StartVolumePairingResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) TestConnectEnsemble(request sftypes.TestConnectEnsembleRequest) (sftypes.TestConnectEnsembleResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("TestConnectEnsemble", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.TestConnectEnsembleResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) TestConnectMvip(request sftypes.TestConnectMvipRequest) (sftypes.TestConnectMvipResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("TestConnectMvip", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.TestConnectMvipResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) TestConnectSvip(request sftypes.TestConnectSvipRequest) (sftypes.TestConnectSvipResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("TestConnectSvip", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.TestConnectSvipResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) TestDrives(request sftypes.TestDrivesRequest) (sftypes.TestDrivesResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("TestDrives", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.TestDrivesResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) TestLdapAuthentication(request sftypes.TestLdapAuthenticationRequest) (sftypes.TestLdapAuthenticationResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("TestLdapAuthentication", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.TestLdapAuthenticationResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) TestPing(request sftypes.TestPingRequest) (sftypes.TestPingResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("TestPing", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.TestPingResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}

func (c *Client) UpdateBulkVolumeStatus(request sftypes.UpdateBulkVolumeStatusRequest) (sftypes.UpdateBulkVolumeStatusResult, error) {
	// Get a list of all active volumes on the Cluster
	response, err := c.SendRequest("UpdateBulkVolumeStatus", request)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	var result sftypes.UpdateBulkVolumeStatusResult
	decoder, err := GetDecoder(&result)
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	err = decoder.Decode(response["result"])
	if err != nil {
		log.Fatal("Err: %v", err)
	}
	return result, nil
}
