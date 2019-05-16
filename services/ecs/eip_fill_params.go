
package ecs

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

// EipFillParams invokes the ecs.EipFillParams API synchronously
// api document: https://help.aliyun.com/api/ecs/eipfillparams.html
func (client *Client) EipFillParams(request *EipFillParamsRequest) (response *EipFillParamsResponse, err error) {
response = CreateEipFillParamsResponse()
err = client.DoAction(request, response)
return
}

// EipFillParamsWithChan invokes the ecs.EipFillParams API asynchronously
// api document: https://help.aliyun.com/api/ecs/eipfillparams.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) EipFillParamsWithChan(request *EipFillParamsRequest) (<-chan *EipFillParamsResponse, <-chan error) {
responseChan := make(chan *EipFillParamsResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.EipFillParams(request)
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

// EipFillParamsWithCallback invokes the ecs.EipFillParams API asynchronously
// api document: https://help.aliyun.com/api/ecs/eipfillparams.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) EipFillParamsWithCallback(request *EipFillParamsRequest, callback func(response *EipFillParamsResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *EipFillParamsResponse
var err error
defer close(result)
response, err = client.EipFillParams(request)
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

// EipFillParamsRequest is the request struct for api EipFillParams
type EipFillParamsRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    Data     string `position:"Query" name:"data"`
                    ClientToken     string `position:"Query" name:"ClientToken"`
                    UserCidr     string `position:"Query" name:"UserCidr"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    OwnerAccount     string `position:"Query" name:"OwnerAccount"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
}


// EipFillParamsResponse is the response struct for api EipFillParams
type EipFillParamsResponse struct {
*responses.BaseResponse
            RequestId     string `json:"requestId" xml:"requestId"`
            Data     string `json:"data" xml:"data"`
            Code     string `json:"code" xml:"code"`
            Success     bool `json:"success" xml:"success"`
            Message     string `json:"message" xml:"message"`
}

// CreateEipFillParamsRequest creates a request to invoke EipFillParams API
func CreateEipFillParamsRequest() (request *EipFillParamsRequest) {
request = &EipFillParamsRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Ecs", "2014-05-26", "EipFillParams", "", "")
return
}

// CreateEipFillParamsResponse creates a response to parse from EipFillParams response
func CreateEipFillParamsResponse() (response *EipFillParamsResponse) {
response = &EipFillParamsResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


