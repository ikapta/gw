FROM golang

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY="https://goproxy.cn,direct"

WORKDIR /build

COPY . .

RUN go build -o front_api_gw .

WORKDIR /dist

RUN cp /build/front_api_gw .

EXPOSE 8003

CMD ["/dist/front_api_gw"]