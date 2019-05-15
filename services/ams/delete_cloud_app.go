
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

// DeleteCloudApp invokes the ams.DeleteCloudApp API synchronously
// api document: https://help.aliyun.com/api/ams/deletecloudapp.html
func (client *Client) DeleteCloudApp(request *DeleteCloudAppRequest) (response *DeleteCloudAppResponse, err error) {
response = CreateDeleteCloudAppResponse()
err = client.DoAction(request, response)
return
}

// DeleteCloudAppWithChan invokes the ams.DeleteCloudApp API asynchronously
// api document: https://help.aliyun.com/api/ams/deletecloudapp.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteCloudAppWithChan(request *DeleteCloudAppRequest) (<-chan *DeleteCloudAppResponse, <-chan error) {
responseChan := make(chan *DeleteCloudAppResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.DeleteCloudApp(request)
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

// DeleteCloudAppWithCallback invokes the ams.DeleteCloudApp API asynchronously
// api document: https://help.aliyun.com/api/ams/deletecloudapp.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteCloudAppWithCallback(request *DeleteCloudAppRequest, callback func(response *DeleteCloudAppResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *DeleteCloudAppResponse
var err error
defer close(result)
response, err = client.DeleteCloudApp(request)
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

// DeleteCloudAppRequest is the request struct for api DeleteCloudApp
type DeleteCloudAppRequest struct {
*requests.RpcRequest
                    AppId     requests.Integer `position:"Query" name:"AppId"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
}


// DeleteCloudAppResponse is the response struct for api DeleteCloudApp
type DeleteCloudAppResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteCloudAppRequest creates a request to invoke DeleteCloudApp API
func CreateDeleteCloudAppRequest() (request *DeleteCloudAppRequest) {
request = &DeleteCloudAppRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("AMS", "2016-02-01", "DeleteCloudApp", "", "")
return
}

// CreateDeleteCloudAppResponse creates a response to parse from DeleteCloudApp response
func CreateDeleteCloudAppResponse() (response *DeleteCloudAppResponse) {
response = &DeleteCloudAppResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


