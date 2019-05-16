
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

// AttachVolume_GatedLaunch invokes the ecs.AttachVolume_GatedLaunch API synchronously
// api document: https://help.aliyun.com/api/ecs/attachvolume_gatedlaunch.html
func (client *Client) AttachVolume_GatedLaunch(request *AttachVolume_GatedLaunchRequest) (response *AttachVolume_GatedLaunchResponse, err error) {
response = CreateAttachVolume_GatedLaunchResponse()
err = client.DoAction(request, response)
return
}

// AttachVolume_GatedLaunchWithChan invokes the ecs.AttachVolume_GatedLaunch API asynchronously
// api document: https://help.aliyun.com/api/ecs/attachvolume_gatedlaunch.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AttachVolume_GatedLaunchWithChan(request *AttachVolume_GatedLaunchRequest) (<-chan *AttachVolume_GatedLaunchResponse, <-chan error) {
responseChan := make(chan *AttachVolume_GatedLaunchResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.AttachVolume_GatedLaunch(request)
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

// AttachVolume_GatedLaunchWithCallback invokes the ecs.AttachVolume_GatedLaunch API asynchronously
// api document: https://help.aliyun.com/api/ecs/attachvolume_gatedlaunch.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AttachVolume_GatedLaunchWithCallback(request *AttachVolume_GatedLaunchRequest, callback func(response *AttachVolume_GatedLaunchResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *AttachVolume_GatedLaunchResponse
var err error
defer close(result)
response, err = client.AttachVolume_GatedLaunch(request)
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

// AttachVolume_GatedLaunchRequest is the request struct for api AttachVolume_GatedLaunch
type AttachVolume_GatedLaunchRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    OwnerAccount     string `position:"Query" name:"OwnerAccount"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
                    InstanceId     string `position:"Query" name:"InstanceId"`
                    VolumeId     string `position:"Query" name:"VolumeId"`
}


// AttachVolume_GatedLaunchResponse is the response struct for api AttachVolume_GatedLaunch
type AttachVolume_GatedLaunchResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
}

// CreateAttachVolume_GatedLaunchRequest creates a request to invoke AttachVolume_GatedLaunch API
func CreateAttachVolume_GatedLaunchRequest() (request *AttachVolume_GatedLaunchRequest) {
request = &AttachVolume_GatedLaunchRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Ecs", "2014-05-26", "AttachVolume_GatedLaunch", "", "")
return
}

// CreateAttachVolume_GatedLaunchResponse creates a response to parse from AttachVolume_GatedLaunch response
func CreateAttachVolume_GatedLaunchResponse() (response *AttachVolume_GatedLaunchResponse) {
response = &AttachVolume_GatedLaunchResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


