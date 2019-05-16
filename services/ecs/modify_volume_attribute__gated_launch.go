
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

// ModifyVolumeAttribute_GatedLaunch invokes the ecs.ModifyVolumeAttribute_GatedLaunch API synchronously
// api document: https://help.aliyun.com/api/ecs/modifyvolumeattribute_gatedlaunch.html
func (client *Client) ModifyVolumeAttribute_GatedLaunch(request *ModifyVolumeAttribute_GatedLaunchRequest) (response *ModifyVolumeAttribute_GatedLaunchResponse, err error) {
response = CreateModifyVolumeAttribute_GatedLaunchResponse()
err = client.DoAction(request, response)
return
}

// ModifyVolumeAttribute_GatedLaunchWithChan invokes the ecs.ModifyVolumeAttribute_GatedLaunch API asynchronously
// api document: https://help.aliyun.com/api/ecs/modifyvolumeattribute_gatedlaunch.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyVolumeAttribute_GatedLaunchWithChan(request *ModifyVolumeAttribute_GatedLaunchRequest) (<-chan *ModifyVolumeAttribute_GatedLaunchResponse, <-chan error) {
responseChan := make(chan *ModifyVolumeAttribute_GatedLaunchResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.ModifyVolumeAttribute_GatedLaunch(request)
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

// ModifyVolumeAttribute_GatedLaunchWithCallback invokes the ecs.ModifyVolumeAttribute_GatedLaunch API asynchronously
// api document: https://help.aliyun.com/api/ecs/modifyvolumeattribute_gatedlaunch.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyVolumeAttribute_GatedLaunchWithCallback(request *ModifyVolumeAttribute_GatedLaunchRequest, callback func(response *ModifyVolumeAttribute_GatedLaunchResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *ModifyVolumeAttribute_GatedLaunchResponse
var err error
defer close(result)
response, err = client.ModifyVolumeAttribute_GatedLaunch(request)
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

// ModifyVolumeAttribute_GatedLaunchRequest is the request struct for api ModifyVolumeAttribute_GatedLaunch
type ModifyVolumeAttribute_GatedLaunchRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    Description     string `position:"Query" name:"Description"`
                    VolumeName     string `position:"Query" name:"VolumeName"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    OwnerAccount     string `position:"Query" name:"OwnerAccount"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
                    VolumeId     string `position:"Query" name:"VolumeId"`
}


// ModifyVolumeAttribute_GatedLaunchResponse is the response struct for api ModifyVolumeAttribute_GatedLaunch
type ModifyVolumeAttribute_GatedLaunchResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyVolumeAttribute_GatedLaunchRequest creates a request to invoke ModifyVolumeAttribute_GatedLaunch API
func CreateModifyVolumeAttribute_GatedLaunchRequest() (request *ModifyVolumeAttribute_GatedLaunchRequest) {
request = &ModifyVolumeAttribute_GatedLaunchRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Ecs", "2014-05-26", "ModifyVolumeAttribute_GatedLaunch", "", "")
return
}

// CreateModifyVolumeAttribute_GatedLaunchResponse creates a response to parse from ModifyVolumeAttribute_GatedLaunch response
func CreateModifyVolumeAttribute_GatedLaunchResponse() (response *ModifyVolumeAttribute_GatedLaunchResponse) {
response = &ModifyVolumeAttribute_GatedLaunchResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


