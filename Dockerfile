FROM golang:1.14.12-stretch
RUN mkdir -p "$GOPATH/src/github.com/kardiachain/go-kardiamain"
WORKDIR /go/src/github.com/kardiachain/go-kardiamain
RUN apt-get update && apt-get install -y libzmq3-dev
ADD . .
WORKDIR /go/src/github.com/kardiachain/go-kardiamain/cmd
RUN go install
WORKDIR /go/src/github.com/kardiachain/go-kardiamain/dualnode/eth/eth_client
RUN go install
WORKDIR /go/bin
RUN mkdir -p /go/bin/cfg
COPY cmd/cfg /go/bin/cfg
ENTRYPOINT ["./cmd"]
