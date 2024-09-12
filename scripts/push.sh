#!/bin/bash
(cd probe && wash build)
(cd control-center && wash build)
wash push --insecure --allow-latest 127.0.0.1:5001/lwc_probe:latest probe/build/lwc_probe_s.wasm
wash push --insecure --allow-latest 127.0.0.1:5001/lwc_control_center:latest control-center/build/lwc_control_center_s.wasm
