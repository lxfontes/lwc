#!/bin/bash
curl 'localhost:30000/?ttl=10000&route=special'
curl 'localhost:30000/?ttl=10000&route=default'

curl 'localhost:30000/?ttl=10000&route=special'
curl 'localhost:30000/?ttl=10000&route=default'

curl 'localhost:30000/?ttl=10000&route=special'
curl 'localhost:30000/?ttl=10000&route=default'
