
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

// DescribeSQLDiagnosisList invokes the rds.DescribeSQLDiagnosisList API synchronously
// api document: https://help.aliyun.com/api/rds/describesqldiagnosislist.html
func (client *Client) DescribeSQLDiagnosisList(request *DescribeSQLDiagnosisListRequest) (response *DescribeSQLDiagnosisListResponse, err error) {
response = CreateDescribeSQLDiagnosisListResponse()
err = client.DoAction(request, response)
return
}

// DescribeSQLDiagnosisListWithChan invokes the rds.DescribeSQLDiagnosisList API asynchronously
// api document: https://help.aliyun.com/api/rds/describesqldiagnosislist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSQLDiagnosisListWithChan(request *DescribeSQLDiagnosisListRequest) (<-chan *DescribeSQLDiagnosisListResponse, <-chan error) {
responseChan := make(chan *DescribeSQLDiagnosisListResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.DescribeSQLDiagnosisList(request)
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

// DescribeSQLDiagnosisListWithCallback invokes the rds.DescribeSQLDiagnosisList API asynchronously
// api document: https://help.aliyun.com/api/rds/describesqldiagnosislist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSQLDiagnosisListWithCallback(request *DescribeSQLDiagnosisListRequest, callback func(response *DescribeSQLDiagnosisListResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *DescribeSQLDiagnosisListResponse
var err error
defer close(result)
response, err = client.DescribeSQLDiagnosisList(request)
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

// DescribeSQLDiagnosisListRequest is the request struct for api DescribeSQLDiagnosisList
type DescribeSQLDiagnosisListRequest struct {
*requests.RpcRequest
                    DBInstanceId     string `position:"Query" name:"DBInstanceId"`
}


// DescribeSQLDiagnosisListResponse is the response struct for api DescribeSQLDiagnosisList
type DescribeSQLDiagnosisListResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
                    SQLDiagList []SQLDiag  `json:"SQLDiagList" xml:"SQLDiagList"`
}

// CreateDescribeSQLDiagnosisListRequest creates a request to invoke DescribeSQLDiagnosisList API
func CreateDescribeSQLDiagnosisListRequest() (request *DescribeSQLDiagnosisListRequest) {
request = &DescribeSQLDiagnosisListRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Rds", "2014-08-15", "DescribeSQLDiagnosisList", "rds", "openAPI")
return
}

// CreateDescribeSQLDiagnosisListResponse creates a response to parse from DescribeSQLDiagnosisList response
func CreateDescribeSQLDiagnosisListResponse() (response *DescribeSQLDiagnosisListResponse) {
response = &DescribeSQLDiagnosisListResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

