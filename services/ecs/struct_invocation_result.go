
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

// InvocationResult is a nested struct in ecs response
type InvocationResult struct {
            CommandId     string `json:"CommandId" xml:"CommandId"`
            InvokeId     string `json:"InvokeId" xml:"InvokeId"`
            InstanceId     string `json:"InstanceId" xml:"InstanceId"`
            StartTime     string `json:"StartTime" xml:"StartTime"`
            FinishedTime     string `json:"FinishedTime" xml:"FinishedTime"`
            Output     string `json:"Output" xml:"Output"`
            InvokeRecordStatus     string `json:"InvokeRecordStatus" xml:"InvokeRecordStatus"`
            ExitCode     int `json:"ExitCode" xml:"ExitCode"`
}
