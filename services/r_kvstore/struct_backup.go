
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

// Backup is a nested struct in r_kvstore response
type Backup struct {
            BackupId     int `json:"BackupId" xml:"BackupId"`
            BackupDBNames     string `json:"BackupDBNames" xml:"BackupDBNames"`
            BackupStatus     string `json:"BackupStatus" xml:"BackupStatus"`
            BackupStartTime     string `json:"BackupStartTime" xml:"BackupStartTime"`
            BackupEndTime     string `json:"BackupEndTime" xml:"BackupEndTime"`
            BackupType     string `json:"BackupType" xml:"BackupType"`
            BackupMode     string `json:"BackupMode" xml:"BackupMode"`
            BackupMethod     string `json:"BackupMethod" xml:"BackupMethod"`
            BackupDownloadURL     string `json:"BackupDownloadURL" xml:"BackupDownloadURL"`
            BackupSize     int64 `json:"BackupSize" xml:"BackupSize"`
            EngineVersion     string `json:"EngineVersion" xml:"EngineVersion"`
            NodeInstanceId     string `json:"NodeInstanceId" xml:"NodeInstanceId"`
            BackupIntranetDownloadURL     string `json:"BackupIntranetDownloadURL" xml:"BackupIntranetDownloadURL"`
}
