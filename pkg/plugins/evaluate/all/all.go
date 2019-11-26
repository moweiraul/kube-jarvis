package all

import (
	"github.com/RayHuangCN/kube-jarvis/pkg/plugins/evaluate"
	"github.com/RayHuangCN/kube-jarvis/pkg/plugins/evaluate/sum"
)

func init() {
	evaluate.Add(sum.EvaluatorType, evaluate.Factory{
		Creator: sum.NewEvaluator,
	})
}
