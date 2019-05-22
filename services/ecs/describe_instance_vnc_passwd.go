
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

// DescribeInstanceVncPasswd invokes the ecs.DescribeInstanceVncPasswd API synchronously
// api document: https://help.aliyun.com/api/ecs/describeinstancevncpasswd.html
func (client *Client) DescribeInstanceVncPasswd(request *DescribeInstanceVncPasswdRequest) (response *DescribeInstanceVncPasswdResponse, err error) {
response = CreateDescribeInstanceVncPasswdResponse()
err = client.DoAction(request, response)
return
}

// DescribeInstanceVncPasswdWithChan invokes the ecs.DescribeInstanceVncPasswd API asynchronously
// api document: https://help.aliyun.com/api/ecs/describeinstancevncpasswd.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeInstanceVncPasswdWithChan(request *DescribeInstanceVncPasswdRequest) (<-chan *DescribeInstanceVncPasswdResponse, <-chan error) {
responseChan := make(chan *DescribeInstanceVncPasswdResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.DescribeInstanceVncPasswd(request)
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

// DescribeInstanceVncPasswdWithCallback invokes the ecs.DescribeInstanceVncPasswd API asynchronously
// api document: https://help.aliyun.com/api/ecs/describeinstancevncpasswd.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeInstanceVncPasswdWithCallback(request *DescribeInstanceVncPasswdRequest, callback func(response *DescribeInstanceVncPasswdResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *DescribeInstanceVncPasswdResponse
var err error
defer close(result)
response, err = client.DescribeInstanceVncPasswd(request)
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

// DescribeInstanceVncPasswdRequest is the request struct for api DescribeInstanceVncPasswd
type DescribeInstanceVncPasswdRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    OwnerAccount     string `position:"Query" name:"OwnerAccount"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
                    InstanceId     string `position:"Query" name:"InstanceId"`
}


// DescribeInstanceVncPasswdResponse is the response struct for api DescribeInstanceVncPasswd
type DescribeInstanceVncPasswdResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            VncPasswd     string `json:"VncPasswd" xml:"VncPasswd"`
}

// CreateDescribeInstanceVncPasswdRequest creates a request to invoke DescribeInstanceVncPasswd API
func CreateDescribeInstanceVncPasswdRequest() (request *DescribeInstanceVncPasswdRequest) {
request = &DescribeInstanceVncPasswdRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeInstanceVncPasswd", "", "")
return
}

// CreateDescribeInstanceVncPasswdResponse creates a response to parse from DescribeInstanceVncPasswd response
func CreateDescribeInstanceVncPasswdResponse() (response *DescribeInstanceVncPasswdResponse) {
response = &DescribeInstanceVncPasswdResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


