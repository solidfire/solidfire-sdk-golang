package sftypes

type Account struct {
	AccountID       int64       `json:"accountID"`
	Username        string      `json:"username"`
	Status          string      `json:"status"`
	Volumes         []int64     `json:"volumes"`
	InitiatorSecret string      `json:"initiatorSecret,omitempty"`
	TargetSecret    string      `json:"targetSecret,omitempty"`
	Attributes      interface{} `json:"attributes,omitempty"`
}

type AddedNode struct {
	NodeID        int64 `json:"nodeID"`
	PendingNodeID int64 `json:"pendingNodeID"`
}

type AddressBlock struct {
	Start string `json:"start"`
	Size  int64  `json:"size"`
}

type BackupTarget struct {
	Name           string      `json:"name"`
	BackupTargetID int64       `json:"backupTargetID"`
	Attributes     interface{} `json:"attributes,omitempty"`
}

type BulkVolumeJob struct {
	BulkVolumeID    int64       `json:"bulkVolumeID"`
	CreateTime      string      `json:"createTime"`
	ElapsedTime     int64       `json:"elapsedTime"`
	Format          string      `json:"format"`
	Key             string      `json:"key"`
	PercentComplete int64       `json:"percentComplete"`
	RemainingTime   int64       `json:"remainingTime"`
	SrcVolumeID     int64       `json:"srcVolumeID"`
	Status          string      `json:"status"`
	Script          string      `json:"script"`
	SnapshotID      int64       `json:"snapshotID"`
	Type            string      `json:"type"`
	Attributes      interface{} `json:"attributes"`
}

type CloneMultipleVolumeParams struct {
	VolumeID     int64       `json:"volumeID"`
	Access       string      `json:"access,omitempty"`
	Name         string      `json:"name,omitempty"`
	NewAccountID int64       `json:"newAccountID,omitempty"`
	NewSize      int64       `json:"newSize,omitempty"`
	Attributes   interface{} `json:"attributes,omitempty"`
}

type ClusterAdmin struct {
	Access         []string    `json:"access"`
	ClusterAdminID int64       `json:"clusterAdminID"`
	Username       string      `json:"username"`
	Attributes     interface{} `json:"attributes"`
}

type ClusterCapacity struct {
	ActiveBlockSpace             int64  `json:"activeBlockSpace"`
	ActiveSessions               int64  `json:"activeSessions"`
	AverageIOPS                  int64  `json:"averageIOPS"`
	ClusterRecentIOSize          int64  `json:"clusterRecentIOSize"`
	CurrentIOPS                  int64  `json:"currentIOPS"`
	MaxIOPS                      int64  `json:"maxIOPS"`
	MaxOverProvisionableSpace    int64  `json:"maxOverProvisionableSpace"`
	MaxProvisionedSpace          int64  `json:"maxProvisionedSpace"`
	MaxUsedMetadataSpace         int64  `json:"maxUsedMetadataSpace"`
	MaxUsedSpace                 int64  `json:"maxUsedSpace"`
	NonZeroBlocks                int64  `json:"nonZeroBlocks"`
	PeakActiveSessions           int64  `json:"peakActiveSessions"`
	PeakIOPS                     int64  `json:"peakIOPS"`
	ProvisionedSpace             int64  `json:"provisionedSpace"`
	SnapshotNonZeroBlocks        int64  `json:"snapshotNonZeroBlocks"`
	Timestamp                    string `json:"timestamp"`
	TotalOps                     int64  `json:"totalOps"`
	UniqueBlocks                 int64  `json:"uniqueBlocks"`
	UniqueBlocksUsedSpace        int64  `json:"uniqueBlocksUsedSpace"`
	UsedMetadataSpace            int64  `json:"usedMetadataSpace"`
	UsedMetadataSpaceInSnapshots int64  `json:"usedMetadataSpaceInSnapshots"`
	UsedSpace                    int64  `json:"usedSpace"`
	ZeroBlocks                   int64  `json:"zeroBlocks"`
}

type ClusterConfig struct {
	Cipi          string   `json:"cipi,omitempty"`
	Cluster       string   `json:"cluster,omitempty"`
	Ensemble      []string `json:"ensemble,omitempty"`
	Mipi          string   `json:"mipi,omitempty"`
	Name          string   `json:"name,omitempty"`
	NodeID        int64    `json:"nodeID,omitempty"`
	PendingNodeID int64    `json:"pendingNodeID,omitempty"`
	Role          string   `json:"role,omitempty"`
	Sipi          string   `json:"sipi,omitempty"`
	State         string   `json:"state,omitempty"`
}

type ClusterFaultInfo struct {
	Severity            string      `json:"severity"`
	Type                string      `json:"type"`
	Code                string      `json:"code"`
	Details             string      `json:"details"`
	NodeHardwareFaultID int64       `json:"nodeHardwareFaultID"`
	NodeID              int64       `json:"nodeID"`
	ServiceID           int64       `json:"serviceID"`
	DriveID             int64       `json:"driveID"`
	Resolved            bool        `json:"resolved"`
	ClusterFaultID      int64       `json:"clusterFaultID"`
	Date                string      `json:"date"`
	ResolvedDate        string      `json:"resolvedDate"`
	Data                interface{} `json:"data"`
}

type ClusterInfo struct {
	Attributes            interface{} `json:"attributes"`
	EncryptionAtRestState string      `json:"encryptionAtRestState"`
	Ensemble              []string    `json:"ensemble"`
	Mvip                  string      `json:"mvip"`
	MvipNodeID            int64       `json:"mvipNodeID"`
	Name                  string      `json:"name"`
	RepCount              int64       `json:"repCount"`
	State                 string      `json:"state"`
	Svip                  string      `json:"svip"`
	SvipNodeID            int64       `json:"svipNodeID"`
	UniqueID              string      `json:"uniqueID"`
	Uuid                  string      `json:"uuid"`
}

type ClusterStats struct {
	ClusterUtilization float64 `json:"clusterUtilization"`
	ClientQueueDepth   int64   `json:"clientQueueDepth"`
	ReadBytes          int64   `json:"readBytes"`
	ReadOps            int64   `json:"readOps"`
	Timestamp          string  `json:"timestamp"`
	WriteBytes         int64   `json:"writeBytes"`
	WriteOps           int64   `json:"writeOps"`
}

type ClusterVersionInfo struct {
	NodeID               int64  `json:"nodeID"`
	NodeVersion          string `json:"nodeVersion"`
	NodeInternalRevision string `json:"nodeInternalRevision"`
}

type Config struct {
	Cluster ClusterConfig `json:"cluster"`
	Network Network       `json:"network"`
}

type DriveHardware struct {
	CanonicalName          string `json:"canonicalName"`
	Connected              bool   `json:"connected"`
	Dev                    int64  `json:"dev"`
	DevPath                string `json:"devPath"`
	DriveType              string `json:"driveType"`
	LifeRemainingPercent   int64  `json:"lifeRemainingPercent"`
	LifetimeReadBytes      int64  `json:"lifetimeReadBytes"`
	LifetimeWriteBytes     int64  `json:"lifetimeWriteBytes"`
	Name                   string `json:"name"`
	Path                   string `json:"path"`
	PathLink               string `json:"pathLink"`
	PowerOnHours           int64  `json:"powerOnHours"`
	Product                string `json:"product"`
	ReallocatedSectors     int64  `json:"reallocatedSectors"`
	ReserveCapacityPercent int64  `json:"reserveCapacityPercent"`
	ScsiCompatId           string `json:"scsiCompatId"`
	ScsiState              string `json:"scsiState"`
	SecurityAtMaximum      bool   `json:"securityAtMaximum"`
	SecurityEnabled        bool   `json:"securityEnabled"`
	SecurityFrozen         bool   `json:"securityFrozen"`
	SecurityLocked         bool   `json:"securityLocked"`
	SecuritySupported      bool   `json:"securitySupported"`
	Serial                 string `json:"serial"`
	Size                   int64  `json:"size"`
	Slot                   int64  `json:"slot"`
	SmartSsdWriteCapable   bool   `json:"smartSsdWriteCapable,omitempty"`
	Uuid                   string `json:"uuid"`
	Vendor                 string `json:"vendor"`
	Version                string `json:"version"`
}

type DriveHardwareInfo struct {
	Description              string `json:"description"`
	Dev                      string `json:"dev"`
	Devpath                  string `json:"devpath"`
	DriveSecurityAtMaximum   bool   `json:"driveSecurityAtMaximum"`
	DriveSecurityFrozen      bool   `json:"driveSecurityFrozen"`
	DriveSecurityLocked      bool   `json:"driveSecurityLocked"`
	Logicalname              string `json:"logicalname"`
	Product                  string `json:"product"`
	SecurityFeatureEnabled   bool   `json:"securityFeatureEnabled"`
	SecurityFeatureSupported bool   `json:"securityFeatureSupported"`
	Serial                   string `json:"serial"`
	Size                     int64  `json:"size"`
	Uuid                     string `json:"uuid"`
	Version                  string `json:"version"`
}

type DriveInfo struct {
	Capacity   int64       `json:"capacity"`
	DriveID    int64       `json:"driveID"`
	NodeID     int64       `json:"nodeID"`
	Serial     string      `json:"serial"`
	Slot       int64       `json:"slot"`
	Status     string      `json:"status"`
	Type       string      `json:"type"`
	Attributes interface{} `json:"attributes"`
}

type DriveStats struct {
	ActiveSessions         int64  `json:"activeSessions"`
	FailedDieCount         int64  `json:"failedDieCount"`
	LifeRemainingPercent   int64  `json:"lifeRemainingPercent"`
	LifetimeReadBytes      int64  `json:"lifetimeReadBytes"`
	LifetimeWriteBytes     int64  `json:"lifetimeWriteBytes"`
	PowerOnHours           int64  `json:"powerOnHours"`
	ReadBytes              int64  `json:"readBytes"`
	ReadOps                int64  `json:"readOps"`
	ReallocatedSectors     int64  `json:"reallocatedSectors"`
	ReserveCapacityPercent int64  `json:"reserveCapacityPercent"`
	Timestamp              string `json:"timestamp"`
	TotalCapacity          int64  `json:"totalCapacity"`
	UsedCapacity           int64  `json:"usedCapacity,omitempty"`
	UsedMemory             int64  `json:"usedMemory"`
	WriteBytes             int64  `json:"writeBytes"`
	WriteOps               int64  `json:"writeOps"`
}

type DrivesHardware struct {
	DriveHardware []DriveHardware `json:"driveHardware"`
}

type EventInfo struct {
	EventID       int64       `json:"eventID"`
	Severity      int64       `json:"severity"`
	EventInfoType string      `json:"eventInfoType"`
	Message       string      `json:"message"`
	ServiceID     int64       `json:"serviceID"`
	NodeID        int64       `json:"nodeID"`
	DriveID       int64       `json:"driveID"`
	TimeOfReport  string      `json:"timeOfReport"`
	TimeOfPublish string      `json:"timeOfPublish"`
	Details       interface{} `json:"details"`
}

type FibreChannelPortInfo struct {
	Firmware  string `json:"firmware"`
	HbaPort   int64  `json:"hbaPort"`
	Model     string `json:"model"`
	NPortID   string `json:"nPortID"`
	PciSlot   int64  `json:"pciSlot"`
	Serial    string `json:"serial"`
	Speed     string `json:"speed"`
	State     string `json:"state"`
	SwitchWwn string `json:"switchWwn"`
	Wwnn      string `json:"wwnn"`
	Wwpn      string `json:"wwpn"`
}

type FibreChannelPortList struct {
	FibreChannelPorts []FibreChannelPortInfo `json:"fibreChannelPorts"`
}

type FibreChannelSession struct {
	InitiatorWWPN       string `json:"initiatorWWPN"`
	NodeID              int64  `json:"nodeID"`
	ServiceID           int64  `json:"serviceID"`
	TargetWWPN          string `json:"targetWWPN"`
	VolumeAccessGroupID int64  `json:"volumeAccessGroupID"`
}

type GroupCloneVolumeMember struct {
	VolumeID    int64 `json:"volumeID"`
	SrcVolumeID int64 `json:"srcVolumeID"`
}

type GroupSnapshot struct {
	GroupSnapshotID   int64                  `json:"groupSnapshotID"`
	GroupSnapshotUUID string                 `json:"groupSnapshotUUID"`
	Members           []GroupSnapshotMembers `json:"members"`
	Name              string                 `json:"name"`
	CreateTime        string                 `json:"createTime"`
	Status            string                 `json:"status"`
	Attributes        interface{}            `json:"attributes"`
}

type GroupSnapshotMembers struct {
	VolumeID     int64  `json:"volumeID"`
	SnapshotID   int64  `json:"snapshotID"`
	SnapshotUUID string `json:"snapshotUUID"`
	Checksum     string `json:"checksum"`
}

type ISCSISession struct {
	AccountID        int64  `json:"accountID"`
	AccountName      string `json:"accountName"`
	DriveID          int64  `json:"driveID"`
	InitiatorIP      string `json:"initiatorIP"`
	InitiatorName    string `json:"initiatorName"`
	NodeID           int64  `json:"nodeID"`
	ServiceID        int64  `json:"serviceID"`
	SessionID        int64  `json:"sessionID"`
	TargetName       string `json:"targetName"`
	TargetIP         string `json:"targetIP"`
	VirtualNetworkID string `json:"virtualNetworkID"`
	VolumeID         int64  `json:"volumeID"`
}

type LdapConfiguration struct {
	AuthType                string   `json:"authType"`
	Enabled                 bool     `json:"enabled"`
	GroupSearchBaseDN       string   `json:"groupSearchBaseDN"`
	GroupSearchCustomFilter string   `json:"groupSearchCustomFilter"`
	GroupSearchType         string   `json:"groupSearchType"`
	SearchBindDN            string   `json:"searchBindDN"`
	ServerURIs              []string `json:"serverURIs"`
	UserDNTemplate          string   `json:"userDNTemplate"`
	UserSearchBaseDN        string   `json:"userSearchBaseDN"`
	UserSearchFilter        string   `json:"userSearchFilter"`
}

type LunAssignment struct {
	VolumeID int64 `json:"volumeID"`
	Lun      int64 `json:"lun"`
}

type MetadataHosts struct {
	DeadSecondaries []int64 `json:"deadSecondaries"`
	LiveSecondaries []int64 `json:"liveSecondaries"`
	Primary         int64   `json:"primary"`
}

type Network struct {
	Bond10G NetworkConfig `json:"Bond10G,omitempty"`
	Bond1G  NetworkConfig `json:"Bond1G,omitempty"`
}

type NetworkConfig struct {
	Default              bool            `json:"#default,omitempty"`
	Address              string          `json:"address,omitempty"`
	Auto                 bool            `json:"auto,omitempty"`
	Bonddowndelay        int64           `json:"bond-downdelay,omitempty"`
	Bondfail_over_mac    string          `json:"bond-fail_over_mac,omitempty"`
	Bondprimary_reselect string          `json:"bond-primary_reselect,omitempty"`
	Bondlacp_rate        string          `json:"bond-lacp_rate,omitempty"`
	Bondmiimon           int64           `json:"bond-miimon,omitempty"`
	Bondmode             string          `json:"bond-mode,omitempty"`
	Bondslaves           string          `json:"bond-slaves,omitempty"`
	Bondupdelay          int64           `json:"bond-updelay,omitempty"`
	Broadcast            string          `json:"broadcast,omitempty"`
	Dnsnameservers       string          `json:"dns-nameservers,omitempty"`
	Dnssearch            string          `json:"dns-search,omitempty"`
	Family               string          `json:"family,omitempty"`
	Gateway              string          `json:"gateway,omitempty"`
	MacAddress           string          `json:"macAddress,omitempty"`
	MacAddressPermanent  string          `json:"macAddressPermanent,omitempty"`
	Method               string          `json:"method,omitempty"`
	Mtu                  string          `json:"mtu,omitempty"`
	Netmask              string          `json:"netmask,omitempty"`
	Network              string          `json:"network,omitempty"`
	Physical             PhysicalAdapter `json:"physical,omitempty"`
	Routes               []string        `json:"routes,omitempty"`
	Status               string          `json:"status,omitempty"`
	SymmetricRouteRules  []string        `json:"symmetricRouteRules,omitempty"`
	UpAndRunning         bool            `json:"upAndRunning,omitempty"`
}

type NewDrive struct {
	DriveID int64 `json:"driveID"`
}

type Node struct {
	NodeID                      int64       `json:"nodeID"`
	AssociatedMasterServiceID   int64       `json:"associatedMasterServiceID"`
	AssociatedFServiceID        int64       `json:"associatedFServiceID"`
	FibreChannelTargetPortGroup string      `json:"fibreChannelTargetPortGroup"`
	Name                        string      `json:"name"`
	PlatformInfo                Platform    `json:"platformInfo"`
	SoftwareVersion             string      `json:"softwareVersion"`
	Cip                         string      `json:"cip"`
	Cipi                        string      `json:"cipi"`
	Mip                         string      `json:"mip"`
	Mipi                        string      `json:"mipi"`
	Sip                         string      `json:"sip"`
	Sipi                        string      `json:"sipi"`
	Uuid                        string      `json:"uuid"`
	Attributes                  interface{} `json:"attributes"`
}

type NodeDriveHardware struct {
	NodeID int64          `json:"nodeID"`
	Result DrivesHardware `json:"result"`
}

type NodeStatsInfo struct {
	CBytesIn                  int64  `json:"cBytesIn"`
	CBytesOut                 int64  `json:"cBytesOut"`
	Cpu                       int64  `json:"cpu"`
	MBytesIn                  int64  `json:"mBytesIn"`
	MBytesOut                 int64  `json:"mBytesOut"`
	NetworkUtilizationCluster int64  `json:"networkUtilizationCluster"`
	NetworkUtilizationStorage int64  `json:"networkUtilizationStorage"`
	NodeID                    int64  `json:"nodeID"`
	SBytesIn                  int64  `json:"sBytesIn"`
	SBytesOut                 int64  `json:"sBytesOut"`
	Timestamp                 string `json:"timestamp"`
	UsedMemory                int64  `json:"usedMemory"`
}

type NodeStatsNodes struct {
	Nodes []NodeStatsInfo `json:"nodes"`
}

type PairedCluster struct {
	ClusterName     string `json:"clusterName"`
	ClusterPairID   int64  `json:"clusterPairID"`
	ClusterPairUUID string `json:"clusterPairUUID"`
	Latency         int64  `json:"latency"`
	Mvip            string `json:"mvip"`
	Status          string `json:"status"`
	Version         string `json:"version"`
}

type PendingNode struct {
	PendingNodeID   int64    `json:"pendingNodeID"`
	AssignedNodeID  int64    `json:"AssignedNodeID"`
	Name            string   `json:"name"`
	Compatible      bool     `json:"compatible"`
	PlatformInfo    Platform `json:"platformInfo"`
	Cip             string   `json:"cip"`
	Cipi            string   `json:"cipi"`
	Mip             string   `json:"mip"`
	Mipi            string   `json:"mipi"`
	Sip             string   `json:"sip"`
	Sipi            string   `json:"sipi"`
	SoftwareVersion string   `json:"softwareVersion"`
	Uuid            string   `json:"uuid"`
}

type PhysicalAdapter struct {
	Address             string `json:"address,omitempty"`
	MacAddress          string `json:"macAddress,omitempty"`
	MacAddressPermanent string `json:"macAddressPermanent,omitempty"`
	Mtu                 string `json:"mtu,omitempty"`
	Netmask             string `json:"netmask,omitempty"`
	Network             string `json:"network,omitempty"`
	UpAndRunning        bool   `json:"upAndRunning,omitempty"`
}

type Platform struct {
	NodeType     string `json:"nodeType"`
	ChassisType  string `json:"chassisType"`
	CpuModel     string `json:"cpuModel"`
	NodeMemoryGB int64  `json:"nodeMemoryGB"`
}

type QoS struct {
	MinIOPS   int64 `json:"minIOPS,omitempty"`
	MaxIOPS   int64 `json:"maxIOPS,omitempty"`
	BurstIOPS int64 `json:"burstIOPS,omitempty"`
	BurstTime int64 `json:"burstTime,omitempty"`
}

type RemoteReplication struct {
	Mode                string              `json:"mode"`
	PauseLimit          int64               `json:"pauseLimit"`
	RemoteServiceID     int64               `json:"remoteServiceID"`
	ResumeDetails       string              `json:"resumeDetails"`
	SnapshotReplication SnapshotReplication `json:"snapshotReplication"`
	State               string              `json:"state"`
	StateDetails        string              `json:"stateDetails"`
}

type ResetDriveDetails struct {
	Drive      string `json:"drive"`
	ReturnCode int64  `json:"returnCode"`
	Stderr     string `json:"stderr"`
	Stdout     string `json:"stdout"`
}

type ResetDrivesDetails struct {
	Drives []ResetDriveDetails `json:"drives"`
}

type Schedule struct {
	Frequency        Frequency    `json:"frequency"`
	HasError         bool         `json:"hasError,omitempty"`
	LastRunStatus    string       `json:"lastRunStatus"`
	LastRunTimeStart string       `json:"lastRunTimeStart"`
	Paused           bool         `json:"paused,omitempty"`
	Recurring        bool         `json:"recurring,omitempty"`
	RunNextInterval  bool         `json:"runNextInterval,omitempty"`
	ScheduleID       int64        `json:"scheduleID,omitempty"`
	ScheduleInfo     ScheduleInfo `json:"scheduleInfo"`
	Name             string       `json:"name"`
	StartingDate     string       `json:"startingDate"`
	ToBeDeleted      bool         `json:"toBeDeleted,omitempty"`
}

type ScheduleInfo struct {
	VolumeIDs               []int64 `json:"volumeIDs,omitempty"`
	SnapshotName            string  `json:"snapshotName,omitempty"`
	EnableRemoteReplication bool    `json:"enableRemoteReplication,omitempty"`
	Retention               string  `json:"retention,omitempty"`
}

type Snapshot struct {
	SnapshotID              int64       `json:"snapshotID"`
	VolumeID                int64       `json:"volumeID"`
	Name                    string      `json:"name"`
	Checksum                string      `json:"checksum"`
	EnableRemoteReplication bool        `json:"enableRemoteReplication"`
	ExpirationReason        string      `json:"expirationReason"`
	ExpirationTime          string      `json:"expirationTime"`
	RemoteStatuses          string      `json:"remoteStatuses"`
	Status                  string      `json:"status"`
	SnapshotUUID            string      `json:"snapshotUUID"`
	TotalSize               int64       `json:"totalSize"`
	GroupID                 int64       `json:"groupID,omitempty"`
	GroupSnapshotUUID       string      `json:"groupSnapshotUUID"`
	CreateTime              string      `json:"createTime"`
	Attributes              interface{} `json:"attributes"`
}

type SnapshotReplication struct {
	State        string `json:"state"`
	StateDetails string `json:"stateDetails"`
}

type SnmpNetwork struct {
	Access    string `json:"access"`
	Cidr      int64  `json:"cidr"`
	Community string `json:"community"`
	Network   string `json:"network"`
}

type SnmpTrapRecipient struct {
	Host      string `json:"host"`
	Community string `json:"community"`
	Port      int64  `json:"port"`
}

type SnmpV3UsmUser struct {
	Access     string `json:"access"`
	Name       string `json:"name"`
	Password   string `json:"password"`
	Passphrase string `json:"passphrase"`
	SecLevel   string `json:"secLevel"`
}

type SoftwareVersionInfo struct {
	CurrentVersion string `json:"currentVersion"`
	NodeID         int64  `json:"nodeID"`
	PackageName    string `json:"packageName"`
	PendingVersion string `json:"pendingVersion"`
	StartTime      string `json:"startTime"`
}

type TestConnectEnsembleDetails struct {
	Nodes interface{} `json:"nodes"`
}

type TestConnectMvipDetails struct {
	PingBytes interface{} `json:"pingBytes"`
	Mvip      string      `json:"mvip"`
	Connected bool        `json:"connected"`
}

type TestConnectSvipDetails struct {
	PingBytes interface{} `json:"pingBytes"`
	Svip      string      `json:"svip"`
	Connected bool        `json:"connected"`
}

type VirtualNetwork struct {
	VirtualNetworkID  int64          `json:"virtualNetworkID"`
	VirtualNetworkTag int64          `json:"virtualNetworkTag"`
	AddressBlocks     []AddressBlock `json:"addressBlocks"`
	Name              string         `json:"name"`
	Netmask           string         `json:"netmask"`
	Svip              string         `json:"svip"`
	Gateway           string         `json:"gateway,omitempty"`
	Namespace         bool           `json:"namespace,omitempty"`
	Attributes        interface{}    `json:"attributes"`
}

type Volume struct {
	VolumeID           int64        `json:"volumeID"`
	Name               string       `json:"name"`
	AccountID          int64        `json:"accountID"`
	CreateTime         string       `json:"createTime"`
	Status             string       `json:"status"`
	Access             string       `json:"access"`
	Enable512e         bool         `json:"enable512e"`
	Iqn                string       `json:"iqn"`
	ScsiEUIDeviceID    string       `json:"scsiEUIDeviceID"`
	ScsiNAADeviceID    string       `json:"scsiNAADeviceID"`
	Qos                VolumeQOS    `json:"qos"`
	VolumeAccessGroups []int64      `json:"volumeAccessGroups"`
	VolumePairs        []VolumePair `json:"volumePairs"`
	DeleteTime         string       `json:"deleteTime,omitempty"`
	PurgeTime          string       `json:"purgeTime,omitempty"`
	SliceCount         int64        `json:"sliceCount"`
	TotalSize          int64        `json:"totalSize"`
	BlockSize          int64        `json:"blockSize"`
	VirtualVolumeID    string       `json:"virtualVolumeID"`
	Attributes         interface{}  `json:"attributes"`
}

type VolumeAccessGroup struct {
	VolumeAccessGroupID int64    `json:"volumeAccessGroupID"`
	Name                string   `json:"name"`
	Initiators          []string `json:"initiators"`
	Volumes             []int64  `json:"volumes"`
}

type VolumeAccessGroupLunAssignments struct {
	VolumeAccessGroupID   int64           `json:"volumeAccessGroupID"`
	LunAssignments        []LunAssignment `json:"lunAssignments"`
	DeletedLunAssignments []LunAssignment `json:"deletedLunAssignments"`
}

type VolumePair struct {
	ClusterPairID     int64             `json:"clusterPairID"`
	RemoteVolumeID    int64             `json:"remoteVolumeID"`
	RemoteSliceID     int64             `json:"remoteSliceID"`
	RemoteVolumeName  string            `json:"remoteVolumeName"`
	VolumePairUUID    string            `json:"volumePairUUID"`
	RemoteReplication RemoteReplication `json:"remoteReplication"`
}

type VolumeQOS struct {
	MinIOPS   int64 `json:"minIOPS"`
	MaxIOPS   int64 `json:"maxIOPS"`
	BurstIOPS int64 `json:"burstIOPS"`
	BurstTime int64 `json:"burstTime"`
	Curve     int64 `json:"curve"`
}

type VolumeStats struct {
	AccountID            int64         `json:"accountID"`
	ActualIOPS           int64         `json:"actualIOPS"`
	AsyncDelay           string        `json:"asyncDelay,omitempty"`
	AverageIOPSize       int64         `json:"averageIOPSize"`
	BurstIOPSCredit      int64         `json:"burstIOPSCredit"`
	ClientQueueDepth     int64         `json:"clientQueueDepth"`
	DesiredMetadataHosts MetadataHosts `json:"desiredMetadataHosts"`
	LatencyUSec          int64         `json:"latencyUSec"`
	MetadataHosts        MetadataHosts `json:"metadataHosts"`
	NonZeroBlocks        int64         `json:"nonZeroBlocks"`
	ReadBytes            int64         `json:"readBytes"`
	ReadLatencyUSec      int64         `json:"readLatencyUSec"`
	ReadOps              int64         `json:"readOps"`
	Throttle             float64       `json:"throttle"`
	Timestamp            string        `json:"timestamp"`
	TotalLatencyUSec     int64         `json:"totalLatencyUSec"`
	UnalignedReads       int64         `json:"unalignedReads"`
	UnalignedWrites      int64         `json:"unalignedWrites"`
	VolumeAccessGroups   []int64       `json:"volumeAccessGroups"`
	VolumeID             int64         `json:"volumeID"`
	VolumeSize           int64         `json:"volumeSize"`
	VolumeUtilization    float64       `json:"volumeUtilization"`
	WriteBytes           int64         `json:"writeBytes"`
	WriteLatencyUSec     int64         `json:"writeLatencyUSec"`
	WriteOps             int64         `json:"writeOps"`
	ZeroBlocks           int64         `json:"zeroBlocks"`
}
