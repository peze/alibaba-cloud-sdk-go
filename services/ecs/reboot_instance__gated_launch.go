
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

// RebootInstance_GatedLaunch invokes the ecs.RebootInstance_GatedLaunch API synchronously
// api document: https://help.aliyun.com/api/ecs/rebootinstance_gatedlaunch.html
func (client *Client) RebootInstance_GatedLaunch(request *RebootInstance_GatedLaunchRequest) (response *RebootInstance_GatedLaunchResponse, err error) {
response = CreateRebootInstance_GatedLaunchResponse()
err = client.DoAction(request, response)
return
}

// RebootInstance_GatedLaunchWithChan invokes the ecs.RebootInstance_GatedLaunch API asynchronously
// api document: https://help.aliyun.com/api/ecs/rebootinstance_gatedlaunch.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RebootInstance_GatedLaunchWithChan(request *RebootInstance_GatedLaunchRequest) (<-chan *RebootInstance_GatedLaunchResponse, <-chan error) {
responseChan := make(chan *RebootInstance_GatedLaunchResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.RebootInstance_GatedLaunch(request)
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

// RebootInstance_GatedLaunchWithCallback invokes the ecs.RebootInstance_GatedLaunch API asynchronously
// api document: https://help.aliyun.com/api/ecs/rebootinstance_gatedlaunch.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RebootInstance_GatedLaunchWithCallback(request *RebootInstance_GatedLaunchRequest, callback func(response *RebootInstance_GatedLaunchResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *RebootInstance_GatedLaunchResponse
var err error
defer close(result)
response, err = client.RebootInstance_GatedLaunch(request)
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

// RebootInstance_GatedLaunchRequest is the request struct for api RebootInstance_GatedLaunch
type RebootInstance_GatedLaunchRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    ForceStop     requests.Boolean `position:"Query" name:"ForceStop"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    OwnerAccount     string `position:"Query" name:"OwnerAccount"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
                    InstanceId     string `position:"Query" name:"InstanceId"`
}


// RebootInstance_GatedLaunchResponse is the response struct for api RebootInstance_GatedLaunch
type RebootInstance_GatedLaunchResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
}

// CreateRebootInstance_GatedLaunchRequest creates a request to invoke RebootInstance_GatedLaunch API
func CreateRebootInstance_GatedLaunchRequest() (request *RebootInstance_GatedLaunchRequest) {
request = &RebootInstance_GatedLaunchRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Ecs", "2014-05-26", "RebootInstance_GatedLaunch", "", "")
return
}

// CreateRebootInstance_GatedLaunchResponse creates a response to parse from RebootInstance_GatedLaunch response
func CreateRebootInstance_GatedLaunchResponse() (response *RebootInstance_GatedLaunchResponse) {
response = &RebootInstance_GatedLaunchResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


