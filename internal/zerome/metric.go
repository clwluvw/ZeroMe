package zerome

import (
	"time"

	"github.com/clwluvw/zerome/internal/querier"
	"github.com/clwluvw/zerome/internal/writer"
)

type Metric struct {
	Name     string        `yaml:"name"`
	Interval time.Duration `yaml:"interval"`
	Querier  string        `yaml:"querier"`
	Writer   string        `yaml:"writer"`
	UpLabels []string      `yaml:"up_labels"`
	querier  querier.Querier
	writer   writer.Writer
}

func (m *Metric) SetQuerier(q querier.Querier) {
	m.querier = q
}

func (m *Metric) SetWriter(w writer.Writer) {
	m.writer = w
}
