
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

// PutVersionDoc invokes the acs.PutVersionDoc API synchronously
// api document: https://help.aliyun.com/api/acs/putversiondoc.html
func (client *Client) PutVersionDoc(request *PutVersionDocRequest) (response *PutVersionDocResponse, err error) {
response = CreatePutVersionDocResponse()
err = client.DoAction(request, response)
return
}

// PutVersionDocWithChan invokes the acs.PutVersionDoc API asynchronously
// api document: https://help.aliyun.com/api/acs/putversiondoc.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PutVersionDocWithChan(request *PutVersionDocRequest) (<-chan *PutVersionDocResponse, <-chan error) {
responseChan := make(chan *PutVersionDocResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.PutVersionDoc(request)
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

// PutVersionDocWithCallback invokes the acs.PutVersionDoc API asynchronously
// api document: https://help.aliyun.com/api/acs/putversiondoc.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PutVersionDocWithCallback(request *PutVersionDocRequest, callback func(response *PutVersionDocResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *PutVersionDocResponse
var err error
defer close(result)
response, err = client.PutVersionDoc(request)
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

// PutVersionDocRequest is the request struct for api PutVersionDoc
type PutVersionDocRequest struct {
*requests.RoaRequest
                    ContentLength     requests.Integer `position:"Header" name:"Content-Length"`
                    BodyContent     string `position:"Body" name:"BodyContent"`
                    ContentMD5     string `position:"Header" name:"Content-MD5"`
                    ProductName     string `position:"Path" name:"ProductName"`
                    ContentType     string `position:"Header" name:"Content-Type"`
                    Language     string `position:"Path" name:"Language"`
                    VersionName     string `position:"Path" name:"VersionName"`
                    Accept     string `position:"Header" name:"Accept"`
}


// PutVersionDocResponse is the response struct for api PutVersionDoc
type PutVersionDocResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
}

// CreatePutVersionDocRequest creates a request to invoke PutVersionDoc API
func CreatePutVersionDocRequest() (request *PutVersionDocRequest) {
request = &PutVersionDocRequest{
RoaRequest: &requests.RoaRequest{},
}
request.InitWithApiInfo("Acs", "2015-01-01", "PutVersionDoc", "/VersionDoc/[ProductName]/[VersionName]/[Language]", "", "")
request.Method = requests.PUT
return
}

// CreatePutVersionDocResponse creates a response to parse from PutVersionDoc response
func CreatePutVersionDocResponse() (response *PutVersionDocResponse) {
response = &PutVersionDocResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

