FROM golang:1.20.6-bullseye AS build-env

WORKDIR /go/src/github.com/xnephilim/black

COPY . .

RUN make build

FROM golang:1.20.6-bullseye

RUN apt-get update  \ 
&& apt-get install ca-certificates jq=1.6-2.1 -y --no-install-recommends

WORKDIR /root

COPY --from=build-env /go/src/github.com/xnephilim/black/build/black /usr/bin/black

EXPOSE 26656 26657 1317 9090 8545 8546

CMD ["black"]
