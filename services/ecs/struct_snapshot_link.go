
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

// SnapshotLink is a nested struct in ecs response
type SnapshotLink struct {
            SnapshotLinkId     string `json:"SnapshotLinkId" xml:"SnapshotLinkId"`
            RegionId     string `json:"RegionId" xml:"RegionId"`
            InstanceId     string `json:"InstanceId" xml:"InstanceId"`
            InstanceName     string `json:"InstanceName" xml:"InstanceName"`
            SourceDiskId     string `json:"SourceDiskId" xml:"SourceDiskId"`
            SourceDiskName     string `json:"SourceDiskName" xml:"SourceDiskName"`
            SourceDiskSize     int `json:"SourceDiskSize" xml:"SourceDiskSize"`
            SourceDiskType     string `json:"SourceDiskType" xml:"SourceDiskType"`
            TotalSize     int `json:"TotalSize" xml:"TotalSize"`
            TotalCount     int `json:"TotalCount" xml:"TotalCount"`
}
