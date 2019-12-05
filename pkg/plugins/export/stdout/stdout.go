package stdout

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/RayHuangCN/kube-jarvis/pkg/plugins/export"

	"github.com/RayHuangCN/kube-jarvis/pkg/plugins/diagnose"
	"github.com/fatih/color"
)

const (
	// ExporterType is type name of this Exporter
	ExporterType = "stdout"
)

// Exporter just print information to logger with a simple format
type Exporter struct {
	export.Collector
	*export.MetaData
}

// NewExporter return a stdout Exporter
func NewExporter(m *export.MetaData) export.Exporter {
	e := &Exporter{
		MetaData: m,
	}

	e.Format = "fmt"
	return e
}

// CoordinateBegin export information about coordinator Run begin
func (e *Exporter) CoordinateBegin(ctx context.Context) error {
	if e.Format != "fmt" {
		e.Collector.Output = []io.Writer{os.Stdout}
		return e.Collector.CoordinateBegin(ctx)
	}

	fmt.Println("===================================================================")
	fmt.Println("                       kube-jarivs                                 ")
	fmt.Println("===================================================================")
	return nil
}

// DiagnosticBegin export information about a Diagnostic begin
func (e *Exporter) DiagnosticBegin(ctx context.Context, dia diagnose.Diagnostic) error {
	if e.Format != "fmt" {
		return e.Collector.DiagnosticBegin(ctx, dia)
	}

	fmt.Println("Diagnostic report")
	fmt.Printf("    Type : %s\n", dia.Meta().Type)
	fmt.Printf("    Name : %s\n", dia.Meta().Name)
	fmt.Printf("- ----- results ----------------\n")
	return nil
}

// DiagnosticFinish export information about a Diagnostic finished
func (e *Exporter) DiagnosticFinish(ctx context.Context, dia diagnose.Diagnostic) error {
	if e.Format != "fmt" {
		return e.Collector.DiagnosticFinish(ctx, dia)
	}

	fmt.Println("===================================================================")
	return nil
}

// DiagnosticResult export information about one diagnose.Result
func (e *Exporter) DiagnosticResult(ctx context.Context, dia diagnose.Diagnostic, result *diagnose.Result) error {
	if e.Format != "fmt" {
		return e.Collector.DiagnosticResult(ctx, dia, result)
	}

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
		pt("    Describe : %s\n", result.Desc)
		pt("    Proposal : %s\n", result.Proposal)
	}
	fmt.Printf("- -----------------------------\n")
	return nil
}

// CoordinateFinish export information about coordinator Run finished
func (e *Exporter) CoordinateFinish(ctx context.Context) error {
	if e.Format != "fmt" {
		if err := e.Collector.CoordinateFinish(ctx); err != nil {
			return err
		}
	}
	fmt.Println("===================================================================")
	return nil
}
