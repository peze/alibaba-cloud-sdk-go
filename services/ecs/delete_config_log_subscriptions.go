
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

// DeleteConfigLogSubscriptions invokes the ecs.DeleteConfigLogSubscriptions API synchronously
// api document: https://help.aliyun.com/api/ecs/deleteconfiglogsubscriptions.html
func (client *Client) DeleteConfigLogSubscriptions(request *DeleteConfigLogSubscriptionsRequest) (response *DeleteConfigLogSubscriptionsResponse, err error) {
response = CreateDeleteConfigLogSubscriptionsResponse()
err = client.DoAction(request, response)
return
}

// DeleteConfigLogSubscriptionsWithChan invokes the ecs.DeleteConfigLogSubscriptions API asynchronously
// api document: https://help.aliyun.com/api/ecs/deleteconfiglogsubscriptions.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteConfigLogSubscriptionsWithChan(request *DeleteConfigLogSubscriptionsRequest) (<-chan *DeleteConfigLogSubscriptionsResponse, <-chan error) {
responseChan := make(chan *DeleteConfigLogSubscriptionsResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.DeleteConfigLogSubscriptions(request)
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

// DeleteConfigLogSubscriptionsWithCallback invokes the ecs.DeleteConfigLogSubscriptions API asynchronously
// api document: https://help.aliyun.com/api/ecs/deleteconfiglogsubscriptions.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteConfigLogSubscriptionsWithCallback(request *DeleteConfigLogSubscriptionsRequest, callback func(response *DeleteConfigLogSubscriptionsResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *DeleteConfigLogSubscriptionsResponse
var err error
defer close(result)
response, err = client.DeleteConfigLogSubscriptions(request)
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

// DeleteConfigLogSubscriptionsRequest is the request struct for api DeleteConfigLogSubscriptions
type DeleteConfigLogSubscriptionsRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    OwnerAccount     string `position:"Query" name:"OwnerAccount"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
                    Name  *[]string `position:"Query" name:"Name"  type:"Repeated"`
}


// DeleteConfigLogSubscriptionsResponse is the response struct for api DeleteConfigLogSubscriptions
type DeleteConfigLogSubscriptionsResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteConfigLogSubscriptionsRequest creates a request to invoke DeleteConfigLogSubscriptions API
func CreateDeleteConfigLogSubscriptionsRequest() (request *DeleteConfigLogSubscriptionsRequest) {
request = &DeleteConfigLogSubscriptionsRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Ecs", "2016-03-14", "DeleteConfigLogSubscriptions", "", "")
return
}

// CreateDeleteConfigLogSubscriptionsResponse creates a response to parse from DeleteConfigLogSubscriptions response
func CreateDeleteConfigLogSubscriptionsResponse() (response *DeleteConfigLogSubscriptionsResponse) {
response = &DeleteConfigLogSubscriptionsResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


