
package aas

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

// UpdateAccessKeyStatusForAccount invokes the aas.UpdateAccessKeyStatusForAccount API synchronously
// api document: https://help.aliyun.com/api/aas/updateaccesskeystatusforaccount.html
func (client *Client) UpdateAccessKeyStatusForAccount(request *UpdateAccessKeyStatusForAccountRequest) (response *UpdateAccessKeyStatusForAccountResponse, err error) {
response = CreateUpdateAccessKeyStatusForAccountResponse()
err = client.DoAction(request, response)
return
}

// UpdateAccessKeyStatusForAccountWithChan invokes the aas.UpdateAccessKeyStatusForAccount API asynchronously
// api document: https://help.aliyun.com/api/aas/updateaccesskeystatusforaccount.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateAccessKeyStatusForAccountWithChan(request *UpdateAccessKeyStatusForAccountRequest) (<-chan *UpdateAccessKeyStatusForAccountResponse, <-chan error) {
responseChan := make(chan *UpdateAccessKeyStatusForAccountResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.UpdateAccessKeyStatusForAccount(request)
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

// UpdateAccessKeyStatusForAccountWithCallback invokes the aas.UpdateAccessKeyStatusForAccount API asynchronously
// api document: https://help.aliyun.com/api/aas/updateaccesskeystatusforaccount.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateAccessKeyStatusForAccountWithCallback(request *UpdateAccessKeyStatusForAccountRequest, callback func(response *UpdateAccessKeyStatusForAccountResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *UpdateAccessKeyStatusForAccountResponse
var err error
defer close(result)
response, err = client.UpdateAccessKeyStatusForAccount(request)
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

// UpdateAccessKeyStatusForAccountRequest is the request struct for api UpdateAccessKeyStatusForAccount
type UpdateAccessKeyStatusForAccountRequest struct {
*requests.RpcRequest
                    AKStatus     string `position:"Query" name:"AKStatus"`
                    AKId     string `position:"Query" name:"AKId"`
                    PK     string `position:"Query" name:"PK"`
}


// UpdateAccessKeyStatusForAccountResponse is the response struct for api UpdateAccessKeyStatusForAccount
type UpdateAccessKeyStatusForAccountResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            PK     string `json:"PK" xml:"PK"`
            Result     string `json:"Result" xml:"Result"`
}

// CreateUpdateAccessKeyStatusForAccountRequest creates a request to invoke UpdateAccessKeyStatusForAccount API
func CreateUpdateAccessKeyStatusForAccountRequest() (request *UpdateAccessKeyStatusForAccountRequest) {
request = &UpdateAccessKeyStatusForAccountRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Aas", "2015-07-01", "UpdateAccessKeyStatusForAccount", "", "")
return
}

// CreateUpdateAccessKeyStatusForAccountResponse creates a response to parse from UpdateAccessKeyStatusForAccount response
func CreateUpdateAccessKeyStatusForAccountResponse() (response *UpdateAccessKeyStatusForAccountResponse) {
response = &UpdateAccessKeyStatusForAccountResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

