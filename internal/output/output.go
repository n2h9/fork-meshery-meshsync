package output

import (
	"github.com/layer5io/meshkit/broker"
	"github.com/n2h9/fork-meshery-meshsync/internal/config"
	"github.com/n2h9/fork-meshery-meshsync/pkg/model"
)

type Writer interface {
	Write(
		obj model.KubernetesResource,
		evtype broker.EventType,
		config config.PipelineConfig,
	) error
}
