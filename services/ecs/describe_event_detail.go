
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

// DescribeEventDetail invokes the ecs.DescribeEventDetail API synchronously
// api document: https://help.aliyun.com/api/ecs/describeeventdetail.html
func (client *Client) DescribeEventDetail(request *DescribeEventDetailRequest) (response *DescribeEventDetailResponse, err error) {
response = CreateDescribeEventDetailResponse()
err = client.DoAction(request, response)
return
}

// DescribeEventDetailWithChan invokes the ecs.DescribeEventDetail API asynchronously
// api document: https://help.aliyun.com/api/ecs/describeeventdetail.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeEventDetailWithChan(request *DescribeEventDetailRequest) (<-chan *DescribeEventDetailResponse, <-chan error) {
responseChan := make(chan *DescribeEventDetailResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.DescribeEventDetail(request)
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

// DescribeEventDetailWithCallback invokes the ecs.DescribeEventDetail API asynchronously
// api document: https://help.aliyun.com/api/ecs/describeeventdetail.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeEventDetailWithCallback(request *DescribeEventDetailRequest, callback func(response *DescribeEventDetailResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *DescribeEventDetailResponse
var err error
defer close(result)
response, err = client.DescribeEventDetail(request)
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

// DescribeEventDetailRequest is the request struct for api DescribeEventDetail
type DescribeEventDetailRequest struct {
*requests.RpcRequest
                    EventId     string `position:"Query" name:"EventId"`
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
}


// DescribeEventDetailResponse is the response struct for api DescribeEventDetail
type DescribeEventDetailResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            ResourceId     string `json:"ResourceId" xml:"ResourceId"`
            EventType     string `json:"EventType" xml:"EventType"`
            EventCategory     string `json:"EventCategory" xml:"EventCategory"`
            Status     string `json:"Status" xml:"Status"`
            SupportModify     string `json:"SupportModify" xml:"SupportModify"`
            PlanTime     string `json:"PlanTime" xml:"PlanTime"`
            ExpireTime     string `json:"ExpireTime" xml:"ExpireTime"`
            EventId     string `json:"EventId" xml:"EventId"`
            StartTime     string `json:"StartTime" xml:"StartTime"`
            EndTime     string `json:"EndTime" xml:"EndTime"`
            EffectTime     string `json:"EffectTime" xml:"EffectTime"`
            LimitTime     string `json:"LimitTime" xml:"LimitTime"`
            Mark     string `json:"Mark" xml:"Mark"`
}

// CreateDescribeEventDetailRequest creates a request to invoke DescribeEventDetail API
func CreateDescribeEventDetailRequest() (request *DescribeEventDetailRequest) {
request = &DescribeEventDetailRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeEventDetail", "ecs", "openAPI")
return
}

// CreateDescribeEventDetailResponse creates a response to parse from DescribeEventDetail response
func CreateDescribeEventDetailResponse() (response *DescribeEventDetailResponse) {
response = &DescribeEventDetailResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


