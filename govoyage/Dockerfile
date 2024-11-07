FROM golang:1.23-alpine AS builder
RUN apk add -U tzdata
RUN apk --update add ca-certificates
WORKDIR /app
COPY . .
COPY ../go.* .
RUN go mod tidy
RUN make app-linux


FROM scratch
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /app/bin/govoyage-server .
EXPOSE 8080 8090
CMD ["/govoyage-server"]
