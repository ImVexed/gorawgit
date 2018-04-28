FROM alpine:latest

COPY main /main
COPY ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

EXPOSE 80

ENTRYPOINT [ "/main" ]