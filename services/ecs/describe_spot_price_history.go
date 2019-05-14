
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

// DescribeSpotPriceHistory invokes the ecs.DescribeSpotPriceHistory API synchronously
// api document: https://help.aliyun.com/api/ecs/describespotpricehistory.html
func (client *Client) DescribeSpotPriceHistory(request *DescribeSpotPriceHistoryRequest) (response *DescribeSpotPriceHistoryResponse, err error) {
response = CreateDescribeSpotPriceHistoryResponse()
err = client.DoAction(request, response)
return
}

// DescribeSpotPriceHistoryWithChan invokes the ecs.DescribeSpotPriceHistory API asynchronously
// api document: https://help.aliyun.com/api/ecs/describespotpricehistory.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSpotPriceHistoryWithChan(request *DescribeSpotPriceHistoryRequest) (<-chan *DescribeSpotPriceHistoryResponse, <-chan error) {
responseChan := make(chan *DescribeSpotPriceHistoryResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.DescribeSpotPriceHistory(request)
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

// DescribeSpotPriceHistoryWithCallback invokes the ecs.DescribeSpotPriceHistory API asynchronously
// api document: https://help.aliyun.com/api/ecs/describespotpricehistory.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSpotPriceHistoryWithCallback(request *DescribeSpotPriceHistoryRequest, callback func(response *DescribeSpotPriceHistoryResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *DescribeSpotPriceHistoryResponse
var err error
defer close(result)
response, err = client.DescribeSpotPriceHistory(request)
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

// DescribeSpotPriceHistoryRequest is the request struct for api DescribeSpotPriceHistory
type DescribeSpotPriceHistoryRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    IoOptimized     string `position:"Query" name:"IoOptimized"`
                    NetworkType     string `position:"Query" name:"NetworkType"`
                    StartTime     string `position:"Query" name:"StartTime"`
                    InstanceType     string `position:"Query" name:"InstanceType"`
                    Offset     requests.Integer `position:"Query" name:"Offset"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    OwnerAccount     string `position:"Query" name:"OwnerAccount"`
                    EndTime     string `position:"Query" name:"EndTime"`
                    OSType     string `position:"Query" name:"OSType"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
                    ZoneId     string `position:"Query" name:"ZoneId"`
}


// DescribeSpotPriceHistoryResponse is the response struct for api DescribeSpotPriceHistory
type DescribeSpotPriceHistoryResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            NextOffset     int `json:"NextOffset" xml:"NextOffset"`
            Currency     string `json:"Currency" xml:"Currency"`
                    SpotPrices SpotPrices `json:"SpotPrices" xml:"SpotPrices"`
}

// CreateDescribeSpotPriceHistoryRequest creates a request to invoke DescribeSpotPriceHistory API
func CreateDescribeSpotPriceHistoryRequest() (request *DescribeSpotPriceHistoryRequest) {
request = &DescribeSpotPriceHistoryRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeSpotPriceHistory", "ecs", "openAPI")
return
}

// CreateDescribeSpotPriceHistoryResponse creates a response to parse from DescribeSpotPriceHistory response
func CreateDescribeSpotPriceHistoryResponse() (response *DescribeSpotPriceHistoryResponse) {
response = &DescribeSpotPriceHistoryResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


