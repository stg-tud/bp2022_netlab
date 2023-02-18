// Package eventgeneratortypes holds all supported types of eventgenerator configurations such as MessageEventGenerator and MessageBurstGenerator.
package eventgeneratortypes

type EventGeneratorType interface {
	String() string
}
