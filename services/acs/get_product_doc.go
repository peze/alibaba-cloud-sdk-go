
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

// GetProductDoc invokes the acs.GetProductDoc API synchronously
// api document: https://help.aliyun.com/api/acs/getproductdoc.html
func (client *Client) GetProductDoc(request *GetProductDocRequest) (response *GetProductDocResponse, err error) {
response = CreateGetProductDocResponse()
err = client.DoAction(request, response)
return
}

// GetProductDocWithChan invokes the acs.GetProductDoc API asynchronously
// api document: https://help.aliyun.com/api/acs/getproductdoc.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetProductDocWithChan(request *GetProductDocRequest) (<-chan *GetProductDocResponse, <-chan error) {
responseChan := make(chan *GetProductDocResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.GetProductDoc(request)
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

// GetProductDocWithCallback invokes the acs.GetProductDoc API asynchronously
// api document: https://help.aliyun.com/api/acs/getproductdoc.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetProductDocWithCallback(request *GetProductDocRequest, callback func(response *GetProductDocResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *GetProductDocResponse
var err error
defer close(result)
response, err = client.GetProductDoc(request)
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

// GetProductDocRequest is the request struct for api GetProductDoc
type GetProductDocRequest struct {
*requests.RoaRequest
                    ProductName     string `position:"Path" name:"ProductName"`
                    Language     string `position:"Path" name:"Language"`
                    Accept     string `position:"Header" name:"Accept"`
}


// GetProductDocResponse is the response struct for api GetProductDoc
type GetProductDocResponse struct {
*responses.BaseResponse
            Name     string `json:"name" xml:"name"`
            Title     string `json:"title" xml:"title"`
            Language     string `json:"language" xml:"language"`
            Description     string `json:"Description" xml:"Description"`
}

// CreateGetProductDocRequest creates a request to invoke GetProductDoc API
func CreateGetProductDocRequest() (request *GetProductDocRequest) {
request = &GetProductDocRequest{
RoaRequest: &requests.RoaRequest{},
}
request.InitWithApiInfo("Acs", "2015-01-01", "GetProductDoc", "/ProductDoc/[ProductName]/[Language]", "", "")
request.Method = requests.GET
return
}

// CreateGetProductDocResponse creates a response to parse from GetProductDoc response
func CreateGetProductDocResponse() (response *GetProductDocResponse) {
response = &GetProductDocResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

