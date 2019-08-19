package slb

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

// AddDomainWafWhiteList invokes the slb.AddDomainWafWhiteList API synchronously
// api document: https://help.aliyun.com/api/slb/adddomainwafwhitelist.html
func (client *Client) AddDomainWafWhiteList(request *AddDomainWafWhiteListRequest) (response *AddDomainWafWhiteListResponse, err error) {
	response = CreateAddDomainWafWhiteListResponse()
	err = client.DoAction(request, response)
	return
}

// AddDomainWafWhiteListWithChan invokes the slb.AddDomainWafWhiteList API asynchronously
// api document: https://help.aliyun.com/api/slb/adddomainwafwhitelist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddDomainWafWhiteListWithChan(request *AddDomainWafWhiteListRequest) (<-chan *AddDomainWafWhiteListResponse, <-chan error) {
	responseChan := make(chan *AddDomainWafWhiteListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddDomainWafWhiteList(request)
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

// AddDomainWafWhiteListWithCallback invokes the slb.AddDomainWafWhiteList API asynchronously
// api document: https://help.aliyun.com/api/slb/adddomainwafwhitelist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddDomainWafWhiteListWithCallback(request *AddDomainWafWhiteListRequest, callback func(response *AddDomainWafWhiteListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddDomainWafWhiteListResponse
		var err error
		defer close(result)
		response, err = client.AddDomainWafWhiteList(request)
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

// AddDomainWafWhiteListRequest is the request struct for api AddDomainWafWhiteList
type AddDomainWafWhiteListRequest struct {
	*requests.RpcRequest
	AccessKeyId          string           `position:"Query" name:"access_key_id"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	DomainName           string           `position:"Query" name:"DomainName"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	List                 string           `position:"Query" name:"List"`
}

// AddDomainWafWhiteListResponse is the response struct for api AddDomainWafWhiteList
type AddDomainWafWhiteListResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAddDomainWafWhiteListRequest creates a request to invoke AddDomainWafWhiteList API
func CreateAddDomainWafWhiteListRequest() (request *AddDomainWafWhiteListRequest) {
	request = &AddDomainWafWhiteListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "AddDomainWafWhiteList", "slb", "openAPI")
	return
}

// CreateAddDomainWafWhiteListResponse creates a response to parse from AddDomainWafWhiteList response
func CreateAddDomainWafWhiteListResponse() (response *AddDomainWafWhiteListResponse) {
	response = &AddDomainWafWhiteListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}