
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

// GetControlPolicyByProduct invokes the acs.GetControlPolicyByProduct API synchronously
// api document: https://help.aliyun.com/api/acs/getcontrolpolicybyproduct.html
func (client *Client) GetControlPolicyByProduct(request *GetControlPolicyByProductRequest) (response *GetControlPolicyByProductResponse, err error) {
response = CreateGetControlPolicyByProductResponse()
err = client.DoAction(request, response)
return
}

// GetControlPolicyByProductWithChan invokes the acs.GetControlPolicyByProduct API asynchronously
// api document: https://help.aliyun.com/api/acs/getcontrolpolicybyproduct.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetControlPolicyByProductWithChan(request *GetControlPolicyByProductRequest) (<-chan *GetControlPolicyByProductResponse, <-chan error) {
responseChan := make(chan *GetControlPolicyByProductResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.GetControlPolicyByProduct(request)
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

// GetControlPolicyByProductWithCallback invokes the acs.GetControlPolicyByProduct API asynchronously
// api document: https://help.aliyun.com/api/acs/getcontrolpolicybyproduct.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetControlPolicyByProductWithCallback(request *GetControlPolicyByProductRequest, callback func(response *GetControlPolicyByProductResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *GetControlPolicyByProductResponse
var err error
defer close(result)
response, err = client.GetControlPolicyByProduct(request)
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

// GetControlPolicyByProductRequest is the request struct for api GetControlPolicyByProduct
type GetControlPolicyByProductRequest struct {
*requests.RoaRequest
                    Product     string `position:"Path" name:"Product"`
                    Accept     string `position:"Header" name:"Accept"`
}


// GetControlPolicyByProductResponse is the response struct for api GetControlPolicyByProduct
type GetControlPolicyByProductResponse struct {
*responses.BaseResponse
                ControlPolicys ControlPolicys `json:"ControlPolicys" xml:"ControlPolicys"`
}

// CreateGetControlPolicyByProductRequest creates a request to invoke GetControlPolicyByProduct API
func CreateGetControlPolicyByProductRequest() (request *GetControlPolicyByProductRequest) {
request = &GetControlPolicyByProductRequest{
RoaRequest: &requests.RoaRequest{},
}
request.InitWithApiInfo("Acs", "2015-01-01", "GetControlPolicyByProduct", "/ControlPolicy/[Product]", "", "")
request.Method = requests.GET
return
}

// CreateGetControlPolicyByProductResponse creates a response to parse from GetControlPolicyByProduct response
func CreateGetControlPolicyByProductResponse() (response *GetControlPolicyByProductResponse) {
response = &GetControlPolicyByProductResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

