
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

// CreateMulOrderForBuy invokes the rds.CreateMulOrderForBuy API synchronously
// api document: https://help.aliyun.com/api/rds/createmulorderforbuy.html
func (client *Client) CreateMulOrderForBuy(request *CreateMulOrderForBuyRequest) (response *CreateMulOrderForBuyResponse, err error) {
response = CreateCreateMulOrderForBuyResponse()
err = client.DoAction(request, response)
return
}

// CreateMulOrderForBuyWithChan invokes the rds.CreateMulOrderForBuy API asynchronously
// api document: https://help.aliyun.com/api/rds/createmulorderforbuy.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateMulOrderForBuyWithChan(request *CreateMulOrderForBuyRequest) (<-chan *CreateMulOrderForBuyResponse, <-chan error) {
responseChan := make(chan *CreateMulOrderForBuyResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.CreateMulOrderForBuy(request)
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

// CreateMulOrderForBuyWithCallback invokes the rds.CreateMulOrderForBuy API asynchronously
// api document: https://help.aliyun.com/api/rds/createmulorderforbuy.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateMulOrderForBuyWithCallback(request *CreateMulOrderForBuyRequest, callback func(response *CreateMulOrderForBuyResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *CreateMulOrderForBuyResponse
var err error
defer close(result)
response, err = client.CreateMulOrderForBuy(request)
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

// CreateMulOrderForBuyRequest is the request struct for api CreateMulOrderForBuy
type CreateMulOrderForBuyRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    ClientToken     string `position:"Query" name:"ClientToken"`
                    Param     string `position:"Query" name:"Param"`
                    AgentId     string `position:"Query" name:"AgentId"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    OwnerAccount     string `position:"Query" name:"OwnerAccount"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
}


// CreateMulOrderForBuyResponse is the response struct for api CreateMulOrderForBuy
type CreateMulOrderForBuyResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            DBInstanceId     string `json:"DBInstanceId" xml:"DBInstanceId"`
            OrderId     int `json:"OrderId" xml:"OrderId"`
}

// CreateCreateMulOrderForBuyRequest creates a request to invoke CreateMulOrderForBuy API
func CreateCreateMulOrderForBuyRequest() (request *CreateMulOrderForBuyRequest) {
request = &CreateMulOrderForBuyRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Rds", "2014-08-15", "CreateMulOrderForBuy", "rds", "openAPI")
return
}

// CreateCreateMulOrderForBuyResponse creates a response to parse from CreateMulOrderForBuy response
func CreateCreateMulOrderForBuyResponse() (response *CreateMulOrderForBuyResponse) {
response = &CreateMulOrderForBuyResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

