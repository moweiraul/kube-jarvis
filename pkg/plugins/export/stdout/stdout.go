package stdout

import (
	"context"
	"fmt"

	"github.com/RayHuangCN/kube-jarvis/pkg/plugins/export"

	"github.com/RayHuangCN/kube-jarvis/pkg/plugins/diagnose"
	"github.com/RayHuangCN/kube-jarvis/pkg/plugins/evaluate"
	"github.com/fatih/color"
)

const (
	// ExporterType is type name of this Exporter
	ExporterType = "stdout"
)

// Exporter just print information to logger with a simple format
type Exporter struct {
	*export.MetaData
}

// NewExporter return a stdout Exporter
func NewExporter(m *export.MetaData) export.Exporter {
	return &Exporter{
		MetaData: m,
	}
}

// CoordinateBegin export information about coordinator Run begin
func (e *Exporter) CoordinateBegin(ctx context.Context) error {
	fmt.Println("===================================================================")
	fmt.Println("                       kube-jarivs                                 ")
	fmt.Println("===================================================================")
	return nil
}

// DiagnosticBegin export information about a Diagnostic begin
func (e *Exporter) DiagnosticBegin(ctx context.Context, dia diagnose.Diagnostic) error {
	fmt.Println("Diagnostic report")
	fmt.Printf("    Type : %s\n", dia.Meta().Type)
	fmt.Printf("    Name : %s\n", dia.Meta().Name)
	fmt.Printf("- ----- results ----------------\n")
	return nil
}

// DiagnosticFinish export information about a Diagnostic finished
func (e *Exporter) DiagnosticFinish(ctx context.Context, dia diagnose.Diagnostic) error {
	fmt.Printf("Diagnostic Score : %.2f/%.2f\n", dia.Meta().Score, dia.Meta().TotalScore)
	fmt.Println("===================================================================")
	return nil
}

// DiagnosticResult export information about one diagnose.Result
func (e *Exporter) DiagnosticResult(ctx context.Context, dia diagnose.Diagnostic, result *diagnose.Result) error {
	if result.Error != nil {
		color.HiRed("[!!ERROR] %s\n", result.Error.Error())
	} else {
		var pt func(format string, a ...interface{})
		switch result.Level {
		case diagnose.HealthyLevelGood:
			pt = color.Green
		case diagnose.HealthyLevelWarn:
			pt = color.Yellow
		case diagnose.HealthyLevelRisk:
			pt = color.Red
		case diagnose.HealthyLevelSerious:
			pt = color.HiRed
		default:
			pt = func(format string, a ...interface{}) {
				fmt.Printf(format, a...)
			}
		}
		pt("[%s] %s -> %s\n", result.Level, result.Title, result.ObjName)
		pt("    Score : -%.2f\n", result.Score)
		pt("    Describe : %s\n", result.Desc)
		pt("    Proposal : %s\n", result.Proposal)
	}
	fmt.Printf("- -----------------------------\n")
	return nil
}

// EvaluationBegin export information about a Evaluator begin
func (e *Exporter) EvaluationBegin(ctx context.Context, eva evaluate.Evaluator) error {
	fmt.Println("Evaluation report")
	fmt.Printf("    Type : %s\n", eva.Meta().Type)
	fmt.Printf("    Name : %s\n", eva.Meta().Name)
	fmt.Printf("- ----- result -----------------\n")
	return nil
}

// EvaluationFinish export information about a Evaluator finish
func (e *Exporter) EvaluationFinish(ctx context.Context, eva evaluate.Evaluator) error {
	fmt.Println("===================================================================")
	return nil
}

// EvaluationResult export information about a Evaluator result
func (e *Exporter) EvaluationResult(ctx context.Context, eva evaluate.Evaluator, result *evaluate.Result) error {
	fmt.Printf("[%s]\n", result.Name)
	fmt.Printf("    Describe : %s\n", result.Desc)
	return nil
}

// CoordinateFinish export information about coordinator Run finished
func (e *Exporter) CoordinateFinish(ctx context.Context) error {
	fmt.Println("===================================================================")
	return nil
}
