
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

// AppendReplicaInstance invokes the r_kvstore.AppendReplicaInstance API synchronously
// api document: https://help.aliyun.com/api/r-kvstore/appendreplicainstance.html
func (client *Client) AppendReplicaInstance(request *AppendReplicaInstanceRequest) (response *AppendReplicaInstanceResponse, err error) {
response = CreateAppendReplicaInstanceResponse()
err = client.DoAction(request, response)
return
}

// AppendReplicaInstanceWithChan invokes the r_kvstore.AppendReplicaInstance API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/appendreplicainstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AppendReplicaInstanceWithChan(request *AppendReplicaInstanceRequest) (<-chan *AppendReplicaInstanceResponse, <-chan error) {
responseChan := make(chan *AppendReplicaInstanceResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.AppendReplicaInstance(request)
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

// AppendReplicaInstanceWithCallback invokes the r_kvstore.AppendReplicaInstance API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/appendreplicainstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AppendReplicaInstanceWithCallback(request *AppendReplicaInstanceRequest, callback func(response *AppendReplicaInstanceResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *AppendReplicaInstanceResponse
var err error
defer close(result)
response, err = client.AppendReplicaInstance(request)
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

// AppendReplicaInstanceRequest is the request struct for api AppendReplicaInstance
type AppendReplicaInstanceRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    ClientToken     string `position:"Query" name:"ClientToken"`
                    NetworkType     string `position:"Query" name:"NetworkType"`
                    SecurityToken     string `position:"Query" name:"SecurityToken"`
                    ReplicaId     string `position:"Query" name:"ReplicaId"`
                    Period     requests.Integer `position:"Query" name:"Period"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    OwnerAccount     string `position:"Query" name:"OwnerAccount"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
                    VSwitchId     string `position:"Query" name:"VSwitchId"`
                    InstanceName     string `position:"Query" name:"InstanceName"`
                    VpcId     string `position:"Query" name:"VpcId"`
                    ZoneId     string `position:"Query" name:"ZoneId"`
                    ChargeType     string `position:"Query" name:"ChargeType"`
}


// AppendReplicaInstanceResponse is the response struct for api AppendReplicaInstance
type AppendReplicaInstanceResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            InstanceId     string `json:"InstanceId" xml:"InstanceId"`
            OrderId     int64 `json:"OrderId" xml:"OrderId"`
            ReplicaId     string `json:"ReplicaId" xml:"ReplicaId"`
}

// CreateAppendReplicaInstanceRequest creates a request to invoke AppendReplicaInstance API
func CreateAppendReplicaInstanceRequest() (request *AppendReplicaInstanceRequest) {
request = &AppendReplicaInstanceRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("R-kvstore", "2015-01-01", "AppendReplicaInstance", "kvstore", "openAPI")
return
}

// CreateAppendReplicaInstanceResponse creates a response to parse from AppendReplicaInstance response
func CreateAppendReplicaInstanceResponse() (response *AppendReplicaInstanceResponse) {
response = &AppendReplicaInstanceResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

