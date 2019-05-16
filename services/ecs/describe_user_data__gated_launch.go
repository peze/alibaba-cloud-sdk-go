
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

// DescribeUserData_GatedLaunch invokes the ecs.DescribeUserData_GatedLaunch API synchronously
// api document: https://help.aliyun.com/api/ecs/describeuserdata_gatedlaunch.html
func (client *Client) DescribeUserData_GatedLaunch(request *DescribeUserData_GatedLaunchRequest) (response *DescribeUserData_GatedLaunchResponse, err error) {
response = CreateDescribeUserData_GatedLaunchResponse()
err = client.DoAction(request, response)
return
}

// DescribeUserData_GatedLaunchWithChan invokes the ecs.DescribeUserData_GatedLaunch API asynchronously
// api document: https://help.aliyun.com/api/ecs/describeuserdata_gatedlaunch.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeUserData_GatedLaunchWithChan(request *DescribeUserData_GatedLaunchRequest) (<-chan *DescribeUserData_GatedLaunchResponse, <-chan error) {
responseChan := make(chan *DescribeUserData_GatedLaunchResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.DescribeUserData_GatedLaunch(request)
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

// DescribeUserData_GatedLaunchWithCallback invokes the ecs.DescribeUserData_GatedLaunch API asynchronously
// api document: https://help.aliyun.com/api/ecs/describeuserdata_gatedlaunch.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeUserData_GatedLaunchWithCallback(request *DescribeUserData_GatedLaunchRequest, callback func(response *DescribeUserData_GatedLaunchResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *DescribeUserData_GatedLaunchResponse
var err error
defer close(result)
response, err = client.DescribeUserData_GatedLaunch(request)
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

// DescribeUserData_GatedLaunchRequest is the request struct for api DescribeUserData_GatedLaunch
type DescribeUserData_GatedLaunchRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
                    InstanceId     string `position:"Query" name:"InstanceId"`
}


// DescribeUserData_GatedLaunchResponse is the response struct for api DescribeUserData_GatedLaunch
type DescribeUserData_GatedLaunchResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            RegionId     string `json:"RegionId" xml:"RegionId"`
            InstanceId     string `json:"InstanceId" xml:"InstanceId"`
            UserData     string `json:"UserData" xml:"UserData"`
}

// CreateDescribeUserData_GatedLaunchRequest creates a request to invoke DescribeUserData_GatedLaunch API
func CreateDescribeUserData_GatedLaunchRequest() (request *DescribeUserData_GatedLaunchRequest) {
request = &DescribeUserData_GatedLaunchRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeUserData_GatedLaunch", "", "")
return
}

// CreateDescribeUserData_GatedLaunchResponse creates a response to parse from DescribeUserData_GatedLaunch response
func CreateDescribeUserData_GatedLaunchResponse() (response *DescribeUserData_GatedLaunchResponse) {
response = &DescribeUserData_GatedLaunchResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


