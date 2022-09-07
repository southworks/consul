#!/usr/bin/env bash
mkdir output

if [ $1 == 0 ]
then

    go test -v -timeout=30m -tags integration ./test/integration/connect/envoy -run="TestEnvoy/case-badauthz" -win=true > output/case-badauthz.txt
    echo "completado 33%"
    go test -v -timeout=30m -tags integration ./test/integration/connect/envoy -run="TestEnvoy/case-basic" -win=true > output/case-basic.txt
    echo "completado 66%"
    go test -v -timeout=30m -tags integration ./test/integration/connect/envoy -run="TestEnvoy/case-centralconf" -win=true > output/case-centralconf.txt
    echo "completado 100%"

elif [ $1 == 1 ]
then
    go test -v -timeout=30m -tags integration ./test/integration/connect/envoy -run="TestEnvoy/case-consul-exec" -win=true > output/case-consul-exec.txt
    echo "completado 20%"
    go test -v -timeout=30m -tags integration ./test/integration/connect/envoy -run="TestEnvoy/case-expose-checks" -win=true > output/case-expose-checks.txt
    echo "completado 40%"
    go test -v -timeout=30m -tags integration ./test/integration/connect/envoy -run="TestEnvoy/case-I7-intentions" -win=true > output/case-I7-intentions.txt
    echo "completado 60%"
    go test -v -timeout=30m -tags integration ./test/integration/connect/envoy -run="TestEnvoy/case-prometheus" -win=true > output/case-prometheus.txt
    echo "completado 80%"
    go test -v -timeout=30m -tags integration ./test/integration/connect/envoy -run="TestEnvoy/case-upstream-config" -win=true > output/case-upstream-config.txt
    echo "completado 100%"

elif [ $1 == 2 ]
then
    go test -v -timeout=30m -tags integration ./test/integration/connect/envoy -run="TestEnvoy/case-zipkin" -win=true > output/case-zipkin.txt
    echo "completado 10%"
    go test -v -timeout=30m -tags integration ./test/integration/connect/envoy -run="TestEnvoy/case-cfg-resolver-defaultsubset" -win=true > output/case-cfg-resolver-defaultsubset.txt
    echo "completado 20%"
    go test -v -timeout=30m -tags integration ./test/integration/connect/envoy -run="TestEnvoy/case-cfg-resolver-features" -win=true > output/case-cfg-resolver-features.txt
    echo "completado 30%"
    go test -v -timeout=30m -tags integration ./test/integration/connect/envoy -run="TestEnvoy/case-cfg-resolver-subset-onlypassing" -win=true > output/case-cfg-resolver-subset-onlypassing.txt
    echo "completado 40%"
    go test -v -timeout=30m -tags integration ./test/integration/connect/envoy -run="TestEnvoy/case-cfg-resolver-subset-redirect" -win=true > output/case-cfg-resolver-subset-redirect.txt
    echo "completado 50%"
    go test -v -timeout=30m -tags integration ./test/integration/connect/envoy -run="TestEnvoy/case-cfg-resolver-svc-failover" -win=true > output/case-cfg-resolver-svc-failover.txt
    echo "completado 60%"
    go test -v -timeout=30m -tags integration ./test/integration/connect/envoy -run="TestEnvoy/case-cfg-resolver-svc-redirect-http" -win=true > output/case-cfg-resolver-svc-redirect-http.txt
    echo "completado 70%"
    go test -v -timeout=30m -tags integration ./test/integration/connect/envoy -run="TestEnvoy/case-cfg-resolver-svc-redirect-tcp" -win=true > output/case-cfg-resolver-svc-redirect-tcp.txt
    echo "completado 80%"
    go test -v -timeout=30m -tags integration ./test/integration/connect/envoy -run="TestEnvoy/case-cfg-router-features" -win=true > output/case-cfg-router-features.txt
    echo "completado 90%"
    go test -v -timeout=30m -tags integration ./test/integration/connect/envoy -run="TestEnvoy/case-cfg-splitter-features" -win=true > output/case-cfg-splitter-features.txt
    echo "completado 100%"

elif [ $1 == 3 ]
then
    go test -v -timeout=30m -tags integration ./test/integration/connect/envoy -run="TestEnvoy/case-ingress-gateway-grpc" -win=true > output/case-ingress-gateway-grpc.txt
    echo "completado 16%"
    go test -v -timeout=30m -tags integration ./test/integration/connect/envoy -run="TestEnvoy/case-ingress-gateway-http" -win=true > output/case-ingress-gateway-http.txt
    echo "completado 32%"
    go test -v -timeout=30m -tags integration ./test/integration/connect/envoy -run="TestEnvoy/case-ingress-gateway-multiple-services" -win=true > output/case-ingress-gateway-multiple-services.txt
    echo "completado 48%"
    go test -v -timeout=30m -tags integration ./test/integration/connect/envoy -run="TestEnvoy/case-ingress-gateway-sds" -win=true > output/case-ingress-gateway-sds.txt
    echo "completado 66%"
    go test -v -timeout=30m -tags integration ./test/integration/connect/envoy -run="TestEnvoy/case-ingress-gateway-simple" -win=true > output/case-ingress-gateway-simple.txt
    echo "completado 83%"
    go test -v -timeout=30m -tags integration ./test/integration/connect/envoy -run="TestEnvoy/case-ingress-gateway-tls" -win=true > output/case-ingress-gateway-tls.txt
    echo "completado 100%"

elif [ $1 == 4 ]
then
    go test -v -timeout=30m -tags integration ./test/integration/connect/envoy -run="TestEnvoy/case-gateway-without-services" -win=true > output/case-gateway-without-services.txt
    echo "completado 20%"
    go test -v -timeout=30m -tags integration ./test/integration/connect/envoy -run="TestEnvoy/case-terminating-gateway-hostnames" -win=true > output/case-terminating-gateway-hostnames.txt
    echo "completado 40%"
    go test -v -timeout=30m -tags integration ./test/integration/connect/envoy -run="TestEnvoy/case-terminating-gateway-simple" -win=true > output/case-terminating-gateway-simple.txt
    echo "completado 60%"
    go test -v -timeout=30m -tags integration ./test/integration/connect/envoy -run="TestEnvoy/case-terminating-gateway-subsets" -win=true > output/case-terminating-gateway-subsets.txt
    echo "completado 80%"
    go test -v -timeout=30m -tags integration ./test/integration/connect/envoy -run="TestEnvoy/case-terminating-gateway-without-services" -win=true > output/case-terminating-gateway-without-services.txt
    echo "completado 100%"

elif [ $1 == 5 ]
then
    go test -v -timeout=30m -tags integration ./test/integration/connect/envoy -run="TestEnvoy/case-dogstatsd-udp" -win=true > output/case-dogstatsd-udp.txt
    echo "completado 16%"
    go test -v -timeout=30m -tags integration ./test/integration/connect/envoy -run="TestEnvoy/case-grpc" -win=true > output/case-grpc.txt
    echo "completado 32%"
    go test -v -timeout=30m -tags integration ./test/integration/connect/envoy -run="TestEnvoy/case-http" -win=true > output/case-http.txt
    echo "completado 48%"
    go test -v -timeout=30m -tags integration ./test/integration/connect/envoy -run="TestEnvoy/case-http-badauthz" -win=true > output/case-http-badauthz.txt
    echo "completado 66%"
    go test -v -timeout=30m -tags integration ./test/integration/connect/envoy -run="TestEnvoy/case-statsd-udp" -win=true > output/case-statsd-udp.txt
    echo "completado 83%"
    go test -v -timeout=30m -tags integration ./test/integration/connect/envoy -run="TestEnvoy/case-stats-proxy" -win=true > output/case-stats-proxy.txt
    echo "completado 100%"
else
    go test -v -timeout=30m -tags integration ./test/integration/connect/envoy -run="TestEnvoy/case-badauthz" -win=true > output/case-badauthz.txt
    go test -v -timeout=30m -tags integration ./test/integration/connect/envoy -run="TestEnvoy/case-basic" -win=true > output/case-basic.txt
    go test -v -timeout=30m -tags integration ./test/integration/connect/envoy -run="TestEnvoy/case-centralconf" -win=true > output/case-centralconf.txt
    go test -v -timeout=30m -tags integration ./test/integration/connect/envoy -run="TestEnvoy/case-consul-exec" -win=true > output/case-consul-exec.txt
    go test -v -timeout=30m -tags integration ./test/integration/connect/envoy -run="TestEnvoy/case-expose-checks" -win=true > output/case-expose-checks.txt
    go test -v -timeout=30m -tags integration ./test/integration/connect/envoy -run="TestEnvoy/case-l7-intentions" -win=true > output/case-l7-intentions.txt
    go test -v -timeout=30m -tags integration ./test/integration/connect/envoy -run="TestEnvoy/case-prometheus" -win=true > output/case-prometheus.txt
    go test -v -timeout=30m -tags integration ./test/integration/connect/envoy -run="TestEnvoy/case-upstream-config" -win=true > output/case-upstream-config.txt
    go test -v -timeout=30m -tags integration ./test/integration/connect/envoy -run="TestEnvoy/case-zipkin" -win=true > output/case-zipkin.txt
    go test -v -timeout=30m -tags integration ./test/integration/connect/envoy -run="TestEnvoy/case-cfg-resolver-defaultsubset" -win=true > output/case-cfg-resolver-defaultsubset.txt
    go test -v -timeout=30m -tags integration ./test/integration/connect/envoy -run="TestEnvoy/case-cfg-resolver-features" -win=true > output/case-cfg-resolver-features.txt
    go test -v -timeout=30m -tags integration ./test/integration/connect/envoy -run="TestEnvoy/case-cfg-resolver-subset-onlypassing" -win=true > output/case-cfg-resolver-subset-onlypassing.txt
    go test -v -timeout=30m -tags integration ./test/integration/connect/envoy -run="TestEnvoy/case-cfg-resolver-subset-redirect" -win=true > output/case-cfg-resolver-subset-redirect.txt
    go test -v -timeout=30m -tags integration ./test/integration/connect/envoy -run="TestEnvoy/case-cfg-resolver-svc-failover" -win=true > output/case-cfg-resolver-svc-failover.txt
    go test -v -timeout=30m -tags integration ./test/integration/connect/envoy -run="TestEnvoy/case-cfg-resolver-svc-redirect-http" -win=true > output/case-cfg-resolver-svc-redirect-http.txt
    go test -v -timeout=30m -tags integration ./test/integration/connect/envoy -run="TestEnvoy/case-cfg-resolver-svc-redirect-tcp" -win=true > output/case-cfg-resolver-svc-redirect-tcp.txt
    go test -v -timeout=30m -tags integration ./test/integration/connect/envoy -run="TestEnvoy/case-cfg-router-features" -win=true > output/case-cfg-router-features.txt
    go test -v -timeout=30m -tags integration ./test/integration/connect/envoy -run="TestEnvoy/case-cfg-splitter-features" -win=true > output/case-cfg-splitter-features.txt
    go test -v -timeout=30m -tags integration ./test/integration/connect/envoy -run="TestEnvoy/case-ingress-gateway-grpc" -win=true > output/case-ingress-gateway-grpc.txt
    go test -v -timeout=30m -tags integration ./test/integration/connect/envoy -run="TestEnvoy/case-ingress-gateway-http" -win=true > output/case-ingress-gateway-http.txt
    go test -v -timeout=30m -tags integration ./test/integration/connect/envoy -run="TestEnvoy/case-ingress-gateway-multiple-services" -win=true > output/case-ingress-gateway-multiple-services.txt
    go test -v -timeout=30m -tags integration ./test/integration/connect/envoy -run="TestEnvoy/case-ingress-gateway-sds" -win=true > output/case-ingress-gateway-sds.txt
    go test -v -timeout=30m -tags integration ./test/integration/connect/envoy -run="TestEnvoy/case-ingress-gateway-simple" -win=true > output/case-ingress-gateway-simple.txt
    go test -v -timeout=30m -tags integration ./test/integration/connect/envoy -run="TestEnvoy/case-ingress-gateway-tls" -win=true > output/case-ingress-gateway-tls.txt
    go test -v -timeout=30m -tags integration ./test/integration/connect/envoy -run="TestEnvoy/case-gateway-without-services" -win=true > output/case-gateway-without-services.txt
    go test -v -timeout=30m -tags integration ./test/integration/connect/envoy -run="TestEnvoy/case-terminating-gateway-hostnames" -win=true > output/case-terminating-gateway-hostnames.txt
    go test -v -timeout=30m -tags integration ./test/integration/connect/envoy -run="TestEnvoy/case-terminating-gateway-simple" -win=true > output/case-terminating-gateway-simple.txt
    go test -v -timeout=30m -tags integration ./test/integration/connect/envoy -run="TestEnvoy/case-terminating-gateway-subsets" -win=true > output/case-terminating-gateway-subsets.txt
    go test -v -timeout=30m -tags integration ./test/integration/connect/envoy -run="TestEnvoy/case-terminating-gateway-without-services" -win=true > output/case-terminating-gateway-without-services.txt
    go test -v -timeout=30m -tags integration ./test/integration/connect/envoy -run="TestEnvoy/case-dogstatsd-udp" -win=true > output/case-dogstatsd-udp.txt
    go test -v -timeout=30m -tags integration ./test/integration/connect/envoy -run="TestEnvoy/case-grpc" -win=true > output/case-grpc.txt
    go test -v -timeout=30m -tags integration ./test/integration/connect/envoy -run="TestEnvoy/case-http" -win=true > output/case-http.txt
    go test -v -timeout=30m -tags integration ./test/integration/connect/envoy -run="TestEnvoy/case-http-badauthz" -win=true > output/case-http-badauthz.txt
    go test -v -timeout=30m -tags integration ./test/integration/connect/envoy -run="TestEnvoy/case-statsd-udp" -win=true > output/case-statsd-udp.txt
    go test -v -timeout=30m -tags integration ./test/integration/connect/envoy -run="TestEnvoy/case-stats-proxy" -win=true > output/case-stats-proxy.txt
    echo "completado 100%"
fi

echo "Completed test"