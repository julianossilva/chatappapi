FROM ubuntu:22.04

USER root

RUN export DEBIAN_FRONTEND=noninteractive
RUN apt-get update 
RUN apt-get -y install --no-install-recommends \
    git \
    curl \
    golang-go \
    nodejs npm

RUN go install github.com/mitranim/gow@latest

ARG USERNAME=vscode
ARG GROUP=vscode
ARG UID=1000
ARG GID=1000
RUN groupadd -g ${GID} ${GROUP}
RUN useradd -u ${UID} -g ${GROUP} -s /bin/sh -m ${USERNAME} 

USER ${UID}:${GID}
#USER ${USERNAME}

WORKDIR /app 

#Run command as non root
#RUN command... 