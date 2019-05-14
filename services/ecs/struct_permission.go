
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

// Permission is a nested struct in ecs response
type Permission struct {
            IpProtocol     string `json:"IpProtocol" xml:"IpProtocol"`
            PortRange     string `json:"PortRange" xml:"PortRange"`
            SourcePortRange     string `json:"SourcePortRange" xml:"SourcePortRange"`
            SourceGroupId     string `json:"SourceGroupId" xml:"SourceGroupId"`
            SourceGroupName     string `json:"SourceGroupName" xml:"SourceGroupName"`
            SourceCidrIp     string `json:"SourceCidrIp" xml:"SourceCidrIp"`
            Ipv6SourceCidrIp     string `json:"Ipv6SourceCidrIp" xml:"Ipv6SourceCidrIp"`
            Policy     string `json:"Policy" xml:"Policy"`
            NicType     string `json:"NicType" xml:"NicType"`
            SourceGroupOwnerAccount     string `json:"SourceGroupOwnerAccount" xml:"SourceGroupOwnerAccount"`
            DestGroupId     string `json:"DestGroupId" xml:"DestGroupId"`
            DestGroupName     string `json:"DestGroupName" xml:"DestGroupName"`
            DestCidrIp     string `json:"DestCidrIp" xml:"DestCidrIp"`
            Ipv6DestCidrIp     string `json:"Ipv6DestCidrIp" xml:"Ipv6DestCidrIp"`
            DestGroupOwnerAccount     string `json:"DestGroupOwnerAccount" xml:"DestGroupOwnerAccount"`
            Priority     string `json:"Priority" xml:"Priority"`
            Direction     string `json:"Direction" xml:"Direction"`
            Description     string `json:"Description" xml:"Description"`
            CreateTime     string `json:"CreateTime" xml:"CreateTime"`
}
