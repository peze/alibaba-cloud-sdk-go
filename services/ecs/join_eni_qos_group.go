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

// JoinEniQosGroup invokes the ecs.JoinEniQosGroup API synchronously
// api document: https://help.aliyun.com/api/ecs/joineniqosgroup.html
func (client *Client) JoinEniQosGroup(request *JoinEniQosGroupRequest) (response *JoinEniQosGroupResponse, err error) {
	response = CreateJoinEniQosGroupResponse()
	err = client.DoAction(request, response)
	return
}

// JoinEniQosGroupWithChan invokes the ecs.JoinEniQosGroup API asynchronously
// api document: https://help.aliyun.com/api/ecs/joineniqosgroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) JoinEniQosGroupWithChan(request *JoinEniQosGroupRequest) (<-chan *JoinEniQosGroupResponse, <-chan error) {
	responseChan := make(chan *JoinEniQosGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.JoinEniQosGroup(request)
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

// JoinEniQosGroupWithCallback invokes the ecs.JoinEniQosGroup API asynchronously
// api document: https://help.aliyun.com/api/ecs/joineniqosgroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) JoinEniQosGroupWithCallback(request *JoinEniQosGroupRequest, callback func(response *JoinEniQosGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *JoinEniQosGroupResponse
		var err error
		defer close(result)
		response, err = client.JoinEniQosGroup(request)
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

// JoinEniQosGroupRequest is the request struct for api JoinEniQosGroup
type JoinEniQosGroupRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	QosGroupName         string           `position:"Query" name:"QosGroupName"`
	NetworkInterfaceId   string           `position:"Query" name:"NetworkInterfaceId"`
}

// JoinEniQosGroupResponse is the response struct for api JoinEniQosGroup
type JoinEniQosGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateJoinEniQosGroupRequest creates a request to invoke JoinEniQosGroup API
func CreateJoinEniQosGroupRequest() (request *JoinEniQosGroupRequest) {
	request = &JoinEniQosGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2016-03-14", "JoinEniQosGroup", "", "")
	return
}

// CreateJoinEniQosGroupResponse creates a response to parse from JoinEniQosGroup response
func CreateJoinEniQosGroupResponse() (response *JoinEniQosGroupResponse) {
	response = &JoinEniQosGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
