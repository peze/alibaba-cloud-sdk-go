
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

// RevokeSecurityGroup_GatedLaunch invokes the ecs.RevokeSecurityGroup_GatedLaunch API synchronously
// api document: https://help.aliyun.com/api/ecs/revokesecuritygroup_gatedlaunch.html
func (client *Client) RevokeSecurityGroup_GatedLaunch(request *RevokeSecurityGroup_GatedLaunchRequest) (response *RevokeSecurityGroup_GatedLaunchResponse, err error) {
response = CreateRevokeSecurityGroup_GatedLaunchResponse()
err = client.DoAction(request, response)
return
}

// RevokeSecurityGroup_GatedLaunchWithChan invokes the ecs.RevokeSecurityGroup_GatedLaunch API asynchronously
// api document: https://help.aliyun.com/api/ecs/revokesecuritygroup_gatedlaunch.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RevokeSecurityGroup_GatedLaunchWithChan(request *RevokeSecurityGroup_GatedLaunchRequest) (<-chan *RevokeSecurityGroup_GatedLaunchResponse, <-chan error) {
responseChan := make(chan *RevokeSecurityGroup_GatedLaunchResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.RevokeSecurityGroup_GatedLaunch(request)
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

// RevokeSecurityGroup_GatedLaunchWithCallback invokes the ecs.RevokeSecurityGroup_GatedLaunch API asynchronously
// api document: https://help.aliyun.com/api/ecs/revokesecuritygroup_gatedlaunch.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RevokeSecurityGroup_GatedLaunchWithCallback(request *RevokeSecurityGroup_GatedLaunchRequest, callback func(response *RevokeSecurityGroup_GatedLaunchResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *RevokeSecurityGroup_GatedLaunchResponse
var err error
defer close(result)
response, err = client.RevokeSecurityGroup_GatedLaunch(request)
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

// RevokeSecurityGroup_GatedLaunchRequest is the request struct for api RevokeSecurityGroup_GatedLaunch
type RevokeSecurityGroup_GatedLaunchRequest struct {
*requests.RpcRequest
                    NicType     string `position:"Query" name:"NicType"`
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    SecurityGroupId     string `position:"Query" name:"SecurityGroupId"`
                    SourceGroupOwnerId     requests.Integer `position:"Query" name:"SourceGroupOwnerId"`
                    SourceGroupOwnerAccount     string `position:"Query" name:"SourceGroupOwnerAccount"`
                    Policy     string `position:"Query" name:"Policy"`
                    PortRange     string `position:"Query" name:"PortRange"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    IpProtocol     string `position:"Query" name:"IpProtocol"`
                    SourceCidrIp     string `position:"Query" name:"SourceCidrIp"`
                    OwnerAccount     string `position:"Query" name:"OwnerAccount"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
                    SourceGroupId     string `position:"Query" name:"SourceGroupId"`
}


// RevokeSecurityGroup_GatedLaunchResponse is the response struct for api RevokeSecurityGroup_GatedLaunch
type RevokeSecurityGroup_GatedLaunchResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
}

// CreateRevokeSecurityGroup_GatedLaunchRequest creates a request to invoke RevokeSecurityGroup_GatedLaunch API
func CreateRevokeSecurityGroup_GatedLaunchRequest() (request *RevokeSecurityGroup_GatedLaunchRequest) {
request = &RevokeSecurityGroup_GatedLaunchRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Ecs", "2014-05-26", "RevokeSecurityGroup_GatedLaunch", "", "")
return
}

// CreateRevokeSecurityGroup_GatedLaunchResponse creates a response to parse from RevokeSecurityGroup_GatedLaunch response
func CreateRevokeSecurityGroup_GatedLaunchResponse() (response *RevokeSecurityGroup_GatedLaunchResponse) {
response = &RevokeSecurityGroup_GatedLaunchResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


