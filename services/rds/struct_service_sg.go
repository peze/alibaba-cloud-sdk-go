package rds

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

// ServiceSg is a nested struct in rds response
type ServiceSg struct {
	CreatedAt string `json:"CreatedAt" xml:"CreatedAt"`
	UpdateAt  string `json:"UpdateAt" xml:"UpdateAt"`
	GroupName string `json:"GroupName" xml:"GroupName"`
	NetType   string `json:"NetType" xml:"NetType"`
	ServiceId string `json:"ServiceId" xml:"ServiceId"`
	Ips       string `json:"Ips" xml:"Ips"`
	Enabled   string `json:"Enabled" xml:"Enabled"`
}