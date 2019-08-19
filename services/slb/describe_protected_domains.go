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

// DescribeProtectedDomains invokes the slb.DescribeProtectedDomains API synchronously
// api document: https://help.aliyun.com/api/slb/describeprotecteddomains.html
func (client *Client) DescribeProtectedDomains(request *DescribeProtectedDomainsRequest) (response *DescribeProtectedDomainsResponse, err error) {
	response = CreateDescribeProtectedDomainsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeProtectedDomainsWithChan invokes the slb.DescribeProtectedDomains API asynchronously
// api document: https://help.aliyun.com/api/slb/describeprotecteddomains.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeProtectedDomainsWithChan(request *DescribeProtectedDomainsRequest) (<-chan *DescribeProtectedDomainsResponse, <-chan error) {
	responseChan := make(chan *DescribeProtectedDomainsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeProtectedDomains(request)
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

// DescribeProtectedDomainsWithCallback invokes the slb.DescribeProtectedDomains API asynchronously
// api document: https://help.aliyun.com/api/slb/describeprotecteddomains.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeProtectedDomainsWithCallback(request *DescribeProtectedDomainsRequest, callback func(response *DescribeProtectedDomainsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeProtectedDomainsResponse
		var err error
		defer close(result)
		response, err = client.DescribeProtectedDomains(request)
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

// DescribeProtectedDomainsRequest is the request struct for api DescribeProtectedDomains
type DescribeProtectedDomainsRequest struct {
	*requests.RpcRequest
	AccessKeyId          string           `position:"Query" name:"access_key_id"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeProtectedDomainsResponse is the response struct for api DescribeProtectedDomains
type DescribeProtectedDomainsResponse struct {
	*responses.BaseResponse
	RequestId string  `json:"RequestId" xml:"RequestId"`
	Domains   Domains `json:"Domains" xml:"Domains"`
}

// CreateDescribeProtectedDomainsRequest creates a request to invoke DescribeProtectedDomains API
func CreateDescribeProtectedDomainsRequest() (request *DescribeProtectedDomainsRequest) {
	request = &DescribeProtectedDomainsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "DescribeProtectedDomains", "slb", "openAPI")
	return
}

// CreateDescribeProtectedDomainsResponse creates a response to parse from DescribeProtectedDomains response
func CreateDescribeProtectedDomainsResponse() (response *DescribeProtectedDomainsResponse) {
	response = &DescribeProtectedDomainsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}