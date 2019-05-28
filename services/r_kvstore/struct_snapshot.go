
package r_kvstore

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

// Snapshot is a nested struct in r_kvstore response
type Snapshot struct {
            SnapshotId     string `json:"SnapshotId" xml:"SnapshotId"`
            SnapshotName     string `json:"SnapshotName" xml:"SnapshotName"`
            InstanceId     string `json:"InstanceId" xml:"InstanceId"`
            CreateTime     string `json:"CreateTime" xml:"CreateTime"`
            Memory     int64 `json:"Memory" xml:"Memory"`
            RdbSize     int64 `json:"RdbSize" xml:"RdbSize"`
            Status     string `json:"Status" xml:"Status"`
            Type     string `json:"Type" xml:"Type"`
            OssDownloadInPath     string `json:"OssDownloadInPath" xml:"OssDownloadInPath"`
            OssDownloadOutPath     string `json:"OssDownloadOutPath" xml:"OssDownloadOutPath"`
}
