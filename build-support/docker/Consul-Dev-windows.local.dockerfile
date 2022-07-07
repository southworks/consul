FROM mcr.microsoft.com/windows/servercore:1809

RUN ["powershell", "Set-ExecutionPolicy", "Bypass", "-Scope", "Process", "-Force;"]
RUN ["powershell", "iex ((New-Object System.Net.WebClient).DownloadString('https://chocolatey.org/install.ps1'))"]

RUN choco install git.install -yf

RUN mkdir consul
RUN mkdir C:\\consul\\data
RUN mkdir C:\\consul\\config

EXPOSE 8300
EXPOSE 8301 8301/udp 8302 8302/udp
EXPOSE 8500 8600 8600/udp

COPY dist/ C:\\consul

ENV PATH C:\\Program Files\\Git\\bin;C:\\consul;%PATH%

COPY .release/docker/docker-entrypoint-windows.sh C: 

ENTRYPOINT ["bash.exe", "C:\\docker-entrypoint-windows.sh"]

# TODO Comment out until you have a working in Windows
# CMD ["agent" "-dev" "-client" "0.0.0.0"]

