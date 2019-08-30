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

// ModifyInstanceCapacityReservationAttributes invokes the ecs.ModifyInstanceCapacityReservationAttributes API synchronously
// api document: https://help.aliyun.com/api/ecs/modifyinstancecapacityreservationattributes.html
func (client *Client) ModifyInstanceCapacityReservationAttributes(request *ModifyInstanceCapacityReservationAttributesRequest) (response *ModifyInstanceCapacityReservationAttributesResponse, err error) {
	response = CreateModifyInstanceCapacityReservationAttributesResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyInstanceCapacityReservationAttributesWithChan invokes the ecs.ModifyInstanceCapacityReservationAttributes API asynchronously
// api document: https://help.aliyun.com/api/ecs/modifyinstancecapacityreservationattributes.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyInstanceCapacityReservationAttributesWithChan(request *ModifyInstanceCapacityReservationAttributesRequest) (<-chan *ModifyInstanceCapacityReservationAttributesResponse, <-chan error) {
	responseChan := make(chan *ModifyInstanceCapacityReservationAttributesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyInstanceCapacityReservationAttributes(request)
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

// ModifyInstanceCapacityReservationAttributesWithCallback invokes the ecs.ModifyInstanceCapacityReservationAttributes API asynchronously
// api document: https://help.aliyun.com/api/ecs/modifyinstancecapacityreservationattributes.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyInstanceCapacityReservationAttributesWithCallback(request *ModifyInstanceCapacityReservationAttributesRequest, callback func(response *ModifyInstanceCapacityReservationAttributesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyInstanceCapacityReservationAttributesResponse
		var err error
		defer close(result)
		response, err = client.ModifyInstanceCapacityReservationAttributes(request)
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

// ModifyInstanceCapacityReservationAttributesRequest is the request struct for api ModifyInstanceCapacityReservationAttributes
type ModifyInstanceCapacityReservationAttributesRequest struct {
	*requests.RpcRequest
	ResourceOwnerId               requests.Integer `position:"Query" name:"ResourceOwnerId"`
	CapacityReservationId         string           `position:"Query" name:"CapacityReservationId"`
	ResourceOwnerAccount          string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                  string           `position:"Query" name:"OwnerAccount"`
	OwnerId                       requests.Integer `position:"Query" name:"OwnerId"`
	CapacityReservationPreference string           `position:"Query" name:"CapacityReservationPreference"`
	InstanceId                    string           `position:"Query" name:"InstanceId"`
}

// ModifyInstanceCapacityReservationAttributesResponse is the response struct for api ModifyInstanceCapacityReservationAttributes
type ModifyInstanceCapacityReservationAttributesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyInstanceCapacityReservationAttributesRequest creates a request to invoke ModifyInstanceCapacityReservationAttributes API
func CreateModifyInstanceCapacityReservationAttributesRequest() (request *ModifyInstanceCapacityReservationAttributesRequest) {
	request = &ModifyInstanceCapacityReservationAttributesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2016-03-14", "ModifyInstanceCapacityReservationAttributes", "", "")
	return
}

// CreateModifyInstanceCapacityReservationAttributesResponse creates a response to parse from ModifyInstanceCapacityReservationAttributes response
func CreateModifyInstanceCapacityReservationAttributesResponse() (response *ModifyInstanceCapacityReservationAttributesResponse) {
	response = &ModifyInstanceCapacityReservationAttributesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
