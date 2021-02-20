FROM gitpod/workspace-full

USER gitpod

# Install custom tools, runtime, etc. using apt-get
# For example, the command below would install "bastet" - a command line tetris clone:
#
# RUN sudo apt-get -q update && \
#     sudo apt-get install -yq bastet && \
#     sudo rm -rf /var/lib/apt/lists/*
#
# More information: https://www.gitpod.io/docs/config-docker/
RUN sudo apt update && \
    sudo apt install -yq info bc && \
    sudo rm -rf /var/lib/apt/lists/*

RUN go get golang.org/dl/go1.16 && go1.16 download && go mod tidy
ENV GOROOT=/home/gitpod/sdk/go1.16
ENV PATH=$GOROOT/bin:$PATH
