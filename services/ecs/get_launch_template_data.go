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

// GetLaunchTemplateData invokes the ecs.GetLaunchTemplateData API synchronously
// api document: https://help.aliyun.com/api/ecs/getlaunchtemplatedata.html
func (client *Client) GetLaunchTemplateData(request *GetLaunchTemplateDataRequest) (response *GetLaunchTemplateDataResponse, err error) {
	response = CreateGetLaunchTemplateDataResponse()
	err = client.DoAction(request, response)
	return
}

// GetLaunchTemplateDataWithChan invokes the ecs.GetLaunchTemplateData API asynchronously
// api document: https://help.aliyun.com/api/ecs/getlaunchtemplatedata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetLaunchTemplateDataWithChan(request *GetLaunchTemplateDataRequest) (<-chan *GetLaunchTemplateDataResponse, <-chan error) {
	responseChan := make(chan *GetLaunchTemplateDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetLaunchTemplateData(request)
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

// GetLaunchTemplateDataWithCallback invokes the ecs.GetLaunchTemplateData API asynchronously
// api document: https://help.aliyun.com/api/ecs/getlaunchtemplatedata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetLaunchTemplateDataWithCallback(request *GetLaunchTemplateDataRequest, callback func(response *GetLaunchTemplateDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetLaunchTemplateDataResponse
		var err error
		defer close(result)
		response, err = client.GetLaunchTemplateData(request)
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

// GetLaunchTemplateDataRequest is the request struct for api GetLaunchTemplateData
type GetLaunchTemplateDataRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// GetLaunchTemplateDataResponse is the response struct for api GetLaunchTemplateData
type GetLaunchTemplateDataResponse struct {
	*responses.BaseResponse
	RequestId          string             `json:"RequestId" xml:"RequestId"`
	LaunchTemplateData LaunchTemplateData `json:"LaunchTemplateData" xml:"LaunchTemplateData"`
}

// CreateGetLaunchTemplateDataRequest creates a request to invoke GetLaunchTemplateData API
func CreateGetLaunchTemplateDataRequest() (request *GetLaunchTemplateDataRequest) {
	request = &GetLaunchTemplateDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2016-03-14", "GetLaunchTemplateData", "", "")
	return
}

// CreateGetLaunchTemplateDataResponse creates a response to parse from GetLaunchTemplateData response
func CreateGetLaunchTemplateDataResponse() (response *GetLaunchTemplateDataResponse) {
	response = &GetLaunchTemplateDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
