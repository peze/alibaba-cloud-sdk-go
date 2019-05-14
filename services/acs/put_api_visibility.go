
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

// PutApiVisibility invokes the acs.PutApiVisibility API synchronously
// api document: https://help.aliyun.com/api/acs/putapivisibility.html
func (client *Client) PutApiVisibility(request *PutApiVisibilityRequest) (response *PutApiVisibilityResponse, err error) {
response = CreatePutApiVisibilityResponse()
err = client.DoAction(request, response)
return
}

// PutApiVisibilityWithChan invokes the acs.PutApiVisibility API asynchronously
// api document: https://help.aliyun.com/api/acs/putapivisibility.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PutApiVisibilityWithChan(request *PutApiVisibilityRequest) (<-chan *PutApiVisibilityResponse, <-chan error) {
responseChan := make(chan *PutApiVisibilityResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.PutApiVisibility(request)
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

// PutApiVisibilityWithCallback invokes the acs.PutApiVisibility API asynchronously
// api document: https://help.aliyun.com/api/acs/putapivisibility.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PutApiVisibilityWithCallback(request *PutApiVisibilityRequest, callback func(response *PutApiVisibilityResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *PutApiVisibilityResponse
var err error
defer close(result)
response, err = client.PutApiVisibility(request)
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

// PutApiVisibilityRequest is the request struct for api PutApiVisibility
type PutApiVisibilityRequest struct {
*requests.RoaRequest
                    BodyContent     string `position:"Body" name:"BodyContent"`
                    ApiName     string `position:"Path" name:"ApiName"`
                    ProductName     string `position:"Path" name:"ProductName"`
                    VersionName     string `position:"Path" name:"VersionName"`
                    Accept     string `position:"Header" name:"Accept"`
}


// PutApiVisibilityResponse is the response struct for api PutApiVisibility
type PutApiVisibilityResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
}

// CreatePutApiVisibilityRequest creates a request to invoke PutApiVisibility API
func CreatePutApiVisibilityRequest() (request *PutApiVisibilityRequest) {
request = &PutApiVisibilityRequest{
RoaRequest: &requests.RoaRequest{},
}
request.InitWithApiInfo("Acs", "2015-01-01", "PutApiVisibility", "/ApiVisibility/[ProductName]/[VersionName]/[ApiName]", "", "")
request.Method = requests.PUT
return
}

// CreatePutApiVisibilityResponse creates a response to parse from PutApiVisibility response
func CreatePutApiVisibilityResponse() (response *PutApiVisibilityResponse) {
response = &PutApiVisibilityResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


