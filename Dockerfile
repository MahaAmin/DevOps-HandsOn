FROM golang:1.13-alpine AS builder

WORKDIR /go/src/app

# copy be-svc/ and fe-svc/ from host to container
COPY . .

# build be-svc/ in be-bin output file
RUN cd be-svc && go build -o be-bin && cd ..

# build fe-svc/ in fe-bin output file
RUN cd fe-svc && go build -o fe-bin && cd ..

#############################################################
FROM alpine:3.11

WORKDIR /src/app

# copy binary files from build stage 
COPY --from=builder /go/src/app/be-svc/be-bin .
COPY --from=builder /go/src/app/fe-svc/fe-bin .

# run only the backend package
ENTRYPOINT [ "./be-bin" ]