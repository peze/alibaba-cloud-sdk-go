
package acs

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

// GetApiBackup invokes the acs.GetApiBackup API synchronously
// api document: https://help.aliyun.com/api/acs/getapibackup.html
func (client *Client) GetApiBackup(request *GetApiBackupRequest) (response *GetApiBackupResponse, err error) {
response = CreateGetApiBackupResponse()
err = client.DoAction(request, response)
return
}

// GetApiBackupWithChan invokes the acs.GetApiBackup API asynchronously
// api document: https://help.aliyun.com/api/acs/getapibackup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetApiBackupWithChan(request *GetApiBackupRequest) (<-chan *GetApiBackupResponse, <-chan error) {
responseChan := make(chan *GetApiBackupResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.GetApiBackup(request)
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

// GetApiBackupWithCallback invokes the acs.GetApiBackup API asynchronously
// api document: https://help.aliyun.com/api/acs/getapibackup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetApiBackupWithCallback(request *GetApiBackupRequest, callback func(response *GetApiBackupResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *GetApiBackupResponse
var err error
defer close(result)
response, err = client.GetApiBackup(request)
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

// GetApiBackupRequest is the request struct for api GetApiBackup
type GetApiBackupRequest struct {
*requests.RoaRequest
                    ApiName     string `position:"Path" name:"ApiName"`
                    BackupId     string `position:"Query" name:"BackupId"`
                    ProductName     string `position:"Path" name:"ProductName"`
                    VersionName     string `position:"Path" name:"VersionName"`
                    Accept     string `position:"Header" name:"Accept"`
}


// GetApiBackupResponse is the response struct for api GetApiBackup
type GetApiBackupResponse struct {
*responses.BaseResponse
            Name     string `json:"name" xml:"name"`
            Status     string `json:"status" xml:"status"`
            Product     string `json:"product" xml:"product"`
            Version     string `json:"version" xml:"version"`
            Visibility     string `json:"visibility" xml:"visibility"`
            ControlPolicy     string `json:"controlPolicy" xml:"controlPolicy"`
            ResultMapping     string `json:"ResultMapping" xml:"ResultMapping"`
            IsvProtocol IsvProtocol  `json:"IsvProtocol" xml:"IsvProtocol"`
            IspProtocol IspProtocol  `json:"IspProtocol" xml:"IspProtocol"`
            ErrorMapping ErrorMapping  `json:"ErrorMapping" xml:"ErrorMapping"`
            FlowControl FlowControl  `json:"FlowControl" xml:"FlowControl"`
                    Parameters ParametersInGetApiBackup `json:"Parameters" xml:"Parameters"`
}

// CreateGetApiBackupRequest creates a request to invoke GetApiBackup API
func CreateGetApiBackupRequest() (request *GetApiBackupRequest) {
request = &GetApiBackupRequest{
RoaRequest: &requests.RoaRequest{},
}
request.InitWithApiInfo("Acs", "2015-01-01", "GetApiBackup", "/ApiBackUp/[ProductName]/[VersionName]/[ApiName]", "", "")
request.Method = requests.GET
return
}

// CreateGetApiBackupResponse creates a response to parse from GetApiBackup response
func CreateGetApiBackupResponse() (response *GetApiBackupResponse) {
response = &GetApiBackupResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


