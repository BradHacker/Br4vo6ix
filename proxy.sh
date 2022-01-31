#!/bin/bash

if [ ! -d "out/scripts" ]; then
  mkdir out/scripts
fi

for ip in ${REDTEAM_IPS//,/ }
do
  if [ -f "out/scripts/$ip-proxy.sh" ]; then
    rm "out/scripts/$ip-proxy.sh"
  fi

  touch "out/scripts/$ip-proxy.sh"

  for p in ${PORTS//,/ }
  do
    echo "socat TCP4-LISTEN:$p,fork TCP:$PROXY_IP:4444 &" >> "out/scripts/$ip-proxy.sh"
  done
done
  