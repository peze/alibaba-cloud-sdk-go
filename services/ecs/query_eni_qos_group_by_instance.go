
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

// QueryEniQosGroupByInstance invokes the ecs.QueryEniQosGroupByInstance API synchronously
// api document: https://help.aliyun.com/api/ecs/queryeniqosgroupbyinstance.html
func (client *Client) QueryEniQosGroupByInstance(request *QueryEniQosGroupByInstanceRequest) (response *QueryEniQosGroupByInstanceResponse, err error) {
response = CreateQueryEniQosGroupByInstanceResponse()
err = client.DoAction(request, response)
return
}

// QueryEniQosGroupByInstanceWithChan invokes the ecs.QueryEniQosGroupByInstance API asynchronously
// api document: https://help.aliyun.com/api/ecs/queryeniqosgroupbyinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryEniQosGroupByInstanceWithChan(request *QueryEniQosGroupByInstanceRequest) (<-chan *QueryEniQosGroupByInstanceResponse, <-chan error) {
responseChan := make(chan *QueryEniQosGroupByInstanceResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.QueryEniQosGroupByInstance(request)
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

// QueryEniQosGroupByInstanceWithCallback invokes the ecs.QueryEniQosGroupByInstance API asynchronously
// api document: https://help.aliyun.com/api/ecs/queryeniqosgroupbyinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryEniQosGroupByInstanceWithCallback(request *QueryEniQosGroupByInstanceRequest, callback func(response *QueryEniQosGroupByInstanceResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *QueryEniQosGroupByInstanceResponse
var err error
defer close(result)
response, err = client.QueryEniQosGroupByInstance(request)
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

// QueryEniQosGroupByInstanceRequest is the request struct for api QueryEniQosGroupByInstance
type QueryEniQosGroupByInstanceRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
                    InstanceId     string `position:"Query" name:"InstanceId"`
}


// QueryEniQosGroupByInstanceResponse is the response struct for api QueryEniQosGroupByInstance
type QueryEniQosGroupByInstanceResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
                Data Data `json:"Data" xml:"Data"`
}

// CreateQueryEniQosGroupByInstanceRequest creates a request to invoke QueryEniQosGroupByInstance API
func CreateQueryEniQosGroupByInstanceRequest() (request *QueryEniQosGroupByInstanceRequest) {
request = &QueryEniQosGroupByInstanceRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Ecs", "2016-03-14", "QueryEniQosGroupByInstance", "", "")
return
}

// CreateQueryEniQosGroupByInstanceResponse creates a response to parse from QueryEniQosGroupByInstance response
func CreateQueryEniQosGroupByInstanceResponse() (response *QueryEniQosGroupByInstanceResponse) {
response = &QueryEniQosGroupByInstanceResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

