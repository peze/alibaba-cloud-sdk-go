
package bss

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

// VnoBatchRefundOrder invokes the bss.VnoBatchRefundOrder API synchronously
// api document: https://help.aliyun.com/api/bss/vnobatchrefundorder.html
func (client *Client) VnoBatchRefundOrder(request *VnoBatchRefundOrderRequest) (response *VnoBatchRefundOrderResponse, err error) {
response = CreateVnoBatchRefundOrderResponse()
err = client.DoAction(request, response)
return
}

// VnoBatchRefundOrderWithChan invokes the bss.VnoBatchRefundOrder API asynchronously
// api document: https://help.aliyun.com/api/bss/vnobatchrefundorder.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) VnoBatchRefundOrderWithChan(request *VnoBatchRefundOrderRequest) (<-chan *VnoBatchRefundOrderResponse, <-chan error) {
responseChan := make(chan *VnoBatchRefundOrderResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.VnoBatchRefundOrder(request)
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

// VnoBatchRefundOrderWithCallback invokes the bss.VnoBatchRefundOrder API asynchronously
// api document: https://help.aliyun.com/api/bss/vnobatchrefundorder.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) VnoBatchRefundOrderWithCallback(request *VnoBatchRefundOrderRequest, callback func(response *VnoBatchRefundOrderResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *VnoBatchRefundOrderResponse
var err error
defer close(result)
response, err = client.VnoBatchRefundOrder(request)
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

// VnoBatchRefundOrderRequest is the request struct for api VnoBatchRefundOrder
type VnoBatchRefundOrderRequest struct {
*requests.RpcRequest
                    ParamStr     string `position:"Query" name:"paramStr"`
}


// VnoBatchRefundOrderResponse is the response struct for api VnoBatchRefundOrder
type VnoBatchRefundOrderResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            Success     bool `json:"Success" xml:"Success"`
            Code     string `json:"Code" xml:"Code"`
            Message     string `json:"Message" xml:"Message"`
            Data     string `json:"Data" xml:"Data"`
}

// CreateVnoBatchRefundOrderRequest creates a request to invoke VnoBatchRefundOrder API
func CreateVnoBatchRefundOrderRequest() (request *VnoBatchRefundOrderRequest) {
request = &VnoBatchRefundOrderRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Bss", "2014-07-14", "VnoBatchRefundOrder", "", "")
return
}

// CreateVnoBatchRefundOrderResponse creates a response to parse from VnoBatchRefundOrder response
func CreateVnoBatchRefundOrderResponse() (response *VnoBatchRefundOrderResponse) {
response = &VnoBatchRefundOrderResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


