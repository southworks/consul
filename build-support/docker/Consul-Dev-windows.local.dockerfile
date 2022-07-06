FROM mcr.microsoft.com/windows/servercore:1809

RUN mkdir consul

COPY dist/ C:\\consul

ENV PATH C:\\consul;%PATH%
