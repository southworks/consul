#!/bin/bash

# initialize the outputs for each dc
for dc in primary secondary; do
    rm -rf "workdir/${dc}/tls"
    mkdir -p "workdir/${dc}/tls"
done

container="consul-envoy-integ-tls-init--${CASE_NAME}"

scriptlet="
cd C:\\workdir ;
mkdir out ;
cd out ;
consul tls ca create ;
consul tls cert create -dc=primary -server -node=pri ;
consul tls cert create -dc=secondary -server -node=sec ;
"

docker.exe rm -f "$container" &>/dev/null || true
docker.exe run -i --net=none --name="$container" -v envoy_workdir_new:C:\\workdir windows/consul-dev sh -c "${scriptlet}"

# Private keys have 600 perms but tests are run as another user
chmod 666 workdir/primary/tls/primary-server-consul-0-key.pem
chmod 666 workdir/secondary/tls/secondary-server-consul-0-key.pem

docker.exe rm -f "$container" >/dev/null || true
