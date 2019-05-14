
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

// ResizeDisk_GatedLaunch invokes the ecs.ResizeDisk_GatedLaunch API synchronously
// api document: https://help.aliyun.com/api/ecs/resizedisk_gatedlaunch.html
func (client *Client) ResizeDisk_GatedLaunch(request *ResizeDisk_GatedLaunchRequest) (response *ResizeDisk_GatedLaunchResponse, err error) {
response = CreateResizeDisk_GatedLaunchResponse()
err = client.DoAction(request, response)
return
}

// ResizeDisk_GatedLaunchWithChan invokes the ecs.ResizeDisk_GatedLaunch API asynchronously
// api document: https://help.aliyun.com/api/ecs/resizedisk_gatedlaunch.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ResizeDisk_GatedLaunchWithChan(request *ResizeDisk_GatedLaunchRequest) (<-chan *ResizeDisk_GatedLaunchResponse, <-chan error) {
responseChan := make(chan *ResizeDisk_GatedLaunchResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.ResizeDisk_GatedLaunch(request)
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

// ResizeDisk_GatedLaunchWithCallback invokes the ecs.ResizeDisk_GatedLaunch API asynchronously
// api document: https://help.aliyun.com/api/ecs/resizedisk_gatedlaunch.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ResizeDisk_GatedLaunchWithCallback(request *ResizeDisk_GatedLaunchRequest, callback func(response *ResizeDisk_GatedLaunchResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *ResizeDisk_GatedLaunchResponse
var err error
defer close(result)
response, err = client.ResizeDisk_GatedLaunch(request)
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

// ResizeDisk_GatedLaunchRequest is the request struct for api ResizeDisk_GatedLaunch
type ResizeDisk_GatedLaunchRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    ClientToken     string `position:"Query" name:"ClientToken"`
                    DiskId     string `position:"Query" name:"DiskId"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    OwnerAccount     string `position:"Query" name:"OwnerAccount"`
                    NewSize     requests.Integer `position:"Query" name:"NewSize"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
}


// ResizeDisk_GatedLaunchResponse is the response struct for api ResizeDisk_GatedLaunch
type ResizeDisk_GatedLaunchResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
}

// CreateResizeDisk_GatedLaunchRequest creates a request to invoke ResizeDisk_GatedLaunch API
func CreateResizeDisk_GatedLaunchRequest() (request *ResizeDisk_GatedLaunchRequest) {
request = &ResizeDisk_GatedLaunchRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Ecs", "2014-05-26", "ResizeDisk_GatedLaunch", "ecs", "openAPI")
return
}

// CreateResizeDisk_GatedLaunchResponse creates a response to parse from ResizeDisk_GatedLaunch response
func CreateResizeDisk_GatedLaunchResponse() (response *ResizeDisk_GatedLaunchResponse) {
response = &ResizeDisk_GatedLaunchResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

