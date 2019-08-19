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

// DescribeDomainWafBlackList invokes the slb.DescribeDomainWafBlackList API synchronously
// api document: https://help.aliyun.com/api/slb/describedomainwafblacklist.html
func (client *Client) DescribeDomainWafBlackList(request *DescribeDomainWafBlackListRequest) (response *DescribeDomainWafBlackListResponse, err error) {
	response = CreateDescribeDomainWafBlackListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDomainWafBlackListWithChan invokes the slb.DescribeDomainWafBlackList API asynchronously
// api document: https://help.aliyun.com/api/slb/describedomainwafblacklist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDomainWafBlackListWithChan(request *DescribeDomainWafBlackListRequest) (<-chan *DescribeDomainWafBlackListResponse, <-chan error) {
	responseChan := make(chan *DescribeDomainWafBlackListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDomainWafBlackList(request)
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

// DescribeDomainWafBlackListWithCallback invokes the slb.DescribeDomainWafBlackList API asynchronously
// api document: https://help.aliyun.com/api/slb/describedomainwafblacklist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDomainWafBlackListWithCallback(request *DescribeDomainWafBlackListRequest, callback func(response *DescribeDomainWafBlackListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDomainWafBlackListResponse
		var err error
		defer close(result)
		response, err = client.DescribeDomainWafBlackList(request)
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

// DescribeDomainWafBlackListRequest is the request struct for api DescribeDomainWafBlackList
type DescribeDomainWafBlackListRequest struct {
	*requests.RpcRequest
	AccessKeyId          string           `position:"Query" name:"access_key_id"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	DomainName           string           `position:"Query" name:"DomainName"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeDomainWafBlackListResponse is the response struct for api DescribeDomainWafBlackList
type DescribeDomainWafBlackListResponse struct {
	*responses.BaseResponse
	RequestId    string                                   `json:"RequestId" xml:"RequestId"`
	WafBlackList WafBlackListInDescribeDomainWafBlackList `json:"WafBlackList" xml:"WafBlackList"`
}

// CreateDescribeDomainWafBlackListRequest creates a request to invoke DescribeDomainWafBlackList API
func CreateDescribeDomainWafBlackListRequest() (request *DescribeDomainWafBlackListRequest) {
	request = &DescribeDomainWafBlackListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "DescribeDomainWafBlackList", "slb", "openAPI")
	return
}

// CreateDescribeDomainWafBlackListResponse creates a response to parse from DescribeDomainWafBlackList response
func CreateDescribeDomainWafBlackListResponse() (response *DescribeDomainWafBlackListResponse) {
	response = &DescribeDomainWafBlackListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}