
package rds

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

// BatchGrantAccountPrivilege invokes the rds.BatchGrantAccountPrivilege API synchronously
// api document: https://help.aliyun.com/api/rds/batchgrantaccountprivilege.html
func (client *Client) BatchGrantAccountPrivilege(request *BatchGrantAccountPrivilegeRequest) (response *BatchGrantAccountPrivilegeResponse, err error) {
response = CreateBatchGrantAccountPrivilegeResponse()
err = client.DoAction(request, response)
return
}

// BatchGrantAccountPrivilegeWithChan invokes the rds.BatchGrantAccountPrivilege API asynchronously
// api document: https://help.aliyun.com/api/rds/batchgrantaccountprivilege.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) BatchGrantAccountPrivilegeWithChan(request *BatchGrantAccountPrivilegeRequest) (<-chan *BatchGrantAccountPrivilegeResponse, <-chan error) {
responseChan := make(chan *BatchGrantAccountPrivilegeResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.BatchGrantAccountPrivilege(request)
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

// BatchGrantAccountPrivilegeWithCallback invokes the rds.BatchGrantAccountPrivilege API asynchronously
// api document: https://help.aliyun.com/api/rds/batchgrantaccountprivilege.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) BatchGrantAccountPrivilegeWithCallback(request *BatchGrantAccountPrivilegeRequest, callback func(response *BatchGrantAccountPrivilegeResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *BatchGrantAccountPrivilegeResponse
var err error
defer close(result)
response, err = client.BatchGrantAccountPrivilege(request)
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

// BatchGrantAccountPrivilegeRequest is the request struct for api BatchGrantAccountPrivilege
type BatchGrantAccountPrivilegeRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    AccountPrivilege     string `position:"Query" name:"AccountPrivilege"`
                    AccountName     string `position:"Query" name:"AccountName"`
                    DBInstanceId     string `position:"Query" name:"DBInstanceId"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    OwnerAccount     string `position:"Query" name:"OwnerAccount"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
                    DBName     string `position:"Query" name:"DBName"`
}


// BatchGrantAccountPrivilegeResponse is the response struct for api BatchGrantAccountPrivilege
type BatchGrantAccountPrivilegeResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
}

// CreateBatchGrantAccountPrivilegeRequest creates a request to invoke BatchGrantAccountPrivilege API
func CreateBatchGrantAccountPrivilegeRequest() (request *BatchGrantAccountPrivilegeRequest) {
request = &BatchGrantAccountPrivilegeRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Rds", "2014-08-15", "BatchGrantAccountPrivilege", "rds", "openAPI")
return
}

// CreateBatchGrantAccountPrivilegeResponse creates a response to parse from BatchGrantAccountPrivilege response
func CreateBatchGrantAccountPrivilegeResponse() (response *BatchGrantAccountPrivilegeResponse) {
response = &BatchGrantAccountPrivilegeResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

