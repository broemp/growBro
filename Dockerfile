FROM --platform=${BUILDPLATFORM:-linux/amd64} golang:1.22.1 as BUILDER

ARG TARGETPLATFORM
ARG BUILDPLATFORM
ARG TARGETOS
ARG TARGETARCH

ARG Version
ARG GitCommit

ENV CGO_ENABLED=0
ENV GO111MODULE=on

ARG Version
ARG GitCommit

WORKDIR /app

# Setup Env
RUN <<EOF
apt-get update
apt-get install -y nodejs npm
npm install -g pnpm
go install github.com/a-h/templ/cmd/templ@latest
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
EOF

# Setup Tailwindcss
COPY package.json pnpm-lock.yaml ./
RUN pnpm install

# Copy Go mod 
COPY go.mod go.sum ./
RUN <<EOF
go mod tidy
go mod download
EOF

# Copy remaining files
COPY . .

# Setup Project
RUN <<EOF
templ generate view
sqlc generate
pnpx tailwindcss  -i view/css/input.css -o public/styles.css --minify
EOF

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
