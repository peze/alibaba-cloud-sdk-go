
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

// GetApi invokes the acs.GetApi API synchronously
// api document: https://help.aliyun.com/api/acs/getapi.html
func (client *Client) GetApi(request *GetApiRequest) (response *GetApiResponse, err error) {
response = CreateGetApiResponse()
err = client.DoAction(request, response)
return
}

// GetApiWithChan invokes the acs.GetApi API asynchronously
// api document: https://help.aliyun.com/api/acs/getapi.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetApiWithChan(request *GetApiRequest) (<-chan *GetApiResponse, <-chan error) {
responseChan := make(chan *GetApiResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.GetApi(request)
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

// GetApiWithCallback invokes the acs.GetApi API asynchronously
// api document: https://help.aliyun.com/api/acs/getapi.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetApiWithCallback(request *GetApiRequest, callback func(response *GetApiResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *GetApiResponse
var err error
defer close(result)
response, err = client.GetApi(request)
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

// GetApiRequest is the request struct for api GetApi
type GetApiRequest struct {
*requests.RoaRequest
                    ApiName     string `position:"Path" name:"ApiName"`
                    ProductName     string `position:"Path" name:"ProductName"`
                    VersionName     string `position:"Path" name:"VersionName"`
                    Accept     string `position:"Header" name:"Accept"`
}


// GetApiResponse is the response struct for api GetApi
type GetApiResponse struct {
*responses.BaseResponse
            Name     string `json:"name" xml:"name"`
            AliasName     string `json:"aliasName" xml:"aliasName"`
            Status     string `json:"status" xml:"status"`
            Product     string `json:"product" xml:"product"`
            Version     string `json:"version" xml:"version"`
            Visibility     string `json:"visibility" xml:"visibility"`
            ControlPolicy     string `json:"controlPolicy" xml:"controlPolicy"`
            AuthType     string `json:"authType" xml:"authType"`
            GatedLaunch     string `json:"gatedLaunch" xml:"gatedLaunch"`
            GatedLaunchPolicy     string `json:"gatedLaunchPolicy" xml:"gatedLaunchPolicy"`
            ShowJsonItemName     string `json:"showJsonItemName" xml:"showJsonItemName"`
            ParameterType     string `json:"parameterType" xml:"parameterType"`
            KeepClientResourceOwnerId     string `json:"keepClientResourceOwnerId" xml:"keepClientResourceOwnerId"`
            ResponseLog     string `json:"responseLog" xml:"responseLog"`
            RoaRequestBodyLog     string `json:"roaRequestBodyLog" xml:"roaRequestBodyLog"`
            IsolationType     string `json:"isolationType" xml:"isolationType"`
            ResultMapping     string `json:"ResultMapping" xml:"ResultMapping"`
            Parameters     string `json:"Parameters" xml:"Parameters"`
            IsvProtocol IsvProtocol  `json:"IsvProtocol" xml:"IsvProtocol"`
            IspProtocol IspProtocol  `json:"IspProtocol" xml:"IspProtocol"`
            ErrorMapping ErrorMapping  `json:"ErrorMapping" xml:"ErrorMapping"`
            FlowControl FlowControl  `json:"FlowControl" xml:"FlowControl"`
            ApiSpec ApiSpec  `json:"ApiSpec" xml:"ApiSpec"`
}

// CreateGetApiRequest creates a request to invoke GetApi API
func CreateGetApiRequest() (request *GetApiRequest) {
request = &GetApiRequest{
RoaRequest: &requests.RoaRequest{},
}
request.InitWithApiInfo("Acs", "2015-01-01", "GetApi", "/Api/[ProductName]/[VersionName]/[ApiName]", "", "")
request.Method = requests.GET
return
}

// CreateGetApiResponse creates a response to parse from GetApi response
func CreateGetApiResponse() (response *GetApiResponse) {
response = &GetApiResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


