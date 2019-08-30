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

// DescribeTransitionVSwitches invokes the ecs.DescribeTransitionVSwitches API synchronously
// api document: https://help.aliyun.com/api/ecs/describetransitionvswitches.html
func (client *Client) DescribeTransitionVSwitches(request *DescribeTransitionVSwitchesRequest) (response *DescribeTransitionVSwitchesResponse, err error) {
	response = CreateDescribeTransitionVSwitchesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeTransitionVSwitchesWithChan invokes the ecs.DescribeTransitionVSwitches API asynchronously
// api document: https://help.aliyun.com/api/ecs/describetransitionvswitches.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeTransitionVSwitchesWithChan(request *DescribeTransitionVSwitchesRequest) (<-chan *DescribeTransitionVSwitchesResponse, <-chan error) {
	responseChan := make(chan *DescribeTransitionVSwitchesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeTransitionVSwitches(request)
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

// DescribeTransitionVSwitchesWithCallback invokes the ecs.DescribeTransitionVSwitches API asynchronously
// api document: https://help.aliyun.com/api/ecs/describetransitionvswitches.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeTransitionVSwitchesWithCallback(request *DescribeTransitionVSwitchesRequest, callback func(response *DescribeTransitionVSwitchesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeTransitionVSwitchesResponse
		var err error
		defer close(result)
		response, err = client.DescribeTransitionVSwitches(request)
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

// DescribeTransitionVSwitchesRequest is the request struct for api DescribeTransitionVSwitches
type DescribeTransitionVSwitchesRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	VpcId                string           `position:"Query" name:"VpcId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceType         string           `position:"Query" name:"ResourceType"`
}

// DescribeTransitionVSwitchesResponse is the response struct for api DescribeTransitionVSwitches
type DescribeTransitionVSwitchesResponse struct {
	*responses.BaseResponse
	RequestId     string        `json:"RequestId" xml:"RequestId"`
	Code          string        `json:"Code" xml:"Code"`
	Success       bool          `json:"Success" xml:"Success"`
	VSwitchModels VSwitchModels `json:"VSwitchModels" xml:"VSwitchModels"`
}

// CreateDescribeTransitionVSwitchesRequest creates a request to invoke DescribeTransitionVSwitches API
func CreateDescribeTransitionVSwitchesRequest() (request *DescribeTransitionVSwitchesRequest) {
	request = &DescribeTransitionVSwitchesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2016-03-14", "DescribeTransitionVSwitches", "", "")
	return
}

// CreateDescribeTransitionVSwitchesResponse creates a response to parse from DescribeTransitionVSwitches response
func CreateDescribeTransitionVSwitchesResponse() (response *DescribeTransitionVSwitchesResponse) {
	response = &DescribeTransitionVSwitchesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
