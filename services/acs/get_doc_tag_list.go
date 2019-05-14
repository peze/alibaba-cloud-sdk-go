
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

// GetDocTagList invokes the acs.GetDocTagList API synchronously
// api document: https://help.aliyun.com/api/acs/getdoctaglist.html
func (client *Client) GetDocTagList(request *GetDocTagListRequest) (response *GetDocTagListResponse, err error) {
response = CreateGetDocTagListResponse()
err = client.DoAction(request, response)
return
}

// GetDocTagListWithChan invokes the acs.GetDocTagList API asynchronously
// api document: https://help.aliyun.com/api/acs/getdoctaglist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetDocTagListWithChan(request *GetDocTagListRequest) (<-chan *GetDocTagListResponse, <-chan error) {
responseChan := make(chan *GetDocTagListResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.GetDocTagList(request)
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

// GetDocTagListWithCallback invokes the acs.GetDocTagList API asynchronously
// api document: https://help.aliyun.com/api/acs/getdoctaglist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetDocTagListWithCallback(request *GetDocTagListRequest, callback func(response *GetDocTagListResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *GetDocTagListResponse
var err error
defer close(result)
response, err = client.GetDocTagList(request)
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

// GetDocTagListRequest is the request struct for api GetDocTagList
type GetDocTagListRequest struct {
*requests.RoaRequest
                    ProductName     string `position:"Path" name:"ProductName"`
                    Accept     string `position:"Header" name:"Accept"`
}


// GetDocTagListResponse is the response struct for api GetDocTagList
type GetDocTagListResponse struct {
*responses.BaseResponse
                    DoctTagList DoctTagList `json:"DoctTagList" xml:"DoctTagList"`
}

// CreateGetDocTagListRequest creates a request to invoke GetDocTagList API
func CreateGetDocTagListRequest() (request *GetDocTagListRequest) {
request = &GetDocTagListRequest{
RoaRequest: &requests.RoaRequest{},
}
request.InitWithApiInfo("Acs", "2015-01-01", "GetDocTagList", "/DocTag/[ProductName]", "", "")
request.Method = requests.GET
return
}

// CreateGetDocTagListResponse creates a response to parse from GetDocTagList response
func CreateGetDocTagListResponse() (response *GetDocTagListResponse) {
response = &GetDocTagListResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


