package extensions

import (
	"fmt"
)

// EventTracer interface allows setting extension for event context.
type EventTracer interface {
	SetExtension(k string, v interface{}) error
}

const (
	traceParentName = "traceparent"
	traceStateName  = "tracestate"
)

var attributes = make(map[string]interface{})

// TraceParent returns traceparent attribute value
func TraceParent() string {
	return fmt.Sprintf("%v", attributes[traceParentName])
}

// SetTraceParent sets traceparent attribute value.
func SetTraceParent(traceparent string) {
	attributes[traceParentName] = traceparent
}

// TraceState returns tracestate attribute value
func TraceState() string {
	return fmt.Sprintf("%v", attributes[traceStateName])
}

// SetTraceState sets tracestate attribute value.
func SetTraceState(tracestate string) {
	attributes[traceStateName] = tracestate
}

// AddTracingAttributes adds the tracing attributes traceparent and tracestate to the CloudEvents context.
func AddTracingAttributes(ec EventTracer) error {
	if _, ok := attributes[traceParentName]; ok {
		for k, v := range attributes {
			if err := ec.SetExtension(k, v); err != nil {
				return err
			}
		}
	}
	return nil
}
