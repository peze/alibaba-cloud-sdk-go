
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

// RollbackVolume invokes the ecs.RollbackVolume API synchronously
// api document: https://help.aliyun.com/api/ecs/rollbackvolume.html
func (client *Client) RollbackVolume(request *RollbackVolumeRequest) (response *RollbackVolumeResponse, err error) {
response = CreateRollbackVolumeResponse()
err = client.DoAction(request, response)
return
}

// RollbackVolumeWithChan invokes the ecs.RollbackVolume API asynchronously
// api document: https://help.aliyun.com/api/ecs/rollbackvolume.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RollbackVolumeWithChan(request *RollbackVolumeRequest) (<-chan *RollbackVolumeResponse, <-chan error) {
responseChan := make(chan *RollbackVolumeResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.RollbackVolume(request)
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

// RollbackVolumeWithCallback invokes the ecs.RollbackVolume API asynchronously
// api document: https://help.aliyun.com/api/ecs/rollbackvolume.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RollbackVolumeWithCallback(request *RollbackVolumeRequest, callback func(response *RollbackVolumeResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *RollbackVolumeResponse
var err error
defer close(result)
response, err = client.RollbackVolume(request)
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

// RollbackVolumeRequest is the request struct for api RollbackVolume
type RollbackVolumeRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    SnapshotId     string `position:"Query" name:"SnapshotId"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    OwnerAccount     string `position:"Query" name:"OwnerAccount"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
                    VolumeId     string `position:"Query" name:"VolumeId"`
}


// RollbackVolumeResponse is the response struct for api RollbackVolume
type RollbackVolumeResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
}

// CreateRollbackVolumeRequest creates a request to invoke RollbackVolume API
func CreateRollbackVolumeRequest() (request *RollbackVolumeRequest) {
request = &RollbackVolumeRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Ecs", "2014-05-26", "RollbackVolume", "ecs", "openAPI")
return
}

// CreateRollbackVolumeResponse creates a response to parse from RollbackVolume response
func CreateRollbackVolumeResponse() (response *RollbackVolumeResponse) {
response = &RollbackVolumeResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

