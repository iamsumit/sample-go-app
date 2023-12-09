package otel

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	api "go.opentelemetry.io/otel/metric"
)

// Records an increment to the counter with the given label values.
func (c Counter) Record(ctx context.Context, incr float64, labelValues ...string) {
	c.counter.Add(ctx, incr, api.WithAttributes(labeler(c.labels, labelValues)...))
}

// labeler takes a list of keys and a list of values and returns a list of attribute values of the same length as keys
// and each element having a Key equal to key[i] and Value as a string equal to v[i] or the empty string if v[i] is
// out of bounds.
func labeler(keys, values []string) []attribute.KeyValue {
	out := make([]attribute.KeyValue, 0, len(keys))

	for i, k := range keys {
		var v string
		if i < len(values) {
			v = values[i]
		}
		out = append(out, attribute.String(k, v))
	}
	return out
}
