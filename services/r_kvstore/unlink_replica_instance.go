
package r_kvstore

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

// UnlinkReplicaInstance invokes the r_kvstore.UnlinkReplicaInstance API synchronously
// api document: https://help.aliyun.com/api/r-kvstore/unlinkreplicainstance.html
func (client *Client) UnlinkReplicaInstance(request *UnlinkReplicaInstanceRequest) (response *UnlinkReplicaInstanceResponse, err error) {
response = CreateUnlinkReplicaInstanceResponse()
err = client.DoAction(request, response)
return
}

// UnlinkReplicaInstanceWithChan invokes the r_kvstore.UnlinkReplicaInstance API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/unlinkreplicainstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UnlinkReplicaInstanceWithChan(request *UnlinkReplicaInstanceRequest) (<-chan *UnlinkReplicaInstanceResponse, <-chan error) {
responseChan := make(chan *UnlinkReplicaInstanceResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.UnlinkReplicaInstance(request)
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

// UnlinkReplicaInstanceWithCallback invokes the r_kvstore.UnlinkReplicaInstance API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/unlinkreplicainstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UnlinkReplicaInstanceWithCallback(request *UnlinkReplicaInstanceRequest, callback func(response *UnlinkReplicaInstanceResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *UnlinkReplicaInstanceResponse
var err error
defer close(result)
response, err = client.UnlinkReplicaInstance(request)
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

// UnlinkReplicaInstanceRequest is the request struct for api UnlinkReplicaInstance
type UnlinkReplicaInstanceRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    SecurityToken     string `position:"Query" name:"SecurityToken"`
                    ReplicaId     string `position:"Query" name:"ReplicaId"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    OwnerAccount     string `position:"Query" name:"OwnerAccount"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
                    InstanceId     string `position:"Query" name:"InstanceId"`
}


// UnlinkReplicaInstanceResponse is the response struct for api UnlinkReplicaInstance
type UnlinkReplicaInstanceResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            InstanceId     string `json:"InstanceId" xml:"InstanceId"`
            WorkflowId     string `json:"WorkflowId" xml:"WorkflowId"`
            ReplicaId     string `json:"ReplicaId" xml:"ReplicaId"`
}

// CreateUnlinkReplicaInstanceRequest creates a request to invoke UnlinkReplicaInstance API
func CreateUnlinkReplicaInstanceRequest() (request *UnlinkReplicaInstanceRequest) {
request = &UnlinkReplicaInstanceRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("R-kvstore", "2015-01-01", "UnlinkReplicaInstance", "kvstore", "openAPI")
return
}

// CreateUnlinkReplicaInstanceResponse creates a response to parse from UnlinkReplicaInstance response
func CreateUnlinkReplicaInstanceResponse() (response *UnlinkReplicaInstanceResponse) {
response = &UnlinkReplicaInstanceResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

