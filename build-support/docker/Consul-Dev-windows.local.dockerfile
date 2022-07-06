FROM mcr.microsoft.com/windows/servercore:1809

RUN mkdir consul
RUN mkdir C:\\consul\\data
RUN mkdir C:\\consul\\config

EXPOSE 8300
EXPOSE 8301 8301/udp 8302 8302/udp
EXPOSE 8500 8600 8600/udp

COPY dist/ C:\\consul

ENV PATH C:\\consul;%PATH%

COPY .release/docker/docker-entrypoint-windows.sh C:

# TODO Comment out until you have a working docker-entrypoint-windows.sh
# ENTRYPOINT ["docker-entrypoint-windows.sh"]
# CMD ["agent" "-dev" "-client" "0.0.0.0"]