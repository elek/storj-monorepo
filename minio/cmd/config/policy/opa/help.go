/*
 * MinIO Cloud Storage, (C) 2019,2020 MinIO, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package opa

import "storj.io/minio/cmd/config"

// Help template for OPA policy feature.
var (
	Help = config.HelpKVS{
		config.HelpKV{
			Key:         URL,
			Description: `[DEPRECATED] OPA HTTP(s) endpoint e.g. "http://localhost:8181/v1/data/httpapi/authz/allow"`,
			Type:        "url",
		},
		config.HelpKV{
			Key:         AuthToken,
			Description: "[DEPRECATED] authorization token for OPA endpoint",
			Optional:    true,
			Type:        "string",
		},
		config.HelpKV{
			Key:         config.Comment,
			Description: config.DefaultComment,
			Optional:    true,
			Type:        "sentence",
		},
	}
)
