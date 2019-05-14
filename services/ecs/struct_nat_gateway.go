
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

// NatGateway is a nested struct in ecs response
type NatGateway struct {
            NatGatewayId     string `json:"NatGatewayId" xml:"NatGatewayId"`
            RegionId     string `json:"RegionId" xml:"RegionId"`
            Name     string `json:"Name" xml:"Name"`
            Description     string `json:"Description" xml:"Description"`
            VpcId     string `json:"VpcId" xml:"VpcId"`
            Spec     string `json:"Spec" xml:"Spec"`
            InstanceChargeType     string `json:"InstanceChargeType" xml:"InstanceChargeType"`
            BusinessStatus     string `json:"BusinessStatus" xml:"BusinessStatus"`
            CreationTime     string `json:"CreationTime" xml:"CreationTime"`
            Status     string `json:"Status" xml:"Status"`
                ForwardTableIds ForwardTableIdsInDescribeNatGateways `json:"ForwardTableIds" xml:"ForwardTableIds"`
                BandwidthPackageIds BandwidthPackageIdsInDescribeNatGateways `json:"BandwidthPackageIds" xml:"BandwidthPackageIds"`
}
