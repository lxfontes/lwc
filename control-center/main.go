package main

//go:generate wit-bindgen-go generate --world control-center --out internal ../wit

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/bytecodealliance/wasm-tools-go/cm"
	packetlooper "github.com/lxfontes/lwc/control-center/internal/wasmcloud/lwc/packet-looper"
	"go.wasmcloud.dev/component/net/wasihttp"
)

func init() {
	wasihttp.Handle(http.HandlerFunc(blastHandler))
}

type blastResponse struct {
	Packet uint64 `json:"packet"`
	Route  string `json:"route"`
	Ttl    uint64 `json:"ttl"`
	Chaos  uint16 `json:"chaos"`
}

func blastHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "invalid query params", http.StatusBadRequest)
		return
	}

	// Determine route
	route := "default"
	if v := r.FormValue("route"); v != "" {
		route = v
	}

	// Determine ttl ( time to live )
	rawTtl := "100"
	if v := r.FormValue("ttl"); v != "" {
		rawTtl = v
	}
	ttl, err := strconv.ParseUint(rawTtl, 10, 64)
	if err != nil {
		http.Error(w, "invalid ttl", http.StatusBadRequest)
		return
	}

	// Determine chaos
	rawChaos := "0"
	if v := r.FormValue("chaos"); v != "" {
		rawChaos = v
	}
	chaos, err := strconv.ParseUint(rawChaos, 10, 16)
	if err != nil {
		http.Error(w, "invalid chaos", http.StatusBadRequest)
		return
	}

	// Generate a unique packet identifier
	packetId := time.Now().UnixNano()

	// Inject packet into the system
	packetlooper.Process(uint64(packetId), ttl, uint16(chaos), cm.Some(route))

	resp := blastResponse{
		Packet: uint64(packetId),
		Route:  route,
		Ttl:    ttl,
		Chaos:  uint16(chaos),
	}

	json.NewEncoder(w).Encode(&resp)
}

func main() {}
