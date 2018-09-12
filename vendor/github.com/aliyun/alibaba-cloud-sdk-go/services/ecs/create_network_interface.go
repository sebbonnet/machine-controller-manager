package ecs

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

// CreateNetworkInterface invokes the ecs.CreateNetworkInterface API synchronously
// api document: https://help.aliyun.com/api/ecs/createnetworkinterface.html
func (client *Client) CreateNetworkInterface(request *CreateNetworkInterfaceRequest) (response *CreateNetworkInterfaceResponse, err error) {
	response = CreateCreateNetworkInterfaceResponse()
	err = client.DoAction(request, response)
	return
}

// CreateNetworkInterfaceWithChan invokes the ecs.CreateNetworkInterface API asynchronously
// api document: https://help.aliyun.com/api/ecs/createnetworkinterface.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateNetworkInterfaceWithChan(request *CreateNetworkInterfaceRequest) (<-chan *CreateNetworkInterfaceResponse, <-chan error) {
	responseChan := make(chan *CreateNetworkInterfaceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateNetworkInterface(request)
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

// CreateNetworkInterfaceWithCallback invokes the ecs.CreateNetworkInterface API asynchronously
// api document: https://help.aliyun.com/api/ecs/createnetworkinterface.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateNetworkInterfaceWithCallback(request *CreateNetworkInterfaceRequest, callback func(response *CreateNetworkInterfaceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateNetworkInterfaceResponse
		var err error
		defer close(result)
		response, err = client.CreateNetworkInterface(request)
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

// CreateNetworkInterfaceRequest is the request struct for api CreateNetworkInterface
type CreateNetworkInterfaceRequest struct {
	*requests.RpcRequest
}

// CreateNetworkInterfaceResponse is the response struct for api CreateNetworkInterface
type CreateNetworkInterfaceResponse struct {
	*responses.BaseResponse
	RequestId          string `json:"RequestId" xml:"RequestId"`
	NetworkInterfaceId string `json:"NetworkInterfaceId" xml:"NetworkInterfaceId"`
}

// CreateCreateNetworkInterfaceRequest creates a request to invoke CreateNetworkInterface API
func CreateCreateNetworkInterfaceRequest() (request *CreateNetworkInterfaceRequest) {
	request = &CreateNetworkInterfaceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "CreateNetworkInterface", "ecs", "openAPI")
	return
}

// CreateCreateNetworkInterfaceResponse creates a response to parse from CreateNetworkInterface response
func CreateCreateNetworkInterfaceResponse() (response *CreateNetworkInterfaceResponse) {
	response = &CreateNetworkInterfaceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}