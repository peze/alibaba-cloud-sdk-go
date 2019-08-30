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

// DescribeInstanceTopology invokes the ecs.DescribeInstanceTopology API synchronously
// api document: https://help.aliyun.com/api/ecs/describeinstancetopology.html
func (client *Client) DescribeInstanceTopology(request *DescribeInstanceTopologyRequest) (response *DescribeInstanceTopologyResponse, err error) {
	response = CreateDescribeInstanceTopologyResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeInstanceTopologyWithChan invokes the ecs.DescribeInstanceTopology API asynchronously
// api document: https://help.aliyun.com/api/ecs/describeinstancetopology.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeInstanceTopologyWithChan(request *DescribeInstanceTopologyRequest) (<-chan *DescribeInstanceTopologyResponse, <-chan error) {
	responseChan := make(chan *DescribeInstanceTopologyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeInstanceTopology(request)
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

// DescribeInstanceTopologyWithCallback invokes the ecs.DescribeInstanceTopology API asynchronously
// api document: https://help.aliyun.com/api/ecs/describeinstancetopology.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeInstanceTopologyWithCallback(request *DescribeInstanceTopologyRequest, callback func(response *DescribeInstanceTopologyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeInstanceTopologyResponse
		var err error
		defer close(result)
		response, err = client.DescribeInstanceTopology(request)
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

// DescribeInstanceTopologyRequest is the request struct for api DescribeInstanceTopology
type DescribeInstanceTopologyRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	InstanceIds          string           `position:"Query" name:"InstanceIds"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeInstanceTopologyResponse is the response struct for api DescribeInstanceTopology
type DescribeInstanceTopologyResponse struct {
	*responses.BaseResponse
	RequestId string    `json:"RequestId" xml:"RequestId"`
	Topologys Topologys `json:"Topologys" xml:"Topologys"`
}

// CreateDescribeInstanceTopologyRequest creates a request to invoke DescribeInstanceTopology API
func CreateDescribeInstanceTopologyRequest() (request *DescribeInstanceTopologyRequest) {
	request = &DescribeInstanceTopologyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeInstanceTopology", "", "")
	return
}

// CreateDescribeInstanceTopologyResponse creates a response to parse from DescribeInstanceTopology response
func CreateDescribeInstanceTopologyResponse() (response *DescribeInstanceTopologyResponse) {
	response = &DescribeInstanceTopologyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
