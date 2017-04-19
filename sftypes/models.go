package sftypes

type Account struct {
	AccountID          int64       `json:"accountID"`
	Username           string      `json:"username"`
	Status             string      `json:"status"`
	Volumes            []int64     `json:"volumes"`
	InitiatorSecret    CHAPSecret  `json:"initiatorSecret,omitempty"`
	TargetSecret       CHAPSecret  `json:"targetSecret,omitempty"`
	StorageContainerID string      `json:"storageContainerID,omitempty"`
	Attributes         interface{} `json:"attributes,omitempty"`
}

type AddedNode struct {
	NodeID          int64    `json:"nodeID,omitempty"`
	PendingNodeID   int64    `json:"pendingNodeID"`
	ActiveNodeKey   string   `json:"activeNodeKey,omitempty"`
	AssignedNodeID  int64    `json:"assignedNodeID,omitempty"`
	AsyncHandle     int64    `json:"asyncHandle,omitempty"`
	Cip             string   `json:"cip,omitempty"`
	Mip             string   `json:"mip,omitempty"`
	PlatformInfo    Platform `json:"platformInfo,omitempty"`
	Sip             string   `json:"sip,omitempty"`
	SoftwareVersion string   `json:"softwareVersion,omitempty"`
}

type AddressBlock struct {
	Start     string `json:"start"`
	Size      int64  `json:"size"`
	Available string `json:"available"`
}

type AsyncHandle struct {
	AsyncResultID  int64       `json:"asyncResultID"`
	Completed      bool        `json:"completed"`
	CreateTime     string      `json:"createTime"`
	LastUpdateTime string      `json:"lastUpdateTime"`
	ResultType     string      `json:"resultType"`
	Success        bool        `json:"success"`
	Data           interface{} `json:"data"`
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
	Script          string      `json:"script,omitempty"`
	SnapshotID      int64       `json:"snapshotID,omitempty"`
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
	AuthMethod     string      `json:"authMethod"`
	Access         []string    `json:"access"`
	ClusterAdminID int64       `json:"clusterAdminID"`
	Username       string      `json:"username"`
	Attributes     interface{} `json:"attributes,omitempty"`
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
	Cipi              string   `json:"cipi,omitempty"`
	Cluster           string   `json:"cluster,omitempty"`
	Ensemble          []string `json:"ensemble,omitempty"`
	Mipi              string   `json:"mipi,omitempty"`
	Name              string   `json:"name,omitempty"`
	NodeID            int64    `json:"nodeID,omitempty"`
	PendingNodeID     int64    `json:"pendingNodeID,omitempty"`
	Role              string   `json:"role,omitempty"`
	Sipi              string   `json:"sipi,omitempty"`
	State             string   `json:"state,omitempty"`
	EncryptionCapable bool     `json:"encryptionCapable,omitempty"`
	HasLocalAdmin     bool     `json:"hasLocalAdmin,omitempty"`
	Version           string   `json:"version,omitempty"`
}

type ClusterFaultInfo struct {
	DriveIDs            []int64     `json:"driveIDs,omitempty"`
	NetworkInterface    string      `json:"networkInterface,omitempty"`
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
	Data                interface{} `json:"data,omitempty"`
}

type ClusterHardwareInfo struct {
	Drives DriveHardwareInfo `json:"drives"`
	Nodes  interface{}       `json:"nodes"`
}

type ClusterInfo struct {
	MvipInterface         string      `json:"mvipInterface,omitempty"`
	MvipVlanTag           string      `json:"mvipVlanTag,omitempty"`
	SvipInterface         string      `json:"svipInterface,omitempty"`
	SvipVlanTag           string      `json:"svipVlanTag,omitempty"`
	EncryptionAtRestState string      `json:"encryptionAtRestState"`
	Ensemble              []string    `json:"ensemble"`
	Mvip                  string      `json:"mvip"`
	MvipNodeID            int64       `json:"mvipNodeID"`
	Name                  string      `json:"name"`
	RepCount              int64       `json:"repCount"`
	Svip                  string      `json:"svip"`
	SvipNodeID            int64       `json:"svipNodeID"`
	UniqueID              string      `json:"uniqueID"`
	Uuid                  string      `json:"uuid"`
	Attributes            interface{} `json:"attributes"`
}

type ClusterStats struct {
	ClusterUtilization   float64 `json:"clusterUtilization"`
	ClientQueueDepth     int64   `json:"clientQueueDepth"`
	ReadBytes            int64   `json:"readBytes"`
	ReadOps              int64   `json:"readOps"`
	Timestamp            string  `json:"timestamp"`
	WriteBytes           int64   `json:"writeBytes"`
	WriteOps             int64   `json:"writeOps"`
	ActualIOPS           int64   `json:"actualIOPS,omitempty"`
	AverageIOPSize       int64   `json:"averageIOPSize,omitempty"`
	LatencyUSec          int64   `json:"latencyUSec,omitempty"`
	ReadBytesLastSample  int64   `json:"readBytesLastSample,omitempty"`
	ReadLatencyUSec      int64   `json:"readLatencyUSec,omitempty"`
	ReadOpsLastSample    int64   `json:"readOpsLastSample,omitempty"`
	SamplePeriodMsec     int64   `json:"samplePeriodMsec,omitempty"`
	UnalignedReads       int64   `json:"unalignedReads,omitempty"`
	UnalignedWrites      int64   `json:"unalignedWrites,omitempty"`
	WriteBytesLastSample int64   `json:"writeBytesLastSample,omitempty"`
	WriteLatencyUSec     int64   `json:"writeLatencyUSec,omitempty"`
	WriteOpsLastSample   int64   `json:"writeOpsLastSample,omitempty"`
}

type ClusterVersionInfo struct {
	NodeID               int64  `json:"nodeID"`
	NodeVersion          string `json:"nodeVersion"`
	NodeInternalRevision string `json:"nodeInternalRevision"`
}

type Config struct {
	Cluster ClusterConfig `json:"cluster"`
	Network NetworkParams `json:"network"`
}

type CreateInitiator struct {
	Name                string      `json:"name"`
	Alias               string      `json:"alias,omitempty"`
	VolumeAccessGroupID int64       `json:"volumeAccessGroupID,omitempty"`
	Attributes          interface{} `json:"attributes,omitempty"`
}

type DetailedService struct {
	Service Service `json:"service"`
	Node    Node    `json:"node"`
	Drive   Drive   `json:"drive,omitempty"`
	Drives  []Drive `json:"drives"`
}

type Drive struct {
	DriveID                   int64       `json:"driveID"`
	NodeID                    int64       `json:"nodeID"`
	AssignedService           int64       `json:"assignedService,omitempty"`
	AsyncResultIDs            []int64     `json:"asyncResultIDs"`
	Capacity                  int64       `json:"capacity"`
	Serial                    string      `json:"serial"`
	Slot                      int64       `json:"slot,omitempty"`
	DriveStatus               string      `json:"driveStatus"`
	DriveType                 string      `json:"driveType"`
	ReservedSliceFileCapacity int64       `json:"reservedSliceFileCapacity,omitempty"`
	CustomerSliceFileCapacity int64       `json:"customerSliceFileCapacity,omitempty"`
	Attributes                interface{} `json:"attributes"`
}

type DriveConfigInfo struct {
	CanonicalName        string `json:"canonicalName"`
	Connected            bool   `json:"connected"`
	Dev                  int64  `json:"dev"`
	DevPath              string `json:"devPath"`
	DriveType            string `json:"driveType"`
	Product              string `json:"product"`
	Name                 string `json:"name"`
	Path                 string `json:"path"`
	PathLink             string `json:"pathLink"`
	ScsiCompatId         string `json:"scsiCompatId"`
	SmartSsdWriteCapable bool   `json:"smartSsdWriteCapable,omitempty"`
	SecurityEnabled      bool   `json:"securityEnabled"`
	SecurityFrozen       bool   `json:"securityFrozen"`
	SecurityLocked       bool   `json:"securityLocked"`
	SecuritySupported    bool   `json:"securitySupported"`
	Size                 int64  `json:"size"`
	Slot                 int64  `json:"slot"`
	Uuid                 string `json:"uuid"`
	Vendor               string `json:"vendor"`
	Version              string `json:"version"`
	SecurityAtMaximum    bool   `json:"securityAtMaximum"`
	Serial               string `json:"serial"`
	ScsiState            string `json:"scsiState"`
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
	ScsiCompatID             string `json:"scsiCompatID"`
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
	ActiveSessions         int64  `json:"activeSessions,omitempty"`
	DriveID                int64  `json:"driveID,omitempty"`
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

type DrivesConfigInfo struct {
	Drives           []DriveConfigInfo `json:"drives"`
	NumBlockActual   int64             `json:"numBlockActual"`
	NumBlockExpected int64             `json:"numBlockExpected"`
	NumSliceActual   int64             `json:"numSliceActual"`
	NumSliceExpected int64             `json:"numSliceExpected"`
	NumTotalActual   int64             `json:"numTotalActual"`
	NumTotalExpected int64             `json:"numTotalExpected"`
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
	DriveIDs      []int64     `json:"driveIDs"`
	TimeOfReport  string      `json:"timeOfReport"`
	TimeOfPublish string      `json:"timeOfPublish"`
	Details       interface{} `json:"details,omitempty"`
}

type FeatureObject struct {
	Enabled bool   `json:"enabled"`
	Feature string `json:"feature"`
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
	VolumeAccessGroupID int64  `json:"volumeAccessGroupID,omitempty"`
}

type GetIpmiInfoNodesResultObject struct {
	IpmiInfo IpmiInfo `json:"ipmiInfo"`
}

type GetOriginNode struct {
	NodeID int64               `json:"nodeID"`
	Result GetOriginNodeResult `json:"result"`
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
	VolumeID                int64         `json:"volumeID"`
	SnapshotID              int64         `json:"snapshotID"`
	SnapshotUUID            string        `json:"snapshotUUID"`
	Checksum                string        `json:"checksum"`
	Attributes              interface{}   `json:"attributes,omitempty"`
	CreateTime              string        `json:"createTime,omitempty"`
	EnableRemoteReplication bool          `json:"enableRemoteReplication,omitempty"`
	ExpirationReason        string        `json:"expirationReason,omitempty"`
	ExpirationTime          string        `json:"expirationTime,omitempty"`
	GroupID                 int64         `json:"groupID,omitempty"`
	GroupSnapshotUUID       string        `json:"groupSnapshotUUID,omitempty"`
	Name                    string        `json:"name,omitempty"`
	RemoteStatus            string        `json:"remoteStatus,omitempty"`
	RemoteStatuses          []interface{} `json:"remoteStatuses,omitempty"`
	Status                  string        `json:"status,omitempty"`
	TotalSize               int64         `json:"totalSize,omitempty"`
	VirtualVolumeID         int64         `json:"virtualVolumeID,omitempty"`
	VolumePairUUID          string        `json:"volumePairUUID,omitempty"`
}

type ISCSISession struct {
	DriveIDs           []int64   `json:"driveIDs,omitempty"`
	AccountID          int64     `json:"accountID"`
	Initiator          Initiator `json:"initiator,omitempty"`
	AccountName        string    `json:"accountName"`
	DriveID            int64     `json:"driveID"`
	InitiatorIP        string    `json:"initiatorIP"`
	InitiatorPortName  string    `json:"initiatorPortName"`
	TargetPortName     string    `json:"targetPortName"`
	InitiatorName      string    `json:"initiatorName"`
	NodeID             int64     `json:"nodeID"`
	ServiceID          int64     `json:"serviceID"`
	SessionID          int64     `json:"sessionID"`
	TargetName         string    `json:"targetName"`
	TargetIP           string    `json:"targetIP"`
	VirtualNetworkID   int64     `json:"virtualNetworkID"`
	VolumeID           int64     `json:"volumeID"`
	CreateTime         string    `json:"createTime"`
	VolumeInstance     int64     `json:"volumeInstance"`
	InitiatorSessionID int64     `json:"initiatorSessionID"`
}

type Initiator struct {
	Alias              string      `json:"alias"`
	InitiatorID        int64       `json:"initiatorID"`
	InitiatorName      string      `json:"initiatorName"`
	VolumeAccessGroups []int64     `json:"volumeAccessGroups"`
	Attributes         interface{} `json:"attributes"`
}

type IpmiInfo struct {
	Sensors []interface{} `json:"sensors"`
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

type LoggingServer struct {
	Host string `json:"host"`
	Port int64  `json:"port"`
}

type LoginSessionInfo struct {
	Timeout string `json:"timeout"`
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

type ModifyInitiator struct {
	InitiatorID         int64       `json:"initiatorID"`
	Alias               string      `json:"alias,omitempty"`
	VolumeAccessGroupID int64       `json:"volumeAccessGroupID,omitempty"`
	Attributes          interface{} `json:"attributes,omitempty"`
}

type Network struct {
	Bond10G NetworkConfig `json:"Bond10G,omitempty"`
	Bond1G  NetworkConfig `json:"Bond1G,omitempty"`
	Eth0    NetworkConfig `json:"eth0,omitempty"`
	Eth1    NetworkConfig `json:"eth1,omitempty"`
	Eth2    NetworkConfig `json:"eth2,omitempty"`
	Eth3    NetworkConfig `json:"eth3,omitempty"`
	Lo      NetworkConfig `json:"lo,omitempty"`
}

type NetworkConfig struct {
	Default              bool            `json:"#default,omitempty"`
	Bondmaster           string          `json:"bond-master,omitempty"`
	VirtualNetworkTag    string          `json:"virtualNetworkTag,omitempty"`
	Address              string          `json:"address,omitempty"`
	Auto                 bool            `json:"auto,omitempty"`
	Bonddowndelay        string          `json:"bond-downdelay,omitempty"`
	Bondfail_over_mac    string          `json:"bond-fail_over_mac,omitempty"`
	Bondprimary_reselect string          `json:"bond-primary_reselect,omitempty"`
	Bondlacp_rate        string          `json:"bond-lacp_rate,omitempty"`
	Bondmiimon           string          `json:"bond-miimon,omitempty"`
	Bondmode             string          `json:"bond-mode,omitempty"`
	Bondslaves           string          `json:"bond-slaves,omitempty"`
	Bondupdelay          string          `json:"bond-updelay,omitempty"`
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
	Routes               []interface{}   `json:"routes,omitempty"`
	Status               string          `json:"status,omitempty"`
	SymmetricRouteRules  []string        `json:"symmetricRouteRules,omitempty"`
	UpAndRunning         bool            `json:"upAndRunning,omitempty"`
	Bondxmit_hash_policy string          `json:"bond-xmit_hash_policy,omitempty"`
	Bondad_num_ports     string          `json:"bond-ad_num_ports,omitempty"`
}

type NetworkConfigParams struct {
	Default              bool            `json:"#default,omitempty"`
	Bondmaster           string          `json:"bond-master,omitempty"`
	VirtualNetworkTag    string          `json:"virtualNetworkTag,omitempty"`
	Address              string          `json:"address,omitempty"`
	Auto                 bool            `json:"auto,omitempty"`
	Bonddowndelay        string          `json:"bond-downdelay,omitempty"`
	Bondfail_over_mac    string          `json:"bond-fail_over_mac,omitempty"`
	Bondprimary_reselect string          `json:"bond-primary_reselect,omitempty"`
	Bondlacp_rate        string          `json:"bond-lacp_rate,omitempty"`
	Bondmiimon           string          `json:"bond-miimon,omitempty"`
	Bondmode             string          `json:"bond-mode,omitempty"`
	Bondslaves           string          `json:"bond-slaves,omitempty"`
	Bondupdelay          string          `json:"bond-updelay,omitempty"`
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
	Routes               []interface{}   `json:"routes,omitempty"`
	Status               string          `json:"status,omitempty"`
	SymmetricRouteRules  []string        `json:"symmetricRouteRules,omitempty"`
	UpAndRunning         bool            `json:"upAndRunning,omitempty"`
}

type NetworkInterface struct {
	Address           string `json:"address"`
	Broadcast         string `json:"broadcast"`
	MacAddress        string `json:"macAddress"`
	Mtu               int64  `json:"mtu"`
	Name              string `json:"name"`
	Netmask           string `json:"netmask"`
	Status            string `json:"status"`
	Type              string `json:"type"`
	VirtualNetworkTag int64  `json:"virtualNetworkTag"`
	Namespace         bool   `json:"namespace,omitempty"`
}

type NetworkParams struct {
	Bond10G NetworkConfigParams `json:"Bond10G,omitempty"`
	Bond1G  NetworkConfigParams `json:"Bond1G,omitempty"`
	Eth0    NetworkConfigParams `json:"eth0,omitempty"`
	Eth1    NetworkConfigParams `json:"eth1,omitempty"`
	Eth2    NetworkConfigParams `json:"eth2,omitempty"`
	Eth3    NetworkConfigParams `json:"eth3,omitempty"`
	Lo      NetworkConfigParams `json:"lo,omitempty"`
}

type NewDrive struct {
	DriveID int64  `json:"driveID"`
	Type    string `json:"type,omitempty"`
}

type Node struct {
	NodeID                      int64                   `json:"nodeID"`
	AssociatedMasterServiceID   int64                   `json:"associatedMasterServiceID"`
	AssociatedFServiceID        int64                   `json:"associatedFServiceID"`
	FibreChannelTargetPortGroup string                  `json:"fibreChannelTargetPortGroup,omitempty"`
	Name                        string                  `json:"name"`
	PlatformInfo                Platform                `json:"platformInfo"`
	SoftwareVersion             string                  `json:"softwareVersion"`
	Cip                         string                  `json:"cip"`
	Cipi                        string                  `json:"cipi"`
	Mip                         string                  `json:"mip"`
	Mipi                        string                  `json:"mipi"`
	Sip                         string                  `json:"sip"`
	Sipi                        string                  `json:"sipi"`
	Uuid                        string                  `json:"uuid"`
	VirtualNetworks             []VirtualNetworkAddress `json:"virtualNetworks"`
	Attributes                  interface{}             `json:"attributes"`
}

type NodeDriveHardware struct {
	NodeID int64          `json:"nodeID"`
	Result DrivesHardware `json:"result"`
}

type NodeStateInfo struct {
	Cluster string `json:"cluster"`
	State   string `json:"state"`
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

type NodeWaitingToJoin struct {
	Name          string `json:"name,omitempty"`
	Version       string `json:"version"`
	NodeID        int64  `json:"nodeID,omitempty"`
	PendingNodeID int64  `json:"pendingNodeID,omitempty"`
	Mip           string `json:"mip,omitempty"`
	Cip           string `json:"cip,omitempty"`
	Sip           string `json:"sip,omitempty"`
	Compatible    bool   `json:"compatible"`
	ChassisType   string `json:"chassisType,omitempty"`
	Hostname      string `json:"hostname,omitempty"`
	NodeType      string `json:"nodeType,omitempty"`
}

type Origin struct {
	Signature        Signature `json:"&lt;signature&gt;"`
	Contractdate     string    `json:"contract-date"`
	Contractname     string    `json:"contract-name"`
	Contractquantity int64     `json:"contract-quantity"`
	Contracttype     string    `json:"contract-type"`
	Integrator       string    `json:"integrator"`
	Location         string    `json:"location"`
	Organization     string    `json:"organization"`
	Type             string    `json:"type"`
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

type PendingActiveNode struct {
	ActiveNodeKey   string   `json:"activeNodeKey"`
	AssignedNodeID  int64    `json:"assignedNodeID"`
	AsyncHandle     int64    `json:"asyncHandle"`
	Cip             string   `json:"cip"`
	Mip             string   `json:"mip"`
	PendingNodeID   int64    `json:"pendingNodeID"`
	PlatformInfo    Platform `json:"platformInfo"`
	Sip             string   `json:"sip"`
	SoftwareVersion string   `json:"softwareVersion"`
}

type PendingNode struct {
	PendingNodeID   int64    `json:"pendingNodeID"`
	AssignedNodeID  int64    `json:"assignedNodeID"`
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

type PendingOperation struct {
	Pending   bool   `json:"pending"`
	Operation string `json:"operation"`
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
	NodeType              string `json:"nodeType"`
	ChassisType           string `json:"chassisType"`
	CpuModel              string `json:"cpuModel"`
	NodeMemoryGB          int64  `json:"nodeMemoryGB"`
	PlatformConfigVersion string `json:"platformConfigVersion,omitempty"`
}

type ProtocolEndpoint struct {
	ProtocolEndpointID    string `json:"protocolEndpointID"`
	ProtocolEndpointState string `json:"protocolEndpointState"`
	ProviderType          string `json:"providerType"`
	PrimaryProviderID     int64  `json:"primaryProviderID"`
	SecondaryProviderID   int64  `json:"secondaryProviderID"`
	ScsiNAADeviceID       string `json:"scsiNAADeviceID"`
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

type ResetNodeDetails struct {
	RtfiInfo RtfiInfo `json:"rtfiInfo"`
}

type RtfiInfo struct {
	Mipi             string      `json:"mipi,omitempty"`
	Generation       interface{} `json:"generation"`
	StatusUrlLogfile string      `json:"statusUrlLogfile,omitempty"`
	Build            string      `json:"build"`
	StatusUrlAll     string      `json:"statusUrlAll"`
	GenerationNext   int64       `json:"generationNext,omitempty"`
	Mip              string      `json:"mip,omitempty"`
	StatusUrlCurrent string      `json:"statusUrlCurrent"`
	Options          interface{} `json:"options,omitempty"`
}

type Schedule struct {
	LastRunTimeStarted string       `json:"lastRunTimeStarted,omitempty"`
	HasError           bool         `json:"hasError,omitempty"`
	ScheduleInfo       ScheduleInfo `json:"scheduleInfo"`
	RunNextInterval    bool         `json:"runNextInterval,omitempty"`
	Name               string       `json:"name"`
	LastRunStatus      string       `json:"lastRunStatus,omitempty"`
	ScheduleID         int64        `json:"scheduleID,omitempty"`
	Paused             bool         `json:"paused,omitempty"`
	ToBeDeleted        bool         `json:"toBeDeleted,omitempty"`
	Frequency          Frequency    `json:"frequency"`
	StartingDate       string       `json:"startingDate,omitempty"`
	Recurring          bool         `json:"recurring,omitempty"`
}

type ScheduleInfo struct {
	SnapshotName            string  `json:"snapshotName,omitempty"`
	EnableRemoteReplication bool    `json:"enableRemoteReplication,omitempty"`
	VolumeIDs               []int64 `json:"volumeIDs,omitempty"`
	Retention               string  `json:"retention,omitempty"`
}

type Service struct {
	ServiceID        int64   `json:"serviceID"`
	ServiceType      string  `json:"serviceType"`
	NodeID           int64   `json:"nodeID"`
	AssociatedBV     int64   `json:"associatedBV,omitempty"`
	AssociatedTS     int64   `json:"associatedTS,omitempty"`
	AssociatedVS     int64   `json:"associatedVS,omitempty"`
	AsyncResultIDs   []int64 `json:"asyncResultIDs"`
	DriveID          int64   `json:"driveID,omitempty"`
	FirstTimeStartup bool    `json:"firstTimeStartup"`
	IpcPort          int64   `json:"ipcPort"`
	IscsiPort        int64   `json:"iscsiPort"`
	Status           string  `json:"status"`
	StartedDriveIDs  []int64 `json:"startedDriveIDs"`
	DriveIDs         []int64 `json:"driveIDs"`
}

type Signature struct {
	Data    string `json:"data"`
	Pubkey  string `json:"pubkey"`
	Version int64  `json:"version"`
}

type Snapshot struct {
	SnapshotID              int64                  `json:"snapshotID"`
	VolumeID                int64                  `json:"volumeID"`
	Name                    string                 `json:"name"`
	Checksum                string                 `json:"checksum"`
	EnableRemoteReplication bool                   `json:"enableRemoteReplication"`
	ExpirationReason        string                 `json:"expirationReason"`
	ExpirationTime          string                 `json:"expirationTime,omitempty"`
	RemoteStatuses          []SnapshotRemoteStatus `json:"remoteStatuses,omitempty"`
	Status                  string                 `json:"status"`
	SnapshotUUID            string                 `json:"snapshotUUID"`
	TotalSize               int64                  `json:"totalSize"`
	GroupID                 int64                  `json:"groupID,omitempty"`
	GroupSnapshotUUID       string                 `json:"groupSnapshotUUID"`
	CreateTime              string                 `json:"createTime"`
	VirtualVolumeID         string                 `json:"virtualVolumeID,omitempty"`
	Attributes              interface{}            `json:"attributes"`
}

type SnapshotRemoteStatus struct {
	RemoteStatus   string `json:"remoteStatus"`
	VolumePairUUID string `json:"volumePairUUID"`
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

type StorageContainer struct {
	Name                 string   `json:"name"`
	StorageContainerID   string   `json:"storageContainerID"`
	AccountID            int64    `json:"accountID"`
	ProtocolEndpointType string   `json:"protocolEndpointType"`
	InitiatorSecret      string   `json:"initiatorSecret"`
	TargetSecret         string   `json:"targetSecret"`
	Status               string   `json:"status"`
	VirtualVolumes       []string `json:"virtualVolumes,omitempty"`
}

type SupportBundleDetails struct {
	BundleName string   `json:"bundleName"`
	ExtraArgs  string   `json:"extraArgs"`
	Files      []string `json:"files"`
	Url        []string `json:"url"`
	Output     string   `json:"output"`
	TimeoutSec int64    `json:"timeoutSec"`
}

type SyncJob struct {
	BytesPerSecond  float64 `json:"bytesPerSecond"`
	CurrentBytes    int64   `json:"currentBytes"`
	DstServiceID    int64   `json:"dstServiceID"`
	ElapsedTime     float64 `json:"elapsedTime"`
	PercentComplete float64 `json:"percentComplete"`
	RemainingTime   float64 `json:"remainingTime"`
	SliceID         int64   `json:"sliceID"`
	SrcServiceID    int64   `json:"srcServiceID"`
	TotalBytes      int64   `json:"totalBytes"`
	Type            string  `json:"type"`
	CloneID         int64   `json:"cloneID"`
	DstVolumeID     int64   `json:"dstVolumeID"`
	NodeID          int64   `json:"nodeID"`
	SnapshotID      int64   `json:"snapshotID"`
	SrcVolumeID     int64   `json:"srcVolumeID"`
	BlocksPerSecond float64 `json:"blocksPerSecond"`
	Stage           string  `json:"stage"`
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
	Attributes        interface{}    `json:"attributes,omitempty"`
}

type VirtualNetworkAddress struct {
	VirtualNetworkID int64  `json:"virtualNetworkID"`
	Address          string `json:"address"`
}

type VirtualVolumeBinding struct {
	ProtocolEndpointID       string `json:"protocolEndpointID"`
	ProtocolEndpointInBandID string `json:"protocolEndpointInBandID"`
	ProtocolEndpointType     string `json:"protocolEndpointType"`
	VirtualVolumeBindingID   int64  `json:"virtualVolumeBindingID"`
	VirtualVolumeHostID      string `json:"virtualVolumeHostID"`
	VirtualVolumeID          string `json:"virtualVolumeID"`
	VirtualVolumeSecondaryID string `json:"virtualVolumeSecondaryID"`
}

type VirtualVolumeHost struct {
	VirtualVolumeHostID        string   `json:"virtualVolumeHostID"`
	ClusterID                  string   `json:"clusterID"`
	VisibleProtocolEndpointIDs []string `json:"visibleProtocolEndpointIDs"`
	Bindings                   []int64  `json:"bindings"`
	InitiatorNames             []string `json:"initiatorNames"`
	HostAddress                string   `json:"hostAddress"`
}

type VirtualVolumeInfo struct {
	VirtualVolumeID       string           `json:"virtualVolumeID"`
	ParentVirtualVolumeID string           `json:"parentVirtualVolumeID"`
	StorageContainer      StorageContainer `json:"storageContainer"`
	VolumeID              int64            `json:"volumeID"`
	SnapshotID            int64            `json:"snapshotID"`
	VirtualVolumeType     string           `json:"virtualVolumeType"`
	Status                string           `json:"status"`
	Bindings              []int64          `json:"bindings"`
	Children              []string         `json:"children"`
	Metadata              interface{}      `json:"metadata"`
	SnapshotInfo          Snapshot         `json:"snapshotInfo,omitempty"`
	VolumeInfo            Volume           `json:"volumeInfo,omitempty"`
	Descendants           []int64          `json:"descendants,omitempty"`
}

type VirtualVolumeStats struct {
	AccountID            int64         `json:"accountID"`
	ActualIOPS           int64         `json:"actualIOPS,omitempty"`
	AsyncDelay           string        `json:"asyncDelay,omitempty"`
	AverageIOPSize       int64         `json:"averageIOPSize,omitempty"`
	BurstIOPSCredit      int64         `json:"burstIOPSCredit,omitempty"`
	ClientQueueDepth     int64         `json:"clientQueueDepth,omitempty"`
	DesiredMetadataHosts MetadataHosts `json:"desiredMetadataHosts,omitempty"`
	LatencyUSec          int64         `json:"latencyUSec,omitempty"`
	MetadataHosts        MetadataHosts `json:"metadataHosts,omitempty"`
	NonZeroBlocks        int64         `json:"nonZeroBlocks"`
	ReadBytes            int64         `json:"readBytes"`
	ReadLatencyUSec      int64         `json:"readLatencyUSec,omitempty"`
	ReadOps              int64         `json:"readOps"`
	Throttle             float64       `json:"throttle,omitempty"`
	Timestamp            string        `json:"timestamp"`
	TotalLatencyUSec     int64         `json:"totalLatencyUSec,omitempty"`
	UnalignedReads       int64         `json:"unalignedReads"`
	UnalignedWrites      int64         `json:"unalignedWrites"`
	VolumeAccessGroups   []int64       `json:"volumeAccessGroups"`
	VolumeID             int64         `json:"volumeID"`
	VolumeSize           int64         `json:"volumeSize"`
	VolumeUtilization    float64       `json:"volumeUtilization,omitempty"`
	WriteBytes           int64         `json:"writeBytes"`
	WriteLatencyUSec     int64         `json:"writeLatencyUSec,omitempty"`
	WriteOps             int64         `json:"writeOps"`
	ZeroBlocks           int64         `json:"zeroBlocks"`
	WriteBytesLastSample int64         `json:"writeBytesLastSample,omitempty"`
	SamplePeriodMSec     int64         `json:"samplePeriodMSec,omitempty"`
	ReadBytesLastSample  int64         `json:"readBytesLastSample,omitempty"`
	ReadOpsLastSample    int64         `json:"readOpsLastSample,omitempty"`
	WriteOpsLastSample   int64         `json:"writeOpsLastSample,omitempty"`
	VirtualVolumeID      string        `json:"virtualVolumeID,omitempty"`
}

type VirtualVolumeTask struct {
	VirtualVolumeTaskID  string      `json:"virtualVolumeTaskID"`
	VirtualvolumeID      string      `json:"virtualvolumeID"`
	CloneVirtualVolumeID string      `json:"cloneVirtualVolumeID"`
	Status               string      `json:"status"`
	Operation            string      `json:"operation"`
	VirtualVolumeHostID  string      `json:"virtualVolumeHostID"`
	ParentMetadata       interface{} `json:"parentMetadata"`
	ParentTotalSize      int64       `json:"parentTotalSize"`
	ParentUsedSize       int64       `json:"parentUsedSize"`
	Cancelled            bool        `json:"cancelled"`
}

type Volume struct {
	VolumeID           int64        `json:"volumeID"`
	Name               string       `json:"name"`
	AccountID          int64        `json:"accountID"`
	CreateTime         string       `json:"createTime"`
	Status             string       `json:"status"`
	Access             string       `json:"access"`
	Enable512e         bool         `json:"enable512e"`
	Iqn                string       `json:"iqn,omitempty"`
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
	VirtualVolumeID    string       `json:"virtualVolumeID,omitempty"`
	Attributes         interface{}  `json:"attributes"`
}

type VolumeAccessGroup struct {
	DeletedVolumes      []int64     `json:"deletedVolumes"`
	VolumeAccessGroupID int64       `json:"volumeAccessGroupID"`
	Name                string      `json:"name"`
	InitiatorIDs        []int64     `json:"initiatorIDs"`
	Initiators          []string    `json:"initiators"`
	Volumes             []int64     `json:"volumes"`
	Attributes          interface{} `json:"attributes"`
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
	ActualIOPS           int64         `json:"actualIOPS,omitempty"`
	AverageIOPSize       int64         `json:"averageIOPSize,omitempty"`
	BurstIOPSCredit      int64         `json:"burstIOPSCredit,omitempty"`
	ClientQueueDepth     int64         `json:"clientQueueDepth,omitempty"`
	LatencyUSec          int64         `json:"latencyUSec,omitempty"`
	MetadataHosts        MetadataHosts `json:"metadataHosts,omitempty"`
	NonZeroBlocks        int64         `json:"nonZeroBlocks"`
	ReadBytes            int64         `json:"readBytes"`
	ReadLatencyUSec      int64         `json:"readLatencyUSec,omitempty"`
	ReadOps              int64         `json:"readOps"`
	Throttle             float64       `json:"throttle,omitempty"`
	Timestamp            string        `json:"timestamp"`
	TotalLatencyUSec     int64         `json:"totalLatencyUSec,omitempty"`
	UnalignedReads       int64         `json:"unalignedReads"`
	UnalignedWrites      int64         `json:"unalignedWrites"`
	VolumeAccessGroups   []int64       `json:"volumeAccessGroups"`
	VolumeID             int64         `json:"volumeID"`
	VolumeSize           int64         `json:"volumeSize"`
	VolumeUtilization    float64       `json:"volumeUtilization,omitempty"`
	WriteBytes           int64         `json:"writeBytes"`
	WriteLatencyUSec     int64         `json:"writeLatencyUSec,omitempty"`
	WriteOps             int64         `json:"writeOps"`
	ZeroBlocks           int64         `json:"zeroBlocks"`
	WriteBytesLastSample int64         `json:"writeBytesLastSample,omitempty"`
	SamplePeriodMSec     int64         `json:"samplePeriodMSec,omitempty"`
	ReadBytesLastSample  int64         `json:"readBytesLastSample,omitempty"`
	ReadOpsLastSample    int64         `json:"readOpsLastSample,omitempty"`
	WriteOpsLastSample   int64         `json:"writeOpsLastSample,omitempty"`
}
