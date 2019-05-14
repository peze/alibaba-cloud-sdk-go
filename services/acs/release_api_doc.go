
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

// ReleaseApiDoc invokes the acs.ReleaseApiDoc API synchronously
// api document: https://help.aliyun.com/api/acs/releaseapidoc.html
func (client *Client) ReleaseApiDoc(request *ReleaseApiDocRequest) (response *ReleaseApiDocResponse, err error) {
response = CreateReleaseApiDocResponse()
err = client.DoAction(request, response)
return
}

// ReleaseApiDocWithChan invokes the acs.ReleaseApiDoc API asynchronously
// api document: https://help.aliyun.com/api/acs/releaseapidoc.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ReleaseApiDocWithChan(request *ReleaseApiDocRequest) (<-chan *ReleaseApiDocResponse, <-chan error) {
responseChan := make(chan *ReleaseApiDocResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.ReleaseApiDoc(request)
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

// ReleaseApiDocWithCallback invokes the acs.ReleaseApiDoc API asynchronously
// api document: https://help.aliyun.com/api/acs/releaseapidoc.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ReleaseApiDocWithCallback(request *ReleaseApiDocRequest, callback func(response *ReleaseApiDocResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *ReleaseApiDocResponse
var err error
defer close(result)
response, err = client.ReleaseApiDoc(request)
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

// ReleaseApiDocRequest is the request struct for api ReleaseApiDoc
type ReleaseApiDocRequest struct {
*requests.RoaRequest
                    ApiName     string `position:"Path" name:"ApiName"`
                    ProductName     string `position:"Path" name:"ProductName"`
                    Language     string `position:"Path" name:"Language"`
                    VersionName     string `position:"Path" name:"VersionName"`
                    Accept     string `position:"Header" name:"Accept"`
}


// ReleaseApiDocResponse is the response struct for api ReleaseApiDoc
type ReleaseApiDocResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
}

// CreateReleaseApiDocRequest creates a request to invoke ReleaseApiDoc API
func CreateReleaseApiDocRequest() (request *ReleaseApiDocRequest) {
request = &ReleaseApiDocRequest{
RoaRequest: &requests.RoaRequest{},
}
request.InitWithApiInfo("Acs", "2015-01-01", "ReleaseApiDoc", "/ReleaseApiDoc/[ProductName]/[VersionName]/[ApiName]/[Language]", "", "")
request.Method = requests.PUT
return
}

// CreateReleaseApiDocResponse creates a response to parse from ReleaseApiDoc response
func CreateReleaseApiDocResponse() (response *ReleaseApiDocResponse) {
response = &ReleaseApiDocResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

