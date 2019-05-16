
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

// CreateMulOrderForDefer invokes the rds.CreateMulOrderForDefer API synchronously
// api document: https://help.aliyun.com/api/rds/createmulorderfordefer.html
func (client *Client) CreateMulOrderForDefer(request *CreateMulOrderForDeferRequest) (response *CreateMulOrderForDeferResponse, err error) {
response = CreateCreateMulOrderForDeferResponse()
err = client.DoAction(request, response)
return
}

// CreateMulOrderForDeferWithChan invokes the rds.CreateMulOrderForDefer API asynchronously
// api document: https://help.aliyun.com/api/rds/createmulorderfordefer.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateMulOrderForDeferWithChan(request *CreateMulOrderForDeferRequest) (<-chan *CreateMulOrderForDeferResponse, <-chan error) {
responseChan := make(chan *CreateMulOrderForDeferResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.CreateMulOrderForDefer(request)
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

// CreateMulOrderForDeferWithCallback invokes the rds.CreateMulOrderForDefer API asynchronously
// api document: https://help.aliyun.com/api/rds/createmulorderfordefer.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateMulOrderForDeferWithCallback(request *CreateMulOrderForDeferRequest, callback func(response *CreateMulOrderForDeferResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *CreateMulOrderForDeferResponse
var err error
defer close(result)
response, err = client.CreateMulOrderForDefer(request)
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

// CreateMulOrderForDeferRequest is the request struct for api CreateMulOrderForDefer
type CreateMulOrderForDeferRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    ClientToken     string `position:"Query" name:"ClientToken"`
                    DBInstanceId     string `position:"Query" name:"DBInstanceId"`
                    BusinessInfo     string `position:"Query" name:"BusinessInfo"`
                    AutoPay     requests.Boolean `position:"Query" name:"AutoPay"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    Resource     string `position:"Query" name:"Resource"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
}


// CreateMulOrderForDeferResponse is the response struct for api CreateMulOrderForDefer
type CreateMulOrderForDeferResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            DBInstanceId     string `json:"DBInstanceId" xml:"DBInstanceId"`
            OrderId     int `json:"OrderId" xml:"OrderId"`
}

// CreateCreateMulOrderForDeferRequest creates a request to invoke CreateMulOrderForDefer API
func CreateCreateMulOrderForDeferRequest() (request *CreateMulOrderForDeferRequest) {
request = &CreateMulOrderForDeferRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Rds", "2014-08-15", "CreateMulOrderForDefer", "rds", "openAPI")
return
}

// CreateCreateMulOrderForDeferResponse creates a response to parse from CreateMulOrderForDefer response
func CreateCreateMulOrderForDeferResponse() (response *CreateMulOrderForDeferResponse) {
response = &CreateMulOrderForDeferResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

