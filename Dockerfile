FROM alpine:3.18

WORKDIR /app

COPY . .

RUN chmod +x /app/bin/main

CMD ["./bin/main"]
