ARG CONSUL_IMAGE_VERSION=1.12.0
FROM mcr.microsoft.com/windows/servercore:1809

RUN ["powershell", "Set-ExecutionPolicy", "Bypass", "-Scope", "Process", "-Force;"]
RUN ["powershell", "iex ((New-Object System.Net.WebClient).DownloadString('https://chocolatey.org/install.ps1'))"]

RUN choco install git.install -yf

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

ENV PATH C:\\Program Files\\Git\\bin;C:\\consul;%PATH%

COPY .release\docker\docker-entrypoint-windows.sh C:

ENTRYPOINT ["bash.exe", "C:\\docker-entrypoint-windows.sh"]
