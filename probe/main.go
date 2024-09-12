package main

//go:generate wit-bindgen-go generate --world probe --out internal ../wit

import (
	"math/rand/v2"
	"time"

	"github.com/bytecodealliance/wasm-tools-go/cm"
	packetlooper "github.com/lxfontes/lwc/probe/internal/wasmcloud/lwc/packet-looper"
	"go.wasmcloud.dev/component"
	"go.wasmcloud.dev/component/lattice"
)

func probeReceive(packet uint64, ttl uint64, chaos uint16, route cm.Option[string]) {
	// Pick Route
	target_route := "default"
	if !route.None() {
		target_route = *route.Some()
	}

	// Compute arrival time
	srcTime := time.Unix(0, int64(packet))
	delta := time.Since(srcTime).Milliseconds()
	logger := component.ContextLogger("probe").With("packet", packet, "ttl", ttl, "route", target_route, "delta", delta)

	logger.Info("packet seen", "action", "seen")

	// Final destination
	if ttl == 0 {
		logger.Info("packet expired", "action", "expired")
		return
	}

	// Forward to next hop
	lattice.SetLinkName(target_route, cm.ToList([]lattice.CallTargetInterface{
		lattice.NewCallTargetInterface("wasmcloud", "lwc", "packet-looper"),
	}))
	packetlooper.Process(packet, ttl-1, chaos, route)

	// We introduce chaos by randomly panicking
	// This is to demonstrate how wasmCloud deals with faults in a high traffic scenario
	// We panic after the packet has been forwarded to the next hop so it doesn't interrupt the test
	rng := rand.IntN(100)
	if rng < int(chaos) {
		panic("chaos")
	}
}

func init() {
	packetlooper.Exports.Process = probeReceive
}

func main() {}
