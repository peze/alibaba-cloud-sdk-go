
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

// AddTags_GatedLaunch invokes the ecs.AddTags_GatedLaunch API synchronously
// api document: https://help.aliyun.com/api/ecs/addtags_gatedlaunch.html
func (client *Client) AddTags_GatedLaunch(request *AddTags_GatedLaunchRequest) (response *AddTags_GatedLaunchResponse, err error) {
response = CreateAddTags_GatedLaunchResponse()
err = client.DoAction(request, response)
return
}

// AddTags_GatedLaunchWithChan invokes the ecs.AddTags_GatedLaunch API asynchronously
// api document: https://help.aliyun.com/api/ecs/addtags_gatedlaunch.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddTags_GatedLaunchWithChan(request *AddTags_GatedLaunchRequest) (<-chan *AddTags_GatedLaunchResponse, <-chan error) {
responseChan := make(chan *AddTags_GatedLaunchResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.AddTags_GatedLaunch(request)
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

// AddTags_GatedLaunchWithCallback invokes the ecs.AddTags_GatedLaunch API asynchronously
// api document: https://help.aliyun.com/api/ecs/addtags_gatedlaunch.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddTags_GatedLaunchWithCallback(request *AddTags_GatedLaunchRequest, callback func(response *AddTags_GatedLaunchResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *AddTags_GatedLaunchResponse
var err error
defer close(result)
response, err = client.AddTags_GatedLaunch(request)
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

// AddTags_GatedLaunchRequest is the request struct for api AddTags_GatedLaunch
type AddTags_GatedLaunchRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    Tag  *[]AddTags_GatedLaunchTag `position:"Query" name:"Tag"  type:"Repeated"`
                    ResourceId     string `position:"Query" name:"ResourceId"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
                    ResourceType     string `position:"Query" name:"ResourceType"`
}

// AddTags_GatedLaunchTag is a repeated param struct in AddTags_GatedLaunchRequest
type AddTags_GatedLaunchTag struct{
        Value     string `name:"Value"`
        Key     string `name:"Key"`
}

// AddTags_GatedLaunchResponse is the response struct for api AddTags_GatedLaunch
type AddTags_GatedLaunchResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
}

// CreateAddTags_GatedLaunchRequest creates a request to invoke AddTags_GatedLaunch API
func CreateAddTags_GatedLaunchRequest() (request *AddTags_GatedLaunchRequest) {
request = &AddTags_GatedLaunchRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Ecs", "2014-05-26", "AddTags_GatedLaunch", "ecs", "openAPI")
return
}

// CreateAddTags_GatedLaunchResponse creates a response to parse from AddTags_GatedLaunch response
func CreateAddTags_GatedLaunchResponse() (response *AddTags_GatedLaunchResponse) {
response = &AddTags_GatedLaunchResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

