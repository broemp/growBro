FROM golang:1.22.2 as GENERATOR

WORKDIR /app

COPY . .

RUN <<EOF
go install github.com/a-h/templ/cmd/templ@latest
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
templ generate view
sqlc generate
EOF

FROM node:21 as CSS

WORKDIR /app

COPY package.json package.json
RUN npm install

COPY ./view/css/input.css ./view/css/input.css
RUN npx tailwindcss  -i view/css/input.css -o public/styles.css --minify

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
COPY --from=GENERATOR /app /app
COPY --from=CSS /app/public/styles.css ./public/

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
