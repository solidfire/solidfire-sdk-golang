package sftypes

type AddAccountResult struct {
	AccountID int64 `json:"accountID"`
}

type AddClusterAdminResult struct {
	ClusterAdminID int64 `json:"clusterAdminID"`
}

type AddNodesResult struct {
	Nodes []AddedNode `json:"nodes"`
}

type AddVirtualNetworkResult struct {
	VirtualNetworkID int64 `json:"virtualNetworkID"`
}

type AsyncHandleResult struct {
	AsyncHandle int64 `json:"asyncHandle"`
}

type AsyncResult struct {
	Message string `json:"message"`
}

type CloneMultipleVolumesResult struct {
	AsyncHandle  int64                    `json:"asyncHandle"`
	GroupCloneID int64                    `json:"groupCloneID"`
	Members      []GroupCloneVolumeMember `json:"members"`
}

type CloneVolumeResult struct {
	CloneID     int64 `json:"cloneID"`
	VolumeID    int64 `json:"volumeID"`
	AsyncHandle int64 `json:"asyncHandle"`
}

type CompleteClusterPairingResult struct {
	ClusterPairID int64 `json:"clusterPairID"`
}

type CreateBackupTargetResult struct {
	BackupTargetID int64 `json:"backupTargetID"`
}

type CreateGroupSnapshotResult struct {
	GroupSnapshotID int64                  `json:"groupSnapshotID"`
	Members         []GroupSnapshotMembers `json:"members"`
}

type CreateScheduleResult struct {
	ScheduleID int64 `json:"scheduleID"`
}

type CreateSnapshotResult struct {
	SnapshotID int64  `json:"snapshotID"`
	Checksum   string `json:"checksum"`
}

type CreateVolumeAccessGroupResult struct {
	VolumeAccessGroupID int64 `json:"volumeAccessGroupID"`
}

type CreateVolumeResult struct {
	VolumeID int64 `json:"volumeID"`
	Curve    int64 `json:"curve"`
}

type FibreChannelPortInfoResult struct {
	Result FibreChannelPortList `json:"result"`
}

type GetAPIResult struct {
	CurrentVersion    float64   `json:"currentVersion"`
	SupportedVersions []float64 `json:"supportedVersions"`
}

type GetAccountResult struct {
	Account Account `json:"account"`
}

type GetAsyncResultResult struct {
	Result AsyncResult `json:"result"`
	Status string      `json:"status"`
}

type GetBackupTargetResult struct {
	BackupTarget BackupTarget `json:"backupTarget"`
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

type GetClusterInfoResult struct {
	ClusterInfo ClusterInfo `json:"clusterInfo"`
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

type GetDriveHardwareInfoResult struct {
	DriveHardwareInfo DriveHardwareInfo `json:"driveHardwareInfo"`
}

type GetDriveStatsResult struct {
	DriveStats DriveStats `json:"driveStats"`
}

type GetEfficiencyResult struct {
	Compression      float64 `json:"compression"`
	Deduplication    float64 `json:"deduplication"`
	ThinProvisioning float64 `json:"thinProvisioning"`
	Timestamp        string  `json:"timestamp"`
	MissingVolumes   []int64 `json:"missingVolumes"`
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
	InitiatorsPerVolumeAccessGroupCountMax int64 `json:"initiatorsPerVolumeAccessGroupCountMax"`
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
}

type GetNetworkConfigResult struct {
	Network Network `json:"network"`
}

type GetNodeStatsResult struct {
	NodeStats NodeStatsInfo `json:"nodeStats"`
}

type GetScheduleResult struct {
	Schedule Schedule `json:"schedule"`
}

type GetSnmpACLResult struct {
	Networks []SnmpNetwork   `json:"networks"`
	UsmUsers []SnmpV3UsmUser `json:"usmUsers"`
}

type GetSnmpInfoResult struct {
	Networks      []SnmpNetwork   `json:"networks"`
	Enabled       bool            `json:"enabled"`
	SnmpV3Enabled bool            `json:"snmpV3Enabled"`
	UsmUsers      []SnmpV3UsmUser `json:"usmUsers"`
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

type GetVolumeAccessGroupLunAssignmentsResult struct {
	VolumeAccessGroupLunAssignments VolumeAccessGroupLunAssignments `json:"volumeAccessGroupLunAssignments"`
}

type GetVolumeEfficiencyResult struct {
	Compression      float64 `json:"compression"`
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
	Nodes        []Node        `json:"nodes"`
	PendingNodes []PendingNode `json:"pendingNodes"`
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

type ListNodeFibreChannelPortInfoResult struct {
	Nodes []NodeFibreChannelPortInfoResult `json:"nodes"`
}

type ListNodeStatsResult struct {
	NodeStats NodeStatsNodes `json:"nodeStats"`
}

type ListPendingNodesResult struct {
	PendingNodes []PendingNode `json:"pendingNodes"`
}

type ListSchedulesResult struct {
	Schedules []Schedule `json:"schedules"`
}

type ListSnapshotsResult struct {
	Snapshots []Snapshot `json:"snapshots"`
}

type ListTestsResult struct {
	Tests []string `json:"tests"`
}

type ListUtilitiesResult struct {
	Utilities []string `json:"utilities"`
}

type ListVirtualNetworksResult struct {
	VirtualNetworks []VirtualNetwork `json:"virtualNetworks"`
}

type ListVolumeAccessGroupsResult struct {
	VolumeAccessGroups []VolumeAccessGroup `json:"volumeAccessGroups"`
}

type ListVolumeStatsByAccountResult struct {
	VolumeStats []VolumeStats `json:"volumeStats"`
}

type ListVolumeStatsByVolumeAccessGroupResult struct {
	VolumeStats []VolumeStats `json:"volumeStats"`
}

type ListVolumeStatsByVolumeResult struct {
	VolumeStats []VolumeStats `json:"volumeStats"`
}

type ListVolumesForAccountResult struct {
	Volumes []Volume `json:"volumes"`
}

type ListVolumesResult struct {
	Volumes []Volume `json:"volumes"`
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

type ModifyVolumeResult struct {
	Curve int64 `json:"curve"`
}

type NodeFibreChannelPortInfoResult struct {
	NodeID int64                `json:"nodeID"`
	Result FibreChannelPortList `json:"result"`
}

type ResetDrivesResult struct {
	Details ResetDrivesDetails `json:"details"`
}

type SetClusterConfigResult struct {
	Cluster ClusterConfig `json:"cluster"`
}

type SetConfigResult struct {
	Config Config `json:"config"`
}

type SetNetworkConfigResult struct {
	Network Network `json:"network"`
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
	Details string `json:"details"`
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
