ARG CONSUL_IMAGE_VERSION=1.12.0
FROM mcr.microsoft.com/windows/servercore:1809

RUN mkdir C:\\consul
RUN mkdir -p C:\\consul\\data
RUN mkdir -p C:\\consul\\config

VOLUME [/consul/data]

EXPOSE 8300
EXPOSE 8301 8301/udp 8302 8302/udp
EXPOSE 8500 8600 8600/udp

ENV CONSUL_VERSION=1.12.0
ENV CONSUL_URL=https://releases.hashicorp.com/consul/${CONSUL_VERSION}/consul_${CONSUL_VERSION}_windows_amd64.zip

RUN curl %CONSUL_URL% -L -o consul.zip
RUN tar -xf consul.zip -C consul

ENV PATH C:\\consul;%PATH%

COPY .release\docker\docker-entrypoint-windows.sh C:

# TODO Comment out until you have a working docker-entrypoint-windows.sh
# ENTRYPOINT ["docker-entrypoint-windows.sh"]
# CMD ["agent" "-dev" "-client" "0.0.0.0"]