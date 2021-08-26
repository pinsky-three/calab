#!/bin/bash

docker run --rm \
  -p 7880:7880 \
  -p 7881:7881 \
  -p 7882:7882/udp \
  -v "$PWD"/config.yaml:/var/config/config.yaml \
  -e LIVEKIT_KEYS="APIwLeah7g4fuLYDYAJeaKsSE: 8nTlwISkb-63DPP7OH4e.nw.J44JjicvZDiz8J59EoQ+" \
  livekit/livekit-server \
  --dev \
  --config=/var/config/config.yaml \
  --node-ip=127.0.0.1