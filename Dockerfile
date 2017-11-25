FROM alpine:3.4

RUN apk add --update ca-certificates bash && \
    rm -rf /var/cache/apk/*

COPY ./blockathon-backend-linux-amd64 /app/
EXPOSE 3000

CMD ["/app/blockathon-backend-linux-amd64"]
