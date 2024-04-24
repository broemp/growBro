# Dockerfile
FROM alpine:latest
COPY goreleaser-docker-manifest-actions-example \
  /usr/bin/goreleaser-docker-manifest-actions-example
ENTRYPOINT ["/usr/bin/goreleaser-docker-manifest-actions-example"]
EXPOSE 3000
