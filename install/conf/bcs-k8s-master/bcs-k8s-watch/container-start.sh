#!/bin/bash

module="bcs-k8s-watch"

cd /data/bcs/${module}
chmod +x ${module}

#check configuration render
if [ $BCS_CONFIG_TYPE == "render" ] then
  cat ${module}.json.template | envsubst | tee ${module}.json
fi

#ready to start
/data/bcs/${module}/${module} $@
