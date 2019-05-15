
package ams

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

// PermanentDeleteApp invokes the ams.PermanentDeleteApp API synchronously
// api document: https://help.aliyun.com/api/ams/permanentdeleteapp.html
func (client *Client) PermanentDeleteApp(request *PermanentDeleteAppRequest) (response *PermanentDeleteAppResponse, err error) {
response = CreatePermanentDeleteAppResponse()
err = client.DoAction(request, response)
return
}

// PermanentDeleteAppWithChan invokes the ams.PermanentDeleteApp API asynchronously
// api document: https://help.aliyun.com/api/ams/permanentdeleteapp.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PermanentDeleteAppWithChan(request *PermanentDeleteAppRequest) (<-chan *PermanentDeleteAppResponse, <-chan error) {
responseChan := make(chan *PermanentDeleteAppResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.PermanentDeleteApp(request)
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

// PermanentDeleteAppWithCallback invokes the ams.PermanentDeleteApp API asynchronously
// api document: https://help.aliyun.com/api/ams/permanentdeleteapp.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PermanentDeleteAppWithCallback(request *PermanentDeleteAppRequest, callback func(response *PermanentDeleteAppResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *PermanentDeleteAppResponse
var err error
defer close(result)
response, err = client.PermanentDeleteApp(request)
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

// PermanentDeleteAppRequest is the request struct for api PermanentDeleteApp
type PermanentDeleteAppRequest struct {
*requests.RpcRequest
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
                    AppKey     requests.Integer `position:"Query" name:"AppKey"`
}


// PermanentDeleteAppResponse is the response struct for api PermanentDeleteApp
type PermanentDeleteAppResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            Code     string `json:"Code" xml:"Code"`
}

// CreatePermanentDeleteAppRequest creates a request to invoke PermanentDeleteApp API
func CreatePermanentDeleteAppRequest() (request *PermanentDeleteAppRequest) {
request = &PermanentDeleteAppRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("AMS", "2018-02-12", "PermanentDeleteApp", "", "")
return
}

// CreatePermanentDeleteAppResponse creates a response to parse from PermanentDeleteApp response
func CreatePermanentDeleteAppResponse() (response *PermanentDeleteAppResponse) {
response = &PermanentDeleteAppResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


