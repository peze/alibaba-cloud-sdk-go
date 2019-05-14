
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

// CreateDisk_GatedLaunch invokes the ecs.CreateDisk_GatedLaunch API synchronously
// api document: https://help.aliyun.com/api/ecs/createdisk_gatedlaunch.html
func (client *Client) CreateDisk_GatedLaunch(request *CreateDisk_GatedLaunchRequest) (response *CreateDisk_GatedLaunchResponse, err error) {
response = CreateCreateDisk_GatedLaunchResponse()
err = client.DoAction(request, response)
return
}

// CreateDisk_GatedLaunchWithChan invokes the ecs.CreateDisk_GatedLaunch API asynchronously
// api document: https://help.aliyun.com/api/ecs/createdisk_gatedlaunch.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateDisk_GatedLaunchWithChan(request *CreateDisk_GatedLaunchRequest) (<-chan *CreateDisk_GatedLaunchResponse, <-chan error) {
responseChan := make(chan *CreateDisk_GatedLaunchResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.CreateDisk_GatedLaunch(request)
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

// CreateDisk_GatedLaunchWithCallback invokes the ecs.CreateDisk_GatedLaunch API asynchronously
// api document: https://help.aliyun.com/api/ecs/createdisk_gatedlaunch.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateDisk_GatedLaunchWithCallback(request *CreateDisk_GatedLaunchRequest, callback func(response *CreateDisk_GatedLaunchResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *CreateDisk_GatedLaunchResponse
var err error
defer close(result)
response, err = client.CreateDisk_GatedLaunch(request)
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

// CreateDisk_GatedLaunchRequest is the request struct for api CreateDisk_GatedLaunch
type CreateDisk_GatedLaunchRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    SnapshotId     string `position:"Query" name:"SnapshotId"`
                    ClientToken     string `position:"Query" name:"ClientToken"`
                    Description     string `position:"Query" name:"Description"`
                    DiskName     string `position:"Query" name:"DiskName"`
                    DiskCategory     string `position:"Query" name:"DiskCategory"`
                    Tag  *[]CreateDisk_GatedLaunchTag `position:"Query" name:"Tag"  type:"Repeated"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    OwnerAccount     string `position:"Query" name:"OwnerAccount"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
                    Size     requests.Integer `position:"Query" name:"Size"`
                    ZoneId     string `position:"Query" name:"ZoneId"`
}

// CreateDisk_GatedLaunchTag is a repeated param struct in CreateDisk_GatedLaunchRequest
type CreateDisk_GatedLaunchTag struct{
        Value     string `name:"Value"`
        Key     string `name:"Key"`
}

// CreateDisk_GatedLaunchResponse is the response struct for api CreateDisk_GatedLaunch
type CreateDisk_GatedLaunchResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            DiskId     string `json:"DiskId" xml:"DiskId"`
}

// CreateCreateDisk_GatedLaunchRequest creates a request to invoke CreateDisk_GatedLaunch API
func CreateCreateDisk_GatedLaunchRequest() (request *CreateDisk_GatedLaunchRequest) {
request = &CreateDisk_GatedLaunchRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Ecs", "2014-05-26", "CreateDisk_GatedLaunch", "ecs", "openAPI")
return
}

// CreateCreateDisk_GatedLaunchResponse creates a response to parse from CreateDisk_GatedLaunch response
func CreateCreateDisk_GatedLaunchResponse() (response *CreateDisk_GatedLaunchResponse) {
response = &CreateDisk_GatedLaunchResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

