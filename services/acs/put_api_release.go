
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

// PutApiRelease invokes the acs.PutApiRelease API synchronously
// api document: https://help.aliyun.com/api/acs/putapirelease.html
func (client *Client) PutApiRelease(request *PutApiReleaseRequest) (response *PutApiReleaseResponse, err error) {
response = CreatePutApiReleaseResponse()
err = client.DoAction(request, response)
return
}

// PutApiReleaseWithChan invokes the acs.PutApiRelease API asynchronously
// api document: https://help.aliyun.com/api/acs/putapirelease.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PutApiReleaseWithChan(request *PutApiReleaseRequest) (<-chan *PutApiReleaseResponse, <-chan error) {
responseChan := make(chan *PutApiReleaseResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.PutApiRelease(request)
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

// PutApiReleaseWithCallback invokes the acs.PutApiRelease API asynchronously
// api document: https://help.aliyun.com/api/acs/putapirelease.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PutApiReleaseWithCallback(request *PutApiReleaseRequest, callback func(response *PutApiReleaseResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *PutApiReleaseResponse
var err error
defer close(result)
response, err = client.PutApiRelease(request)
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

// PutApiReleaseRequest is the request struct for api PutApiRelease
type PutApiReleaseRequest struct {
*requests.RoaRequest
                    ContentLength     requests.Integer `position:"Header" name:"Content-Length"`
                    BodyContent     string `position:"Body" name:"BodyContent"`
                    ApiName     string `position:"Path" name:"ApiName"`
                    ContentMD5     string `position:"Header" name:"Content-MD5"`
                    ProductName     string `position:"Path" name:"ProductName"`
                    ContentType     string `position:"Header" name:"Content-Type"`
                    VersionName     string `position:"Path" name:"VersionName"`
                    Accept     string `position:"Header" name:"Accept"`
}


// PutApiReleaseResponse is the response struct for api PutApiRelease
type PutApiReleaseResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
}

// CreatePutApiReleaseRequest creates a request to invoke PutApiRelease API
func CreatePutApiReleaseRequest() (request *PutApiReleaseRequest) {
request = &PutApiReleaseRequest{
RoaRequest: &requests.RoaRequest{},
}
request.InitWithApiInfo("Acs", "2015-01-01", "PutApiRelease", "/Release/[ProductName]/[VersionName]/[ApiName]", "", "")
request.Method = requests.PUT
return
}

// CreatePutApiReleaseResponse creates a response to parse from PutApiRelease response
func CreatePutApiReleaseResponse() (response *PutApiReleaseResponse) {
response = &PutApiReleaseResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


