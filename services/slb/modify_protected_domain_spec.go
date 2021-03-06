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

// ModifyProtectedDomainSpec invokes the slb.ModifyProtectedDomainSpec API synchronously
// api document: https://help.aliyun.com/api/slb/modifyprotecteddomainspec.html
func (client *Client) ModifyProtectedDomainSpec(request *ModifyProtectedDomainSpecRequest) (response *ModifyProtectedDomainSpecResponse, err error) {
	response = CreateModifyProtectedDomainSpecResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyProtectedDomainSpecWithChan invokes the slb.ModifyProtectedDomainSpec API asynchronously
// api document: https://help.aliyun.com/api/slb/modifyprotecteddomainspec.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyProtectedDomainSpecWithChan(request *ModifyProtectedDomainSpecRequest) (<-chan *ModifyProtectedDomainSpecResponse, <-chan error) {
	responseChan := make(chan *ModifyProtectedDomainSpecResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyProtectedDomainSpec(request)
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

// ModifyProtectedDomainSpecWithCallback invokes the slb.ModifyProtectedDomainSpec API asynchronously
// api document: https://help.aliyun.com/api/slb/modifyprotecteddomainspec.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyProtectedDomainSpecWithCallback(request *ModifyProtectedDomainSpecRequest, callback func(response *ModifyProtectedDomainSpecResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyProtectedDomainSpecResponse
		var err error
		defer close(result)
		response, err = client.ModifyProtectedDomainSpec(request)
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

// ModifyProtectedDomainSpecRequest is the request struct for api ModifyProtectedDomainSpec
type ModifyProtectedDomainSpecRequest struct {
	*requests.RpcRequest
	AccessKeyId          string           `position:"Query" name:"access_key_id"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	DomainName           string           `position:"Query" name:"DomainName"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	WafPolicyLevel       string           `position:"Query" name:"WafPolicyLevel"`
}

// ModifyProtectedDomainSpecResponse is the response struct for api ModifyProtectedDomainSpec
type ModifyProtectedDomainSpecResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyProtectedDomainSpecRequest creates a request to invoke ModifyProtectedDomainSpec API
func CreateModifyProtectedDomainSpecRequest() (request *ModifyProtectedDomainSpecRequest) {
	request = &ModifyProtectedDomainSpecRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "ModifyProtectedDomainSpec", "slb", "openAPI")
	return
}

// CreateModifyProtectedDomainSpecResponse creates a response to parse from ModifyProtectedDomainSpec response
func CreateModifyProtectedDomainSpecResponse() (response *ModifyProtectedDomainSpecResponse) {
	response = &ModifyProtectedDomainSpecResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
