package zerome

import (
	"testing"

	"github.com/clwluvw/zerome/internal/querier/promquerier"
	"github.com/clwluvw/zerome/internal/writer/promremotewriter"
	"github.com/stretchr/testify/require"
)

func TestSetQuerier(t *testing.T) {
	t.Parallel()

	m := Metric{}
	q := &promquerier.PromQuerier{}
	m.SetQuerier(q)
	require.Equal(t, q, m.querier)
}

func TestSetWriter(t *testing.T) {
	t.Parallel()

	m := Metric{}
	w := &promremotewriter.PromRemoteWrite{}
	m.SetWriter(w)
	require.Equal(t, w, m.writer)
}
