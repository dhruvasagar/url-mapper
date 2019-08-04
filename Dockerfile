FROM golang:1.12.7 as builder
COPY . .
RUN make linux

FROM alpine:3.9
WORKDIR /root/
COPY --from=builder url-mapper .
CMD ["./url-mapper"]
