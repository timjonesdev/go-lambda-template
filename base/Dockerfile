FROM golang:1.12 as builder

COPY / /app/

WORKDIR /app

RUN GIT_TERMINAL_PROMPT=1 GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build --installsuffix cgo --ldflags="-s" -o {{exe_name}}

FROM {{docker_base}}

COPY --from=builder /app/{{exe_name}} /app/{{exe_name}}

WORKDIR /app

ENTRYPOINT ["./{{exe_name}}"]