// Copyright 2019 The Terraformer Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package mikrotik

import (
	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/ddelnano/terraform-provider-mikrotik/client"
)

type MikrotikService struct {
	terraform_utils.Service
}

func (m *MikrotikService) generateClient() client.Mikrotik {
	return client.NewClient(
		m.Args["host"].(string),
		m.Args["user"].(string),
		m.Args["password"].(string),
	)
}