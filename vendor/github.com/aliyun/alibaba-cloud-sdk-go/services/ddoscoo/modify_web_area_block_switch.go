package ddoscoo

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

// ModifyWebAreaBlockSwitch invokes the ddoscoo.ModifyWebAreaBlockSwitch API synchronously
// api document: https://help.aliyun.com/api/ddoscoo/modifywebareablockswitch.html
func (client *Client) ModifyWebAreaBlockSwitch(request *ModifyWebAreaBlockSwitchRequest) (response *ModifyWebAreaBlockSwitchResponse, err error) {
	response = CreateModifyWebAreaBlockSwitchResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyWebAreaBlockSwitchWithChan invokes the ddoscoo.ModifyWebAreaBlockSwitch API asynchronously
// api document: https://help.aliyun.com/api/ddoscoo/modifywebareablockswitch.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyWebAreaBlockSwitchWithChan(request *ModifyWebAreaBlockSwitchRequest) (<-chan *ModifyWebAreaBlockSwitchResponse, <-chan error) {
	responseChan := make(chan *ModifyWebAreaBlockSwitchResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyWebAreaBlockSwitch(request)
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

// ModifyWebAreaBlockSwitchWithCallback invokes the ddoscoo.ModifyWebAreaBlockSwitch API asynchronously
// api document: https://help.aliyun.com/api/ddoscoo/modifywebareablockswitch.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyWebAreaBlockSwitchWithCallback(request *ModifyWebAreaBlockSwitchRequest, callback func(response *ModifyWebAreaBlockSwitchResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyWebAreaBlockSwitchResponse
		var err error
		defer close(result)
		response, err = client.ModifyWebAreaBlockSwitch(request)
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

// ModifyWebAreaBlockSwitchRequest is the request struct for api ModifyWebAreaBlockSwitch
type ModifyWebAreaBlockSwitchRequest struct {
	*requests.RpcRequest
	ResourceGroupId string `position:"Query" name:"ResourceGroupId"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	Domain          string `position:"Query" name:"Domain"`
	Config          string `position:"Query" name:"Config"`
}

// ModifyWebAreaBlockSwitchResponse is the response struct for api ModifyWebAreaBlockSwitch
type ModifyWebAreaBlockSwitchResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyWebAreaBlockSwitchRequest creates a request to invoke ModifyWebAreaBlockSwitch API
func CreateModifyWebAreaBlockSwitchRequest() (request *ModifyWebAreaBlockSwitchRequest) {
	request = &ModifyWebAreaBlockSwitchRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddoscoo", "2020-01-01", "ModifyWebAreaBlockSwitch", "ddoscoo", "openAPI")
	return
}

// CreateModifyWebAreaBlockSwitchResponse creates a response to parse from ModifyWebAreaBlockSwitch response
func CreateModifyWebAreaBlockSwitchResponse() (response *ModifyWebAreaBlockSwitchResponse) {
	response = &ModifyWebAreaBlockSwitchResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}