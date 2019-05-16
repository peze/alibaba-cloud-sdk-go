
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

// QueryEniQosGroupByEni invokes the ecs.QueryEniQosGroupByEni API synchronously
// api document: https://help.aliyun.com/api/ecs/queryeniqosgroupbyeni.html
func (client *Client) QueryEniQosGroupByEni(request *QueryEniQosGroupByEniRequest) (response *QueryEniQosGroupByEniResponse, err error) {
response = CreateQueryEniQosGroupByEniResponse()
err = client.DoAction(request, response)
return
}

// QueryEniQosGroupByEniWithChan invokes the ecs.QueryEniQosGroupByEni API asynchronously
// api document: https://help.aliyun.com/api/ecs/queryeniqosgroupbyeni.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryEniQosGroupByEniWithChan(request *QueryEniQosGroupByEniRequest) (<-chan *QueryEniQosGroupByEniResponse, <-chan error) {
responseChan := make(chan *QueryEniQosGroupByEniResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.QueryEniQosGroupByEni(request)
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

// QueryEniQosGroupByEniWithCallback invokes the ecs.QueryEniQosGroupByEni API asynchronously
// api document: https://help.aliyun.com/api/ecs/queryeniqosgroupbyeni.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryEniQosGroupByEniWithCallback(request *QueryEniQosGroupByEniRequest, callback func(response *QueryEniQosGroupByEniResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *QueryEniQosGroupByEniResponse
var err error
defer close(result)
response, err = client.QueryEniQosGroupByEni(request)
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

// QueryEniQosGroupByEniRequest is the request struct for api QueryEniQosGroupByEni
type QueryEniQosGroupByEniRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
                    NetworkInterfaceId     string `position:"Query" name:"NetworkInterfaceId"`
}


// QueryEniQosGroupByEniResponse is the response struct for api QueryEniQosGroupByEni
type QueryEniQosGroupByEniResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            QosGroupName     string `json:"QosGroupName" xml:"QosGroupName"`
}

// CreateQueryEniQosGroupByEniRequest creates a request to invoke QueryEniQosGroupByEni API
func CreateQueryEniQosGroupByEniRequest() (request *QueryEniQosGroupByEniRequest) {
request = &QueryEniQosGroupByEniRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Ecs", "2016-03-14", "QueryEniQosGroupByEni", "", "")
return
}

// CreateQueryEniQosGroupByEniResponse creates a response to parse from QueryEniQosGroupByEni response
func CreateQueryEniQosGroupByEniResponse() (response *QueryEniQosGroupByEniResponse) {
response = &QueryEniQosGroupByEniResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


