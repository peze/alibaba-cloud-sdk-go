
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

// ReInitVolume_GatedLaunch invokes the ecs.ReInitVolume_GatedLaunch API synchronously
// api document: https://help.aliyun.com/api/ecs/reinitvolume_gatedlaunch.html
func (client *Client) ReInitVolume_GatedLaunch(request *ReInitVolume_GatedLaunchRequest) (response *ReInitVolume_GatedLaunchResponse, err error) {
response = CreateReInitVolume_GatedLaunchResponse()
err = client.DoAction(request, response)
return
}

// ReInitVolume_GatedLaunchWithChan invokes the ecs.ReInitVolume_GatedLaunch API asynchronously
// api document: https://help.aliyun.com/api/ecs/reinitvolume_gatedlaunch.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ReInitVolume_GatedLaunchWithChan(request *ReInitVolume_GatedLaunchRequest) (<-chan *ReInitVolume_GatedLaunchResponse, <-chan error) {
responseChan := make(chan *ReInitVolume_GatedLaunchResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.ReInitVolume_GatedLaunch(request)
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

// ReInitVolume_GatedLaunchWithCallback invokes the ecs.ReInitVolume_GatedLaunch API asynchronously
// api document: https://help.aliyun.com/api/ecs/reinitvolume_gatedlaunch.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ReInitVolume_GatedLaunchWithCallback(request *ReInitVolume_GatedLaunchRequest, callback func(response *ReInitVolume_GatedLaunchResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *ReInitVolume_GatedLaunchResponse
var err error
defer close(result)
response, err = client.ReInitVolume_GatedLaunch(request)
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

// ReInitVolume_GatedLaunchRequest is the request struct for api ReInitVolume_GatedLaunch
type ReInitVolume_GatedLaunchRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    Password     string `position:"Query" name:"Password"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    OwnerAccount     string `position:"Query" name:"OwnerAccount"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
                    VolumeId     string `position:"Query" name:"VolumeId"`
}


// ReInitVolume_GatedLaunchResponse is the response struct for api ReInitVolume_GatedLaunch
type ReInitVolume_GatedLaunchResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
}

// CreateReInitVolume_GatedLaunchRequest creates a request to invoke ReInitVolume_GatedLaunch API
func CreateReInitVolume_GatedLaunchRequest() (request *ReInitVolume_GatedLaunchRequest) {
request = &ReInitVolume_GatedLaunchRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Ecs", "2014-05-26", "ReInitVolume_GatedLaunch", "ecs", "openAPI")
return
}

// CreateReInitVolume_GatedLaunchResponse creates a response to parse from ReInitVolume_GatedLaunch response
func CreateReInitVolume_GatedLaunchResponse() (response *ReInitVolume_GatedLaunchResponse) {
response = &ReInitVolume_GatedLaunchResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


