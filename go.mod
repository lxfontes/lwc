module github.com/lxfontes/lwc

go 1.22.5

require (
	github.com/bytecodealliance/wasm-tools-go v0.2.0
	go.wasmcloud.dev/component v0.0.0-20240911170200-f90fa696ada6
)

require (
	github.com/samber/lo v1.47.0 // indirect
	github.com/samber/slog-common v0.17.1 // indirect
	golang.org/x/text v0.18.0 // indirect
	tinygo.org/x/drivers v0.28.0 // indirect
)

replace go.wasmcloud.dev/component => ../../wasmCloud/component-sdk-go/
