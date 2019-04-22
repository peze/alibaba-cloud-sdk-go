package polardb

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

// CreateDBCluster invokes the polardb.CreateDBCluster API synchronously
// api document: https://help.aliyun.com/api/polardb/createdbcluster.html
func (client *Client) CreateDBCluster(request *CreateDBClusterRequest) (response *CreateDBClusterResponse, err error) {
	response = CreateCreateDBClusterResponse()
	err = client.DoAction(request, response)
	return
}

// CreateDBClusterWithChan invokes the polardb.CreateDBCluster API asynchronously
// api document: https://help.aliyun.com/api/polardb/createdbcluster.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateDBClusterWithChan(request *CreateDBClusterRequest) (<-chan *CreateDBClusterResponse, <-chan error) {
	responseChan := make(chan *CreateDBClusterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateDBCluster(request)
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

// CreateDBClusterWithCallback invokes the polardb.CreateDBCluster API asynchronously
// api document: https://help.aliyun.com/api/polardb/createdbcluster.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateDBClusterWithCallback(request *CreateDBClusterRequest, callback func(response *CreateDBClusterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateDBClusterResponse
		var err error
		defer close(result)
		response, err = client.CreateDBCluster(request)
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

// CreateDBClusterRequest is the request struct for api CreateDBCluster
type CreateDBClusterRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	DBClusterDescription string           `position:"Query" name:"DBClusterDescription"`
	Period               string           `position:"Query" name:"Period"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	UsedTime             string           `position:"Query" name:"UsedTime"`
	ClusterNetworkType   string           `position:"Query" name:"ClusterNetworkType"`
	VSwitchId            string           `position:"Query" name:"VSwitchId"`
	DBNodeClass          string           `position:"Query" name:"DBNodeClass"`
	Engine               string           `position:"Query" name:"Engine"`
	VPCId                string           `position:"Query" name:"VPCId"`
	DBType               string           `position:"Query" name:"DBType"`
	ZoneId               string           `position:"Query" name:"ZoneId"`
	DBVersion            string           `position:"Query" name:"DBVersion"`
	PayType              string           `position:"Query" name:"PayType"`
}

// CreateDBClusterResponse is the response struct for api CreateDBCluster
type CreateDBClusterResponse struct {
	*responses.BaseResponse
	RequestId   string `json:"RequestId" xml:"RequestId"`
	DBClusterId string `json:"DBClusterId" xml:"DBClusterId"`
	OrderId     string `json:"OrderId" xml:"OrderId"`
}

// CreateCreateDBClusterRequest creates a request to invoke CreateDBCluster API
func CreateCreateDBClusterRequest() (request *CreateDBClusterRequest) {
	request = &CreateDBClusterRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardb", "2017-08-01", "CreateDBCluster", "polardb", "openAPI")
	return
}

// CreateCreateDBClusterResponse creates a response to parse from CreateDBCluster response
func CreateCreateDBClusterResponse() (response *CreateDBClusterResponse) {
	response = &CreateDBClusterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
