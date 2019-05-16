
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

// DescribeEniQosGroupInfo invokes the ecs.DescribeEniQosGroupInfo API synchronously
// api document: https://help.aliyun.com/api/ecs/describeeniqosgroupinfo.html
func (client *Client) DescribeEniQosGroupInfo(request *DescribeEniQosGroupInfoRequest) (response *DescribeEniQosGroupInfoResponse, err error) {
response = CreateDescribeEniQosGroupInfoResponse()
err = client.DoAction(request, response)
return
}

// DescribeEniQosGroupInfoWithChan invokes the ecs.DescribeEniQosGroupInfo API asynchronously
// api document: https://help.aliyun.com/api/ecs/describeeniqosgroupinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeEniQosGroupInfoWithChan(request *DescribeEniQosGroupInfoRequest) (<-chan *DescribeEniQosGroupInfoResponse, <-chan error) {
responseChan := make(chan *DescribeEniQosGroupInfoResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.DescribeEniQosGroupInfo(request)
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

// DescribeEniQosGroupInfoWithCallback invokes the ecs.DescribeEniQosGroupInfo API asynchronously
// api document: https://help.aliyun.com/api/ecs/describeeniqosgroupinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeEniQosGroupInfoWithCallback(request *DescribeEniQosGroupInfoRequest, callback func(response *DescribeEniQosGroupInfoResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *DescribeEniQosGroupInfoResponse
var err error
defer close(result)
response, err = client.DescribeEniQosGroupInfo(request)
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

// DescribeEniQosGroupInfoRequest is the request struct for api DescribeEniQosGroupInfo
type DescribeEniQosGroupInfoRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
                    QosGroupName     string `position:"Query" name:"QosGroupName"`
}


// DescribeEniQosGroupInfoResponse is the response struct for api DescribeEniQosGroupInfo
type DescribeEniQosGroupInfoResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            QosGroupInfo QosGroupInfo  `json:"QosGroupInfo" xml:"QosGroupInfo"`
}

// CreateDescribeEniQosGroupInfoRequest creates a request to invoke DescribeEniQosGroupInfo API
func CreateDescribeEniQosGroupInfoRequest() (request *DescribeEniQosGroupInfoRequest) {
request = &DescribeEniQosGroupInfoRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Ecs", "2016-03-14", "DescribeEniQosGroupInfo", "", "")
return
}

// CreateDescribeEniQosGroupInfoResponse creates a response to parse from DescribeEniQosGroupInfo response
func CreateDescribeEniQosGroupInfoResponse() (response *DescribeEniQosGroupInfoResponse) {
response = &DescribeEniQosGroupInfoResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

