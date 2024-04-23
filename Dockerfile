FROM --platform=${BUILDPLATFORM:-linux/amd64} golang:1.22.2 as BUILDER

ARG TARGETPLATFORM
ARG BUILDPLATFORM
ARG TARGETOS
ARG TARGETARCHinput.css

ARG Version
ARG GitCommit

ENV CGO_ENABLED=0
ENV GO111MODULE=on

ARG Version
ARG GitCommit

WORKDIR /app

# Copy Go mod 
COPY go.mod go.sum ./
RUN <<EOF
go mod tidy
go mod download
EOF

# Setup Project
COPY . .

RUN CGO_ENABLED=${CGO_ENABLED} GOOS=${TARGETOS} GOARCH=${TARGETARCH} \
  go build -ldflags "-s -w -X github.com/broemp/growbro/version.Release=${Version} -X github.com/broemp/growbro/version.SHA=${GitCommit}" \
  -o growbro .

#FROM busybox:glibc
FROM --platform=${BUILDPLATFORM:-linux/amd64} gcr.io/distroless/static:nonroot

LABEL org.opencontainers.image.source=https://github.com/inlets/inlets-operator
USER nonroot:nonroot

WORKDIR /app
COPY --from=BUILDER /app/growbro /app/growbro

CMD [ "./growbro" ]
