package capacity

import (
	"context"
	"fmt"
	"github.com/RayHuangCN/kube-jarvis/pkg/logger"
	"github.com/RayHuangCN/kube-jarvis/pkg/plugins"
	"github.com/RayHuangCN/kube-jarvis/pkg/plugins/cluster"
	"github.com/RayHuangCN/kube-jarvis/pkg/plugins/diagnose"
	"github.com/RayHuangCN/kube-jarvis/pkg/translate"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"testing"
)

func TestDiagnostic_diagnoseCapacity(t *testing.T) {
	var cases = []struct {
		Pass         bool
		Err          bool
		Capacities   []Capacity
		NodeTotal    int
		MasterCpu    resource.Quantity
		MasterMemory resource.Quantity
	}{
		// Err
		{
			Err: true,
			Capacities: []Capacity{
				{
					Memory:       resource.MustParse("2Gi"),
					CpuCore:      resource.MustParse("2000m"),
					MaxNodeTotal: 8,
				},
				{
					Memory:       resource.MustParse("2Gi"),
					CpuCore:      resource.MustParse("2000m"),
					MaxNodeTotal: 3,
				},
			},
			NodeTotal:    10,
			MasterCpu:    resource.MustParse("1000m"),
			MasterMemory: resource.MustParse("1Gi"),
		},
		// warn
		{
			Pass: false,
			Capacities: []Capacity{
				{
					Memory:       resource.MustParse("2Gi"),
					CpuCore:      resource.MustParse("2000m"),
					MaxNodeTotal: 8,
				},
				{
					Memory:       resource.MustParse("2Gi"),
					CpuCore:      resource.MustParse("2000m"),
					MaxNodeTotal: 3,
				},
			},
			NodeTotal:    4,
			MasterCpu:    resource.MustParse("1000m"),
			MasterMemory: resource.MustParse("1Gi"),
		},

		// good
		{
			Pass: true,
			Capacities: []Capacity{
				{
					Memory:       resource.MustParse("2Gi"),
					CpuCore:      resource.MustParse("2000m"),
					MaxNodeTotal: 8,
				},
				{
					Memory:       resource.MustParse("2Gi"),
					CpuCore:      resource.MustParse("2000m"),
					MaxNodeTotal: 3,
				},
			},
			NodeTotal:    4,
			MasterCpu:    resource.MustParse("3000m"),
			MasterMemory: resource.MustParse("3Gi"),
		},
	}

	for _, cs := range cases {
		t.Run(fmt.Sprintf("Pass=%v", cs.Pass), func(t *testing.T) {
			res := cluster.NewResources()
			// create master
			node := v1.Node{}
			node.Name = "master"
			node.Labels = map[string]string{
				"node-role.kubernetes.io/master": "true",
			}
			node.Status.Capacity = map[v1.ResourceName]resource.Quantity{
				v1.ResourceCPU:    cs.MasterCpu,
				v1.ResourceMemory: cs.MasterMemory,
			}
			res.Nodes = &v1.NodeList{Items: []v1.Node{node}}

			// create nodes
			for i := 0; i < cs.NodeTotal; i++ {
				node := v1.Node{}
				node.Name = fmt.Sprintf("node-%d", i)
				res.Nodes.Items = append(res.Nodes.Items, node)
			}

			// start diagnostic
			d := NewDiagnostic(&diagnose.MetaData{
				CommonMetaData: plugins.CommonMetaData{
					Translator: translate.NewFake(),
					Logger:     logger.NewLogger(),
					Type:       DiagnosticType,
					Name:       DiagnosticType,
				},
				Catalogue: diagnose.CatalogueMaster,
			}).(*Diagnostic)
			d.Capacities = cs.Capacities

			if err := d.Init(); err != nil {
				t.Fatalf(err.Error())
			}

			d.StartDiagnose(context.Background(), diagnose.StartDiagnoseParam{
				CloudType: "fake",
				Resources: res,
			})
			total := 0
			for {
				r, ok := <-d.result
				if !ok {
					break
				}
				total++
				if cs.Err {
					if r.Error == nil {
						t.Fatalf("want a err")
					}
					return
				}
				if cs.Pass && r.Level != diagnose.HealthyLevelGood {
					t.Fatalf("should return HealthyLevelGood")
				}

				if !cs.Pass && r.Level == diagnose.HealthyLevelGood {
					t.Fatalf("should return not HealthyLevelGood")
				}
			}

			if total != 2 {
				t.Fatalf("should return 2 results but get %d ", total)
			}
		})
	}
}