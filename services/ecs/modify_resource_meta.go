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

// ModifyResourceMeta invokes the ecs.ModifyResourceMeta API synchronously
// api document: https://help.aliyun.com/api/ecs/modifyresourcemeta.html
func (client *Client) ModifyResourceMeta(request *ModifyResourceMetaRequest) (response *ModifyResourceMetaResponse, err error) {
	response = CreateModifyResourceMetaResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyResourceMetaWithChan invokes the ecs.ModifyResourceMeta API asynchronously
// api document: https://help.aliyun.com/api/ecs/modifyresourcemeta.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyResourceMetaWithChan(request *ModifyResourceMetaRequest) (<-chan *ModifyResourceMetaResponse, <-chan error) {
	responseChan := make(chan *ModifyResourceMetaResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyResourceMeta(request)
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

// ModifyResourceMetaWithCallback invokes the ecs.ModifyResourceMeta API asynchronously
// api document: https://help.aliyun.com/api/ecs/modifyresourcemeta.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyResourceMetaWithCallback(request *ModifyResourceMetaRequest, callback func(response *ModifyResourceMetaResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyResourceMetaResponse
		var err error
		defer close(result)
		response, err = client.ModifyResourceMeta(request)
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

// ModifyResourceMetaRequest is the request struct for api ModifyResourceMeta
type ModifyResourceMetaRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer          `position:"Query" name:"ResourceOwnerId"`
	ResourceId           string                    `position:"Query" name:"ResourceId"`
	ResourceOwnerAccount string                    `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string                    `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer          `position:"Query" name:"OwnerId"`
	ResourceType         string                    `position:"Query" name:"ResourceType"`
	Meta                 *[]ModifyResourceMetaMeta `position:"Query" name:"Meta"  type:"Repeated"`
}

// ModifyResourceMetaMeta is a repeated param struct in ModifyResourceMetaRequest
type ModifyResourceMetaMeta struct {
	Key   string `name:"key"`
	Value string `name:"value"`
}

// ModifyResourceMetaResponse is the response struct for api ModifyResourceMeta
type ModifyResourceMetaResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyResourceMetaRequest creates a request to invoke ModifyResourceMeta API
func CreateModifyResourceMetaRequest() (request *ModifyResourceMetaRequest) {
	request = &ModifyResourceMetaRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2016-03-14", "ModifyResourceMeta", "", "")
	return
}

// CreateModifyResourceMetaResponse creates a response to parse from ModifyResourceMeta response
func CreateModifyResourceMetaResponse() (response *ModifyResourceMetaResponse) {
	response = &ModifyResourceMetaResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
