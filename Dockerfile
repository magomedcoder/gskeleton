FROM alpine:3.19 AS builder

ARG GOLANG_VERSION=1.22.6

RUN apk update && \
    apk add --no-cache make gcc openssh bash musl-dev openssl-dev ca-certificates && \
    update-ca-certificates && \
    rm -rf /var/cache/apk/*

RUN wget https://go.dev/dl/go$GOLANG_VERSION.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go$GOLANG_VERSION.linux-amd64.tar.gz && \
    rm go$GOLANG_VERSION.linux-amd64.tar.gz

ENV PATH=$PATH:/usr/local/go/bin

RUN mkdir /usr/src

RUN mkdir /usr/src/gskeleton

WORKDIR /usr/src/gskeleton

COPY . ./

RUN make install

RUN go mod tidy

RUN make build

FROM alpine:3.19

COPY --from=builder /usr/src/gskeleton/build /usr/bin

RUN mkdir /etc/gskeleton

EXPOSE 8000 50051

CMD ["sh"]
