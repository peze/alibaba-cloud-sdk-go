
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

// TerminatePhysicalConnection invokes the ecs.TerminatePhysicalConnection API synchronously
// api document: https://help.aliyun.com/api/ecs/terminatephysicalconnection.html
func (client *Client) TerminatePhysicalConnection(request *TerminatePhysicalConnectionRequest) (response *TerminatePhysicalConnectionResponse, err error) {
response = CreateTerminatePhysicalConnectionResponse()
err = client.DoAction(request, response)
return
}

// TerminatePhysicalConnectionWithChan invokes the ecs.TerminatePhysicalConnection API asynchronously
// api document: https://help.aliyun.com/api/ecs/terminatephysicalconnection.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) TerminatePhysicalConnectionWithChan(request *TerminatePhysicalConnectionRequest) (<-chan *TerminatePhysicalConnectionResponse, <-chan error) {
responseChan := make(chan *TerminatePhysicalConnectionResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.TerminatePhysicalConnection(request)
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

// TerminatePhysicalConnectionWithCallback invokes the ecs.TerminatePhysicalConnection API asynchronously
// api document: https://help.aliyun.com/api/ecs/terminatephysicalconnection.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) TerminatePhysicalConnectionWithCallback(request *TerminatePhysicalConnectionRequest, callback func(response *TerminatePhysicalConnectionResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *TerminatePhysicalConnectionResponse
var err error
defer close(result)
response, err = client.TerminatePhysicalConnection(request)
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

// TerminatePhysicalConnectionRequest is the request struct for api TerminatePhysicalConnection
type TerminatePhysicalConnectionRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    ClientToken     string `position:"Query" name:"ClientToken"`
                    UserCidr     string `position:"Query" name:"UserCidr"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    OwnerAccount     string `position:"Query" name:"OwnerAccount"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
                    PhysicalConnectionId     string `position:"Query" name:"PhysicalConnectionId"`
}


// TerminatePhysicalConnectionResponse is the response struct for api TerminatePhysicalConnection
type TerminatePhysicalConnectionResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
}

// CreateTerminatePhysicalConnectionRequest creates a request to invoke TerminatePhysicalConnection API
func CreateTerminatePhysicalConnectionRequest() (request *TerminatePhysicalConnectionRequest) {
request = &TerminatePhysicalConnectionRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Ecs", "2014-05-26", "TerminatePhysicalConnection", "", "")
return
}

// CreateTerminatePhysicalConnectionResponse creates a response to parse from TerminatePhysicalConnection response
func CreateTerminatePhysicalConnectionResponse() (response *TerminatePhysicalConnectionResponse) {
response = &TerminatePhysicalConnectionResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


