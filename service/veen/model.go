package veen

type VolcResponseMetadata struct {
	RequestId string     `json:"RequestId"`
	Action    string     `json:"Action"`
	Version   string     `json:"Version"`
	Service   string     `json:"Service"`
	Region    string     `json:"Region"`
	Error     *VolcError `json:"Error,omitempty"`
}

type VolcError struct {
	Code    string `json:"Code"`
	Message string `json:"Message"`
}

type CreateCloudServerReq struct {
	// 边缘服务名称
	// in:body
	CloudServerName string `json:"cloudserver_name,omitempty" query:"cloudserver_name"`
	// 镜像 ID
	// in:body
	ImageID string `json:"image_id,omitempty" query:"image_id"`
	// 规格名称
	// in:body
	SpecName string `json:"spec_name,omitempty" query:"spec_name"`
	// 云服务描述
	// in:body
	CloudServerDesc *string `json:"cloud_server_desc,omitempty" query:"cloud_server_desc"`
	// 存储配置
	// in:body
	StorageConfig *StorageConfig `json:"storage_config,omitempty" query:"storage_config"`
	// 网络配置
	// in:body
	NetworkConfig *CloudServerNetworkConfig `json:"network_config,omitempty" query:"network_config"`
	// 密钥配置
	// in:body
	SecretConfig *SecretConfig `json:"secret_config,omitempty" query:"secret_config"`
	// 通用数据
	// in:body
	CustomData *CustomData `json:"custom_data,omitempty" query:"custom_data"`
	// 计费配置
	// in:body
	BillingConfig *CloudServerBillingConfigs `json:"billing_config,omitempty" query:"billing_config"`
	// 实例区域个数
	// in:body
	InstanceAreaNums []*InstanceAreaNum `json:"instance_area_nums,omitempty" query:"instance_area_nums"`
	// 调度策略
	// in:body
	ScheduleStrategy *ScheduleStrategy `json:"schedule_strategy,omitempty" query:"schedule_strategy"`
	// 高级配置
	// in:body
	AdvancedConfiguration *AdvancedConfiguration `json:"advanced_configuration,omitempty" query:"advanced_configuration"`
	// 外挂VPC配置
	// in:body
	AttachedPvcConfig *AttachedPvcConfig `json:"attached_pvc_config,omitempty" query:"attached_pvc_config"`
	// in:body
	DisableVga *bool `json:"disable_vga,omitempty" query:"disable_vga"`
	// in:body
	ImageBootMode *IMAGEBOOTMODE `json:"image_boot_mode,omitempty" query:"image_boot_mode"`
}

type StorageConfig struct {
	SystemDisk     *DiskSpec          `json:"system_disk"`
	DataDisk       *DiskSpec          `json:"data_disk,omitempty"`
	DataLocalDisks []*DiskSpecWithNum `json:"data_local_disks,omitempty"`
	DataDiskList   []*DiskSpec        `json:"data_disk_list,omitempty"`
}

type DiskSpec struct {
	StorageType string `json:"storage_type"`
	Capacity    string `json:"capacity"`
}

type DiskSpecWithNum struct {
	DiskSpec *DiskSpec `json:"disk_spec"`
	Num      int32     `json:"num"`
}

type CloudServerNetworkConfig struct {
	BandwidthPeak               *string `json:"bandwidth_peak,omitempty"`
	InternalBandwidthPeak       *string `json:"internal_bandwidth_peak,omitempty"`
	VfPassthrough               *bool   `json:"vf_passthrough,omitempty"`
	EnableIpv6                  *bool   `json:"enable_ipv6,omitempty"`
	DisableIpv4                 *bool   `json:"disable_ipv4,omitempty"`
	ReverseInterface            *bool   `json:"reverse_interface,omitempty"`
	CustomInternalInterfaceName *string `json:"custom_internal_interface_name,omitempty"`
	CustomExternalInterfaceName *string `json:"custom_external_interface_name,omitempty"`
	SecondaryInternalIPNum      *int32  `json:"secondary_internal_ip_num,omitempty"`
	BoundEipShareBandwidthPeak  *string `json:"bound_eip_share_bandwidth_peak,omitempty"`
}

type SecretConfig struct {
	SecretType SecretType `json:"secret_type"`
	SecretData *string    `json:"secret_data,omitempty"`
	SSHKeyName *string    `json:"ssh_key_name,omitempty"`
}

type SecretType int64

const (
	SecretType_Random SecretType = 1
	SecretType_Custom SecretType = 2
	SecretType_SSH    SecretType = 3
	SecretType_None   SecretType = 4
)

type CustomData struct {
	IsBase64 bool   `json:"is_base64"`
	Data     string `json:"data"`
}

type CloudServerBillingConfigs struct {
	ComputingBillingMethod BILLINGMETHOD `json:"computing_billing_method"`
	BandwidthBillingMethod BILLINGMETHOD `json:"bandwidth_billing_method"`
}

type BILLINGMETHOD = string

const (
	BILLINGMETHODMOUTHPEAK = "MonthlyPeak"

	BILLINGMETHODDAYPEAK = "DailyPeak"

	BILLINGMETHODMOUTHP95 = "MonthlyP95"

	BILLINGMETHODHOURUSED = "HourUsed"
)

type InstanceAreaNum struct {
	AreaName     string       `json:"area_name"`
	Isp          string       `json:"isp"`
	ClusterName  *string      `json:"cluster_name,omitempty"`
	ClusterAlias *string      `json:"cluster_alias,omitempty"`
	VpcIdentity  *string      `json:"vpc_identity,omitempty"`
	Num          int32        `json:"num"`
	Region       *string      `json:"region,omitempty"`
	City         *string      `json:"city,omitempty"`
	DefaultIsp   *string      `json:"default_isp,omitempty"`
	MultiIPMode  *MULTIIPMODE `json:"multi_ip_mode,omitempty"`
}

type MULTIIPMODE = string

const (
	MULTIIPMODESINGLEINTERFACE = "single_interface"

	MULTIIPMODEMULTIINTERFACE = "multi_interface"
)

type ScheduleStrategy struct {
	NetworkStrategy  string `json:"network_strategy"`
	ScheduleStrategy string `json:"schedule_strategy"`
	PriceStrategy    string `json:"price_strategy"`
}

type AdvancedConfiguration struct {
	InstanceName     *string `json:"instance_name,omitempty"`
	InstanceDesc     *string `json:"instance_desc,omitempty"`
	InstanceHostName *string `json:"instance_host_name,omitempty"`
}

type AttachedPvcConfig struct {
	Name string `json:"name"`
}

type IMAGEBOOTMODE = string

const (
	IMAGEBOOTMODEUEFI = "UEFI"

	IMAGEBOOTMODEBIOS = "BIOS"
)

type CreateCloudServerResp struct {
	ResponseMetadata VolcResponseMetadata    `json:"ResponseMetadata"`
	Result           CreateCloudServerResult `json:"Result"`
}

type CreateCloudServerResult struct {
	CloudServerIdentity string `json:"cloud_server_identity"`
}

type ListCloudServersReq struct {
	// 模糊搜索名
	FuzzyName string `json:"fuzzy_name" query:"fuzzy_name"`
	Pagination
}

type Pagination struct {
	// 页码
	Page int32 `json:"page" query:"page"`
	// 每页大小
	Limit int32 `json:"limit" query:"limit"`
	// [DESC=1 ASC=2]
	OrderBy int32 `json:"order_by" query:"order_by"`
}

type ListCloudServersResp struct {
	ResponseMetadata VolcResponseMetadata   `json:"ResponseMetadata"`
	Result           ListCloudServersResult `json:"Result"`
}

type ListCloudServersResult struct {
	CloudServers []*CloudServer `json:"cloud_servers"`
	TotalCount   int64          `json:"total_count"`
}

type CloudServer struct {
	CloudServerIdentity     string                              `json:"cloud_server_identity"`
	Name                    string                              `json:"name"`
	LoadType                *LoadType                           `json:"load_type,omitempty"`
	ServerAreaCount         int32                               `json:"server_area_count"`
	InstanceCount           int64                               `json:"instance_count"`
	ServerAreaLevel         string                              `json:"server_area_level"`
	InstanceStatus          []*CloudServerInstanceStatus        `json:"instance_status"`
	SpecSum                 map[string]int32                    `json:"spec_sum"`
	Spec                    string                              `json:"spec"`
	SpecDisplay             *string                             `json:"spec_display,omitempty"`
	CPU                     string                              `json:"cpu"`
	Mem                     string                              `json:"mem"`
	Image                   *ImageConfig                        `json:"image"`
	Storage                 *StorageConfig                      `json:"storage"`
	Network                 *CloudServerNetworkConfig           `json:"network"`
	Gpu                     *GpuConfig                          `json:"gpu,omitempty"`
	SecretConfig            *SecretConfig                       `json:"secret_config,omitempty"`
	CustomData              *CustomData                         `json:"custom_data,omitempty"`
	BillingConfig           *CloudServerBillingConfigs          `json:"billing_config,omitempty"`
	Desc                    *string                             `json:"desc,omitempty"`
	Arch                    *string                             `json:"arch,omitempty"`
	AdvancedConfiguration   *AdvancedConfiguration              `json:"advanced_configuration,omitempty"`
	ScheduleStrategyConfigs *CloudServerScheduleStrategyConfigs `json:"schedule_strategy_configs"`
	ServerAreas             []*CloudServerArea                  `json:"server_areas"`
	CreateTime              int64                               `json:"create_time"`
	UpdateTime              int64                               `json:"update_time"`
}

type LoadType = string

const (
	LoadTypeVM = "virtual_machine"

	LoadTypeBm = "bare_metal"
)

type CloudServerInstanceStatus struct {
	Status        string `json:"status"`
	InstanceCount int32  `json:"instance_count"`
}

type ImageConfig struct {
	ImageIdentity string         `json:"image_identity"`
	ImageName     string         `json:"image_name"`
	SystemArch    *string        `json:"system_arch,omitempty"`
	SystemType    *string        `json:"system_type,omitempty"`
	SystemBit     *string        `json:"system_bit,omitempty"`
	SystemVersion *string        `json:"system_version,omitempty"`
	Property      *string        `json:"property,omitempty"`
	DisableVga    *bool          `json:"disable_vga,omitempty"`
	ImageBootMode *IMAGEBOOTMODE `json:"image_boot_mode,omitempty"`
}

type GpuConfig struct {
	Gpus []*GpuSpecWithNum `json:"gpus"`
}

type GpuSpecWithNum struct {
	GpuSpec *GpuSpec `json:"gpu_spec"`
	Num     int32    `json:"num"`
}

type GpuSpec struct {
	GpuType string `json:"gpu_type"`
}

type CloudServerScheduleStrategyConfigs struct {
	NetworkStrategy  string `json:"network_strategy"`
	ScheduleStrategy string `json:"schedule_strategy"`
	PriceStrategy    string `json:"price_strategy"`
}

type CloudServerArea struct {
	Area         string  `json:"area"`
	Isp          string  `json:"isp"`
	InstanceNum  int32   `json:"instance_num"`
	ClusterName  *string `json:"cluster_name,omitempty"`
	ClusterAlias *string `json:"cluster_alias,omitempty"`
}

type GetCloudServerReq struct {
	// 云服务 id
	// required:true
	CloudServerID string `json:"cloud_server_id" query:"cloud_server_id"`
}

type GetCloudServerResp struct {
	ResponseMetadata VolcResponseMetadata `json:"ResponseMetadata"`
	Result           GetCloudServerResult `json:"Result"`
}

type GetCloudServerResult struct {
	CloudServer *CloudServer `json:"cloud_server,omitempty"`
}

type StartCloudServerReq struct {
	// 云服务 id
	// required:true
	// in:body
	CloudServerID string `json:"cloud_server_id" query:"cloud_server_id"`
}

type StartCloudServerResp struct {
	ResponseMetadata VolcResponseMetadata   `json:"ResponseMetadata"`
	Result           StartCloudServerResult `json:"Result"`
}

type StartCloudServerResult struct{}

type StopCloudServerReq struct {
	// 云服务 id
	// required:true
	// in:body
	CloudServerID string `json:"cloud_server_id" query:"cloud_server_id"`
}

type StopCloudServerResp struct {
	ResponseMetadata VolcResponseMetadata  `json:"ResponseMetadata"`
	Result           StopCloudServerResult `json:"Result"`
}

type StopCloudServerResult struct{}

type RebootCloudServerReq struct {
	// 云服务 id
	// required:true
	// in:body
	CloudServerID string `json:"cloud_server_id" query:"cloud_server_id"`
}

type RebootCloudServerResp struct {
	ResponseMetadata VolcResponseMetadata    `json:"ResponseMetadata"`
	Result           RebootCloudServerResult `json:"Result"`
}

type RebootCloudServerResult struct{}

type DeleteCloudServerReq struct {
	// 云服务 id
	// required:true
	// in:body
	CloudServerID string `json:"cloud_server_id" query:"cloud_server_id"`
}

type DeleteCloudServerResp struct {
	ResponseMetadata VolcResponseMetadata    `json:"ResponseMetadata"`
	Result           DeleteCloudServerResult `json:"Result"`
}

type DeleteCloudServerResult struct{}

type ListInstanceTypesReq struct {
}

type ListInstanceTypesResp struct {
	ResponseMetadata VolcResponseMetadata    `json:"ResponseMetadata"`
	Result           ListInstanceTypesResult `json:"Result"`
}

type ListInstanceTypesResult struct {
	InstanceTypeConfigs []*InstanceTypeConfig `json:"instance_type_configs"`
}

type InstanceTypeConfig struct {
	InstanceType           string         `json:"instance_type"`
	InstanceTypeFamily     string         `json:"instance_type_family"`
	CPU                    int16          `json:"cpu"`
	Memory                 int16          `json:"memory"`
	Gpu                    int64          `json:"gpu"`
	GpuSpec                string         `json:"gpu_spec"`
	InstanceTypeFamilyName string         `json:"instance_type_family_name"`
	StorageInfo            []*StorageInfo `json:"storage_info,omitempty"`
	Arch                   *INSTANCEARCH  `json:"arch,omitempty"`
}

type StorageInfo struct {
	LocalStorageCategory     string  `json:"local_storage_category"`
	LocalStorageCapacity     int64   `json:"local_storage_capacity"`
	LocalStorageAmount       int64   `json:"local_storage_amount"`
	LocalStorageUnit         string  `json:"local_storage_unit"`
	LocalStorageResourceName *string `json:"local_storage_resource_name,omitempty"`
	Feature                  *int64  `json:"feature,omitempty"`
}

type INSTANCEARCH = string

const (
	INSTANCEARCHX86 = "x86"

	INSTANCEARCHBAREMETAL = "bare_metal"

	INSTANCEARCHARM = "arm"
)

type ListAvailableResourceInfoReq struct {
	//规格
	InstanceType string `json:"instance_type,omitempty" query:"instance_type"`

	//云盘类型：(CloudHDD, CloudSSD)
	CloudDiskType string `json:"cloud_disk_type,omitempty" query:"cloud_disk_type"`
}

type ListAvailableResourceInfoResp struct {
	ResponseMetadata VolcResponseMetadata            `json:"ResponseMetadata"`
	Result           ListAvailableResourceInfoResult `json:"Result"`
}

type ListAvailableResourceInfoResult struct {
	Regions []RegionInfo `json:"regions"`
}

type RegionInfo struct {
	Country      *GeoInfo `json:"country"`
	Area         *GeoInfo `json:"area"`
	City         *GeoInfo `json:"city"`
	Isp          *IspInfo `json:"isp"`
	Cluster      *GeoInfo `json:"cluster"`
	AvailableNum int64    `json:"available_num,omitempty"`
}

type GeoInfo struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	EnName string `json:"en_name"`
}

type IspInfo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CreateInstanceReq struct {
	// in:body
	CloudServerIdentity string `json:"cloud_server_identity" query:"cloud_server_identity"`
	// in:body
	InstanceAreaNums []*InstanceAreaNum `json:"instance_area_nums" query:"instance_area_nums"`
}

type CreateInstanceResp struct {
	ResponseMetadata VolcResponseMetadata `json:"ResponseMetadata"`
	Result           CreateInstanceResult `json:"Result"`
}

type CreateInstanceResult struct {
	InstanceIDList []string `json:"instance_id_list,omitempty"`
}

type ListInstancesReq struct {
	Countries             string `json:"countries,omitempty" query:"countries"`
	Regions               string `json:"regions,omitempty" query:"regions"`
	Cities                string `json:"cities,omitempty" query:"cities"`
	Isps                  string `json:"isps,omitempty" query:"isps"`
	Status                string `json:"status,omitempty" query:"status"`
	ClusterNames          string `json:"cluster_names,omitempty" query:"cluster_names"`
	CloudServerIdentities string `json:"cloud_server_identities,omitempty" query:"cloud_server_identities"`
	Ips                   string `json:"ips,omitempty" query:"ips"`
	Cidrs                 string `json:"cidrs,omitempty" query:"cidrs"`
	SpecNames             string `json:"spec_names,omitempty" query:"spec_names"`
	InstanceIdentities    string `json:"instance_identities,omitempty" query:"instance_identities"`
	InstanceUuids         string `json:"instance_uuids,omitempty" query:"instance_uuids"`
	InstanceNames         string `json:"instance_names,omitempty" query:"instance_names"`
	FuzzyIPWithDots       string `json:"fuzzy_ip_with_dots,omitempty" query:"fuzzy_ip_with_dots"`
	VPCIdentities         string `json:"vpc_identities,omitempty" query:"vpc_identities"`
	Pagination
}

type ListInstancesResp struct {
	ResponseMetadata VolcResponseMetadata `json:"ResponseMetadata"`
	Result           ListInstancesResult  `json:"Result"`
}

type ListInstancesResult struct {
	Instances  []*VeenInstance `json:"instances"`
	TotalCount int64           `json:"total_count"`
}

type VeenInstance struct {
	InstanceIdentity    string         `json:"instance_identity"`
	InstanceUUID        *string        `json:"instance_uuid,omitempty"`
	InstanceName        string         `json:"instance_name"`
	LoadType            *LoadType      `json:"load_type,omitempty"`
	CloudServerIdentity string         `json:"cloud_server_identity"`
	CloudServerName     string         `json:"cloud_server_name"`
	VpcIdentity         string         `json:"vpc_identity"`
	Namespace           *string        `json:"namespace"`
	SubnetCidr          string         `json:"subnet_cidr"`
	Cluster             *VeenCluster   `json:"cluster"`
	Spec                string         `json:"spec"`
	SpecDisplay         *string        `json:"spec_display,omitempty"`
	CPU                 string         `json:"cpu"`
	Mem                 string         `json:"mem"`
	Status              VeenSTATUS     `json:"status"`
	StatusReason        *string        `json:"status_reason,omitempty"`
	Phase               *string        `json:"phase,omitempty"`
	Creator             string         `json:"creator"`
	Image               *ImageConfig   `json:"image"`
	Storage             *StorageConfig `json:"storage"`
	Network             *NetworkConfig `json:"network"`
	Gpu                 *GpuConfig     `json:"gpu,omitempty"`
	Secret              *SecretConfig  `json:"secret,omitempty"`
	Account             *Account       `json:"account,omitempty"`
	InstanceDesc        *string        `json:"instance_desc,omitempty"`
	HostName            *string        `json:"host_name,omitempty"`
	Arch                *string        `json:"arch,omitempty"`
	ImageBootMode       *IMAGEBOOTMODE `json:"image_boot_mode,omitempty"`
	StartTime           int64          `json:"start_time"`
	EndTime             int64          `json:"end_time"`
	CreateTime          int64          `json:"create_time"`
	UpdateTime          int64          `json:"update_time"`
	DeleteTime          *int64         `json:"delete_time,omitempty"`
}

type VeenCluster struct {
	ClusterName string `json:"cluster_name"`
	Country     string `json:"country"`
	Region      string `json:"region"`
	Province    string `json:"province"`
	City        string `json:"city"`
	Isp         string `json:"isp"`
	Level       string `json:"level"`
	Alias       string `json:"alias"`
}

type VeenSTATUS = string

const (
	VEENSTATUSOPENING = "opening"

	VEENSTATUSRUNNING = "running"

	VEENSTATUSSTOPPING = "stopping"

	VEENSTATUSSTOP = "stop"

	VEENSTATUSSTARTING = "starting"

	VEENSTATUSREBOOTING = "rebooting"

	VEENSTATUSTERMINATING = "terminating"

	VEENSTATUSOPENFAIL = "open_fail"
)

type NetworkConfig struct {
	InternalInterface           *NetworkInterfaceConfig   `thrift:"internal_interface,1,required" frugal:"1,required,NetworkInterfaceConfig" json:"internal_interface"`
	ExternalInterface           *NetworkInterfaceConfig   `thrift:"external_interface,2,required" frugal:"2,required,NetworkInterfaceConfig" json:"external_interface"`
	ExternalInterfaces          []*NetworkInterfaceConfig `thrift:"external_interfaces,8,optional" frugal:"8,optional,list<NetworkInterfaceConfig>" json:"external_interfaces,omitempty"`
	VfPassthrough               *bool                     `thrift:"vf_passthrough,3,optional" frugal:"3,optional,bool" json:"vf_passthrough,omitempty"`
	DisableIpv4                 *bool                     `thrift:"disable_ipv4,11,optional" frugal:"11,optional,bool" json:"disable_ipv4,omitempty"`
	EnableIpv6                  *bool                     `thrift:"enable_ipv6,4,optional" frugal:"4,optional,bool" json:"enable_ipv6,omitempty"`
	DefaultIsp                  *string                   `thrift:"default_isp,5,optional" frugal:"5,optional,string" json:"default_isp,omitempty"`
	VlanVfPassthrough           *bool                     `thrift:"vlan_vf_passthrough,6,optional" frugal:"6,optional,bool" json:"vlan_vf_passthrough,omitempty"`
	MultiIPMode                 *MULTIIPMODE              `thrift:"multi_ip_mode,7,optional" frugal:"7,optional,string" json:"multi_ip_mode,omitempty"`
	WantedSecondaryIPNum        *int32                    `thrift:"wanted_secondary_ip_num,9,optional" frugal:"9,optional,i32" json:"wanted_secondary_ip_num,omitempty"`
	ActualSecondaryIPNum        *int32                    `thrift:"actual_secondary_ip_num,10,optional" frugal:"10,optional,i32" json:"actual_secondary_ip_num,omitempty"`
	CustomInternalInterfaceName *string                   `thrift:"custom_internal_interface_name,12,optional" frugal:"12,optional,string" json:"custom_internal_interface_name,omitempty"`
	CustomExternalInterfaceName *string                   `thrift:"custom_external_interface_name,13,optional" frugal:"13,optional,string" json:"custom_external_interface_name,omitempty"`
}

type NetworkInterfaceConfig struct {
	IPAddr             string  `thrift:"ip_addr,1,required" frugal:"1,required,string" json:"ip_addr"`
	Mask               string  `thrift:"mask,2,required" frugal:"2,required,string" json:"mask"`
	IP6Addr            *string `thrift:"ip6_addr,5,optional" frugal:"5,optional,string" json:"ip6_addr,omitempty"`
	Mask6              *string `thrift:"mask6,6,optional" frugal:"6,optional,string" json:"mask6,omitempty"`
	Ips                []*IP   `thrift:"ips,7,optional" frugal:"7,optional,list<IP>" json:"ips,omitempty"`
	MacAddr            string  `thrift:"mac_addr,3,required" frugal:"3,required,string" json:"mac_addr"`
	BandwidthPeak      string  `thrift:"bandwidth_peak,4,required" frugal:"4,required,string" json:"bandwidth_peak"`
	BandwidthPackageID *int64  `thrift:"bandwidth_package_id,8,optional" frugal:"8,optional,i64" json:"bandwidth_package_id,omitempty"`
}

type IP struct {
	Addr      string    `thrift:"addr,1,required" frugal:"1,required,string" json:"addr"`
	Mask      string    `thrift:"mask,2,required" frugal:"2,required,string" json:"mask"`
	Isp       string    `thrift:"isp,3,required" frugal:"3,required,string" json:"isp"`
	IPVersion IPVERSION `thrift:"ip_version,4,required" frugal:"4,required,string" json:"ip_version"`
	Primary   *bool     `thrift:"primary,5,optional" frugal:"5,optional,bool" json:"primary,omitempty"`
}

type IPVERSION = string

const (
	IPVERSION4 = "ipv4"

	IPVERSION6 = "ipv6"
)

type Account struct {
	AccountIdentity int64  `thrift:"account_identity,1,required" frugal:"1,required,i64" json:"account_identity"`
	UserIdentity    *int64 `thrift:"user_identity,2,optional" frugal:"2,optional,i64" json:"user_identity,omitempty"`
}

type GetInstanceReq struct {
	InstanceIdentity string `json:"instance_identity" query:"instance_identity" validate:"required"`
}

type GetInstanceResp struct {
	ResponseMetadata VolcResponseMetadata `json:"ResponseMetadata"`
	Result           GetInstanceResult    `json:"Result"`
}

type GetInstanceResult struct {
	Instance *VeenInstance `json:"instance,omitempty"`
}

type StartInstancesReq struct {
	// in:body
	InstanceIdentities []string `json:"instance_identities" query:"instance_identities" validate:"required,gt=0,dive,required"`
}

type StartInstancesResp struct {
	ResponseMetadata VolcResponseMetadata `json:"ResponseMetadata"`
	Result           StartInstancesResult `json:"Result"`
}

type StartInstancesResult struct{}

type StopInstancesReq struct {
	// in:body
	InstanceIdentities []string `json:"instance_identities" query:"instance_identities" validate:"required,gt=0,dive,required"`
}

type StopInstancesResp struct {
	ResponseMetadata VolcResponseMetadata `json:"ResponseMetadata"`
	Result           StopInstancesResult  `json:"Result"`
}

type StopInstancesResult struct{}

type RebootInstancesReq struct {
	// in:body
	InstanceIdentities []string `json:"instance_identities" query:"instance_identities" validate:"required,gt=0,dive,required"`
}

type RebootInstancesResp struct {
	ResponseMetadata VolcResponseMetadata  `json:"ResponseMetadata"`
	Result           RebootInstancesResult `json:"Result"`
}

type RebootInstancesResult struct{}

type OfflineInstancesReq struct {
	// in:body
	InstanceIdentities []string `json:"instance_identities" query:"instance_identities" validate:"required,gt=0,dive,required"`
	// in:body
	IgnoreRunning *bool `json:"ignore_running,omitempty" query:"ignore_running"`
}

type OfflineInstancesResp struct {
	ResponseMetadata VolcResponseMetadata   `json:"ResponseMetadata"`
	Result           OfflineInstancesResult `json:"Result"`
}

type OfflineInstancesResult struct{}

type SetInstanceNameReq struct {
	// in:body
	InstanceIdentity string `json:"instance_identity" query:"instance_identity" validate:"required"`
	// in:body
	InstanceName string `json:"instance_name" query:"instance_name" validate:"required"`
}

type SetInstanceNameResp struct {
	ResponseMetadata VolcResponseMetadata  `json:"ResponseMetadata"`
	Result           SetInstanceNameResult `json:"Result"`
}

type SetInstanceNameResult struct{}

type ResetLoginCredentialReq struct {
	// in:body
	InstanceIdentity string `json:"instance_identity" query:"instance_identity" validate:"required"`
	// custom: 2, ssh: 3
	// in:body
	SecretType int64 `json:"secret_type" query:"secret_type" validate:"required"`
	// secret_data is key_pair_id when secret_type is ssh.
	// in:body
	SecretData string `json:"secret_data" query:"secret_data" validate:"required"`
}

type ResetLoginCredentialResp struct {
	ResponseMetadata VolcResponseMetadata       `json:"ResponseMetadata"`
	Result           ResetLoginCredentialResult `json:"Result"`
}

type ResetLoginCredentialResult struct{}

type GetInstanceCloudDiskInfoReq struct {
	// required: true
	InstanceIdentity string `json:"instance_identity" query:"instance_identity" validate:"required"`
}

type GetInstanceCloudDiskInfoResp struct {
	ResponseMetadata VolcResponseMetadata           `json:"ResponseMetadata"`
	Result           GetInstanceCloudDiskInfoResult `json:"Result"`
}

type GetInstanceCloudDiskInfoResult struct {
	InstanceIdentity      string           `json:"instance_identity"`
	InstanceName          string           `json:"instance_name"`
	Arch                  string           `json:"arch"`
	SystemCloudDiskInfo   *CloudDiskInfo   `json:"system_cloud_disk_info"`
	DataCloudDiskInfoList []*CloudDiskInfo `json:"data_cloud_disk_info_list"`
}

type CloudDiskInfo struct {
	StorageType string `thrift:"storage_type,1,required" frugal:"1,required,string" json:"storage_type"`
	Capacity    string `thrift:"capacity,2,required" frugal:"2,required,string" json:"capacity"`
	PvcName     string `thrift:"pvc_name,3,required" frugal:"3,required,string" json:"pvc_name"`
	DeviceName  string `thrift:"device_name,4,required" frugal:"4,required,string" json:"device_name"`
}

type ScaleInstanceCloudDiskCapacityReq struct {
	// required: true
	// in: body
	InstanceIdentity string `json:"instance_identity" validate:"required"`
	// required: true
	// in: body
	ScaleSystemCloudDiskInfo *ScaleCloudDiskInfo `json:"scale_system_cloud_disk_info" validate:"required"`
	// required: true
	// in: body
	ScaleDataCloudDiskInfoList []*ScaleCloudDiskInfo `json:"scale_data_cloud_disk_info_list" validate:"required"`
	// required: true
	// in: body
	WithReboot *bool `json:"with_reboot" validate:"required"`
}

type ScaleCloudDiskInfo struct {
	DeviceName string `thrift:"device_name,1,required" frugal:"1,required,string" json:"device_name"`
	Capacity   string `thrift:"capacity,2,required" frugal:"2,required,string" json:"capacity"`
}

type ScaleInstanceCloudDiskCapacityResp struct {
	ResponseMetadata VolcResponseMetadata                 `json:"ResponseMetadata"`
	Result           ScaleInstanceCloudDiskCapacityResult `json:"Result"`
}

type ScaleInstanceCloudDiskCapacityResult struct{}

const (
	ChargeTypeHourUsed ChargeType = "HourUsed"

	EbsTypeData EbsType = "data"

	EbsStatusCreating EbsStatus = "creating"

	EbsStatusDetached EbsStatus = "detached"

	EbsStatusAttaching EbsStatus = "attaching"

	EbsStatusAttached EbsStatus = "attached"

	EbsStatusDetaching EbsStatus = "detaching"

	EbsStatusScaling EbsStatus = "scaling"

	EbsStatusDeleting EbsStatus = "deleting"

	EbsStatusDeleted EbsStatus = "deleted"

	StorageTypeSSD StorageType = "CloudBlockSSD"

	StorageTypeHDD StorageType = "CloudBlockHDD"

	AttachmentTypeVeen AttachmentType = "veen"
)

type ChargeType = string

type EbsType = string

type EbsStatus = string

type StorageType = string

type AttachmentType = string

type EbsPagination struct {
	PageNo   int32 `thrift:"page_no,1,optional" frugal:"1,optional,i32" json:"page_no,omitempty"`
	PageSize int32 `thrift:"page_size,2,optional" frugal:"2,optional,i32" json:"page_size,omitempty"`
}

type EbsOrderOption struct {
	OrderBy string `thrift:"order_by,1,optional" frugal:"1,optional,string" json:"order_by,omitempty"`
	Asc     bool   `thrift:"asc,2,optional" frugal:"2,optional,bool" json:"asc,omitempty"`
}

type EbsAttachment struct {
	EbsID   string         `thrift:"ebs_id,1,required" frugal:"1,required,string" json:"ebs_id"`
	ResType AttachmentType `thrift:"res_type,2,required" frugal:"2,required,string" json:"res_type"`
	ResID   string         `thrift:"res_id,3,required" frugal:"3,required,string" json:"res_id"`
	ResName *string        `thrift:"res_name,4,optional" frugal:"4,optional,string" json:"res_name,omitempty"`
}

type EbsCluster struct {
	ClusterName string `thrift:"cluster_name,1,required" frugal:"1,required,string" json:"cluster_name"`
	Country     string `thrift:"country,2,required" frugal:"2,required,string" json:"country"`
	Region      string `thrift:"region,3,required" frugal:"3,required,string" json:"region"`
	Province    string `thrift:"province,4,required" frugal:"4,required,string" json:"province"`
	City        string `thrift:"city,5,required" frugal:"5,required,string" json:"city"`
	Isp         string `thrift:"isp,6,required" frugal:"6,required,string" json:"isp"`
	Alias       string `thrift:"alias,7,required" frugal:"7,required,string" json:"alias"`
}

type EbsInstance struct {
	AccountID      int64          `thrift:"account_id,1,required" frugal:"1,required,i64" json:"account_id"`
	EbsID          string         `thrift:"ebs_id,2,required" frugal:"2,required,string" json:"ebs_id"`
	EbsName        string         `thrift:"ebs_name,3,required" frugal:"3,required,string" json:"ebs_name"`
	Status         EbsStatus      `thrift:"status,4,required" frugal:"4,required,string" json:"status"`
	Cluster        *EbsCluster    `thrift:"cluster,5,required" frugal:"5,required,EbsCluster" json:"cluster"`
	StorageType    StorageType    `thrift:"storage_type,6,required" frugal:"6,required,string" json:"storage_type"`
	Capacity       string         `thrift:"capacity,7,required" frugal:"7,required,string" json:"capacity"`
	EbsType        EbsType        `thrift:"ebs_type,8,required" frugal:"8,required,string" json:"ebs_type"`
	ChargeType     ChargeType     `thrift:"charge_type,9,required" frugal:"9,required,string" json:"charge_type"`
	Desc           *string        `thrift:"desc,10,optional" frugal:"10,optional,string" json:"desc,omitempty"`
	Attachment     *EbsAttachment `thrift:"attachment,11,optional" frugal:"11,optional,EbsAttachment" json:"attachment,omitempty"`
	DeleteWithVeen *bool          `thrift:"delete_with_veen,12,optional" frugal:"12,optional,bool" json:"delete_with_veen,omitempty"`
	StartTime      *int64         `thrift:"start_time,128,optional" frugal:"128,optional,i64" json:"start_time,omitempty"`
	EndTime        *int64         `thrift:"end_time,129,optional" frugal:"129,optional,i64" json:"end_time,omitempty"`
	CreateTime     int64          `thrift:"create_time,130,required" frugal:"130,required,i64" json:"create_time"`
	UpdateTime     int64          `thrift:"update_time,131,required" frugal:"131,required,i64" json:"update_time"`
}

type CreateEbsInstancesReq struct {
	ClusterName string      `thrift:"cluster_name,3,required" frugal:"3,required,string" json:"cluster_name"`
	ChargeType  ChargeType  `thrift:"charge_type,4,required" frugal:"4,required,string" json:"charge_type"`
	EbsType     EbsType     `thrift:"ebs_type,5,required" frugal:"5,required,string" json:"ebs_type"`
	StorageType StorageType `thrift:"storage_type,6,required" frugal:"6,required,string" json:"storage_type"`
	Capacity    string      `thrift:"capacity,7,required" frugal:"7,required,string" json:"capacity"`
	Number      int64       `thrift:"number,8,required" frugal:"8,required,i64" json:"number"`
	Name        string      `thrift:"name,9,required" frugal:"9,required,string" json:"name"`
	Desc        string      `thrift:"desc,10,optional" frugal:"10,optional,string" json:"desc,omitempty"`
}

type CreateEbsInstancesResp struct {
	ResponseMetadata VolcResponseMetadata     `json:"ResponseMetadata"`
	Result           CreateEbsInstancesResult `json:"Result"`
}

type CreateEbsInstancesResult struct {
	EbsIds []string `thrift:"ebs_ids,1,optional" frugal:"1,optional,list<string>" json:"ebs_ids,omitempty"`
}

type ListEbsInstancesReq struct {
	WithAttachmentInfo  bool           `json:"with_attachment_info,omitempty" query:"with_attachment_info"`
	ResIds              string         `json:"res_ids,omitempty" query:"res_ids"`
	EbsIds              string         `json:"ebs_ids,omitempty" query:"ebs_ids"`
	EbsNames            string         `json:"ebs_names,omitempty" query:"ebs_names"`
	Regions             string         `json:"regions,omitempty" query:"regions"`
	ClusterNames        string         `json:"cluster_names,omitempty" query:"clusters_names"`
	Status              EbsStatus      `json:"status,omitempty" query:"status"`
	EbsType             EbsType        `json:"ebs_type,omitempty" query:"ebs_type"`
	ChargeType          ChargeType     `json:"charge_type,omitempty" query:"charge_type"`
	FuzzyVeenExternalIP string         `json:"fuzzy_veen_external_ip,omitempty" query:"fuzzy_veen_external_ip"`
	PageOption          EbsPagination  `json:"page_option,omitempty" query:"page_option"`
	OrderOption         EbsOrderOption `json:"order_option,omitempty" query:"order_option"`
}

type ListEbsInstancesResp struct {
	ResponseMetadata VolcResponseMetadata   `json:"ResponseMetadata"`
	Result           ListEbsInstancesResult `json:"Result"`
}

type ListEbsInstancesResult struct {
	EbsInstances []*EbsInstance `json:"ebs_instances,omitempty"`
	TotalCnt     *int64         `json:"total_cnt,omitempty"`
}

type GetEbsInstanceReq struct {
	EbsID              string `json:"ebs_id" query:"ebs_id"`
	WithAttachmentInfo bool   `json:"with_attachment_info,omitempty" query:"with_attachment_info"`
}

type GetEbsInstanceResp struct {
	ResponseMetadata VolcResponseMetadata `json:"ResponseMetadata"`
	Result           GetEbsInstanceResult `json:"Result"`
}

type GetEbsInstanceResult struct {
	EbsInstance *EbsInstance `json:"ebs_instance,omitempty"`
}

type ScaleEbsInstanceCapacityReq struct {
	EbsID    string `thrift:"ebs_id,3,required" frugal:"3,required,string" json:"ebs_id"`
	Capacity string `thrift:"capacity,4,required" frugal:"4,required,string" json:"capacity"`
}

type ScaleEbsInstanceCapacityResp struct {
	ResponseMetadata VolcResponseMetadata           `json:"ResponseMetadata"`
	Result           ScaleEbsInstanceCapacityResult `json:"Result"`
}

type ScaleEbsInstanceCapacityResult struct {
}

type AttachEbsReq struct {
	EbsIds  []string       `thrift:"ebs_ids,3,required" frugal:"3,required,list<string>" json:"ebs_ids"`
	ResType AttachmentType `thrift:"res_type,4,required" frugal:"4,required,string" json:"res_type"`
	ResID   string         `thrift:"res_id,5,required" frugal:"5,required,string" json:"res_id"`
}

type AttachEbsResp struct {
	ResponseMetadata VolcResponseMetadata `json:"ResponseMetadata"`
	Result           AttachEbsResult      `json:"Result"`
}

type AttachEbsResult struct {
}

type DetachEbsReq struct {
	EbsIds []string `thrift:"ebs_ids,3,required" frugal:"3,required,list<string>" json:"ebs_ids"`
}

type DetachEbsResp struct {
	ResponseMetadata VolcResponseMetadata `json:"ResponseMetadata"`
	Result           DetachEbsResult      `json:"Result"`
}

type DetachEbsResult struct {
}

type DeleteEbsInstanceReq struct {
	EbsIds []string `thrift:"ebs_ids,3,required" frugal:"3,required,list<string>" json:"ebs_ids"`
}

type DeleteEbsInstanceResp struct {
	ResponseMetadata VolcResponseMetadata    `json:"ResponseMetadata"`
	Result           DeleteEbsInstanceResult `json:"Result"`
}

type DeleteEbsInstanceResult struct {
}
