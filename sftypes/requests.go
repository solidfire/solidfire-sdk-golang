package sftypes

type AddAccountRequest struct {
	Username        string      `json:"username"`
	InitiatorSecret CHAPSecret  `json:"initiatorSecret,omitempty"`
	TargetSecret    CHAPSecret  `json:"targetSecret,omitempty"`
	Attributes      interface{} `json:"attributes,omitempty"`
}

type AddClusterAdminRequest struct {
	Username   string      `json:"username"`
	Password   string      `json:"password"`
	Access     []string    `json:"access"`
	AcceptEula bool        `json:"acceptEula,omitempty"`
	Attributes interface{} `json:"attributes,omitempty"`
}

type AddDrivesRequest struct {
	Drives []NewDrive `json:"drives"`
}

type AddInitiatorsToVolumeAccessGroupRequest struct {
	VolumeAccessGroupID int64    `json:"volumeAccessGroupID"`
	Initiators          []string `json:"initiators"`
}

type AddLdapClusterAdminRequest struct {
	Username   string      `json:"username"`
	Access     []string    `json:"access"`
	AcceptEula bool        `json:"acceptEula,omitempty"`
	Attributes interface{} `json:"attributes,omitempty"`
}

type AddNodesRequest struct {
	PendingNodes []int64 `json:"pendingNodes"`
}

type AddVirtualNetworkRequest struct {
	VirtualNetworkTag int64          `json:"virtualNetworkTag"`
	Name              string         `json:"name"`
	AddressBlocks     []AddressBlock `json:"addressBlocks"`
	Netmask           string         `json:"netmask"`
	Svip              string         `json:"svip"`
	Gateway           string         `json:"gateway,omitempty"`
	Namespace         bool           `json:"namespace,omitempty"`
	Attributes        interface{}    `json:"attributes,omitempty"`
}

type AddVolumesToVolumeAccessGroupRequest struct {
	VolumeAccessGroupID int64   `json:"volumeAccessGroupID"`
	Volumes             []int64 `json:"volumes"`
}

type ClearClusterFaultsRequest struct {
	FaultTypes string `json:"faultTypes,omitempty"`
}

type CloneMultipleVolumesRequest struct {
	Volumes         []CloneMultipleVolumeParams `json:"volumes"`
	Access          string                      `json:"access,omitempty"`
	GroupSnapshotID int64                       `json:"groupSnapshotID,omitempty"`
	NewAccountID    int64                       `json:"newAccountID,omitempty"`
}

type CloneVolumeRequest struct {
	VolumeID     int64       `json:"volumeID"`
	Name         string      `json:"name"`
	NewAccountID int64       `json:"newAccountID,omitempty"`
	NewSize      int64       `json:"newSize,omitempty"`
	Access       string      `json:"access,omitempty"`
	SnapshotID   int64       `json:"snapshotID,omitempty"`
	Attributes   interface{} `json:"attributes,omitempty"`
}

type CompleteClusterPairingRequest struct {
	ClusterPairingKey string `json:"clusterPairingKey"`
}

type CompleteVolumePairingRequest struct {
	VolumePairingKey string `json:"volumePairingKey"`
	VolumeID         int64  `json:"volumeID"`
}

type CreateBackupTargetRequest struct {
	Name       string      `json:"name"`
	Attributes interface{} `json:"attributes,omitempty"`
}

type CreateGroupSnapshotRequest struct {
	Volumes                 []int64     `json:"volumes"`
	Name                    string      `json:"name,omitempty"`
	EnableRemoteReplication bool        `json:"enableRemoteReplication,omitempty"`
	Retention               string      `json:"retention,omitempty"`
	Attributes              interface{} `json:"attributes,omitempty"`
}

type CreateScheduleRequest struct {
	Schedule Schedule `json:"schedule"`
}

type CreateSnapshotRequest struct {
	VolumeID                int64       `json:"volumeID"`
	SnapshotID              int64       `json:"snapshotID,omitempty"`
	Name                    string      `json:"name,omitempty"`
	EnableRemoteReplication bool        `json:"enableRemoteReplication,omitempty"`
	Retention               string      `json:"retention,omitempty"`
	Attributes              interface{} `json:"attributes,omitempty"`
}

type CreateVolumeAccessGroupRequest struct {
	Name               string      `json:"name"`
	Initiators         []string    `json:"initiators,omitempty"`
	Volumes            []int64     `json:"volumes,omitempty"`
	VirtualNetworkID   []int64     `json:"virtualNetworkID,omitempty"`
	VirtualNetworkTags []int64     `json:"virtualNetworkTags,omitempty"`
	Attributes         interface{} `json:"attributes,omitempty"`
}

type CreateVolumeRequest struct {
	Name       string      `json:"name"`
	AccountID  int64       `json:"accountID"`
	TotalSize  int64       `json:"totalSize"`
	Enable512e bool        `json:"enable512e"`
	Qos        QoS         `json:"qos,omitempty"`
	Attributes interface{} `json:"attributes,omitempty"`
}

type DeleteGroupSnapshotRequest struct {
	GroupSnapshotID int64 `json:"groupSnapshotID"`
	SaveMembers     bool  `json:"saveMembers"`
}

type DeleteSnapshotRequest struct {
	SnapshotID int64 `json:"snapshotID"`
}

type DeleteVolumeAccessGroupRequest struct {
	VolumeAccessGroupID int64 `json:"volumeAccessGroupID"`
}

type DeleteVolumeRequest struct {
	VolumeID int64 `json:"volumeID"`
}

type EnableLdapAuthenticationRequest struct {
	AuthType                string   `json:"authType,omitempty"`
	GroupSearchBaseDN       string   `json:"groupSearchBaseDN,omitempty"`
	GroupSearchCustomFilter string   `json:"groupSearchCustomFilter,omitempty"`
	GroupSearchType         string   `json:"groupSearchType,omitempty"`
	SearchBindDN            string   `json:"searchBindDN,omitempty"`
	SearchBindPassword      string   `json:"searchBindPassword,omitempty"`
	ServerURIs              []string `json:"serverURIs"`
	UserDNTemplate          string   `json:"userDNTemplate,omitempty"`
	UserSearchBaseDN        string   `json:"userSearchBaseDN,omitempty"`
	UserSearchFilter        string   `json:"userSearchFilter,omitempty"`
}

type EnableSnmpRequest struct {
	SnmpV3Enabled bool `json:"snmpV3Enabled"`
}

type GetAccountByIDRequest struct {
	AccountID int64 `json:"accountID"`
}

type GetAccountByNameRequest struct {
	Username string `json:"username"`
}

type GetAccountEfficiencyRequest struct {
	AccountID int64 `json:"accountID"`
	Force     bool  `json:"force,omitempty"`
}

type GetAsyncResultRequest struct {
	AsyncHandle int64 `json:"asyncHandle"`
}

type GetBackupTargetRequest struct {
	BackupTargetID int64 `json:"backupTargetID"`
}

type GetDriveHardwareInfoRequest struct {
	DriveID int64 `json:"driveID"`
}

type GetDriveStatsRequest struct {
	DriveID int64 `json:"driveID"`
}

type GetNodeStatsRequest struct {
	NodeID int64 `json:"nodeID"`
}

type GetScheduleRequest struct {
	ScheduleID int64 `json:"scheduleID"`
}

type GetVolumeAccessGroupEfficiencyRequest struct {
	VolumeAccessGroupID int64 `json:"volumeAccessGroupID"`
}

type GetVolumeAccessGroupLunAssignmentsRequest struct {
	VolumeAccessGroupID int64 `json:"volumeAccessGroupID"`
}

type GetVolumeEfficiencyRequest struct {
	VolumeID int64 `json:"volumeID"`
	Force    bool  `json:"force,omitempty"`
}

type GetVolumeStatsRequest struct {
	VolumeID int64 `json:"volumeID"`
}

type ListAccountsRequest struct {
	StartAccountID int64 `json:"startAccountID,omitempty"`
	Limit          int64 `json:"limit,omitempty"`
}

type ListActiveVolumesRequest struct {
	StartVolumeID int64 `json:"startVolumeID,omitempty"`
	Limit         int64 `json:"limit,omitempty"`
}

type ListClusterFaultsRequest struct {
	Exceptions    bool   `json:"exceptions,omitempty"`
	BestPractices bool   `json:"bestPractices,omitempty"`
	Update        bool   `json:"update,omitempty"`
	FaultTypes    string `json:"faultTypes,omitempty"`
}

type ListDriveHardwareRequest struct {
	Force bool `json:"force"`
}

type ListEventsRequest struct {
	MaxEvents      int64  `json:"maxEvents,omitempty"`
	StartEventID   int64  `json:"startEventID,omitempty"`
	EndEventID     int64  `json:"endEventID,omitempty"`
	EventQueueType string `json:"eventQueueType,omitempty"`
}

type ListGroupSnapshotsRequest struct {
	VolumeID int64 `json:"volumeID,omitempty"`
}

type ListNodeFibreChannelPortInfoRequest struct {
	Force bool `json:"force,omitempty"`
}

type ListSnapshotsRequest struct {
	VolumeID int64 `json:"volumeID,omitempty"`
}

type ListVirtualNetworksRequest struct {
	VirtualNetworkID   int64   `json:"virtualNetworkID,omitempty"`
	VirtualNetworkTag  int64   `json:"virtualNetworkTag,omitempty"`
	VirtualNetworkIDs  []int64 `json:"virtualNetworkIDs,omitempty"`
	VirtualNetworkTags []int64 `json:"virtualNetworkTags,omitempty"`
}

type ListVolumeAccessGroupsRequest struct {
	StartVolumeAccessGroupID int64 `json:"startVolumeAccessGroupID,omitempty"`
	Limit                    int64 `json:"limit,omitempty"`
}

type ListVolumeStatsByVolumeAccessGroupRequest struct {
	VolumeAccessGroups []int64 `json:"volumeAccessGroups,omitempty"`
}

type ListVolumesForAccountRequest struct {
	AccountID     int64 `json:"accountID"`
	StartVolumeID int64 `json:"startVolumeID,omitempty"`
	Limit         int64 `json:"limit,omitempty"`
}

type ListVolumesRequest struct {
	StartVolumeID int64   `json:"startVolumeID,omitempty"`
	Limit         int64   `json:"limit,omitempty"`
	VolumeStatus  string  `json:"volumeStatus,omitempty"`
	Accounts      []int64 `json:"accounts,omitempty"`
	IsPaired      bool    `json:"isPaired,omitempty"`
	VolumeIDs     []int64 `json:"volumeIDs,omitempty"`
}

type ModifyAccountRequest struct {
	AccountID       int64       `json:"accountID"`
	Username        string      `json:"username,omitempty"`
	Status          string      `json:"status,omitempty"`
	InitiatorSecret CHAPSecret  `json:"initiatorSecret,omitempty"`
	TargetSecret    CHAPSecret  `json:"targetSecret,omitempty"`
	Attributes      interface{} `json:"attributes,omitempty"`
}

type ModifyBackupTargetRequest struct {
	BackupTargetID int64       `json:"backupTargetID"`
	Name           string      `json:"name,omitempty"`
	Attributes     interface{} `json:"attributes,omitempty"`
}

type ModifyClusterAdminRequest struct {
	ClusterAdminID int64       `json:"clusterAdminID"`
	Password       string      `json:"password,omitempty"`
	Access         []string    `json:"access,omitempty"`
	Attributes     interface{} `json:"attributes,omitempty"`
}

type ModifyClusterFullThresholdRequest struct {
	Stage2AwareThreshold           int64 `json:"stage2AwareThreshold,omitempty"`
	Stage3BlockThresholdPercent    int64 `json:"stage3BlockThresholdPercent,omitempty"`
	MaxMetadataOverProvisionFactor int64 `json:"maxMetadataOverProvisionFactor,omitempty"`
}

type ModifyGroupSnapshotRequest struct {
	GroupSnapshotID         int64  `json:"groupSnapshotID"`
	ExpirationTime          string `json:"expirationTime,omitempty"`
	EnableRemoteReplication bool   `json:"enableRemoteReplication,omitempty"`
}

type ModifyScheduleRequest struct {
	Schedule Schedule `json:"schedule"`
}

type ModifySnapshotRequest struct {
	SnapshotID              int64  `json:"snapshotID"`
	ExpirationTime          string `json:"expirationTime,omitempty"`
	EnableRemoteReplication bool   `json:"enableRemoteReplication,omitempty"`
}

type ModifyVirtualNetworkRequest struct {
	VirtualNetworkID  int64          `json:"virtualNetworkID,omitempty"`
	VirtualNetworkTag int64          `json:"virtualNetworkTag,omitempty"`
	Name              string         `json:"name,omitempty"`
	AddressBlocks     []AddressBlock `json:"addressBlocks,omitempty"`
	Netmask           string         `json:"netmask,omitempty"`
	Svip              string         `json:"svip,omitempty"`
	Gateway           string         `json:"gateway,omitempty"`
	Namespace         bool           `json:"namespace,omitempty"`
	Attributes        interface{}    `json:"attributes,omitempty"`
}

type ModifyVolumeAccessGroupLunAssignmentsRequest struct {
	VolumeAccessGroupID int64           `json:"volumeAccessGroupID"`
	LunAssignments      []LunAssignment `json:"lunAssignments"`
}

type ModifyVolumeAccessGroupRequest struct {
	VolumeAccessGroupID int64       `json:"volumeAccessGroupID"`
	VirtualNetworkID    []int64     `json:"virtualNetworkID,omitempty"`
	VirtualNetworkTags  []int64     `json:"virtualNetworkTags,omitempty"`
	Name                string      `json:"name,omitempty"`
	Initiators          []string    `json:"initiators,omitempty"`
	Volumes             []int64     `json:"volumes,omitempty"`
	Attributes          interface{} `json:"attributes,omitempty"`
}

type ModifyVolumePairRequest struct {
	VolumeID     int64  `json:"volumeID"`
	PausedManual bool   `json:"pausedManual,omitempty"`
	Mode         string `json:"mode,omitempty"`
}

type ModifyVolumeRequest struct {
	VolumeID      int64       `json:"volumeID"`
	AccountID     int64       `json:"accountID,omitempty"`
	Access        string      `json:"access,omitempty"`
	SetCreateTime string      `json:"setCreateTime,omitempty"`
	Qos           QoS         `json:"qos,omitempty"`
	TotalSize     int64       `json:"totalSize,omitempty"`
	Attributes    interface{} `json:"attributes,omitempty"`
}

type PurgeDeletedVolumeRequest struct {
	VolumeID int64 `json:"volumeID"`
}

type RemoveAccountRequest struct {
	AccountID int64 `json:"accountID"`
}

type RemoveBackupTargetRequest struct {
	BackupTargetID int64 `json:"backupTargetID"`
}

type RemoveClusterAdminRequest struct {
	ClusterAdminID int64 `json:"clusterAdminID"`
}

type RemoveClusterPairRequest struct {
	ClusterPairID int64 `json:"clusterPairID"`
}

type RemoveDrivesRequest struct {
	Drives []int64 `json:"drives"`
}

type RemoveInitiatorsFromVolumeAccessGroupRequest struct {
	VolumeAccessGroupID int64    `json:"volumeAccessGroupID"`
	Initiators          []string `json:"initiators"`
}

type RemoveNodesRequest struct {
	Nodes []int64 `json:"nodes"`
}

type RemoveVirtualNetworkRequest struct {
	VirtualNetworkID  int64 `json:"virtualNetworkID,omitempty"`
	VirtualNetworkTag int64 `json:"virtualNetworkTag,omitempty"`
}

type RemoveVolumePairRequest struct {
	VolumeID int64 `json:"volumeID"`
}

type RemoveVolumesFromVolumeAccessGroupRequest struct {
	VolumeAccessGroupID int64   `json:"volumeAccessGroupID"`
	Volumes             []int64 `json:"volumes"`
}

type ResetDrivesRequest struct {
	Drives string `json:"drives"`
	Force  bool   `json:"force"`
}

type RestoreDeletedVolumeRequest struct {
	VolumeID int64 `json:"volumeID"`
}

type RollbackToGroupSnapshotRequest struct {
	GroupSnapshotID  int64       `json:"groupSnapshotID"`
	SaveCurrentState bool        `json:"saveCurrentState"`
	Name             string      `json:"name,omitempty"`
	Attributes       interface{} `json:"attributes,omitempty"`
}

type RollbackToSnapshotRequest struct {
	VolumeID         int64       `json:"volumeID"`
	SnapshotID       int64       `json:"snapshotID"`
	SaveCurrentState bool        `json:"saveCurrentState"`
	Name             string      `json:"name,omitempty"`
	Attributes       interface{} `json:"attributes,omitempty"`
}

type SecureEraseDrivesRequest struct {
	Drives []int64 `json:"drives"`
}

type SetClusterConfigRequest struct {
	Cluster ClusterConfig `json:"cluster"`
}

type SetConfigRequest struct {
	Config Config `json:"config"`
}

type SetNetworkConfigRequest struct {
	Network Network `json:"network"`
}

type SetSnmpACLRequest struct {
	Networks []SnmpNetwork   `json:"networks"`
	UsmUsers []SnmpV3UsmUser `json:"usmUsers"`
}

type SetSnmpInfoRequest struct {
	Networks      []SnmpNetwork   `json:"networks,omitempty"`
	Enabled       bool            `json:"enabled,omitempty"`
	SnmpV3Enabled bool            `json:"snmpV3Enabled,omitempty"`
	UsmUsers      []SnmpV3UsmUser `json:"usmUsers,omitempty"`
}

type SetSnmpTrapInfoRequest struct {
	TrapRecipients                   []SnmpTrapRecipient `json:"trapRecipients"`
	ClusterFaultTrapsEnabled         bool                `json:"clusterFaultTrapsEnabled"`
	ClusterFaultResolvedTrapsEnabled bool                `json:"clusterFaultResolvedTrapsEnabled"`
	ClusterEventTrapsEnabled         bool                `json:"clusterEventTrapsEnabled"`
}

type StartBulkVolumeReadRequest struct {
	VolumeID         int64       `json:"volumeID"`
	Format           string      `json:"format"`
	SnapshotID       int64       `json:"snapshotID,omitempty"`
	Script           string      `json:"script,omitempty"`
	ScriptParameters interface{} `json:"scriptParameters,omitempty"`
	Attributes       interface{} `json:"attributes,omitempty"`
}

type StartBulkVolumeWriteRequest struct {
	VolumeID         int64       `json:"volumeID"`
	Format           string      `json:"format"`
	Script           string      `json:"script,omitempty"`
	ScriptParameters interface{} `json:"scriptParameters,omitempty"`
	Attributes       interface{} `json:"attributes,omitempty"`
}

type StartVolumePairingRequest struct {
	VolumeID int64  `json:"volumeID"`
	Mode     string `json:"mode,omitempty"`
}

type TestConnectEnsembleRequest struct {
	Ensemble string `json:"ensemble,omitempty"`
}

type TestConnectMvipRequest struct {
	Mvip string `json:"mvip,omitempty"`
}

type TestConnectSvipRequest struct {
	Svip string `json:"svip,omitempty"`
}

type TestDrivesRequest struct {
	Minutes int64 `json:"minutes,omitempty"`
	Force   bool  `json:"force"`
}

type TestLdapAuthenticationRequest struct {
	Username          string            `json:"username"`
	Password          string            `json:"password"`
	LdapConfiguration LdapConfiguration `json:"ldapConfiguration,omitempty"`
}

type TestPingRequest struct {
	Attempts        int64  `json:"attempts,omitempty"`
	Hosts           string `json:"hosts,omitempty"`
	TotalTimeoutSec int64  `json:"totalTimeoutSec,omitempty"`
	PacketSize      int64  `json:"packetSize,omitempty"`
	PingTimeoutMsec int64  `json:"pingTimeoutMsec,omitempty"`
}

type UpdateBulkVolumeStatusRequest struct {
	Key             string      `json:"key"`
	Status          string      `json:"status"`
	PercentComplete string      `json:"percentComplete,omitempty"`
	Message         string      `json:"message,omitempty"`
	Attributes      interface{} `json:"attributes,omitempty"`
}
