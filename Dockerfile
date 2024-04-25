FROM ubuntu
WORKDIR /app/
COPY growBro /app/growBro
ENTRYPOINT ["/app/growBro"]
EXPOSE 3000
