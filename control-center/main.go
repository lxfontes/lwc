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
	rawTtl := r.URL.Query().Get("ttl")
	if rawTtl == "" {
		rawTtl = "100"
	}
	ttl, err := strconv.ParseUint(rawTtl, 10, 64)
	if err != nil {
		http.Error(w, "invalid ttl", http.StatusBadRequest)
		return
	}

	rawChaos := r.URL.Query().Get("chaos")
	if rawChaos == "" {
		rawChaos = "0"
	}
	chaos, err := strconv.ParseUint(rawChaos, 10, 16)
	if err != nil {
		http.Error(w, "invalid chaos", http.StatusBadRequest)
		return
	}

	route := r.URL.Query().Get("route")
	if route == "" {
		route = "default"
	}

	packetId := time.Now().UnixNano()
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
