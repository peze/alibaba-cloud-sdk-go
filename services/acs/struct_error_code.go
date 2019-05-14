
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

// ErrorCode is a nested struct in acs response
type ErrorCode struct {
            Type     string `json:"type" xml:"type"`
            Code     string `json:"code" xml:"code"`
            ErrorCode     string `json:"errorCode" xml:"errorCode"`
            HttpStatusCode     int `json:"httpStatusCode" xml:"httpStatusCode"`
            DefaultError     bool `json:"defaultError" xml:"defaultError"`
            Index     string `json:"index" xml:"index"`
            TagId     string `json:"tagId" xml:"tagId"`
            ExtendedErrorCode     string `json:"extendedErrorCode" xml:"extendedErrorCode"`
            ErrorMessage     string `json:"errorMessage" xml:"errorMessage"`
            HttpCode     string `json:"httpCode" xml:"httpCode"`
            Description     string `json:"Description" xml:"Description"`
}
