FROM golang:1.17-alpine AS builder 
WORKDIR /build 
COPY . .
ENV CGO_ENABLED 0
RUN go test -v ./... 
RUN go build -o server cmd/main.go 

FROM alpine 
WORKDIR /app 
COPY --from=builder /build/static/ static/
COPY --from=builder /build/template/ template/
COPY --from=builder /build/config/ config/ 
COPY --from=builder /build/ascii/ ascii/
COPY --from=builder /build/server .
LABEL "authors"="ismanaraev aleka7sk"
LABEL "go version"="1.17"
LABEL "port"="8080"
ENTRYPOINT ["./server"]
