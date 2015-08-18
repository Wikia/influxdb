package datastore

import (
	"github.com/golang/protobuf/proto"
	"github.com/alecthomas/log4go"

	"github.com/Wikia/influxdb/engine"
	"github.com/Wikia/influxdb/protocol"
)

func yieldToProcessor(s *protocol.Series, p engine.Processor, aliases []string) (bool, error) {
	for _, alias := range aliases {
		series := &protocol.Series{
			Name:   proto.String(alias),
			Fields: s.Fields,
			Points: s.Points,
		}
		log4go.Debug("Yielding to %s %s", p.Name(), series)
		if ok, err := p.Yield(series); !ok || err != nil {
			return ok, err
		}
	}
	return true, nil
}
