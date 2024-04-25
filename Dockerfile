FROM ubuntu
WORKDIR /app/
COPY growbro /app/growbro
ENTRYPOINT ["/app/growbro"]
EXPOSE 3000
