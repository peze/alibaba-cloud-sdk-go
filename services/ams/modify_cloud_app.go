
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

// ModifyCloudApp invokes the ams.ModifyCloudApp API synchronously
// api document: https://help.aliyun.com/api/ams/modifycloudapp.html
func (client *Client) ModifyCloudApp(request *ModifyCloudAppRequest) (response *ModifyCloudAppResponse, err error) {
response = CreateModifyCloudAppResponse()
err = client.DoAction(request, response)
return
}

// ModifyCloudAppWithChan invokes the ams.ModifyCloudApp API asynchronously
// api document: https://help.aliyun.com/api/ams/modifycloudapp.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyCloudAppWithChan(request *ModifyCloudAppRequest) (<-chan *ModifyCloudAppResponse, <-chan error) {
responseChan := make(chan *ModifyCloudAppResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.ModifyCloudApp(request)
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

// ModifyCloudAppWithCallback invokes the ams.ModifyCloudApp API asynchronously
// api document: https://help.aliyun.com/api/ams/modifycloudapp.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyCloudAppWithCallback(request *ModifyCloudAppRequest, callback func(response *ModifyCloudAppResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *ModifyCloudAppResponse
var err error
defer close(result)
response, err = client.ModifyCloudApp(request)
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

// ModifyCloudAppRequest is the request struct for api ModifyCloudApp
type ModifyCloudAppRequest struct {
*requests.RpcRequest
                    AppName     string `position:"Query" name:"AppName"`
                    AppId     requests.Integer `position:"Query" name:"AppId"`
                    Description     string `position:"Query" name:"Description"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
}


// ModifyCloudAppResponse is the response struct for api ModifyCloudApp
type ModifyCloudAppResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyCloudAppRequest creates a request to invoke ModifyCloudApp API
func CreateModifyCloudAppRequest() (request *ModifyCloudAppRequest) {
request = &ModifyCloudAppRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("AMS", "2016-02-01", "ModifyCloudApp", "", "")
return
}

// CreateModifyCloudAppResponse creates a response to parse from ModifyCloudApp response
func CreateModifyCloudAppResponse() (response *ModifyCloudAppResponse) {
response = &ModifyCloudAppResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

