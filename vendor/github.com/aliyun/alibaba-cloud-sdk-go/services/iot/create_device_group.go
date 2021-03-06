package iot

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

// CreateDeviceGroup invokes the iot.CreateDeviceGroup API synchronously
// api document: https://help.aliyun.com/api/iot/createdevicegroup.html
func (client *Client) CreateDeviceGroup(request *CreateDeviceGroupRequest) (response *CreateDeviceGroupResponse, err error) {
	response = CreateCreateDeviceGroupResponse()
	err = client.DoAction(request, response)
	return
}

// CreateDeviceGroupWithChan invokes the iot.CreateDeviceGroup API asynchronously
// api document: https://help.aliyun.com/api/iot/createdevicegroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateDeviceGroupWithChan(request *CreateDeviceGroupRequest) (<-chan *CreateDeviceGroupResponse, <-chan error) {
	responseChan := make(chan *CreateDeviceGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateDeviceGroup(request)
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

// CreateDeviceGroupWithCallback invokes the iot.CreateDeviceGroup API asynchronously
// api document: https://help.aliyun.com/api/iot/createdevicegroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateDeviceGroupWithCallback(request *CreateDeviceGroupRequest, callback func(response *CreateDeviceGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateDeviceGroupResponse
		var err error
		defer close(result)
		response, err = client.CreateDeviceGroup(request)
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

// CreateDeviceGroupRequest is the request struct for api CreateDeviceGroup
type CreateDeviceGroupRequest struct {
	*requests.RpcRequest
	SuperGroupId  string `position:"Query" name:"SuperGroupId"`
	GroupName     string `position:"Query" name:"GroupName"`
	GroupDesc     string `position:"Query" name:"GroupDesc"`
	IotInstanceId string `position:"Query" name:"IotInstanceId"`
}

// CreateDeviceGroupResponse is the response struct for api CreateDeviceGroup
type CreateDeviceGroupResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         Data   `json:"Data" xml:"Data"`
}

// CreateCreateDeviceGroupRequest creates a request to invoke CreateDeviceGroup API
func CreateCreateDeviceGroupRequest() (request *CreateDeviceGroupRequest) {
	request = &CreateDeviceGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "CreateDeviceGroup", "iot", "openAPI")
	return
}

// CreateCreateDeviceGroupResponse creates a response to parse from CreateDeviceGroup response
func CreateCreateDeviceGroupResponse() (response *CreateDeviceGroupResponse) {
	response = &CreateDeviceGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
