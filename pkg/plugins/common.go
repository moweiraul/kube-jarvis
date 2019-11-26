package plugins

import (
	"github.com/RayHuangCN/kube-jarvis/pkg/logger"
	"github.com/RayHuangCN/kube-jarvis/pkg/translate"
	"k8s.io/client-go/kubernetes"
)

// CommonMetaData is the common attributes of a plugins
type CommonMetaData struct {
	// Cli is the clientSets of target cluster
	Cli kubernetes.Interface
	// Translator is a translator with diagnostic module context
	Translator translate.Translator
	// Logger is a logger with diagnostic module context
	Logger logger.Logger
	// Type is the type of Diagnostic
	Type string
	// Title is the custom name of Diagnostic
	Name string
	// CloudType indicate the type of cloud provider
	CloudType string
}

// IsSupportedCloud return true if cloud type is supported
func IsSupportedCloud(supported []string, cloud string) bool {
	if len(supported) == 0 {
		return true
	}

	for _, c := range supported {
		if c == cloud {
			return true
		}
	}
	return false
}
