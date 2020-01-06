/*
* Tencent is pleased to support the open source community by making TKEStack
* available.
*
* Copyright (C) 2012-2019 Tencent. All Rights Reserved.
*
* Licensed under the Apache License, Version 2.0 (the “License”); you may not use
* this file except in compliance with the License. You may obtain a copy of the
* License at
*
* https://opensource.org/licenses/Apache-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an “AS IS” BASIS, WITHOUT
* WARRANTIES OF ANY KIND, either express or implied.  See the License for the
* specific language governing permissions and limitations under the License.
 */
package all

import (
	"tkestack.io/kube-jarvis/pkg/plugins/diagnose"
	"tkestack.io/kube-jarvis/pkg/plugins/diagnose/master/capacity"
	"tkestack.io/kube-jarvis/pkg/plugins/diagnose/master/components"
	"tkestack.io/kube-jarvis/pkg/plugins/diagnose/node/ha"
	"tkestack.io/kube-jarvis/pkg/plugins/diagnose/node/status"
	"tkestack.io/kube-jarvis/pkg/plugins/diagnose/node/iptables"
	"tkestack.io/kube-jarvis/pkg/plugins/diagnose/node/sys"
	"tkestack.io/kube-jarvis/pkg/plugins/diagnose/other/example"
	hpaip "tkestack.io/kube-jarvis/pkg/plugins/diagnose/resource/hpa/ip"
	"tkestack.io/kube-jarvis/pkg/plugins/diagnose/resource/workload/affinity"
	"tkestack.io/kube-jarvis/pkg/plugins/diagnose/resource/workload/batch"
	"tkestack.io/kube-jarvis/pkg/plugins/diagnose/resource/workload/healthcheck"
	"tkestack.io/kube-jarvis/pkg/plugins/diagnose/resource/workload/pdb"
	"tkestack.io/kube-jarvis/pkg/plugins/diagnose/resource/workload/requestslimits"
	workloadStatus "tkestack.io/kube-jarvis/pkg/plugins/diagnose/resource/workload/status"
)

func init() {
	addMasterDiagnostics()
	addResourceDiagnostics()
	addOtherDiagnostics()
	addNodeDiagnostics()
	addNodeStatusDiagnostics()
	addWorkloadDiagnostics()
}

func addMasterDiagnostics() {
	diagnose.Add(capacity.DiagnosticType, diagnose.Factory{
		Creator:   capacity.NewDiagnostic,
		Catalogue: diagnose.CatalogueMaster,
	})

	diagnose.Add(components.DiagnosticType, diagnose.Factory{
		Creator:   components.NewDiagnostic,
		Catalogue: diagnose.CatalogueMaster,
	})
}

func addResourceDiagnostics() {
	diagnose.Add(requestslimits.DiagnosticType, diagnose.Factory{
		Creator:   requestslimits.NewDiagnostic,
		Catalogue: diagnose.CatalogueResource,
	})

	diagnose.Add(hpaip.DiagnosticType, diagnose.Factory{
		Creator:   hpaip.NewDiagnostic,
		Catalogue: diagnose.CatalogueResource,
	})

	diagnose.Add(healthcheck.DiagnosticType, diagnose.Factory{
		Creator:   healthcheck.NewDiagnostic,
		Catalogue: diagnose.CatalogueResource,
	})

	diagnose.Add(affinity.DiagnosticType, diagnose.Factory{
		Creator:   affinity.NewDiagnostic,
		Catalogue: diagnose.CatalogueResource,
	})

	diagnose.Add(pdb.DiagnosticType, diagnose.Factory{
		Creator:   pdb.NewDiagnostic,
		Catalogue: diagnose.CatalogueResource,
	})

	diagnose.Add(batch.DiagnosticType, diagnose.Factory{
		Creator:   batch.NewDiagnostic,
		Catalogue: diagnose.CatalogueResource,
	})
}

func addOtherDiagnostics() {
	diagnose.Add(example.DiagnosticType, diagnose.Factory{
		Creator:   example.NewDiagnostic,
		Catalogue: diagnose.CatalogueOther,
	})
}

func addNodeDiagnostics() {
	diagnose.Add(sys.DiagnosticType, diagnose.Factory{
		Creator:   sys.NewDiagnostic,
		Catalogue: diagnose.CatalogueNode,
	})
	diagnose.Add(iptables.DiagnosticType, diagnose.Factory{
		Creator:   iptables.NewDiagnostic,
		Catalogue: diagnose.CatalogueNode,
	})
	diagnose.Add(ha.DiagnosticType, diagnose.Factory{
		Creator:   ha.NewDiagnostic,
		Catalogue: diagnose.CatalogueNode,
	})
}

func addNodeStatusDiagnostics() {
	diagnose.Add(status.DiagnosticType, diagnose.Factory{
		Creator:   status.NewDiagnostic,
		Catalogue: diagnose.CatalogueNode,
	})
}

func addWorkloadDiagnostics() {
	diagnose.Add(workloadStatus.DiagnosticType, diagnose.Factory{
		Creator:   workloadStatus.NewDiagnostic,
		Catalogue: diagnose.CatalogueResource,
	})
}
