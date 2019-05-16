
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

// ReleaseDedicatedHost invokes the ecs.ReleaseDedicatedHost API synchronously
// api document: https://help.aliyun.com/api/ecs/releasededicatedhost.html
func (client *Client) ReleaseDedicatedHost(request *ReleaseDedicatedHostRequest) (response *ReleaseDedicatedHostResponse, err error) {
response = CreateReleaseDedicatedHostResponse()
err = client.DoAction(request, response)
return
}

// ReleaseDedicatedHostWithChan invokes the ecs.ReleaseDedicatedHost API asynchronously
// api document: https://help.aliyun.com/api/ecs/releasededicatedhost.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ReleaseDedicatedHostWithChan(request *ReleaseDedicatedHostRequest) (<-chan *ReleaseDedicatedHostResponse, <-chan error) {
responseChan := make(chan *ReleaseDedicatedHostResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.ReleaseDedicatedHost(request)
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

// ReleaseDedicatedHostWithCallback invokes the ecs.ReleaseDedicatedHost API asynchronously
// api document: https://help.aliyun.com/api/ecs/releasededicatedhost.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ReleaseDedicatedHostWithCallback(request *ReleaseDedicatedHostRequest, callback func(response *ReleaseDedicatedHostResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *ReleaseDedicatedHostResponse
var err error
defer close(result)
response, err = client.ReleaseDedicatedHost(request)
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

// ReleaseDedicatedHostRequest is the request struct for api ReleaseDedicatedHost
type ReleaseDedicatedHostRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    OwnerAccount     string `position:"Query" name:"OwnerAccount"`
                    DedicatedHostId     string `position:"Query" name:"DedicatedHostId"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
}


// ReleaseDedicatedHostResponse is the response struct for api ReleaseDedicatedHost
type ReleaseDedicatedHostResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
}

// CreateReleaseDedicatedHostRequest creates a request to invoke ReleaseDedicatedHost API
func CreateReleaseDedicatedHostRequest() (request *ReleaseDedicatedHostRequest) {
request = &ReleaseDedicatedHostRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Ecs", "2014-05-26", "ReleaseDedicatedHost", "", "")
return
}

// CreateReleaseDedicatedHostResponse creates a response to parse from ReleaseDedicatedHost response
func CreateReleaseDedicatedHostResponse() (response *ReleaseDedicatedHostResponse) {
response = &ReleaseDedicatedHostResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


