
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

// ModifyImageShareGroupPermission_GatedLaunch invokes the ecs.ModifyImageShareGroupPermission_GatedLaunch API synchronously
// api document: https://help.aliyun.com/api/ecs/modifyimagesharegrouppermission_gatedlaunch.html
func (client *Client) ModifyImageShareGroupPermission_GatedLaunch(request *ModifyImageShareGroupPermission_GatedLaunchRequest) (response *ModifyImageShareGroupPermission_GatedLaunchResponse, err error) {
response = CreateModifyImageShareGroupPermission_GatedLaunchResponse()
err = client.DoAction(request, response)
return
}

// ModifyImageShareGroupPermission_GatedLaunchWithChan invokes the ecs.ModifyImageShareGroupPermission_GatedLaunch API asynchronously
// api document: https://help.aliyun.com/api/ecs/modifyimagesharegrouppermission_gatedlaunch.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyImageShareGroupPermission_GatedLaunchWithChan(request *ModifyImageShareGroupPermission_GatedLaunchRequest) (<-chan *ModifyImageShareGroupPermission_GatedLaunchResponse, <-chan error) {
responseChan := make(chan *ModifyImageShareGroupPermission_GatedLaunchResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.ModifyImageShareGroupPermission_GatedLaunch(request)
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

// ModifyImageShareGroupPermission_GatedLaunchWithCallback invokes the ecs.ModifyImageShareGroupPermission_GatedLaunch API asynchronously
// api document: https://help.aliyun.com/api/ecs/modifyimagesharegrouppermission_gatedlaunch.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyImageShareGroupPermission_GatedLaunchWithCallback(request *ModifyImageShareGroupPermission_GatedLaunchRequest, callback func(response *ModifyImageShareGroupPermission_GatedLaunchResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *ModifyImageShareGroupPermission_GatedLaunchResponse
var err error
defer close(result)
response, err = client.ModifyImageShareGroupPermission_GatedLaunch(request)
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

// ModifyImageShareGroupPermission_GatedLaunchRequest is the request struct for api ModifyImageShareGroupPermission_GatedLaunch
type ModifyImageShareGroupPermission_GatedLaunchRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    ImageId     string `position:"Query" name:"ImageId"`
                    AddGroup1     string `position:"Query" name:"AddGroup.1"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    OwnerAccount     string `position:"Query" name:"OwnerAccount"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
                    RemoveGroup1     string `position:"Query" name:"RemoveGroup.1"`
}


// ModifyImageShareGroupPermission_GatedLaunchResponse is the response struct for api ModifyImageShareGroupPermission_GatedLaunch
type ModifyImageShareGroupPermission_GatedLaunchResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyImageShareGroupPermission_GatedLaunchRequest creates a request to invoke ModifyImageShareGroupPermission_GatedLaunch API
func CreateModifyImageShareGroupPermission_GatedLaunchRequest() (request *ModifyImageShareGroupPermission_GatedLaunchRequest) {
request = &ModifyImageShareGroupPermission_GatedLaunchRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Ecs", "2014-05-26", "ModifyImageShareGroupPermission_GatedLaunch", "", "")
return
}

// CreateModifyImageShareGroupPermission_GatedLaunchResponse creates a response to parse from ModifyImageShareGroupPermission_GatedLaunch response
func CreateModifyImageShareGroupPermission_GatedLaunchResponse() (response *ModifyImageShareGroupPermission_GatedLaunchResponse) {
response = &ModifyImageShareGroupPermission_GatedLaunchResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


