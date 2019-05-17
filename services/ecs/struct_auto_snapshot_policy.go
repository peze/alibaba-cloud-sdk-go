
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

// AutoSnapshotPolicy is a nested struct in ecs response
type AutoSnapshotPolicy struct {
            DataDiskPolicyEnabled     string `json:"DataDiskPolicyEnabled" xml:"DataDiskPolicyEnabled"`
            SystemDiskPolicyRetentionLastWeek     string `json:"SystemDiskPolicyRetentionLastWeek" xml:"SystemDiskPolicyRetentionLastWeek"`
            RepeatWeekdays     string `json:"RepeatWeekdays" xml:"RepeatWeekdays"`
            AutoSnapshotPolicyId     string `json:"AutoSnapshotPolicyId" xml:"AutoSnapshotPolicyId"`
            DataDiskPolicyRetentionDays     string `json:"DataDiskPolicyRetentionDays" xml:"DataDiskPolicyRetentionDays"`
            VolumeNums     int `json:"VolumeNums" xml:"VolumeNums"`
            CreationTime     string `json:"CreationTime" xml:"CreationTime"`
            DataDiskPolicyTimePeriod     string `json:"DataDiskPolicyTimePeriod" xml:"DataDiskPolicyTimePeriod"`
            SystemDiskPolicyTimePeriod     string `json:"SystemDiskPolicyTimePeriod" xml:"SystemDiskPolicyTimePeriod"`
            SystemDiskPolicyEnabled     string `json:"SystemDiskPolicyEnabled" xml:"SystemDiskPolicyEnabled"`
            DiskNums     int `json:"DiskNums" xml:"DiskNums"`
            RegionId     string `json:"RegionId" xml:"RegionId"`
            AutoSnapshotPolicyName     string `json:"AutoSnapshotPolicyName" xml:"AutoSnapshotPolicyName"`
            RetentionDays     int `json:"RetentionDays" xml:"RetentionDays"`
            TimePoints     string `json:"TimePoints" xml:"TimePoints"`
            SystemDiskPolicyRetentionDays     string `json:"SystemDiskPolicyRetentionDays" xml:"SystemDiskPolicyRetentionDays"`
            Status     string `json:"Status" xml:"Status"`
            DataDiskPolicyRetentionLastWeek     string `json:"DataDiskPolicyRetentionLastWeek" xml:"DataDiskPolicyRetentionLastWeek"`
}
