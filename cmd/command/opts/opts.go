/*
Copyright 2023 KylinSoft  Co., Ltd.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package opts

var Opts OptionsList

var RootOpts struct {
	LogLevel string
}

type OptionsList struct {
	RootOptDir        string
	Arch              string
	ClusterConfigFile string
	KubeConfigFile    string
	NKD               NKDConfig
	InfraPlatform

	ClusterID string
	Platform  string

	UserName             string
	Password             string
	SSHKey               string
	Master               MasterConfig
	Worker               WorkerConfig
	ExtendCount          uint
	Runtime              string
	ApiServerEndpoint    string
	ImageRegistry        string
	PauseImage           string
	ReleaseImageUrl      string
	KubeVersion          string
	KubernetesAPIVersion uint
	Token                string
	CertificateKey       string
	PreHookScript        string
	PostHookYaml         string

	NetWork NetworkConfig
	Housekeeper
}

type NKDConfig struct {
	Log_Level string
	BootstrapUrl
}

type BootstrapUrl struct {
	BootstrapIgnHost string
	BootstrapIgnPort string
}

type InfraPlatform struct {
	OpenStack
	Libvirt
}

type OpenStack struct {
	UserName          string
	Password          string
	Tenant_Name       string
	Auth_URL          string
	Region            string
	Internal_Network  string
	External_Network  string
	Glance_Name       string
	Availability_Zone string
}

type Libvirt struct {
	URI     string
	OSImage string
	CIDR    string
	Gateway string
}

type MasterConfig struct {
	Hostname []string
	CPU      uint
	RAM      uint
	Disk     uint
	IP       []string
}

type WorkerConfig struct {
	Hostname []string
	CPU      uint
	RAM      uint
	Disk     uint
	IP       []string
}

type NetworkConfig struct {
	ServiceSubnet string
	PodSubnet     string
	Plugin        string
}

type Housekeeper struct {
	DeployHousekeeper  bool
	OperatorImageUrl   string
	ControllerImageUrl string
	KubeVersion        string
	EvictPodForce      bool
	MaxUnavailable     uint
	OSImageURL         string
}
