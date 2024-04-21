FROM golang:1.22.2 as BUILDER

WORKDIR /app
COPY package.json pnpm-lock.yaml ./

# Setup Env
RUN <<EOF
apt-get update
apt-get install -y nodejs npm
npm --version
npm install
EOF

# Copy Go mod 
COPY go.mod go.sum ./
RUN <<EOF
go mod tidy
go mod download
EOF

# Copy remaining files
COPY . .

RUN <<EOF
go install github.com/a-h/templ/cmd/templ@latest
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
templ generate view
sqlc generate
EOF

# Setup Project
RUN <<EOF
npx tailwindcss -i view/css/input.css -o public/styles.css --minify
EOF

RUN CGO_ENABLED=0 go build -o GrowBro .

#FROM busybox:glibc
FROM gcr.io/distroless/static:nonroot

LABEL org.opencontainers.image.source=https://github.com/inlets/inlets-operator
USER nonroot:nonroot

WORKDIR /app
COPY --from=BUILDER /app/GrowBro /app/GrowBro

CMD [ "./GrowBro" ]
