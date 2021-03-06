package ivision

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

// CreateProject invokes the ivision.CreateProject API synchronously
// api document: https://help.aliyun.com/api/ivision/createproject.html
func (client *Client) CreateProject(request *CreateProjectRequest) (response *CreateProjectResponse, err error) {
	response = CreateCreateProjectResponse()
	err = client.DoAction(request, response)
	return
}

// CreateProjectWithChan invokes the ivision.CreateProject API asynchronously
// api document: https://help.aliyun.com/api/ivision/createproject.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateProjectWithChan(request *CreateProjectRequest) (<-chan *CreateProjectResponse, <-chan error) {
	responseChan := make(chan *CreateProjectResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateProject(request)
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

// CreateProjectWithCallback invokes the ivision.CreateProject API asynchronously
// api document: https://help.aliyun.com/api/ivision/createproject.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateProjectWithCallback(request *CreateProjectRequest, callback func(response *CreateProjectResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateProjectResponse
		var err error
		defer close(result)
		response, err = client.CreateProject(request)
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

// CreateProjectRequest is the request struct for api CreateProject
type CreateProjectRequest struct {
	*requests.RpcRequest
	Description string           `position:"Query" name:"Description"`
	ShowLog     string           `position:"Query" name:"ShowLog"`
	OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
	ProType     string           `position:"Query" name:"ProType"`
	Name        string           `position:"Query" name:"Name"`
}

// CreateProjectResponse is the response struct for api CreateProject
type CreateProjectResponse struct {
	*responses.BaseResponse
	RequestId string  `json:"RequestId" xml:"RequestId"`
	Project   Project `json:"Project" xml:"Project"`
}

// CreateCreateProjectRequest creates a request to invoke CreateProject API
func CreateCreateProjectRequest() (request *CreateProjectRequest) {
	request = &CreateProjectRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ivision", "2019-03-08", "CreateProject", "ivision", "openAPI")
	return
}

// CreateCreateProjectResponse creates a response to parse from CreateProject response
func CreateCreateProjectResponse() (response *CreateProjectResponse) {
	response = &CreateProjectResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
