
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

// CreateInstance_GatedLaunch invokes the ecs.CreateInstance_GatedLaunch API synchronously
// api document: https://help.aliyun.com/api/ecs/createinstance_gatedlaunch.html
func (client *Client) CreateInstance_GatedLaunch(request *CreateInstance_GatedLaunchRequest) (response *CreateInstance_GatedLaunchResponse, err error) {
response = CreateCreateInstance_GatedLaunchResponse()
err = client.DoAction(request, response)
return
}

// CreateInstance_GatedLaunchWithChan invokes the ecs.CreateInstance_GatedLaunch API asynchronously
// api document: https://help.aliyun.com/api/ecs/createinstance_gatedlaunch.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateInstance_GatedLaunchWithChan(request *CreateInstance_GatedLaunchRequest) (<-chan *CreateInstance_GatedLaunchResponse, <-chan error) {
responseChan := make(chan *CreateInstance_GatedLaunchResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.CreateInstance_GatedLaunch(request)
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

// CreateInstance_GatedLaunchWithCallback invokes the ecs.CreateInstance_GatedLaunch API asynchronously
// api document: https://help.aliyun.com/api/ecs/createinstance_gatedlaunch.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateInstance_GatedLaunchWithCallback(request *CreateInstance_GatedLaunchRequest, callback func(response *CreateInstance_GatedLaunchResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *CreateInstance_GatedLaunchResponse
var err error
defer close(result)
response, err = client.CreateInstance_GatedLaunch(request)
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

// CreateInstance_GatedLaunchRequest is the request struct for api CreateInstance_GatedLaunch
type CreateInstance_GatedLaunchRequest struct {
*requests.RpcRequest
                    DataDisk3Size     requests.Integer `position:"Query" name:"DataDisk.3.Size"`
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    DataDisk3Category     string `position:"Query" name:"DataDisk.3.Category"`
                    HostName     string `position:"Query" name:"HostName"`
                    Password     string `position:"Query" name:"Password"`
                    DataDisk2SnapshotId     string `position:"Query" name:"DataDisk.2.SnapshotId"`
                    DataDisk3DiskName     string `position:"Query" name:"DataDisk.3.DiskName"`
                    DataDisk4Size     requests.Integer `position:"Query" name:"DataDisk.4.Size"`
                    Tag  *[]CreateInstance_GatedLaunchTag `position:"Query" name:"Tag"  type:"Repeated"`
                    NodeControllerId     string `position:"Query" name:"NodeControllerId"`
                    Period     requests.Integer `position:"Query" name:"Period"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
                    DataDisk2DeleteWithInstance     requests.Boolean `position:"Query" name:"DataDisk.2.DeleteWithInstance"`
                    VSwitchId     string `position:"Query" name:"VSwitchId"`
                    PrivateIpAddress     string `position:"Query" name:"PrivateIpAddress"`
                    SpotStrategy     string `position:"Query" name:"SpotStrategy"`
                    InstanceName     string `position:"Query" name:"InstanceName"`
                    DataDisk3DeleteWithInstance     requests.Boolean `position:"Query" name:"DataDisk.3.DeleteWithInstance"`
                    InternetChargeType     string `position:"Query" name:"InternetChargeType"`
                    DataDisk3Device     string `position:"Query" name:"DataDisk.3.Device"`
                    ZoneId     string `position:"Query" name:"ZoneId"`
                    DataDisk4DeleteWithInstance     requests.Boolean `position:"Query" name:"DataDisk.4.DeleteWithInstance"`
                    DataDisk2DiskName     string `position:"Query" name:"DataDisk.2.DiskName"`
                    InternetMaxBandwidthIn     requests.Integer `position:"Query" name:"InternetMaxBandwidthIn"`
                    DataDisk1DeleteWithInstance     requests.Boolean `position:"Query" name:"DataDisk.1.DeleteWithInstance"`
                    UseAdditionalService     requests.Boolean `position:"Query" name:"UseAdditionalService"`
                    ImageId     string `position:"Query" name:"ImageId"`
                    DataDisk1SnapshotId     string `position:"Query" name:"DataDisk.1.SnapshotId"`
                    DataDisk1Device     string `position:"Query" name:"DataDisk.1.Device"`
                    ClientToken     string `position:"Query" name:"ClientToken"`
                    VlanId     string `position:"Query" name:"VlanId"`
                    DataDisk2Device     string `position:"Query" name:"DataDisk.2.Device"`
                    IoOptimized     string `position:"Query" name:"IoOptimized"`
                    SecurityGroupId     string `position:"Query" name:"SecurityGroupId"`
                    InternetMaxBandwidthOut     requests.Integer `position:"Query" name:"InternetMaxBandwidthOut"`
                    DataDisk4Description     string `position:"Query" name:"DataDisk.4.Description"`
                    Description     string `position:"Query" name:"Description"`
                    DataDisk1DiskName     string `position:"Query" name:"DataDisk.1.DiskName"`
                    DataDisk3Description     string `position:"Query" name:"DataDisk.3.Description"`
                    DataDisk4DiskName     string `position:"Query" name:"DataDisk.4.DiskName"`
                    SystemDiskCategory     string `position:"Query" name:"SystemDisk.Category"`
                    UserData     string `position:"Query" name:"UserData"`
                    DataDisk1Description     string `position:"Query" name:"DataDisk.1.Description"`
                    DataDisk4Category     string `position:"Query" name:"DataDisk.4.Category"`
                    DataDisk2Description     string `position:"Query" name:"DataDisk.2.Description"`
                    InstanceType     string `position:"Query" name:"InstanceType"`
                    DataDisk2Category     string `position:"Query" name:"DataDisk.2.Category"`
                    DataDisk1Size     requests.Integer `position:"Query" name:"DataDisk.1.Size"`
                    InstanceChargeType     string `position:"Query" name:"InstanceChargeType"`
                    DeploymentSetId     string `position:"Query" name:"DeploymentSetId"`
                    InnerIpAddress     string `position:"Query" name:"InnerIpAddress"`
                    DataDisk3SnapshotId     string `position:"Query" name:"DataDisk.3.SnapshotId"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    OwnerAccount     string `position:"Query" name:"OwnerAccount"`
                    SystemDiskDiskName     string `position:"Query" name:"SystemDisk.DiskName"`
                    DataDisk2Size     requests.Integer `position:"Query" name:"DataDisk.2.Size"`
                    ClusterId     string `position:"Query" name:"ClusterId"`
                    DataDisk1Category     string `position:"Query" name:"DataDisk.1.Category"`
                    SystemDiskSize     requests.Integer `position:"Query" name:"SystemDisk.Size"`
                    DataDisk4SnapshotId     string `position:"Query" name:"DataDisk.4.SnapshotId"`
                    DataDisk4Device     string `position:"Query" name:"DataDisk.4.Device"`
                    SystemDiskDescription     string `position:"Query" name:"SystemDisk.Description"`
}

// CreateInstance_GatedLaunchTag is a repeated param struct in CreateInstance_GatedLaunchRequest
type CreateInstance_GatedLaunchTag struct{
        Value     string `name:"Value"`
        Key     string `name:"Key"`
}

// CreateInstance_GatedLaunchResponse is the response struct for api CreateInstance_GatedLaunch
type CreateInstance_GatedLaunchResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            InstanceId     string `json:"InstanceId" xml:"InstanceId"`
}

// CreateCreateInstance_GatedLaunchRequest creates a request to invoke CreateInstance_GatedLaunch API
func CreateCreateInstance_GatedLaunchRequest() (request *CreateInstance_GatedLaunchRequest) {
request = &CreateInstance_GatedLaunchRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Ecs", "2014-05-26", "CreateInstance_GatedLaunch", "", "")
return
}

// CreateCreateInstance_GatedLaunchResponse creates a response to parse from CreateInstance_GatedLaunch response
func CreateCreateInstance_GatedLaunchResponse() (response *CreateInstance_GatedLaunchResponse) {
response = &CreateInstance_GatedLaunchResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


