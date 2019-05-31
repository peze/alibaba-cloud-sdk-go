
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

// GrantAccountPrivilege invokes the r_kvstore.GrantAccountPrivilege API synchronously
// api document: https://help.aliyun.com/api/r-kvstore/grantaccountprivilege.html
func (client *Client) GrantAccountPrivilege(request *GrantAccountPrivilegeRequest) (response *GrantAccountPrivilegeResponse, err error) {
response = CreateGrantAccountPrivilegeResponse()
err = client.DoAction(request, response)
return
}

// GrantAccountPrivilegeWithChan invokes the r_kvstore.GrantAccountPrivilege API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/grantaccountprivilege.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GrantAccountPrivilegeWithChan(request *GrantAccountPrivilegeRequest) (<-chan *GrantAccountPrivilegeResponse, <-chan error) {
responseChan := make(chan *GrantAccountPrivilegeResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.GrantAccountPrivilege(request)
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

// GrantAccountPrivilegeWithCallback invokes the r_kvstore.GrantAccountPrivilege API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/grantaccountprivilege.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GrantAccountPrivilegeWithCallback(request *GrantAccountPrivilegeRequest, callback func(response *GrantAccountPrivilegeResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *GrantAccountPrivilegeResponse
var err error
defer close(result)
response, err = client.GrantAccountPrivilege(request)
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

// GrantAccountPrivilegeRequest is the request struct for api GrantAccountPrivilege
type GrantAccountPrivilegeRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    AccountPrivilege     string `position:"Query" name:"AccountPrivilege"`
                    AccountName     string `position:"Query" name:"AccountName"`
                    SecurityToken     string `position:"Query" name:"SecurityToken"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    OwnerAccount     string `position:"Query" name:"OwnerAccount"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
                    InstanceId     string `position:"Query" name:"InstanceId"`
}


// GrantAccountPrivilegeResponse is the response struct for api GrantAccountPrivilege
type GrantAccountPrivilegeResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
}

// CreateGrantAccountPrivilegeRequest creates a request to invoke GrantAccountPrivilege API
func CreateGrantAccountPrivilegeRequest() (request *GrantAccountPrivilegeRequest) {
request = &GrantAccountPrivilegeRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("R-kvstore", "2015-01-01", "GrantAccountPrivilege", "kvstore", "openAPI")
return
}

// CreateGrantAccountPrivilegeResponse creates a response to parse from GrantAccountPrivilege response
func CreateGrantAccountPrivilegeResponse() (response *GrantAccountPrivilegeResponse) {
response = &GrantAccountPrivilegeResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

