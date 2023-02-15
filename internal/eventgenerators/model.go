// Package eventgenerators holds all supported types of eventgenerator configurations such as MessageEventGenerator and MessageBurstGenerator.
package eventgenerators

type EventGenerator interface {
	String() string
}
