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

// DeploymentSet is a nested struct in ecs response
type DeploymentSet struct {
	DeploymentSetId          string      `json:"DeploymentSetId" xml:"DeploymentSetId"`
	DeploymentSetDescription string      `json:"DeploymentSetDescription" xml:"DeploymentSetDescription"`
	DeploymentSetName        string      `json:"DeploymentSetName" xml:"DeploymentSetName"`
	Strategy                 string      `json:"Strategy" xml:"Strategy"`
	DeploymentStrategy       string      `json:"DeploymentStrategy" xml:"DeploymentStrategy"`
	Domain                   string      `json:"Domain" xml:"Domain"`
	Granularity              string      `json:"Granularity" xml:"Granularity"`
	InstanceAmount           int         `json:"InstanceAmount" xml:"InstanceAmount"`
	CreationTime             string      `json:"CreationTime" xml:"CreationTime"`
	InstanceIds              InstanceIds `json:"InstanceIds" xml:"InstanceIds"`
}
