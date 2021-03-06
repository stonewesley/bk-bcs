/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.,
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under,
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package lib

import (
	"strconv"
	"strings"

	"github.com/emicklei/go-restful"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-storage/storage/operator"
)

// GetQueryParamString get string from rest query parameter
func GetQueryParamString(req *restful.Request, key string) string {
	return req.QueryParameter(key)
}

// GetQueryParamStringArray get string array from restful query parameter
func GetQueryParamStringArray(req *restful.Request, key, sep string) []string {
	s := req.QueryParameter(key)
	if len(s) == 0 {
		return nil
	}
	fields := strings.Split(s, sep)
	return fields
}

// GetQueryParamInt get int from restful query parameter
func GetQueryParamInt(req *restful.Request, key string, defaultValue int) (int, error) {
	s := req.QueryParameter(key)
	if len(s) == 0 {
		return defaultValue, nil
	}
	return strconv.Atoi(s)
}

// GetQueryParamInt64 get int64 from restful query parameter
func GetQueryParamInt64(req *restful.Request, key string, defaultValue int64) (int64, error) {
	s := req.QueryParameter(key)
	if len(s) == 0 {
		return defaultValue, nil
	}
	return strconv.ParseInt(s, 10, 64)
}

// FormatTime format time
func FormatTime(data []operator.M, needTimeFormatList []string) {
	// Some time-field need to be format before return
	for i := range data {
		for _, t := range needTimeFormatList {
			tmp, ok := data[i][t].(primitive.DateTime)
			if !ok {
				continue
			}
			data[i][t] = tmp.Time()
		}
	}
	return
}
