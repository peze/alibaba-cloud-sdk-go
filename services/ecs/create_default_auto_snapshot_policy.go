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

// CreateDefaultAutoSnapshotPolicy invokes the ecs.CreateDefaultAutoSnapshotPolicy API synchronously
// api document: https://help.aliyun.com/api/ecs/createdefaultautosnapshotpolicy.html
func (client *Client) CreateDefaultAutoSnapshotPolicy(request *CreateDefaultAutoSnapshotPolicyRequest) (response *CreateDefaultAutoSnapshotPolicyResponse, err error) {
	response = CreateCreateDefaultAutoSnapshotPolicyResponse()
	err = client.DoAction(request, response)
	return
}

// CreateDefaultAutoSnapshotPolicyWithChan invokes the ecs.CreateDefaultAutoSnapshotPolicy API asynchronously
// api document: https://help.aliyun.com/api/ecs/createdefaultautosnapshotpolicy.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateDefaultAutoSnapshotPolicyWithChan(request *CreateDefaultAutoSnapshotPolicyRequest) (<-chan *CreateDefaultAutoSnapshotPolicyResponse, <-chan error) {
	responseChan := make(chan *CreateDefaultAutoSnapshotPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateDefaultAutoSnapshotPolicy(request)
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

// CreateDefaultAutoSnapshotPolicyWithCallback invokes the ecs.CreateDefaultAutoSnapshotPolicy API asynchronously
// api document: https://help.aliyun.com/api/ecs/createdefaultautosnapshotpolicy.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateDefaultAutoSnapshotPolicyWithCallback(request *CreateDefaultAutoSnapshotPolicyRequest, callback func(response *CreateDefaultAutoSnapshotPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateDefaultAutoSnapshotPolicyResponse
		var err error
		defer close(result)
		response, err = client.CreateDefaultAutoSnapshotPolicy(request)
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

// CreateDefaultAutoSnapshotPolicyRequest is the request struct for api CreateDefaultAutoSnapshotPolicy
type CreateDefaultAutoSnapshotPolicyRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// CreateDefaultAutoSnapshotPolicyResponse is the response struct for api CreateDefaultAutoSnapshotPolicy
type CreateDefaultAutoSnapshotPolicyResponse struct {
	*responses.BaseResponse
	RequestId              string `json:"RequestId" xml:"RequestId"`
	AutoSnapshotPolicyId   string `json:"AutoSnapshotPolicyId" xml:"AutoSnapshotPolicyId"`
	AutoSnapshotPolicyName string `json:"AutoSnapshotPolicyName" xml:"AutoSnapshotPolicyName"`
	TimePoints             string `json:"TimePoints" xml:"TimePoints"`
	RepeatWeekdays         string `json:"RepeatWeekdays" xml:"RepeatWeekdays"`
	RetentionDays          int    `json:"RetentionDays" xml:"RetentionDays"`
}

// CreateCreateDefaultAutoSnapshotPolicyRequest creates a request to invoke CreateDefaultAutoSnapshotPolicy API
func CreateCreateDefaultAutoSnapshotPolicyRequest() (request *CreateDefaultAutoSnapshotPolicyRequest) {
	request = &CreateDefaultAutoSnapshotPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2016-03-14", "CreateDefaultAutoSnapshotPolicy", "", "")
	return
}

// CreateCreateDefaultAutoSnapshotPolicyResponse creates a response to parse from CreateDefaultAutoSnapshotPolicy response
func CreateCreateDefaultAutoSnapshotPolicyResponse() (response *CreateDefaultAutoSnapshotPolicyResponse) {
	response = &CreateDefaultAutoSnapshotPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
