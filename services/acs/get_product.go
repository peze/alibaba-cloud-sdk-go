
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

// GetProduct invokes the acs.GetProduct API synchronously
// api document: https://help.aliyun.com/api/acs/getproduct.html
func (client *Client) GetProduct(request *GetProductRequest) (response *GetProductResponse, err error) {
response = CreateGetProductResponse()
err = client.DoAction(request, response)
return
}

// GetProductWithChan invokes the acs.GetProduct API asynchronously
// api document: https://help.aliyun.com/api/acs/getproduct.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetProductWithChan(request *GetProductRequest) (<-chan *GetProductResponse, <-chan error) {
responseChan := make(chan *GetProductResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.GetProduct(request)
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

// GetProductWithCallback invokes the acs.GetProduct API asynchronously
// api document: https://help.aliyun.com/api/acs/getproduct.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetProductWithCallback(request *GetProductRequest, callback func(response *GetProductResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *GetProductResponse
var err error
defer close(result)
response, err = client.GetProduct(request)
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

// GetProductRequest is the request struct for api GetProduct
type GetProductRequest struct {
*requests.RoaRequest
                    ProductName     string `position:"Path" name:"ProductName"`
                    Accept     string `position:"Header" name:"Accept"`
}


// GetProductResponse is the response struct for api GetProduct
type GetProductResponse struct {
*responses.BaseResponse
            Name     string `json:"name" xml:"name"`
            Domain     string `json:"domain" xml:"domain"`
            Type     string `json:"type" xml:"type"`
                    Versions Versions `json:"Versions" xml:"Versions"`
}

// CreateGetProductRequest creates a request to invoke GetProduct API
func CreateGetProductRequest() (request *GetProductRequest) {
request = &GetProductRequest{
RoaRequest: &requests.RoaRequest{},
}
request.InitWithApiInfo("Acs", "2015-01-01", "GetProduct", "/Product/[ProductName]", "", "")
request.Method = requests.GET
return
}

// CreateGetProductResponse creates a response to parse from GetProduct response
func CreateGetProductResponse() (response *GetProductResponse) {
response = &GetProductResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


