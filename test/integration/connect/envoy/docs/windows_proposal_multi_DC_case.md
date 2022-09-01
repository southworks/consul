# Multi DC case

## Proof of Concept

A Proof of Concept (POC) was carried out, unifying the containers for the test cases where we have multiple Data Centers.

For this particular example, the "case-wanfed-gw" case is used.

### Modifications necessary for operation

The file: "global-setup.sh" of the test was modified so that the container that makes the certificates deposits them directly in a directory of the volume "envoy_workdir"
In addition, the Consul configuration files were modified so that it looks for the certificates in the directory of the volume where they were saved.

### Manual execution scripts for operation

#### Consul-primary

```shell
consul services register workdir\primary\register\service_gateway.hcl
 
consul connect envoy -bootstrap -proxy-id mesh-gateway -envoy-version v1.19.4 -http-addr envoy_consul-primary_1:8500 -grpc-addr envoy_consul-primary_1:8502 -admin-access-log-path C:\\log -admin-bind 0.0.0.0:19000 > C:\workdir\primary\envoy\mesh-gateway-bootstrap.json
 
envoy -c C:\workdir\primary\envoy\mesh-gateway-bootstrap.json -l trace --disable-hot-restart --drain-time-s 1
 
curl -s "http://consul-secondary:8500/v1/catalog/service/consul?dc=secondary" 
 
```

#### Consul-secondary

```shell
consul services register workdir\secondary\register\service_gateway.hcl
 
consul connect envoy -bootstrap -proxy-id mesh-gateway -envoy-version v1.19.4 -http-addr envoy_consul-secondary_1:8500 -grpc-addr envoy_consul-secondary_1:8502 -admin-access-log-path C:\\log -admin-bind 0.0.0.0:19001 > C:\workdir\secondary\envoy\mesh-gateway-bootstrap.json
 
envoy -c C:\workdir\secondary\envoy\mesh-gateway-bootstrap.json -l trace --disable-hot-restart --drain-time-s 1
 
curl -s "http://consul-primary:8500/v1/catalog/service/consul?dc=primary"
```

After the execution of the previous scripts, Consul would be in a green state, connected and ready to execute the tests.
