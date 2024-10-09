FROM gitpod/workspace-base:latest

USER gitpod
ENV GO_VERSION=1.23.2

# For ref, see: https://github.com/gitpod-io/workspace-images/blob/61df77aad71689504112e1087bb7e26d45a43d10/chunks/lang-go/Dockerfile#L10
ENV GOPATH=$HOME/go-packages
ENV GOROOT=$HOME/go
ENV PATH=$GOROOT/bin:$GOPATH/bin:$PATH

RUN curl -fsSL https://dl.google.com/go/go${GO_VERSION}.linux-amd64.tar.gz | tar xzs
RUN go install github.com/uudashr/gopkgs/cmd/gopkgs@v2 
RUN go install github.com/ramya-rao-a/go-outline@latest 
RUN go install github.com/cweill/gotests/gotests@latest 
RUN go install github.com/fatih/gomodifytags@latest 
RUN go install github.com/josharian/impl@latest 
RUN go install github.com/haya14busa/goplay/cmd/goplay@latest 
RUN go install github.com/go-delve/delve/cmd/dlv@latest 
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest 
RUN go install honnef.co/go/tools/cmd/staticcheck@latest 
RUN go install golang.org/x/tools/gopls@latest 
RUN printf '%s\n' 'export GOPATH=/workspace/go' \
                      'export PATH=$GOPATH/bin:$PATH' > $HOME/.bashrc.d/300-go

RUN sudo apt update && sudo apt install -y universal-ctags tree nkf wamerican miller

