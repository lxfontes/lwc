// Code generated by wit-bindgen-go. DO NOT EDIT.

package packetlooper

import (
	"github.com/bytecodealliance/wasm-tools-go/cm"
)

// Exports represents the caller-defined exports from "wasmcloud:lwc/packet-looper".
var Exports struct {
	// Process represents the caller-defined, exported function "process".
	//
	//	process: func(packet: u64, ttl: u64, chaos: u16, route: option<string>)
	Process func(packet uint64, ttl uint64, chaos uint16, route cm.Option[string])
}
