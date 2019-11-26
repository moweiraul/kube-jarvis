package sum

import (
	"context"
	"fmt"

	"github.com/RayHuangCN/kube-jarvis/pkg/plugins/evaluate"

	"github.com/RayHuangCN/kube-jarvis/pkg/plugins/diagnose"
)

const (
	// EvaluatorType is type name of this Evaluator
	EvaluatorType = "sum"
)

// Evaluator sum all diagnostic result score with different healthy level
type Evaluator struct {
	*evaluate.MetaData
	TotalScore   float64
	Score        float64
	WarnScore    float64
	SeriousScore float64
	RiskScore    float64
	ErrorTotal   int
}

// NewEvaluator return a sum Evaluator
func NewEvaluator(param *evaluate.MetaData) evaluate.Evaluator {
	return &Evaluator{
		MetaData: param,
	}
}

// EvaDiagnosticResult evaluate one diagnostic result
func (e *Evaluator) EvaDiagnosticResult(ctx context.Context, dia diagnose.Diagnostic, result *diagnose.Result) error {
	if result.Error != nil {
		e.ErrorTotal++
	} else {
		switch result.Level {
		case diagnose.HealthyLevelWarn:
			e.WarnScore += result.Score
		case diagnose.HealthyLevelRisk:
			e.RiskScore += result.Score
		case diagnose.HealthyLevelSerious:
			e.SeriousScore += result.Score
		}
	}
	return nil
}

// EvaDiagnostic evaluate one diagnostic finish
func (e *Evaluator) EvaDiagnostic(ctx context.Context, dia diagnose.Diagnostic) error {
	e.TotalScore += dia.Meta().TotalScore
	e.Score += dia.Meta().Score
	return nil
}

// Result return a final evaluation result
func (e *Evaluator) Result() *evaluate.Result {
	return &evaluate.Result{
		Name: "score statistics",
		Desc: e.Translator.Message("result", map[string]interface{}{
			"Score":        fmt.Sprintf("%.2f/%.2f", e.Score, e.TotalScore),
			"WarnScore":    fmt.Sprintf("%.2f/%.2f", e.WarnScore, e.TotalScore),
			"RiskScore":    fmt.Sprintf("%.2f/%.2f", e.RiskScore, e.TotalScore),
			"SeriousScore": fmt.Sprintf("%.2f/%.2f", e.SeriousScore, e.TotalScore),
			"ErrorTotal":   e.ErrorTotal,
		}),
	}
}
