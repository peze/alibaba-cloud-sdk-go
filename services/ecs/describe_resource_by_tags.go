
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

// DescribeResourceByTags invokes the ecs.DescribeResourceByTags API synchronously
// api document: https://help.aliyun.com/api/ecs/describeresourcebytags.html
func (client *Client) DescribeResourceByTags(request *DescribeResourceByTagsRequest) (response *DescribeResourceByTagsResponse, err error) {
response = CreateDescribeResourceByTagsResponse()
err = client.DoAction(request, response)
return
}

// DescribeResourceByTagsWithChan invokes the ecs.DescribeResourceByTags API asynchronously
// api document: https://help.aliyun.com/api/ecs/describeresourcebytags.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeResourceByTagsWithChan(request *DescribeResourceByTagsRequest) (<-chan *DescribeResourceByTagsResponse, <-chan error) {
responseChan := make(chan *DescribeResourceByTagsResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.DescribeResourceByTags(request)
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

// DescribeResourceByTagsWithCallback invokes the ecs.DescribeResourceByTags API asynchronously
// api document: https://help.aliyun.com/api/ecs/describeresourcebytags.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeResourceByTagsWithCallback(request *DescribeResourceByTagsRequest, callback func(response *DescribeResourceByTagsResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *DescribeResourceByTagsResponse
var err error
defer close(result)
response, err = client.DescribeResourceByTags(request)
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

// DescribeResourceByTagsRequest is the request struct for api DescribeResourceByTags
type DescribeResourceByTagsRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    PageNumber     requests.Integer `position:"Query" name:"PageNumber"`
                    PageSize     requests.Integer `position:"Query" name:"PageSize"`
                    Tag  *[]DescribeResourceByTagsTag `position:"Query" name:"Tag"  type:"Repeated"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
                    ResourceType     string `position:"Query" name:"ResourceType"`
}

// DescribeResourceByTagsTag is a repeated param struct in DescribeResourceByTagsRequest
type DescribeResourceByTagsTag struct{
        Value     string `name:"Value"`
        Key     string `name:"Key"`
}

// DescribeResourceByTagsResponse is the response struct for api DescribeResourceByTags
type DescribeResourceByTagsResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            PageSize     int `json:"PageSize" xml:"PageSize"`
            PageNumber     int `json:"PageNumber" xml:"PageNumber"`
            TotalCount     int `json:"TotalCount" xml:"TotalCount"`
                    Resources Resources `json:"Resources" xml:"Resources"`
}

// CreateDescribeResourceByTagsRequest creates a request to invoke DescribeResourceByTags API
func CreateDescribeResourceByTagsRequest() (request *DescribeResourceByTagsRequest) {
request = &DescribeResourceByTagsRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeResourceByTags", "ecs", "openAPI")
return
}

// CreateDescribeResourceByTagsResponse creates a response to parse from DescribeResourceByTags response
func CreateDescribeResourceByTagsResponse() (response *DescribeResourceByTagsResponse) {
response = &DescribeResourceByTagsResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


