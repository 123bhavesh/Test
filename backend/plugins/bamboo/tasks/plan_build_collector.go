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

package tasks

import (
	"fmt"

	"github.com/apache/incubator-devlake/core/errors"
	"github.com/apache/incubator-devlake/core/plugin"
	helper "github.com/apache/incubator-devlake/helpers/pluginhelper/api"
)

const RAW_PLAN_BUILD_TABLE = "bamboo_api_plan_builds"

var _ plugin.SubTaskEntryPoint = CollectPlanBuild

func CollectPlanBuild(taskCtx plugin.SubTaskContext) errors.Error {
	rawDataSubTaskArgs, data := CreateRawDataSubTaskArgs(taskCtx, RAW_PLAN_BUILD_TABLE)
	collector, err := helper.NewApiCollector(helper.ApiCollectorArgs{
		RawDataSubTaskArgs: *rawDataSubTaskArgs,
		ApiClient:          data.ApiClient,
		PageSize:           100,
		UrlTemplate:        fmt.Sprintf("result/%s.json", data.Options.PlanKey),
		Query:              QueryForResult,
		GetTotalPages:      GetTotalPagesFromResult,
		ResponseParser:     GetResultsResult,
	})
	if err != nil {
		return err
	}
	return covertError(collector.Execute())
}

var CollectPlanBuildMeta = plugin.SubTaskMeta{
	Name:             "CollectPlanBuild",
	EntryPoint:       CollectPlanBuild,
	EnabledByDefault: true,
	Description:      "Collect PlanBuild data from Bamboo api",
	DomainTypes:      []string{plugin.DOMAIN_TYPE_CICD},
}
