// Package networktypes holds all supported types of network configurations such as Wireless LAN and Switch.
package networktypes

type NetworkType interface {
	String() string
}
