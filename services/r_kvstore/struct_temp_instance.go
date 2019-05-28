
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

// TempInstance is a nested struct in r_kvstore response
type TempInstance struct {
            InstanceId     string `json:"InstanceId" xml:"InstanceId"`
            TempInstanceId     string `json:"TempInstanceId" xml:"TempInstanceId"`
            SnapshotId     string `json:"SnapshotId" xml:"SnapshotId"`
            CreateTime     string `json:"CreateTime" xml:"CreateTime"`
            Domain     string `json:"Domain" xml:"Domain"`
            Status     string `json:"Status" xml:"Status"`
            Memory     int64 `json:"Memory" xml:"Memory"`
            ExpireTime     string `json:"ExpireTime" xml:"ExpireTime"`
}
