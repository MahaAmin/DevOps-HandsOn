FROM golang:1.13-alpine AS builder

WORKDIR /go/src/app

COPY . .

RUN cd be-svc && go build -o be-bin && cd ..

RUN cd fe-svc && go build -o fe-bin && cd ..

FROM alpine:3.11

WORKDIR /src/app

COPY --from=builder /go/src/app/be-svc/be-bin .
COPY --from=builder /go/src/app/fe-svc/fe-bin .

ENTRYPOINT [ "./be-bin" ]