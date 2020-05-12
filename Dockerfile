FROM golang:alpine AS builder

RUN apk add --no-cache gcc libc-dev

WORKDIR /workspace

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a

FROM scratch AS final

WORKDIR /

COPY --from=builder /workspace/secrets-injektor /opt/secrets-injektor

CMD [ "./opt/secrets-injektor" ]
