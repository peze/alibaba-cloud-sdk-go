
package acs

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

// PutApiRollback invokes the acs.PutApiRollback API synchronously
// api document: https://help.aliyun.com/api/acs/putapirollback.html
func (client *Client) PutApiRollback(request *PutApiRollbackRequest) (response *PutApiRollbackResponse, err error) {
response = CreatePutApiRollbackResponse()
err = client.DoAction(request, response)
return
}

// PutApiRollbackWithChan invokes the acs.PutApiRollback API asynchronously
// api document: https://help.aliyun.com/api/acs/putapirollback.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PutApiRollbackWithChan(request *PutApiRollbackRequest) (<-chan *PutApiRollbackResponse, <-chan error) {
responseChan := make(chan *PutApiRollbackResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.PutApiRollback(request)
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

// PutApiRollbackWithCallback invokes the acs.PutApiRollback API asynchronously
// api document: https://help.aliyun.com/api/acs/putapirollback.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PutApiRollbackWithCallback(request *PutApiRollbackRequest, callback func(response *PutApiRollbackResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *PutApiRollbackResponse
var err error
defer close(result)
response, err = client.PutApiRollback(request)
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

// PutApiRollbackRequest is the request struct for api PutApiRollback
type PutApiRollbackRequest struct {
*requests.RoaRequest
                    ContentLength     requests.Integer `position:"Header" name:"Content-Length"`
                    BodyContent     string `position:"Body" name:"BodyContent"`
                    ApiName     string `position:"Path" name:"ApiName"`
                    ContentMD5     string `position:"Header" name:"Content-MD5"`
                    BackupId     requests.Integer `position:"Query" name:"BackupId"`
                    ProductName     string `position:"Path" name:"ProductName"`
                    ContentType     string `position:"Header" name:"Content-Type"`
                    VersionName     string `position:"Path" name:"VersionName"`
                    Accept     string `position:"Header" name:"Accept"`
}


// PutApiRollbackResponse is the response struct for api PutApiRollback
type PutApiRollbackResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
}

// CreatePutApiRollbackRequest creates a request to invoke PutApiRollback API
func CreatePutApiRollbackRequest() (request *PutApiRollbackRequest) {
request = &PutApiRollbackRequest{
RoaRequest: &requests.RoaRequest{},
}
request.InitWithApiInfo("Acs", "2015-01-01", "PutApiRollback", "/Rollback/[ProductName]/[VersionName]/[ApiName]", "", "")
request.Method = requests.PUT
return
}

// CreatePutApiRollbackResponse creates a response to parse from PutApiRollback response
func CreatePutApiRollbackResponse() (response *PutApiRollbackResponse) {
response = &PutApiRollbackResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


