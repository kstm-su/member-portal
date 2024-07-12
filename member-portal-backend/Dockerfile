FROM --platform=$BUILDPLATFORM golang:1.22.4 as build
RUN mkdir /storage

WORKDIR /go/src/github.com/kstm-su/Member-Portal/backend/

COPY ./go.* ./

RUN --mount=type=cache,target=/go/pkg/mod go mod download

ENV GOCACHE=/tmp/go/cache
ENV CGO_ENABLED=0

ARG TARGETOS
ARG TARGETARCH
ENV GOOS=$TARGETOS
ENV GOARCH=$TARGETARCH

COPY . .
RUN --mount=type=cache,target=/go/pkg/mod --mount=type=cache,target=/tmp/go/cache \
    go build -o /Member-Portal

FROM gcr.io/distroless/static-debian11
WORKDIR /app
EXPOSE 8080

COPY --from=build /storage/ /app/storage/
VOLUME /app/storage

COPY --from=build /Member-Portal ./
CMD ["./Member-Portal"]
