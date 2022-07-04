ARG CONSUL_IMAGE_VERSION=1.12.0
FROM mcr.microsoft.com/windows/servercore:1809

RUN mkdir consul

ENV CONSUL_URL=https://releases.hashicorp.com/consul/1.12.0/consul_1.12.0_windows_amd64.zip
RUN curl %CONSUL_URL% -L -o consul.zip
RUN tar -xf consul.zip -C consul

ENV PATH C:\\consul;%PATH%
