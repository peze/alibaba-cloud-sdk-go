
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

import (
"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// CreateAutoSnapshotPolicy_GatedLaunch invokes the ecs.CreateAutoSnapshotPolicy_GatedLaunch API synchronously
// api document: https://help.aliyun.com/api/ecs/createautosnapshotpolicy_gatedlaunch.html
func (client *Client) CreateAutoSnapshotPolicy_GatedLaunch(request *CreateAutoSnapshotPolicy_GatedLaunchRequest) (response *CreateAutoSnapshotPolicy_GatedLaunchResponse, err error) {
response = CreateCreateAutoSnapshotPolicy_GatedLaunchResponse()
err = client.DoAction(request, response)
return
}

// CreateAutoSnapshotPolicy_GatedLaunchWithChan invokes the ecs.CreateAutoSnapshotPolicy_GatedLaunch API asynchronously
// api document: https://help.aliyun.com/api/ecs/createautosnapshotpolicy_gatedlaunch.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateAutoSnapshotPolicy_GatedLaunchWithChan(request *CreateAutoSnapshotPolicy_GatedLaunchRequest) (<-chan *CreateAutoSnapshotPolicy_GatedLaunchResponse, <-chan error) {
responseChan := make(chan *CreateAutoSnapshotPolicy_GatedLaunchResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.CreateAutoSnapshotPolicy_GatedLaunch(request)
if err != nil {
errChan <- err
} else {
responseChan <- response
}
})
if err != nil {
errChan <- err
close(responseChan)
close(errChan)
}
return responseChan, errChan
}

// CreateAutoSnapshotPolicy_GatedLaunchWithCallback invokes the ecs.CreateAutoSnapshotPolicy_GatedLaunch API asynchronously
// api document: https://help.aliyun.com/api/ecs/createautosnapshotpolicy_gatedlaunch.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateAutoSnapshotPolicy_GatedLaunchWithCallback(request *CreateAutoSnapshotPolicy_GatedLaunchRequest, callback func(response *CreateAutoSnapshotPolicy_GatedLaunchResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *CreateAutoSnapshotPolicy_GatedLaunchResponse
var err error
defer close(result)
response, err = client.CreateAutoSnapshotPolicy_GatedLaunch(request)
callback(response, err)
result <- 1
})
if err != nil {
defer close(result)
callback(nil, err)
result <- 0
}
return result
}

// CreateAutoSnapshotPolicy_GatedLaunchRequest is the request struct for api CreateAutoSnapshotPolicy_GatedLaunch
type CreateAutoSnapshotPolicy_GatedLaunchRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    TimePoints     string `position:"Query" name:"timePoints"`
                    RepeatWeekdays     string `position:"Query" name:"repeatWeekdays"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
                    AutoSnapshotPolicyName     string `position:"Query" name:"autoSnapshotPolicyName"`
                    RetentionDays     requests.Integer `position:"Query" name:"retentionDays"`
}


// CreateAutoSnapshotPolicy_GatedLaunchResponse is the response struct for api CreateAutoSnapshotPolicy_GatedLaunch
type CreateAutoSnapshotPolicy_GatedLaunchResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            AutoSnapshotPolicyId     string `json:"AutoSnapshotPolicyId" xml:"AutoSnapshotPolicyId"`
}

// CreateCreateAutoSnapshotPolicy_GatedLaunchRequest creates a request to invoke CreateAutoSnapshotPolicy_GatedLaunch API
func CreateCreateAutoSnapshotPolicy_GatedLaunchRequest() (request *CreateAutoSnapshotPolicy_GatedLaunchRequest) {
request = &CreateAutoSnapshotPolicy_GatedLaunchRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Ecs", "2014-05-26", "CreateAutoSnapshotPolicy_GatedLaunch", "", "")
return
}

// CreateCreateAutoSnapshotPolicy_GatedLaunchResponse creates a response to parse from CreateAutoSnapshotPolicy_GatedLaunch response
func CreateCreateAutoSnapshotPolicy_GatedLaunchResponse() (response *CreateAutoSnapshotPolicy_GatedLaunchResponse) {
response = &CreateAutoSnapshotPolicy_GatedLaunchResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


