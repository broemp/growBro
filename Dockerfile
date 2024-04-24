# Dockerfile
FROM alpine:latest
WORKDIR /app
COPY growbro \
  growbro
ENTRYPOINT ["growbro"]
EXPOSE 3000
