FROM alpine:3.4

RUN apk add --update ca-certificates && \
    rm -rf /var/cache/apk/*

COPY ./blockathon-backend-linux-amd64 /app/bin/blockathon-backend-linux-amd64

CMD ["/app/bin/blockathon-backend-linux-amd64"]
