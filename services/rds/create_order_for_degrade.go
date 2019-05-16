
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

// CreateOrderForDegrade invokes the rds.CreateOrderForDegrade API synchronously
// api document: https://help.aliyun.com/api/rds/createorderfordegrade.html
func (client *Client) CreateOrderForDegrade(request *CreateOrderForDegradeRequest) (response *CreateOrderForDegradeResponse, err error) {
response = CreateCreateOrderForDegradeResponse()
err = client.DoAction(request, response)
return
}

// CreateOrderForDegradeWithChan invokes the rds.CreateOrderForDegrade API asynchronously
// api document: https://help.aliyun.com/api/rds/createorderfordegrade.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateOrderForDegradeWithChan(request *CreateOrderForDegradeRequest) (<-chan *CreateOrderForDegradeResponse, <-chan error) {
responseChan := make(chan *CreateOrderForDegradeResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.CreateOrderForDegrade(request)
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

// CreateOrderForDegradeWithCallback invokes the rds.CreateOrderForDegrade API asynchronously
// api document: https://help.aliyun.com/api/rds/createorderfordegrade.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateOrderForDegradeWithCallback(request *CreateOrderForDegradeRequest, callback func(response *CreateOrderForDegradeResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *CreateOrderForDegradeResponse
var err error
defer close(result)
response, err = client.CreateOrderForDegrade(request)
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

// CreateOrderForDegradeRequest is the request struct for api CreateOrderForDegrade
type CreateOrderForDegradeRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    DBInstanceStorage     requests.Integer `position:"Query" name:"DBInstanceStorage"`
                    ClientToken     string `position:"Query" name:"ClientToken"`
                    EffectiveTime     string `position:"Query" name:"EffectiveTime"`
                    DBInstanceId     string `position:"Query" name:"DBInstanceId"`
                    BusinessInfo     string `position:"Query" name:"BusinessInfo"`
                    AutoPay     requests.Boolean `position:"Query" name:"AutoPay"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    Resource     string `position:"Query" name:"Resource"`
                    CommodityCode     string `position:"Query" name:"CommodityCode"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
                    UsedTime     string `position:"Query" name:"UsedTime"`
                    DBInstanceClass     string `position:"Query" name:"DBInstanceClass"`
                    PromotionCode     string `position:"Query" name:"PromotionCode"`
                    ZoneId     string `position:"Query" name:"ZoneId"`
                    TimeType     string `position:"Query" name:"TimeType"`
                    PayType     string `position:"Query" name:"PayType"`
}


// CreateOrderForDegradeResponse is the response struct for api CreateOrderForDegrade
type CreateOrderForDegradeResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            DBInstanceId     string `json:"DBInstanceId" xml:"DBInstanceId"`
            OrderId     int `json:"OrderId" xml:"OrderId"`
}

// CreateCreateOrderForDegradeRequest creates a request to invoke CreateOrderForDegrade API
func CreateCreateOrderForDegradeRequest() (request *CreateOrderForDegradeRequest) {
request = &CreateOrderForDegradeRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Rds", "2014-08-15", "CreateOrderForDegrade", "rds", "openAPI")
return
}

// CreateCreateOrderForDegradeResponse creates a response to parse from CreateOrderForDegrade response
func CreateCreateOrderForDegradeResponse() (response *CreateOrderForDegradeResponse) {
response = &CreateOrderForDegradeResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

