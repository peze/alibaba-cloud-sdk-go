
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

// PutProduct invokes the acs.PutProduct API synchronously
// api document: https://help.aliyun.com/api/acs/putproduct.html
func (client *Client) PutProduct(request *PutProductRequest) (response *PutProductResponse, err error) {
response = CreatePutProductResponse()
err = client.DoAction(request, response)
return
}

// PutProductWithChan invokes the acs.PutProduct API asynchronously
// api document: https://help.aliyun.com/api/acs/putproduct.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PutProductWithChan(request *PutProductRequest) (<-chan *PutProductResponse, <-chan error) {
responseChan := make(chan *PutProductResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.PutProduct(request)
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

// PutProductWithCallback invokes the acs.PutProduct API asynchronously
// api document: https://help.aliyun.com/api/acs/putproduct.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PutProductWithCallback(request *PutProductRequest, callback func(response *PutProductResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *PutProductResponse
var err error
defer close(result)
response, err = client.PutProduct(request)
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

// PutProductRequest is the request struct for api PutProduct
type PutProductRequest struct {
*requests.RoaRequest
                    ContentLength     requests.Integer `position:"Header" name:"Content-Length"`
                    BodyContent     string `position:"Body" name:"BodyContent"`
                    ContentMD5     string `position:"Header" name:"Content-MD5"`
                    ProductName     string `position:"Path" name:"ProductName"`
                    ContentType     string `position:"Header" name:"Content-Type"`
                    Accept     string `position:"Header" name:"Accept"`
}


// PutProductResponse is the response struct for api PutProduct
type PutProductResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
}

// CreatePutProductRequest creates a request to invoke PutProduct API
func CreatePutProductRequest() (request *PutProductRequest) {
request = &PutProductRequest{
RoaRequest: &requests.RoaRequest{},
}
request.InitWithApiInfo("Acs", "2015-01-01", "PutProduct", "/Product/[ProductName]", "", "")
request.Method = requests.PUT
return
}

// CreatePutProductResponse creates a response to parse from PutProduct response
func CreatePutProductResponse() (response *PutProductResponse) {
response = &PutProductResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

