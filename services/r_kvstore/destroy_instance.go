
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

// DestroyInstance invokes the r_kvstore.DestroyInstance API synchronously
// api document: https://help.aliyun.com/api/r-kvstore/destroyinstance.html
func (client *Client) DestroyInstance(request *DestroyInstanceRequest) (response *DestroyInstanceResponse, err error) {
response = CreateDestroyInstanceResponse()
err = client.DoAction(request, response)
return
}

// DestroyInstanceWithChan invokes the r_kvstore.DestroyInstance API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/destroyinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DestroyInstanceWithChan(request *DestroyInstanceRequest) (<-chan *DestroyInstanceResponse, <-chan error) {
responseChan := make(chan *DestroyInstanceResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.DestroyInstance(request)
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

// DestroyInstanceWithCallback invokes the r_kvstore.DestroyInstance API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/destroyinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DestroyInstanceWithCallback(request *DestroyInstanceRequest, callback func(response *DestroyInstanceResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *DestroyInstanceResponse
var err error
defer close(result)
response, err = client.DestroyInstance(request)
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

// DestroyInstanceRequest is the request struct for api DestroyInstance
type DestroyInstanceRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    SecurityToken     string `position:"Query" name:"SecurityToken"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    OwnerAccount     string `position:"Query" name:"OwnerAccount"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
                    InstanceId     string `position:"Query" name:"InstanceId"`
}


// DestroyInstanceResponse is the response struct for api DestroyInstance
type DestroyInstanceResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
}

// CreateDestroyInstanceRequest creates a request to invoke DestroyInstance API
func CreateDestroyInstanceRequest() (request *DestroyInstanceRequest) {
request = &DestroyInstanceRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("R-kvstore", "2015-01-01", "DestroyInstance", "kvstore", "openAPI")
return
}

// CreateDestroyInstanceResponse creates a response to parse from DestroyInstance response
func CreateDestroyInstanceResponse() (response *DestroyInstanceResponse) {
response = &DestroyInstanceResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


