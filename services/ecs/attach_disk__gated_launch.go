
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

// AttachDisk_GatedLaunch invokes the ecs.AttachDisk_GatedLaunch API synchronously
// api document: https://help.aliyun.com/api/ecs/attachdisk_gatedlaunch.html
func (client *Client) AttachDisk_GatedLaunch(request *AttachDisk_GatedLaunchRequest) (response *AttachDisk_GatedLaunchResponse, err error) {
response = CreateAttachDisk_GatedLaunchResponse()
err = client.DoAction(request, response)
return
}

// AttachDisk_GatedLaunchWithChan invokes the ecs.AttachDisk_GatedLaunch API asynchronously
// api document: https://help.aliyun.com/api/ecs/attachdisk_gatedlaunch.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AttachDisk_GatedLaunchWithChan(request *AttachDisk_GatedLaunchRequest) (<-chan *AttachDisk_GatedLaunchResponse, <-chan error) {
responseChan := make(chan *AttachDisk_GatedLaunchResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.AttachDisk_GatedLaunch(request)
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

// AttachDisk_GatedLaunchWithCallback invokes the ecs.AttachDisk_GatedLaunch API asynchronously
// api document: https://help.aliyun.com/api/ecs/attachdisk_gatedlaunch.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AttachDisk_GatedLaunchWithCallback(request *AttachDisk_GatedLaunchRequest, callback func(response *AttachDisk_GatedLaunchResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *AttachDisk_GatedLaunchResponse
var err error
defer close(result)
response, err = client.AttachDisk_GatedLaunch(request)
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

// AttachDisk_GatedLaunchRequest is the request struct for api AttachDisk_GatedLaunch
type AttachDisk_GatedLaunchRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    DiskId     string `position:"Query" name:"DiskId"`
                    DeleteWithInstance     requests.Boolean `position:"Query" name:"DeleteWithInstance"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    OwnerAccount     string `position:"Query" name:"OwnerAccount"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
                    InstanceId     string `position:"Query" name:"InstanceId"`
                    Device     string `position:"Query" name:"Device"`
}


// AttachDisk_GatedLaunchResponse is the response struct for api AttachDisk_GatedLaunch
type AttachDisk_GatedLaunchResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
}

// CreateAttachDisk_GatedLaunchRequest creates a request to invoke AttachDisk_GatedLaunch API
func CreateAttachDisk_GatedLaunchRequest() (request *AttachDisk_GatedLaunchRequest) {
request = &AttachDisk_GatedLaunchRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Ecs", "2014-05-26", "AttachDisk_GatedLaunch", "", "")
return
}

// CreateAttachDisk_GatedLaunchResponse creates a response to parse from AttachDisk_GatedLaunch response
func CreateAttachDisk_GatedLaunchResponse() (response *AttachDisk_GatedLaunchResponse) {
response = &AttachDisk_GatedLaunchResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


