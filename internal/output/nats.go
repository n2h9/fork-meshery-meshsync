package output

import (
	"github.com/layer5io/meshkit/broker"
	"github.com/n2h9/fork-meshery-meshsync/internal/config"
	"github.com/n2h9/fork-meshery-meshsync/pkg/model"
)

type NatsWriter struct {
	br broker.Handler
}

func NewNatsWriter(br broker.Handler) *NatsWriter {
	return &NatsWriter{
		br: br,
	}
}

func (s *NatsWriter) Write(
	obj model.KubernetesResource,
	evtype broker.EventType,
	config config.PipelineConfig,
) error {
	return s.br.Publish(
		config.PublishTo,
		&broker.Message{
			ObjectType: broker.MeshSync,
			EventType:  evtype,
			Object:     obj,
		},
	)
}
