
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

// QueryPriceForRefund invokes the rds.QueryPriceForRefund API synchronously
// api document: https://help.aliyun.com/api/rds/querypriceforrefund.html
func (client *Client) QueryPriceForRefund(request *QueryPriceForRefundRequest) (response *QueryPriceForRefundResponse, err error) {
response = CreateQueryPriceForRefundResponse()
err = client.DoAction(request, response)
return
}

// QueryPriceForRefundWithChan invokes the rds.QueryPriceForRefund API asynchronously
// api document: https://help.aliyun.com/api/rds/querypriceforrefund.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryPriceForRefundWithChan(request *QueryPriceForRefundRequest) (<-chan *QueryPriceForRefundResponse, <-chan error) {
responseChan := make(chan *QueryPriceForRefundResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.QueryPriceForRefund(request)
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

// QueryPriceForRefundWithCallback invokes the rds.QueryPriceForRefund API asynchronously
// api document: https://help.aliyun.com/api/rds/querypriceforrefund.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryPriceForRefundWithCallback(request *QueryPriceForRefundRequest, callback func(response *QueryPriceForRefundResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *QueryPriceForRefundResponse
var err error
defer close(result)
response, err = client.QueryPriceForRefund(request)
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

// QueryPriceForRefundRequest is the request struct for api QueryPriceForRefund
type QueryPriceForRefundRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    ClientToken     string `position:"Query" name:"ClientToken"`
                    CountryCode     string `position:"Query" name:"CountryCode"`
                    CurrencyCode     string `position:"Query" name:"CurrencyCode"`
                    DBInstanceId     string `position:"Query" name:"DBInstanceId"`
                    BusinessInfo     string `position:"Query" name:"BusinessInfo"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    CommodityCode     string `position:"Query" name:"CommodityCode"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
                    PromotionCode     string `position:"Query" name:"PromotionCode"`
                    OrderType     string `position:"Query" name:"OrderType"`
}


// QueryPriceForRefundResponse is the response struct for api QueryPriceForRefund
type QueryPriceForRefundResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            Order Order  `json:"Order" xml:"Order"`
                    Rules RulesInQueryPriceForRefund `json:"Rules" xml:"Rules"`
}

// CreateQueryPriceForRefundRequest creates a request to invoke QueryPriceForRefund API
func CreateQueryPriceForRefundRequest() (request *QueryPriceForRefundRequest) {
request = &QueryPriceForRefundRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Rds", "2014-08-15", "QueryPriceForRefund", "rds", "openAPI")
return
}

// CreateQueryPriceForRefundResponse creates a response to parse from QueryPriceForRefund response
func CreateQueryPriceForRefundResponse() (response *QueryPriceForRefundResponse) {
response = &QueryPriceForRefundResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

