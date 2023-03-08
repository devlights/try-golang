FROM gitpod/workspace-go:latest

ARG GITPOD_HOME_TMP="/home/gitpod/tmp"
ARG GITPOD_HOME_GO="/home/gitpod/go"
ARG GITPOD_HOME_GO_OLD="/home/gitpod/go-old"
ARG GO_VERSION="1.20.2"

USER gitpod

RUN sudo apt-get update -q && \
    rm -rf ${GITPOD_HOME_TMP} && \
    mkdir -p ${GITPOD_HOME_TMP} && \
    cd ${GITPOD_HOME_TMP} && \
    wget -q -O go.tar.gz https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz && \
    tar zxf go.tar.gz && \
    mv ${GITPOD_HOME_GO} ${GITPOD_HOME_GO_OLD} && \
    mv go /home/gitpod
