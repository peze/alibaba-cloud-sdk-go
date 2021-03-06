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

// DescribeDomainCcWhiteList invokes the slb.DescribeDomainCcWhiteList API synchronously
// api document: https://help.aliyun.com/api/slb/describedomainccwhitelist.html
func (client *Client) DescribeDomainCcWhiteList(request *DescribeDomainCcWhiteListRequest) (response *DescribeDomainCcWhiteListResponse, err error) {
	response = CreateDescribeDomainCcWhiteListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDomainCcWhiteListWithChan invokes the slb.DescribeDomainCcWhiteList API asynchronously
// api document: https://help.aliyun.com/api/slb/describedomainccwhitelist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDomainCcWhiteListWithChan(request *DescribeDomainCcWhiteListRequest) (<-chan *DescribeDomainCcWhiteListResponse, <-chan error) {
	responseChan := make(chan *DescribeDomainCcWhiteListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDomainCcWhiteList(request)
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

// DescribeDomainCcWhiteListWithCallback invokes the slb.DescribeDomainCcWhiteList API asynchronously
// api document: https://help.aliyun.com/api/slb/describedomainccwhitelist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDomainCcWhiteListWithCallback(request *DescribeDomainCcWhiteListRequest, callback func(response *DescribeDomainCcWhiteListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDomainCcWhiteListResponse
		var err error
		defer close(result)
		response, err = client.DescribeDomainCcWhiteList(request)
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

// DescribeDomainCcWhiteListRequest is the request struct for api DescribeDomainCcWhiteList
type DescribeDomainCcWhiteListRequest struct {
	*requests.RpcRequest
	AccessKeyId          string           `position:"Query" name:"access_key_id"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	DomainName           string           `position:"Query" name:"DomainName"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeDomainCcWhiteListResponse is the response struct for api DescribeDomainCcWhiteList
type DescribeDomainCcWhiteListResponse struct {
	*responses.BaseResponse
	RequestId   string                                 `json:"RequestId" xml:"RequestId"`
	CcWhiteList CcWhiteListInDescribeDomainCcWhiteList `json:"CcWhiteList" xml:"CcWhiteList"`
}

// CreateDescribeDomainCcWhiteListRequest creates a request to invoke DescribeDomainCcWhiteList API
func CreateDescribeDomainCcWhiteListRequest() (request *DescribeDomainCcWhiteListRequest) {
	request = &DescribeDomainCcWhiteListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "DescribeDomainCcWhiteList", "slb", "openAPI")
	return
}

// CreateDescribeDomainCcWhiteListResponse creates a response to parse from DescribeDomainCcWhiteList response
func CreateDescribeDomainCcWhiteListResponse() (response *DescribeDomainCcWhiteListResponse) {
	response = &DescribeDomainCcWhiteListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
