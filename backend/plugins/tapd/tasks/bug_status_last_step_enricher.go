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
	"encoding/json"
	"github.com/apache/incubator-devlake/core/dal"
	"github.com/apache/incubator-devlake/core/errors"
	"github.com/apache/incubator-devlake/core/plugin"
	"github.com/apache/incubator-devlake/helpers/pluginhelper/api"
	"github.com/apache/incubator-devlake/plugins/tapd/models"
)

var _ plugin.SubTaskEntryPoint = EnrichBugStatusLastStep

var EnrichBugStatusLastStepMeta = plugin.SubTaskMeta{
	Name:             "enrichBugStatusLastStep",
	EntryPoint:       EnrichBugStatusLastStep,
	EnabledByDefault: true,
	Description:      "Enrich raw data into tool layer table _tool_tapd_bug_status",
	DomainTypes:      []string{plugin.DOMAIN_TYPE_TICKET},
}

func EnrichBugStatusLastStep(taskCtx plugin.SubTaskContext) errors.Error {
	db := taskCtx.GetDal()
	rawDataSubTaskArgs, data := CreateRawDataSubTaskArgs(taskCtx, RAW_BUG_STATUS_LAST_STEP_TABLE)
	extractor, err := api.NewApiExtractor(api.ApiExtractorArgs{
		RawDataSubTaskArgs: *rawDataSubTaskArgs,
		Extract: func(row *api.RawData) ([]interface{}, errors.Error) {
			var bugStatusLastStepRes struct {
				Data   interface{} `json:"data"`
				Status int         `json:"status"`
				Info   string      `json:"info"`
			}
			err := errors.Convert(json.Unmarshal(row.Data, &bugStatusLastStepRes))
			if err != nil {
				return nil, err
			}
			results := make([]interface{}, 0)
			statusList := make([]*models.TapdBugStatus, 0)
			clauses := []dal.Clause{
				dal.Where("connection_id = ? and workspace_id = ?", data.Options.ConnectionId, data.Options.WorkspaceId),
			}
			err = db.All(&statusList, clauses...)
			if err != nil {
				return nil, err
			}

			for _, status := range statusList {
				switch bugStatusLastStepResData := bugStatusLastStepRes.Data.(type) {
				case map[string]interface{}:
					if _, ok := bugStatusLastStepResData[status.EnglishName]; ok {
						status.IsLastStep = true
						results = append(results, status)
					}
				}
			}

			return results, nil
		},
	})

	if err != nil {
		return err
	}

	return extractor.Execute()
}
