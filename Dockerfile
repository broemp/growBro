# Dockerfile
FROM busybox:glibc
WORKDIR /app
COPY growbro \
  growbro
ENTRYPOINT ["growbro"]
EXPOSE 3000
