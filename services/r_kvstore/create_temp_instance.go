
package r_kvstore

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

// CreateTempInstance invokes the r_kvstore.CreateTempInstance API synchronously
// api document: https://help.aliyun.com/api/r-kvstore/createtempinstance.html
func (client *Client) CreateTempInstance(request *CreateTempInstanceRequest) (response *CreateTempInstanceResponse, err error) {
response = CreateCreateTempInstanceResponse()
err = client.DoAction(request, response)
return
}

// CreateTempInstanceWithChan invokes the r_kvstore.CreateTempInstance API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/createtempinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateTempInstanceWithChan(request *CreateTempInstanceRequest) (<-chan *CreateTempInstanceResponse, <-chan error) {
responseChan := make(chan *CreateTempInstanceResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.CreateTempInstance(request)
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

// CreateTempInstanceWithCallback invokes the r_kvstore.CreateTempInstance API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/createtempinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateTempInstanceWithCallback(request *CreateTempInstanceRequest, callback func(response *CreateTempInstanceResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *CreateTempInstanceResponse
var err error
defer close(result)
response, err = client.CreateTempInstance(request)
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

// CreateTempInstanceRequest is the request struct for api CreateTempInstance
type CreateTempInstanceRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    OwnerAccount     string `position:"Query" name:"OwnerAccount"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
}


// CreateTempInstanceResponse is the response struct for api CreateTempInstance
type CreateTempInstanceResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            InstanceId     string `json:"InstanceId" xml:"InstanceId"`
            SnapshotId     string `json:"SnapshotId" xml:"SnapshotId"`
            TempInstanceId     string `json:"TempInstanceId" xml:"TempInstanceId"`
            Status     string `json:"Status" xml:"Status"`
}

// CreateCreateTempInstanceRequest creates a request to invoke CreateTempInstance API
func CreateCreateTempInstanceRequest() (request *CreateTempInstanceRequest) {
request = &CreateTempInstanceRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("R-kvstore", "2015-01-01", "CreateTempInstance", "kvstore", "openAPI")
return
}

// CreateCreateTempInstanceResponse creates a response to parse from CreateTempInstance response
func CreateCreateTempInstanceResponse() (response *CreateTempInstanceResponse) {
response = &CreateTempInstanceResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


