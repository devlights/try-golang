FROM gitpod/workspace-go:latest

USER gitpod

RUN sudo apt update -q && \
    rm -rf /home/gitpod/tmp && \
    mkdir -p /home/gitpod/tmp && \
    cd /home/gitpod/tmp && \
    wget -q -O go.tar.gz https://go.dev/dl/go1.20.1.linux-amd64.tar.gz && \
    tar zxf go.tar.gz && \
    mv /home/gitpod/go /home/gitpod/go-old && \
    mv go /home/gitpod
