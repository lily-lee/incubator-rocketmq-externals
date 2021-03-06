/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package model

import (
	"encoding/json"
	"github.com/apache/incubator-rocketmq-externals/rocketmq-go/api/model"
)

//ConsumerRunningInfo this client's consumer running info
type ConsumerRunningInfo struct {
	Properties map[string]string                           `json:"properties"`
	MqTable    map[rocketmqm.MessageQueue]ProcessQueueInfo `json:"mqTable"`
}

//Encode ConsumerRunningInfo to byte array
func (c *ConsumerRunningInfo) Encode() (jsonByte []byte, err error) {
	mqTableJsonStr := "{"
	first := true
	var keyJson []byte
	var valueJson []byte

	for key, value := range c.MqTable {
		keyJson, err = json.Marshal(key)
		if err != nil {
			return
		}
		valueJson, err = json.Marshal(value)
		if err != nil {
			return
		}
		if first == false {
			mqTableJsonStr = mqTableJsonStr + ","
		}
		mqTableJsonStr = mqTableJsonStr + string(keyJson) + ":" + string(valueJson)
		first = false
	}
	mqTableJsonStr = mqTableJsonStr + "}"
	var propertiesJson []byte
	propertiesJson, err = json.Marshal(c.Properties)
	if err != nil {
		return
	}
	jsonByte = c.formatEncode("\"properties\"", string(propertiesJson), "\"mqTable\"", string(mqTableJsonStr))
	return
}
func (c *ConsumerRunningInfo) formatEncode(kVList ...string) []byte {
	jsonStr := "{"
	first := true
	for i := 0; i+1 < len(kVList); i += 2 {
		if first == false {
			jsonStr = jsonStr + ","
		}
		keyJson := kVList[i]
		valueJson := kVList[i+1]

		jsonStr = jsonStr + string(keyJson) + ":" + string(valueJson)

		first = false
	}
	jsonStr = jsonStr + "}"
	return []byte(jsonStr)

}
