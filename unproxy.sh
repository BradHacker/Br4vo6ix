#!/bin/bash

if [ ! -d "out/scripts" ]; then
  mkdir out/scripts
fi

for ip in ${REDTEAM_IPS//,/ }
do
  if [ -f "out/scripts/$ip-unproxy.sh" ]; then
    rm "out/scripts/$ip-unproxy.sh"
  fi

  touch "out/scripts/$ip-unproxy.sh"

  echo -e "for p in \$(ps -aux | grep socat | awk '{ print \$2 }')\n\
do\n\
  kill -9 \$p\n\
done" >> "out/scripts/$ip-unproxy.sh"
done
  