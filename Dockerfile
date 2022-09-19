FROM --platform=$BUILDPLATFORM golang:1.19.1-alpine3.16 AS builder

WORKDIR /src

COPY . .
RUN go mod download

ARG TARGETOS
ARG TARGETARCH

ARG TARGETPLATFORM
ARG BUILDPLATFORM
RUN echo "Je cours sur $BUILDPLATFORM, je construis pour $TARGETPLATFORM" > /log

RUN GOOS=$TARGETOS GOARCH=$TARGETARCH CGO_ENABLED=0 go build -o /go/bin/api_va -ldflags '-extldflags "-static"' -tags timetzdata
RUN ls -alh /go/bin


FROM scratch

LABEL maintainer="Samuel MICHAUX<samuel.michaux@ik.Me>"

COPY --from=builder /go/bin/api_va /go/bin/api_va

EXPOSE 8080

ENTRYPOINT ["/go/bin/api_va"]