
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

// VerifyPassword invokes the r_kvstore.VerifyPassword API synchronously
// api document: https://help.aliyun.com/api/r-kvstore/verifypassword.html
func (client *Client) VerifyPassword(request *VerifyPasswordRequest) (response *VerifyPasswordResponse, err error) {
response = CreateVerifyPasswordResponse()
err = client.DoAction(request, response)
return
}

// VerifyPasswordWithChan invokes the r_kvstore.VerifyPassword API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/verifypassword.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) VerifyPasswordWithChan(request *VerifyPasswordRequest) (<-chan *VerifyPasswordResponse, <-chan error) {
responseChan := make(chan *VerifyPasswordResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.VerifyPassword(request)
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

// VerifyPasswordWithCallback invokes the r_kvstore.VerifyPassword API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/verifypassword.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) VerifyPasswordWithCallback(request *VerifyPasswordRequest, callback func(response *VerifyPasswordResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *VerifyPasswordResponse
var err error
defer close(result)
response, err = client.VerifyPassword(request)
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

// VerifyPasswordRequest is the request struct for api VerifyPassword
type VerifyPasswordRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    Password     string `position:"Query" name:"Password"`
                    SecurityToken     string `position:"Query" name:"SecurityToken"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    OwnerAccount     string `position:"Query" name:"OwnerAccount"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
                    InstanceId     string `position:"Query" name:"InstanceId"`
}


// VerifyPasswordResponse is the response struct for api VerifyPassword
type VerifyPasswordResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
}

// CreateVerifyPasswordRequest creates a request to invoke VerifyPassword API
func CreateVerifyPasswordRequest() (request *VerifyPasswordRequest) {
request = &VerifyPasswordRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("R-kvstore", "2015-01-01", "VerifyPassword", "kvstore", "openAPI")
return
}

// CreateVerifyPasswordResponse creates a response to parse from VerifyPassword response
func CreateVerifyPasswordResponse() (response *VerifyPasswordResponse) {
response = &VerifyPasswordResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

