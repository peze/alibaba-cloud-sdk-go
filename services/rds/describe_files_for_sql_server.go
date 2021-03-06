package rds

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

// DescribeFilesForSQLServer invokes the rds.DescribeFilesForSQLServer API synchronously
// api document: https://help.aliyun.com/api/rds/describefilesforsqlserver.html
func (client *Client) DescribeFilesForSQLServer(request *DescribeFilesForSQLServerRequest) (response *DescribeFilesForSQLServerResponse, err error) {
	response = CreateDescribeFilesForSQLServerResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeFilesForSQLServerWithChan invokes the rds.DescribeFilesForSQLServer API asynchronously
// api document: https://help.aliyun.com/api/rds/describefilesforsqlserver.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeFilesForSQLServerWithChan(request *DescribeFilesForSQLServerRequest) (<-chan *DescribeFilesForSQLServerResponse, <-chan error) {
	responseChan := make(chan *DescribeFilesForSQLServerResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeFilesForSQLServer(request)
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

// DescribeFilesForSQLServerWithCallback invokes the rds.DescribeFilesForSQLServer API asynchronously
// api document: https://help.aliyun.com/api/rds/describefilesforsqlserver.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeFilesForSQLServerWithCallback(request *DescribeFilesForSQLServerRequest, callback func(response *DescribeFilesForSQLServerResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeFilesForSQLServerResponse
		var err error
		defer close(result)
		response, err = client.DescribeFilesForSQLServer(request)
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

// DescribeFilesForSQLServerRequest is the request struct for api DescribeFilesForSQLServer
type DescribeFilesForSQLServerRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	EndTime              string           `position:"Query" name:"EndTime"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	StartTime            string           `position:"Query" name:"StartTime"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
}

// DescribeFilesForSQLServerResponse is the response struct for api DescribeFilesForSQLServer
type DescribeFilesForSQLServerResponse struct {
	*responses.BaseResponse
	RequestId        string                           `json:"RequestId" xml:"RequestId"`
	DBInstanceId     string                           `json:"DBInstanceId" xml:"DBInstanceId"`
	TotalRecordCount int                              `json:"TotalRecordCount" xml:"TotalRecordCount"`
	PageNumber       int                              `json:"PageNumber" xml:"PageNumber"`
	PageRecordCount  int                              `json:"PageRecordCount" xml:"PageRecordCount"`
	Items            ItemsInDescribeFilesForSQLServer `json:"Items" xml:"Items"`
}

// CreateDescribeFilesForSQLServerRequest creates a request to invoke DescribeFilesForSQLServer API
func CreateDescribeFilesForSQLServerRequest() (request *DescribeFilesForSQLServerRequest) {
	request = &DescribeFilesForSQLServerRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeFilesForSQLServer", "", "")
	return
}

// CreateDescribeFilesForSQLServerResponse creates a response to parse from DescribeFilesForSQLServer response
func CreateDescribeFilesForSQLServerResponse() (response *DescribeFilesForSQLServerResponse) {
	response = &DescribeFilesForSQLServerResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
