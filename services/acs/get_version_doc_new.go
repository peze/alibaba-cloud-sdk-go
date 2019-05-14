
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

// GetVersionDocNew invokes the acs.GetVersionDocNew API synchronously
// api document: https://help.aliyun.com/api/acs/getversiondocnew.html
func (client *Client) GetVersionDocNew(request *GetVersionDocNewRequest) (response *GetVersionDocNewResponse, err error) {
response = CreateGetVersionDocNewResponse()
err = client.DoAction(request, response)
return
}

// GetVersionDocNewWithChan invokes the acs.GetVersionDocNew API asynchronously
// api document: https://help.aliyun.com/api/acs/getversiondocnew.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetVersionDocNewWithChan(request *GetVersionDocNewRequest) (<-chan *GetVersionDocNewResponse, <-chan error) {
responseChan := make(chan *GetVersionDocNewResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.GetVersionDocNew(request)
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

// GetVersionDocNewWithCallback invokes the acs.GetVersionDocNew API asynchronously
// api document: https://help.aliyun.com/api/acs/getversiondocnew.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetVersionDocNewWithCallback(request *GetVersionDocNewRequest, callback func(response *GetVersionDocNewResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *GetVersionDocNewResponse
var err error
defer close(result)
response, err = client.GetVersionDocNew(request)
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

// GetVersionDocNewRequest is the request struct for api GetVersionDocNew
type GetVersionDocNewRequest struct {
*requests.RoaRequest
                    IsDraft     requests.Boolean `position:"Path" name:"IsDraft"`
                    ProductName     string `position:"Path" name:"ProductName"`
                    Language     string `position:"Path" name:"Language"`
                    VersionName     string `position:"Path" name:"VersionName"`
                    Accept     string `position:"Header" name:"Accept"`
}


// GetVersionDocNewResponse is the response struct for api GetVersionDocNew
type GetVersionDocNewResponse struct {
*responses.BaseResponse
            Product     string `json:"product" xml:"product"`
            Version     string `json:"version" xml:"version"`
            Language     string `json:"language" xml:"language"`
            Index     string `json:"index" xml:"index"`
            Description     string `json:"Description" xml:"Description"`
                    ErrorCodeList ErrorCodeListInGetVersionDocNew `json:"ErrorCodeList" xml:"ErrorCodeList"`
}

// CreateGetVersionDocNewRequest creates a request to invoke GetVersionDocNew API
func CreateGetVersionDocNewRequest() (request *GetVersionDocNewRequest) {
request = &GetVersionDocNewRequest{
RoaRequest: &requests.RoaRequest{},
}
request.InitWithApiInfo("Acs", "2015-01-01", "GetVersionDocNew", "/VersionDocNew/[ProductName]/[VersionName]/[Language]/[IsDraft]", "", "")
request.Method = requests.GET
return
}

// CreateGetVersionDocNewResponse creates a response to parse from GetVersionDocNew response
func CreateGetVersionDocNewResponse() (response *GetVersionDocNewResponse) {
response = &GetVersionDocNewResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


