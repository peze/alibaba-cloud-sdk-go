
package bssopenapi

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

// SetResellerUserAlarmThreshold invokes the bssopenapi.SetResellerUserAlarmThreshold API synchronously
// api document: https://help.aliyun.com/api/bssopenapi/setreselleruseralarmthreshold.html
func (client *Client) SetResellerUserAlarmThreshold(request *SetResellerUserAlarmThresholdRequest) (response *SetResellerUserAlarmThresholdResponse, err error) {
response = CreateSetResellerUserAlarmThresholdResponse()
err = client.DoAction(request, response)
return
}

// SetResellerUserAlarmThresholdWithChan invokes the bssopenapi.SetResellerUserAlarmThreshold API asynchronously
// api document: https://help.aliyun.com/api/bssopenapi/setreselleruseralarmthreshold.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetResellerUserAlarmThresholdWithChan(request *SetResellerUserAlarmThresholdRequest) (<-chan *SetResellerUserAlarmThresholdResponse, <-chan error) {
responseChan := make(chan *SetResellerUserAlarmThresholdResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.SetResellerUserAlarmThreshold(request)
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

// SetResellerUserAlarmThresholdWithCallback invokes the bssopenapi.SetResellerUserAlarmThreshold API asynchronously
// api document: https://help.aliyun.com/api/bssopenapi/setreselleruseralarmthreshold.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetResellerUserAlarmThresholdWithCallback(request *SetResellerUserAlarmThresholdRequest, callback func(response *SetResellerUserAlarmThresholdResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *SetResellerUserAlarmThresholdResponse
var err error
defer close(result)
response, err = client.SetResellerUserAlarmThreshold(request)
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

// SetResellerUserAlarmThresholdRequest is the request struct for api SetResellerUserAlarmThreshold
type SetResellerUserAlarmThresholdRequest struct {
*requests.RpcRequest
                    AlarmType     string `position:"Query" name:"AlarmType"`
                    AlarmThresholds     string `position:"Query" name:"AlarmThresholds"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
}


// SetResellerUserAlarmThresholdResponse is the response struct for api SetResellerUserAlarmThreshold
type SetResellerUserAlarmThresholdResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            Code     string `json:"Code" xml:"Code"`
            Message     string `json:"Message" xml:"Message"`
            Success     bool `json:"Success" xml:"Success"`
            Data     bool `json:"Data" xml:"Data"`
}

// CreateSetResellerUserAlarmThresholdRequest creates a request to invoke SetResellerUserAlarmThreshold API
func CreateSetResellerUserAlarmThresholdRequest() (request *SetResellerUserAlarmThresholdRequest) {
request = &SetResellerUserAlarmThresholdRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("BssOpenApi", "2017-12-14", "SetResellerUserAlarmThreshold", "", "")
return
}

// CreateSetResellerUserAlarmThresholdResponse creates a response to parse from SetResellerUserAlarmThreshold response
func CreateSetResellerUserAlarmThresholdResponse() (response *SetResellerUserAlarmThresholdResponse) {
response = &SetResellerUserAlarmThresholdResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


