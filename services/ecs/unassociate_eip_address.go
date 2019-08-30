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

// UnassociateEipAddress invokes the ecs.UnassociateEipAddress API synchronously
// api document: https://help.aliyun.com/api/ecs/unassociateeipaddress.html
func (client *Client) UnassociateEipAddress(request *UnassociateEipAddressRequest) (response *UnassociateEipAddressResponse, err error) {
	response = CreateUnassociateEipAddressResponse()
	err = client.DoAction(request, response)
	return
}

// UnassociateEipAddressWithChan invokes the ecs.UnassociateEipAddress API asynchronously
// api document: https://help.aliyun.com/api/ecs/unassociateeipaddress.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UnassociateEipAddressWithChan(request *UnassociateEipAddressRequest) (<-chan *UnassociateEipAddressResponse, <-chan error) {
	responseChan := make(chan *UnassociateEipAddressResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UnassociateEipAddress(request)
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

// UnassociateEipAddressWithCallback invokes the ecs.UnassociateEipAddress API asynchronously
// api document: https://help.aliyun.com/api/ecs/unassociateeipaddress.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UnassociateEipAddressWithCallback(request *UnassociateEipAddressRequest, callback func(response *UnassociateEipAddressResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UnassociateEipAddressResponse
		var err error
		defer close(result)
		response, err = client.UnassociateEipAddress(request)
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

// UnassociateEipAddressRequest is the request struct for api UnassociateEipAddress
type UnassociateEipAddressRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	InstanceType         string           `position:"Query" name:"InstanceType"`
	AllocationId         string           `position:"Query" name:"AllocationId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// UnassociateEipAddressResponse is the response struct for api UnassociateEipAddress
type UnassociateEipAddressResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUnassociateEipAddressRequest creates a request to invoke UnassociateEipAddress API
func CreateUnassociateEipAddressRequest() (request *UnassociateEipAddressRequest) {
	request = &UnassociateEipAddressRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "UnassociateEipAddress", "", "")
	return
}

// CreateUnassociateEipAddressResponse creates a response to parse from UnassociateEipAddress response
func CreateUnassociateEipAddressResponse() (response *UnassociateEipAddressResponse) {
	response = &UnassociateEipAddressResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
