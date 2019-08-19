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

// DescribeAliCloudCertificates invokes the slb.DescribeAliCloudCertificates API synchronously
// api document: https://help.aliyun.com/api/slb/describealicloudcertificates.html
func (client *Client) DescribeAliCloudCertificates(request *DescribeAliCloudCertificatesRequest) (response *DescribeAliCloudCertificatesResponse, err error) {
	response = CreateDescribeAliCloudCertificatesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAliCloudCertificatesWithChan invokes the slb.DescribeAliCloudCertificates API asynchronously
// api document: https://help.aliyun.com/api/slb/describealicloudcertificates.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAliCloudCertificatesWithChan(request *DescribeAliCloudCertificatesRequest) (<-chan *DescribeAliCloudCertificatesResponse, <-chan error) {
	responseChan := make(chan *DescribeAliCloudCertificatesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAliCloudCertificates(request)
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

// DescribeAliCloudCertificatesWithCallback invokes the slb.DescribeAliCloudCertificates API asynchronously
// api document: https://help.aliyun.com/api/slb/describealicloudcertificates.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAliCloudCertificatesWithCallback(request *DescribeAliCloudCertificatesRequest, callback func(response *DescribeAliCloudCertificatesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAliCloudCertificatesResponse
		var err error
		defer close(result)
		response, err = client.DescribeAliCloudCertificates(request)
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

// DescribeAliCloudCertificatesRequest is the request struct for api DescribeAliCloudCertificates
type DescribeAliCloudCertificatesRequest struct {
	*requests.RpcRequest
	AccessKeyId          string           `position:"Query" name:"access_key_id"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeAliCloudCertificatesResponse is the response struct for api DescribeAliCloudCertificates
type DescribeAliCloudCertificatesResponse struct {
	*responses.BaseResponse
	RequestId            string               `json:"RequestId" xml:"RequestId"`
	AliCloudCertificates AliCloudCertificates `json:"AliCloudCertificates" xml:"AliCloudCertificates"`
}

// CreateDescribeAliCloudCertificatesRequest creates a request to invoke DescribeAliCloudCertificates API
func CreateDescribeAliCloudCertificatesRequest() (request *DescribeAliCloudCertificatesRequest) {
	request = &DescribeAliCloudCertificatesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "DescribeAliCloudCertificates", "slb", "openAPI")
	return
}

// CreateDescribeAliCloudCertificatesResponse creates a response to parse from DescribeAliCloudCertificates response
func CreateDescribeAliCloudCertificatesResponse() (response *DescribeAliCloudCertificatesResponse) {
	response = &DescribeAliCloudCertificatesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}