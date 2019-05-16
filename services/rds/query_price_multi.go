
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

// QueryPriceMulti invokes the rds.QueryPriceMulti API synchronously
// api document: https://help.aliyun.com/api/rds/querypricemulti.html
func (client *Client) QueryPriceMulti(request *QueryPriceMultiRequest) (response *QueryPriceMultiResponse, err error) {
response = CreateQueryPriceMultiResponse()
err = client.DoAction(request, response)
return
}

// QueryPriceMultiWithChan invokes the rds.QueryPriceMulti API asynchronously
// api document: https://help.aliyun.com/api/rds/querypricemulti.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryPriceMultiWithChan(request *QueryPriceMultiRequest) (<-chan *QueryPriceMultiResponse, <-chan error) {
responseChan := make(chan *QueryPriceMultiResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.QueryPriceMulti(request)
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

// QueryPriceMultiWithCallback invokes the rds.QueryPriceMulti API asynchronously
// api document: https://help.aliyun.com/api/rds/querypricemulti.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryPriceMultiWithCallback(request *QueryPriceMultiRequest, callback func(response *QueryPriceMultiResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *QueryPriceMultiResponse
var err error
defer close(result)
response, err = client.QueryPriceMulti(request)
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

// QueryPriceMultiRequest is the request struct for api QueryPriceMulti
type QueryPriceMultiRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    ClientToken     string `position:"Query" name:"ClientToken"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
                    Param     string `position:"Query" name:"Param"`
                    OrderType     string `position:"Query" name:"OrderType"`
}


// QueryPriceMultiResponse is the response struct for api QueryPriceMulti
type QueryPriceMultiResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            Order Order  `json:"Order" xml:"Order"`
                    Rules RulesInQueryPriceMulti `json:"Rules" xml:"Rules"`
}

// CreateQueryPriceMultiRequest creates a request to invoke QueryPriceMulti API
func CreateQueryPriceMultiRequest() (request *QueryPriceMultiRequest) {
request = &QueryPriceMultiRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Rds", "2014-08-15", "QueryPriceMulti", "rds", "openAPI")
return
}

// CreateQueryPriceMultiResponse creates a response to parse from QueryPriceMulti response
func CreateQueryPriceMultiResponse() (response *QueryPriceMultiResponse) {
response = &QueryPriceMultiResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

