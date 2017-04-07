package sftypes

type AddAccountResult struct {
	AccountID int64   `json:"accountID"`
	Account   Account `json:"account,omitempty"`
}

type AddClusterAdminResult struct {
	ClusterAdminID int64 `json:"clusterAdminID"`
}

type AddDrivesResult struct {
	AsyncHandle int64 `json:"asyncHandle,omitempty"`
}

type AddLdapClusterAdminResult struct {
	ClusterAdminID int64 `json:"clusterAdminID,omitempty"`
}

type AddNodesResult struct {
	AutoInstall bool        `json:"autoInstall,omitempty"`
	Nodes       []AddedNode `json:"nodes"`
}

type AddVirtualNetworkResult struct {
	VirtualNetworkID int64 `json:"virtualNetworkID,omitempty"`
}

type AsyncHandleResult struct {
	AsyncHandle int64 `json:"asyncHandle"`
}

type CancelCloneResult struct {
}

type CancelGroupCloneResult struct {
}

type ClearClusterFaultsResult struct {
}

type CloneMultipleVolumesResult struct {
	AsyncHandle  int64                    `json:"asyncHandle"`
	GroupCloneID int64                    `json:"groupCloneID"`
	Members      []GroupCloneVolumeMember `json:"members"`
}

type CloneVolumeResult struct {
	Volume      Volume `json:"volume,omitempty"`
	CloneID     int64  `json:"cloneID"`
	VolumeID    int64  `json:"volumeID"`
	AsyncHandle int64  `json:"asyncHandle"`
}

type CompleteClusterPairingResult struct {
	ClusterPairID int64 `json:"clusterPairID"`
}

type CompleteVolumePairingResult struct {
}

type CopyVolumeResult struct {
	CloneID     int64 `json:"cloneID"`
	AsyncHandle int64 `json:"asyncHandle"`
}

type CreateBackupTargetResult struct {
	BackupTargetID int64 `json:"backupTargetID"`
}

type CreateClusterResult struct {
}

type CreateGroupSnapshotResult struct {
	GroupSnapshot   GroupSnapshot          `json:"groupSnapshot"`
	GroupSnapshotID int64                  `json:"groupSnapshotID"`
	Members         []GroupSnapshotMembers `json:"members"`
}

type CreateInitiatorsResult struct {
	Initiators []Initiator `json:"initiators"`
}

type CreateScheduleResult struct {
	ScheduleID int64 `json:"scheduleID"`
}

type CreateSnapshotResult struct {
	Snapshot   Snapshot `json:"snapshot"`
	SnapshotID int64    `json:"snapshotID"`
	Checksum   string   `json:"checksum"`
}

type CreateStorageContainerResult struct {
	StorageContainer StorageContainer `json:"storageContainer"`
}

type CreateSupportBundleResult struct {
	Details  SupportBundleDetails `json:"details"`
	Duration string               `json:"duration"`
	Result   string               `json:"result"`
}

type CreateVolumeAccessGroupResult struct {
	VolumeAccessGroupID int64             `json:"volumeAccessGroupID"`
	VolumeAccessGroup   VolumeAccessGroup `json:"volumeAccessGroup,omitempty"`
}

type CreateVolumeResult struct {
	Volume   Volume `json:"volume,omitempty"`
	VolumeID int64  `json:"volumeID"`
	Curve    int64  `json:"curve"`
}

type DeleteAllSupportBundlesResult struct {
	Duration string      `json:"duration"`
	Details  interface{} `json:"details"`
	Result   string      `json:"result"`
}

type DeleteGroupSnapshotResult struct {
}

type DeleteInitiatorsResult struct {
}

type DeleteSnapshotResult struct {
}

type DeleteStorageContainerResult struct {
}

type DeleteVolumeAccessGroupResult struct {
}

type DeleteVolumeResult struct {
	Volume Volume `json:"volume,omitempty"`
}

type DeleteVolumesResult struct {
	Volumes []Volume  `json:"volumes"`
	Curve   VolumeQOS `json:"curve,omitempty"`
}

type DisableEncryptionAtRestResult struct {
}

type DisableLdapAuthenticationResult struct {
}

type DisableSnmpResult struct {
}

type EnableEncryptionAtRestResult struct {
}

type EnableFeatureResult struct {
}

type EnableLdapAuthenticationResult struct {
}

type EnableSnmpResult struct {
}

type FibreChannelPortInfoResult struct {
	Result FibreChannelPortList `json:"result"`
}

type GetAPIResult struct {
	SupportedVersions []float64 `json:"supportedVersions"`
	CurrentVersion    float64   `json:"currentVersion"`
}

type GetAccountResult struct {
	Account Account `json:"account"`
}

type GetBackupTargetResult struct {
	BackupTarget BackupTarget `json:"backupTarget"`
}

type GetBootstrapConfigResult struct {
	ClusterName string              `json:"clusterName"`
	NodeName    string              `json:"nodeName"`
	Nodes       []NodeWaitingToJoin `json:"nodes"`
	Version     string              `json:"version"`
}

type GetClusterCapacityResult struct {
	ClusterCapacity ClusterCapacity `json:"clusterCapacity"`
}

type GetClusterConfigResult struct {
	Cluster ClusterConfig `json:"cluster"`
}

type GetClusterFullThresholdResult struct {
	BlockFullness                  string `json:"blockFullness"`
	Fullness                       string `json:"fullness"`
	MaxMetadataOverProvisionFactor int64  `json:"maxMetadataOverProvisionFactor"`
	MetadataFullness               string `json:"metadataFullness"`
	SliceReserveUsedThresholdPct   int64  `json:"sliceReserveUsedThresholdPct"`
	Stage2AwareThreshold           int64  `json:"stage2AwareThreshold"`
	Stage2BlockThresholdBytes      int64  `json:"stage2BlockThresholdBytes"`
	Stage3BlockThresholdBytes      int64  `json:"stage3BlockThresholdBytes"`
	Stage3BlockThresholdPercent    int64  `json:"stage3BlockThresholdPercent"`
	Stage3LowThreshold             int64  `json:"stage3LowThreshold"`
	Stage4CriticalThreshold        int64  `json:"stage4CriticalThreshold"`
	Stage4BlockThresholdBytes      int64  `json:"stage4BlockThresholdBytes"`
	Stage5BlockThresholdBytes      int64  `json:"stage5BlockThresholdBytes"`
	SumTotalClusterBytes           int64  `json:"sumTotalClusterBytes"`
	SumTotalMetadataClusterBytes   int64  `json:"sumTotalMetadataClusterBytes"`
	SumUsedClusterBytes            int64  `json:"sumUsedClusterBytes"`
	SumUsedMetadataClusterBytes    int64  `json:"sumUsedMetadataClusterBytes"`
}

type GetClusterHardwareInfoResult struct {
	ClusterHardwareInfo ClusterHardwareInfo `json:"clusterHardwareInfo"`
}

type GetClusterInfoResult struct {
	ClusterInfo ClusterInfo `json:"clusterInfo"`
}

type GetClusterMasterNodeIDResult struct {
	NodeID int64 `json:"nodeID"`
}

type GetClusterStateResult struct {
	Nodes   []NodeStateResult `json:"nodes,omitempty"`
	Cluster string            `json:"cluster,omitempty"`
	State   string            `json:"state,omitempty"`
}

type GetClusterStatsResult struct {
	ClusterStats ClusterStats `json:"clusterStats"`
}

type GetClusterVersionInfoResult struct {
	ClusterAPIVersion   string               `json:"clusterAPIVersion"`
	ClusterVersion      string               `json:"clusterVersion"`
	ClusterVersionInfo  []ClusterVersionInfo `json:"clusterVersionInfo"`
	SoftwareVersionInfo SoftwareVersionInfo  `json:"softwareVersionInfo"`
}

type GetConfigResult struct {
	Config Config `json:"config"`
}

type GetCurrentClusterAdminResult struct {
	ClusterAdmin ClusterAdmin `json:"clusterAdmin"`
}

type GetDriveConfigResult struct {
	DriveConfig DrivesConfigInfo `json:"driveConfig"`
}

type GetDriveHardwareInfoResult struct {
	DriveHardwareInfo DriveHardwareInfo `json:"driveHardwareInfo"`
}

type GetDriveStatsResult struct {
	DriveStats DriveStats `json:"driveStats"`
}

type GetEfficiencyResult struct {
	Compression      float64 `json:"compression,omitempty"`
	Deduplication    float64 `json:"deduplication,omitempty"`
	ThinProvisioning float64 `json:"thinProvisioning,omitempty"`
	Timestamp        string  `json:"timestamp"`
	MissingVolumes   []int64 `json:"missingVolumes"`
}

type GetFeatureStatusResult struct {
	Features []FeatureObject `json:"features"`
}

type GetHardwareConfigResult struct {
	HardwareConfig interface{} `json:"hardwareConfig"`
}

type GetHardwareInfoResult struct {
	HardwareInfo interface{} `json:"hardwareInfo"`
}

type GetIpmiConfigNodesResult struct {
	NodeID int64       `json:"nodeID"`
	Result interface{} `json:"result"`
}

type GetIpmiConfigResult struct {
	Nodes []GetIpmiConfigNodesResult `json:"nodes"`
}

type GetIpmiInfoNodesResult struct {
	NodeID int64                        `json:"nodeID"`
	Result GetIpmiInfoNodesResultObject `json:"result"`
}

type GetIpmiInfoResult struct {
	Nodes []GetIpmiInfoNodesResult `json:"nodes"`
}

type GetLdapConfigurationResult struct {
	LdapConfiguration LdapConfiguration `json:"ldapConfiguration"`
}

type GetLimitsResult struct {
	AccountCountMax                        int64 `json:"accountCountMax"`
	AccountNameLengthMax                   int64 `json:"accountNameLengthMax"`
	AccountNameLengthMin                   int64 `json:"accountNameLengthMin"`
	BulkVolumeJobsPerNodeMax               int64 `json:"bulkVolumeJobsPerNodeMax"`
	BulkVolumeJobsPerVolumeMax             int64 `json:"bulkVolumeJobsPerVolumeMax"`
	CloneJobsPerVolumeMax                  int64 `json:"cloneJobsPerVolumeMax"`
	ClusterPairsCountMax                   int64 `json:"clusterPairsCountMax"`
	InitiatorNameLengthMax                 int64 `json:"initiatorNameLengthMax"`
	InitiatorCountMax                      int64 `json:"initiatorCountMax"`
	InitiatorsPerVolumeAccessGroupCountMax int64 `json:"initiatorsPerVolumeAccessGroupCountMax"`
	IscsiSessionsFromFibreChannelNodesMax  int64 `json:"iscsiSessionsFromFibreChannelNodesMax"`
	SecretLengthMax                        int64 `json:"secretLengthMax"`
	SecretLengthMin                        int64 `json:"secretLengthMin"`
	SnapshotNameLengthMax                  int64 `json:"snapshotNameLengthMax"`
	SnapshotsPerVolumeMax                  int64 `json:"snapshotsPerVolumeMax"`
	VolumeAccessGroupCountMax              int64 `json:"volumeAccessGroupCountMax"`
	VolumeAccessGroupLunMax                int64 `json:"volumeAccessGroupLunMax"`
	VolumeAccessGroupNameLengthMax         int64 `json:"volumeAccessGroupNameLengthMax"`
	VolumeAccessGroupNameLengthMin         int64 `json:"volumeAccessGroupNameLengthMin"`
	VolumeAccessGroupsPerInitiatorCountMax int64 `json:"volumeAccessGroupsPerInitiatorCountMax"`
	VolumeAccessGroupsPerVolumeCountMax    int64 `json:"volumeAccessGroupsPerVolumeCountMax"`
	InitiatorAliasLengthMax                int64 `json:"initiatorAliasLengthMax"`
	VolumeBurstIOPSMax                     int64 `json:"volumeBurstIOPSMax"`
	VolumeBurstIOPSMin                     int64 `json:"volumeBurstIOPSMin"`
	VolumeCountMax                         int64 `json:"volumeCountMax"`
	VolumeMaxIOPSMax                       int64 `json:"volumeMaxIOPSMax"`
	VolumeMaxIOPSMin                       int64 `json:"volumeMaxIOPSMin"`
	VolumeMinIOPSMax                       int64 `json:"volumeMinIOPSMax"`
	VolumeMinIOPSMin                       int64 `json:"volumeMinIOPSMin"`
	VolumeNameLengthMax                    int64 `json:"volumeNameLengthMax"`
	VolumeNameLengthMin                    int64 `json:"volumeNameLengthMin"`
	VolumeSizeMax                          int64 `json:"volumeSizeMax"`
	VolumeSizeMin                          int64 `json:"volumeSizeMin"`
	VolumesPerAccountCountMax              int64 `json:"volumesPerAccountCountMax"`
	VolumesPerGroupSnapshotMax             int64 `json:"volumesPerGroupSnapshotMax"`
	VolumesPerVolumeAccessGroupCountMax    int64 `json:"volumesPerVolumeAccessGroupCountMax"`
	ClusterAdminAccountMax                 int64 `json:"clusterAdminAccountMax,omitempty"`
	FibreChannelVolumeAccessMax            int64 `json:"fibreChannelVolumeAccessMax,omitempty"`
	VirtualVolumesPerAccountCountMax       int64 `json:"virtualVolumesPerAccountCountMax,omitempty"`
	VirtualVolumeCountMax                  int64 `json:"virtualVolumeCountMax,omitempty"`
}

type GetLoginSessionInfoResult struct {
	LoginSessionInfo LoginSessionInfo `json:"loginSessionInfo"`
}

type GetNetworkConfigResult struct {
	Network Network `json:"network"`
}

type GetNodeHardwareInfoResult struct {
	NodeHardwareInfo interface{} `json:"nodeHardwareInfo"`
}

type GetNodeStatsResult struct {
	NodeStats NodeStatsInfo `json:"nodeStats"`
}

type GetNtpInfoResult struct {
	Broadcastclient bool     `json:"broadcastclient"`
	Servers         []string `json:"servers"`
}

type GetNvramInfoResult struct {
	NvramInfo interface{} `json:"nvramInfo"`
}

type GetOriginNodeResult struct {
	Origin Origin `json:"origin,omitempty"`
}

type GetOriginResult struct {
	Nodes []GetOriginNode `json:"nodes"`
}

type GetPendingOperationResult struct {
	PendingOperation PendingOperation `json:"pendingOperation"`
}

type GetRemoteLoggingHostsResult struct {
	RemoteHosts []LoggingServer `json:"remoteHosts"`
}

type GetScheduleResult struct {
	Schedule Schedule `json:"schedule"`
}

type GetSnmpACLResult struct {
	Networks []SnmpNetwork   `json:"networks,omitempty"`
	UsmUsers []SnmpV3UsmUser `json:"usmUsers,omitempty"`
}

type GetSnmpInfoResult struct {
	Networks      []SnmpNetwork   `json:"networks,omitempty"`
	Enabled       bool            `json:"enabled"`
	SnmpV3Enabled bool            `json:"snmpV3Enabled"`
	UsmUsers      []SnmpV3UsmUser `json:"usmUsers,omitempty"`
}

type GetSnmpStateResult struct {
	Enabled       bool `json:"enabled"`
	SnmpV3Enabled bool `json:"snmpV3Enabled"`
}

type GetSnmpTrapInfoResult struct {
	TrapRecipients                   []SnmpTrapRecipient `json:"trapRecipients"`
	ClusterFaultTrapsEnabled         bool                `json:"clusterFaultTrapsEnabled"`
	ClusterFaultResolvedTrapsEnabled bool                `json:"clusterFaultResolvedTrapsEnabled"`
	ClusterEventTrapsEnabled         bool                `json:"clusterEventTrapsEnabled"`
}

type GetStorageContainerEfficiencyResult struct {
	Compression      float64 `json:"compression"`
	Deduplication    float64 `json:"deduplication"`
	MissingVolumes   []int64 `json:"missingVolumes"`
	ThinProvisioning float64 `json:"thinProvisioning"`
	Timestamp        string  `json:"timestamp"`
}

type GetSystemStatusResult struct {
	RebootRequired bool `json:"rebootRequired"`
}

type GetVirtualVolumeCountResult struct {
	Count int64 `json:"count"`
}

type GetVolumeAccessGroupLunAssignmentsResult struct {
	VolumeAccessGroupLunAssignments VolumeAccessGroupLunAssignments `json:"volumeAccessGroupLunAssignments"`
}

type GetVolumeCountResult struct {
	Count int64 `json:"count"`
}

type GetVolumeEfficiencyResult struct {
	Compression      float64 `json:"compression,omitempty"`
	Deduplication    float64 `json:"deduplication"`
	MissingVolumes   []int64 `json:"missingVolumes"`
	ThinProvisioning float64 `json:"thinProvisioning"`
	Timestamp        string  `json:"timestamp"`
}

type GetVolumeStatsResult struct {
	VolumeStats VolumeStats `json:"volumeStats"`
}

type ListAccountsResult struct {
	Accounts []Account `json:"accounts"`
}

type ListActiveNodesResult struct {
	Nodes []Node `json:"nodes"`
}

type ListActivePairedVolumesResult struct {
	Volumes []Volume `json:"volumes"`
}

type ListActiveVolumesResult struct {
	Volumes []Volume `json:"volumes"`
}

type ListAllNodesResult struct {
	Nodes              []Node              `json:"nodes"`
	PendingNodes       []PendingNode       `json:"pendingNodes"`
	PendingActiveNodes []PendingActiveNode `json:"pendingActiveNodes,omitempty"`
}

type ListAsyncResultsResult struct {
	AsyncHandles []AsyncHandle `json:"asyncHandles"`
}

type ListBackupTargetsResult struct {
	BackupTargets []BackupTarget `json:"backupTargets"`
}

type ListBulkVolumeJobsResult struct {
	BulkVolumeJobs []BulkVolumeJob `json:"bulkVolumeJobs"`
}

type ListClusterAdminsResult struct {
	ClusterAdmins []ClusterAdmin `json:"clusterAdmins"`
}

type ListClusterFaultsResult struct {
	Faults []ClusterFaultInfo `json:"faults"`
}

type ListClusterPairsResult struct {
	ClusterPairs []PairedCluster `json:"clusterPairs"`
}

type ListDeletedVolumesResult struct {
	Volumes []Volume `json:"volumes"`
}

type ListDriveHardwareResult struct {
	Nodes []NodeDriveHardware `json:"nodes"`
}

type ListDriveStatsResult struct {
	DriveStats []DriveStats  `json:"driveStats"`
	Errors     []interface{} `json:"errors"`
}

type ListDrivesResult struct {
	Drives []DriveInfo `json:"drives"`
}

type ListEventsResult struct {
	EventQueueType string      `json:"eventQueueType"`
	Events         []EventInfo `json:"events"`
}

type ListFibreChannelPortInfoResult struct {
	FibreChannelPortInfo FibreChannelPortInfoResult `json:"fibreChannelPortInfo"`
}

type ListFibreChannelSessionsResult struct {
	Sessions []FibreChannelSession `json:"sessions"`
}

type ListGroupSnapshotsResult struct {
	GroupSnapshots []GroupSnapshot `json:"groupSnapshots"`
}

type ListISCSISessionsResult struct {
	Sessions []ISCSISession `json:"sessions"`
}

type ListInitiatorsResult struct {
	Initiators []Initiator `json:"initiators"`
}

type ListNetworkInterfacesResult struct {
	Interfaces []NetworkInterface `json:"interfaces"`
}

type ListNodeFibreChannelPortInfoResult struct {
	Nodes             []NodeFibreChannelPortInfoResult `json:"nodes,omitempty"`
	FibreChannelPorts []FibreChannelPortInfo           `json:"fibreChannelPorts"`
}

type ListNodeStatsResult struct {
	NodeStats NodeStatsNodes `json:"nodeStats"`
}

type ListPendingActiveNodesResult struct {
	PendingActiveNodes []PendingActiveNode `json:"pendingActiveNodes"`
}

type ListPendingNodesResult struct {
	PendingNodes []PendingNode `json:"pendingNodes"`
}

type ListProtocolEndpointsResult struct {
	ProtocolEndpoints []ProtocolEndpoint `json:"protocolEndpoints"`
}

type ListSchedulesResult struct {
	Schedules []Schedule `json:"schedules"`
}

type ListServicesResult struct {
	Services []DetailedService `json:"services"`
}

type ListSnapshotsResult struct {
	Snapshots []Snapshot `json:"snapshots"`
}

type ListStorageContainersResult struct {
	StorageContainers []StorageContainer `json:"storageContainers"`
}

type ListSyncJobsResult struct {
	SyncJobs []SyncJob `json:"syncJobs"`
}

type ListTestsResult struct {
	Tests interface{} `json:"tests"`
}

type ListUtilitiesResult struct {
	Utilities interface{} `json:"utilities"`
}

type ListVirtualNetworksResult struct {
	VirtualNetworks []VirtualNetwork `json:"virtualNetworks"`
}

type ListVirtualVolumeBindingsResult struct {
	Bindings []VirtualVolumeBinding `json:"bindings"`
}

type ListVirtualVolumeHostsResult struct {
	Hosts []VirtualVolumeHost `json:"hosts"`
}

type ListVirtualVolumeTasksResult struct {
	Tasks []VirtualVolumeTask `json:"tasks"`
}

type ListVirtualVolumesResult struct {
	VirtualVolumes      []VirtualVolumeInfo `json:"virtualVolumes"`
	NextVirtualVolumeID string              `json:"nextVirtualVolumeID,omitempty"`
}

type ListVolumeAccessGroupsResult struct {
	VolumeAccessGroups         []VolumeAccessGroup `json:"volumeAccessGroups"`
	VolumeAccessGroupsNotFound []int64             `json:"volumeAccessGroupsNotFound,omitempty"`
}

type ListVolumeStatsByAccountResult struct {
	VolumeStats []VolumeStats `json:"volumeStats"`
}

type ListVolumeStatsByVirtualVolumeResult struct {
	VolumeStats []VolumeStats `json:"volumeStats"`
}

type ListVolumeStatsByVolumeAccessGroupResult struct {
	VolumeStats []VolumeStats `json:"volumeStats"`
}

type ListVolumeStatsByVolumeResult struct {
	VolumeStats []VolumeStats `json:"volumeStats"`
}

type ListVolumeStatsResult struct {
	VolumeStats []VolumeStats `json:"volumeStats"`
}

type ListVolumesForAccountResult struct {
	Volumes []Volume `json:"volumes"`
}

type ListVolumesResult struct {
	Volumes []Volume `json:"volumes"`
}

type ModifyAccountResult struct {
	Account Account `json:"account"`
}

type ModifyBackupTargetResult struct {
}

type ModifyClusterAdminResult struct {
}

type ModifyClusterFullThresholdResult struct {
	BlockFullness                  string `json:"blockFullness"`
	Fullness                       string `json:"fullness"`
	MaxMetadataOverProvisionFactor int64  `json:"maxMetadataOverProvisionFactor"`
	MetadataFullness               string `json:"metadataFullness"`
	SliceReserveUsedThresholdPct   int64  `json:"sliceReserveUsedThresholdPct"`
	Stage2AwareThreshold           int64  `json:"stage2AwareThreshold"`
	Stage2BlockThresholdBytes      int64  `json:"stage2BlockThresholdBytes"`
	Stage3BlockThresholdBytes      int64  `json:"stage3BlockThresholdBytes"`
	Stage3BlockThresholdPercent    int64  `json:"stage3BlockThresholdPercent"`
	Stage3LowThreshold             int64  `json:"stage3LowThreshold"`
	Stage4CriticalThreshold        int64  `json:"stage4CriticalThreshold"`
	Stage4BlockThresholdBytes      int64  `json:"stage4BlockThresholdBytes"`
	Stage5BlockThresholdBytes      int64  `json:"stage5BlockThresholdBytes"`
	SumTotalClusterBytes           int64  `json:"sumTotalClusterBytes"`
	SumTotalMetadataClusterBytes   int64  `json:"sumTotalMetadataClusterBytes"`
	SumUsedClusterBytes            int64  `json:"sumUsedClusterBytes"`
	SumUsedMetadataClusterBytes    int64  `json:"sumUsedMetadataClusterBytes"`
}

type ModifyGroupSnapshotResult struct {
	GroupSnapshot GroupSnapshot `json:"groupSnapshot"`
}

type ModifyInitiatorsResult struct {
	Initiators []Initiator `json:"initiators"`
}

type ModifyScheduleResult struct {
	Schedule Schedule `json:"schedule"`
}

type ModifySnapshotResult struct {
	Snapshot Snapshot `json:"snapshot,omitempty"`
}

type ModifyStorageContainerResult struct {
	StorageContainer StorageContainer `json:"storageContainer"`
}

type ModifyVolumeAccessGroupLunAssignmentsResult struct {
	VolumeAccessGroupLunAssignments VolumeAccessGroupLunAssignments `json:"volumeAccessGroupLunAssignments"`
}

type ModifyVolumeAccessGroupResult struct {
	VolumeAccessGroup VolumeAccessGroup `json:"volumeAccessGroup"`
}

type ModifyVolumePairResult struct {
}

type ModifyVolumeResult struct {
	Volume Volume `json:"volume,omitempty"`
	Curve  QoS    `json:"curve,omitempty"`
}

type ModifyVolumesResult struct {
	Qos     QoS      `json:"qos,omitempty"`
	Volumes []Volume `json:"volumes"`
}

type NodeFibreChannelPortInfoResult struct {
	NodeID int64                `json:"nodeID"`
	Result FibreChannelPortList `json:"result"`
}

type NodeStateResult struct {
	NodeID int64         `json:"nodeID"`
	Result NodeStateInfo `json:"result,omitempty"`
}

type PurgeDeletedVolumeResult struct {
}

type PurgeDeletedVolumesResult struct {
}

type RemoveAccountResult struct {
}

type RemoveBackupTargetResult struct {
}

type RemoveClusterAdminResult struct {
}

type RemoveClusterPairResult struct {
}

type RemoveNodesResult struct {
}

type RemoveVirtualNetworkResult struct {
}

type RemoveVolumePairResult struct {
}

type ResetDrivesResult struct {
	Details ResetDrivesDetails `json:"details"`
}

type ResetNodeResult struct {
	Details  ResetNodeDetails `json:"details,omitempty"`
	Duration string           `json:"duration,omitempty"`
	Result   string           `json:"result,omitempty"`
	RtfiInfo RtfiInfo         `json:"rtfiInfo,omitempty"`
}

type RestoreDeletedVolumeResult struct {
}

type RollbackToGroupSnapshotResult struct {
	GroupSnapshot   GroupSnapshot          `json:"groupSnapshot,omitempty"`
	GroupSnapshotID int64                  `json:"groupSnapshotID,omitempty"`
	Members         []GroupSnapshotMembers `json:"members,omitempty"`
}

type RollbackToSnapshotResult struct {
	Snapshot   Snapshot `json:"snapshot,omitempty"`
	SnapshotID int64    `json:"snapshotID,omitempty"`
	Checksum   string   `json:"checksum,omitempty"`
}

type SetClusterConfigResult struct {
	Cluster ClusterConfig `json:"cluster"`
}

type SetConfigResult struct {
	Config Config `json:"config"`
}

type SetDefaultQoSResult struct {
	MinIOPS   int64 `json:"minIOPS"`
	MaxIOPS   int64 `json:"maxIOPS"`
	BurstIOPS int64 `json:"burstIOPS"`
}

type SetLoginSessionInfoResult struct {
}

type SetNetworkConfigResult struct {
	Network Network `json:"network"`
}

type SetNtpInfoResult struct {
}

type SetRemoteLoggingHostsResult struct {
}

type SetSnmpACLResult struct {
}

type SetSnmpInfoResult struct {
}

type SetSnmpTrapInfoResult struct {
}

type ShutdownResult struct {
	Failed     []int64 `json:"failed"`
	Successful []int64 `json:"successful"`
}

type SnmpSendTestTrapsResult struct {
	Status string `json:"status"`
}

type StartBulkVolumeReadResult struct {
	AsyncHandle int64  `json:"asyncHandle"`
	Key         string `json:"key"`
	Url         string `json:"url"`
}

type StartBulkVolumeWriteResult struct {
	AsyncHandle int64  `json:"asyncHandle"`
	Key         string `json:"key"`
	Url         string `json:"url"`
}

type StartClusterPairingResult struct {
	ClusterPairingKey string `json:"clusterPairingKey"`
	ClusterPairID     int64  `json:"clusterPairID"`
}

type StartVolumePairingResult struct {
	VolumePairingKey string `json:"volumePairingKey"`
}

type TestConnectEnsembleResult struct {
	Details  TestConnectEnsembleDetails `json:"details"`
	Duration string                     `json:"duration"`
	Result   string                     `json:"result"`
}

type TestConnectMvipResult struct {
	Details  TestConnectMvipDetails `json:"details"`
	Duration string                 `json:"duration"`
	Result   string                 `json:"result"`
}

type TestConnectSvipResult struct {
	Details  TestConnectSvipDetails `json:"details"`
	Duration string                 `json:"duration"`
	Result   string                 `json:"result"`
}

type TestDrivesResult struct {
	Details  string `json:"details"`
	Duration string `json:"duration"`
	Result   string `json:"result"`
}

type TestLdapAuthenticationResult struct {
	Groups []string `json:"groups"`
	UserDN string   `json:"userDN"`
}

type TestPingResult struct {
	Result   string      `json:"result"`
	Duration string      `json:"duration"`
	Details  interface{} `json:"details"`
}

type UpdateBulkVolumeStatusResult struct {
	Status     string      `json:"status"`
	Url        string      `json:"url"`
	Attributes interface{} `json:"attributes"`
}
