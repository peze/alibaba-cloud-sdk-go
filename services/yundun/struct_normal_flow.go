
package yundun

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

// NormalFlow is a nested struct in yundun response
type NormalFlow struct {
            Time     int64 `json:"time" xml:"time"`
            BitRecv     int64 `json:"BitRecv" xml:"BitRecv"`
            BitSend     int64 `json:"BitSend" xml:"BitSend"`
            PktRecv     int64 `json:"PktRecv" xml:"PktRecv"`
            PktSend     int64 `json:"PktSend" xml:"PktSend"`
}
