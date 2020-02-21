package r_kvstore

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

// ModifyIntranetAttribute invokes the r_kvstore.ModifyIntranetAttribute API synchronously
// api document: https://help.aliyun.com/api/r-kvstore/modifyintranetattribute.html
func (client *Client) ModifyIntranetAttribute(request *ModifyIntranetAttributeRequest) (response *ModifyIntranetAttributeResponse, err error) {
	response = CreateModifyIntranetAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyIntranetAttributeWithChan invokes the r_kvstore.ModifyIntranetAttribute API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/modifyintranetattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyIntranetAttributeWithChan(request *ModifyIntranetAttributeRequest) (<-chan *ModifyIntranetAttributeResponse, <-chan error) {
	responseChan := make(chan *ModifyIntranetAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyIntranetAttribute(request)
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

// ModifyIntranetAttributeWithCallback invokes the r_kvstore.ModifyIntranetAttribute API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/modifyintranetattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyIntranetAttributeWithCallback(request *ModifyIntranetAttributeRequest, callback func(response *ModifyIntranetAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyIntranetAttributeResponse
		var err error
		defer close(result)
		response, err = client.ModifyIntranetAttribute(request)
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

// ModifyIntranetAttributeRequest is the request struct for api ModifyIntranetAttribute
type ModifyIntranetAttributeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
}

// ModifyIntranetAttributeResponse is the response struct for api ModifyIntranetAttribute
type ModifyIntranetAttributeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyIntranetAttributeRequest creates a request to invoke ModifyIntranetAttribute API
func CreateModifyIntranetAttributeRequest() (request *ModifyIntranetAttributeRequest) {
	request = &ModifyIntranetAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("R-kvstore", "2015-01-01", "ModifyIntranetAttribute", "", "")
	return
}

// CreateModifyIntranetAttributeResponse creates a response to parse from ModifyIntranetAttribute response
func CreateModifyIntranetAttributeResponse() (response *ModifyIntranetAttributeResponse) {
	response = &ModifyIntranetAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
