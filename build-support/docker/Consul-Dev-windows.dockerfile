ARG CONSUL_IMAGE_VERSION=1.12.0
FROM mcr.microsoft.com/windows/servercore:1809

RUN mkdir consul

ENV CONSUL_VERSION=1.12.0
ENV CONSUL_URL=https://releases.hashicorp.com/consul/${CONSUL_VERSION}/consul_${CONSUL_VERSION}_windows_amd64.zip

RUN curl %CONSUL_URL% -L -o consul.zip
RUN tar -xf consul.zip -C consul

ENV PATH C:\\consul;%PATH%
