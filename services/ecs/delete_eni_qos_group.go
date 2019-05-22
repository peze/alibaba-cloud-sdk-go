
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

// DeleteEniQosGroup invokes the ecs.DeleteEniQosGroup API synchronously
// api document: https://help.aliyun.com/api/ecs/deleteeniqosgroup.html
func (client *Client) DeleteEniQosGroup(request *DeleteEniQosGroupRequest) (response *DeleteEniQosGroupResponse, err error) {
response = CreateDeleteEniQosGroupResponse()
err = client.DoAction(request, response)
return
}

// DeleteEniQosGroupWithChan invokes the ecs.DeleteEniQosGroup API asynchronously
// api document: https://help.aliyun.com/api/ecs/deleteeniqosgroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteEniQosGroupWithChan(request *DeleteEniQosGroupRequest) (<-chan *DeleteEniQosGroupResponse, <-chan error) {
responseChan := make(chan *DeleteEniQosGroupResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.DeleteEniQosGroup(request)
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

// DeleteEniQosGroupWithCallback invokes the ecs.DeleteEniQosGroup API asynchronously
// api document: https://help.aliyun.com/api/ecs/deleteeniqosgroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteEniQosGroupWithCallback(request *DeleteEniQosGroupRequest, callback func(response *DeleteEniQosGroupResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *DeleteEniQosGroupResponse
var err error
defer close(result)
response, err = client.DeleteEniQosGroup(request)
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

// DeleteEniQosGroupRequest is the request struct for api DeleteEniQosGroup
type DeleteEniQosGroupRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
                    QosGroupName     string `position:"Query" name:"QosGroupName"`
}


// DeleteEniQosGroupResponse is the response struct for api DeleteEniQosGroup
type DeleteEniQosGroupResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteEniQosGroupRequest creates a request to invoke DeleteEniQosGroup API
func CreateDeleteEniQosGroupRequest() (request *DeleteEniQosGroupRequest) {
request = &DeleteEniQosGroupRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Ecs", "2016-03-14", "DeleteEniQosGroup", "", "")
return
}

// CreateDeleteEniQosGroupResponse creates a response to parse from DeleteEniQosGroup response
func CreateDeleteEniQosGroupResponse() (response *DeleteEniQosGroupResponse) {
response = &DeleteEniQosGroupResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


