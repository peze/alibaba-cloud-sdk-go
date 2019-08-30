package ecs

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// DedicatedHost is a nested struct in ecs response
type DedicatedHost struct {
	DedicatedHostId             string                                              `json:"DedicatedHostId" xml:"DedicatedHostId"`
	AutoPlacement               string                                              `json:"AutoPlacement" xml:"AutoPlacement"`
	RegionId                    string                                              `json:"RegionId" xml:"RegionId"`
	ZoneId                      string                                              `json:"ZoneId" xml:"ZoneId"`
	DedicatedHostName           string                                              `json:"DedicatedHostName" xml:"DedicatedHostName"`
	MachineId                   string                                              `json:"MachineId" xml:"MachineId"`
	Description                 string                                              `json:"Description" xml:"Description"`
	DedicatedHostType           string                                              `json:"DedicatedHostType" xml:"DedicatedHostType"`
	TotalSockets                int                                                 `json:"TotalSockets" xml:"TotalSockets"`
	TotalPhysicalCores          int                                                 `json:"TotalPhysicalCores" xml:"TotalPhysicalCores"`
	PhysicalGpus                int                                                 `json:"PhysicalGpus" xml:"PhysicalGpus"`
	GPUSpec                     string                                              `json:"GPUSpec" xml:"GPUSpec"`
	ActionOnMaintenance         string                                              `json:"ActionOnMaintenance" xml:"ActionOnMaintenance"`
	Status                      string                                              `json:"Status" xml:"Status"`
	CreationTime                string                                              `json:"CreationTime" xml:"CreationTime"`
	ChargeType                  string                                              `json:"ChargeType" xml:"ChargeType"`
	SaleCycle                   string                                              `json:"SaleCycle" xml:"SaleCycle"`
	ExpiredTime                 string                                              `json:"ExpiredTime" xml:"ExpiredTime"`
	AutoReleaseTime             string                                              `json:"AutoReleaseTime" xml:"AutoReleaseTime"`
	ResourceGroupId             string                                              `json:"ResourceGroupId" xml:"ResourceGroupId"`
	SupportInstanceTypeFamilies SupportInstanceTypeFamiliesInDescribeDedicatedHosts `json:"SupportInstanceTypeFamilies" xml:"SupportInstanceTypeFamilies"`
	SupportInstanceTypesList    SupportInstanceTypesListInDescribeDedicatedHosts    `json:"SupportInstanceTypesList" xml:"SupportInstanceTypesList"`
	Capacity                    Capacity                                            `json:"Capacity" xml:"Capacity"`
	NetworkAttributes           NetworkAttributes                                   `json:"NetworkAttributes" xml:"NetworkAttributes"`
	Instances                   Instances                                           `json:"Instances" xml:"Instances"`
	OperationLocks              OperationLocksInDescribeDedicatedHosts              `json:"OperationLocks" xml:"OperationLocks"`
	Tags                        TagsInDescribeDedicatedHosts                        `json:"Tags" xml:"Tags"`
}
