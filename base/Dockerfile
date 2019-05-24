FROM golang:alpine

# Skip known public key check to be able to pull from private repositories 
ENV GIT_SSH_COMMAND "ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no"
RUN apk --no-cache add git ca-certificates openssh \
&& go get github.com/triggermesh/aws-custom-runtime \
&& go get github.com/triggermesh/knative-lambda-runtime/go-1.x \
&& go get github.com/golang/dep/...

WORKDIR /go/src/handler
COPY . .
RUN if [ -f "$HOME/.ssh/id_${SSH_KEY}" ]; then \
    eval "\$(ssh-agent -s)"; \
    ssh-add $HOME/.ssh/id_${SSH_KEY}; \
fi \
&& if [ -f "Gopkg.toml" ]; then dep ensure; fi \
&& go get -v \
&& go install

FROM alpine
WORKDIR /opt
ENV LAMBDA_TASK_ROOT "/opt"
ENV _HANDLER "handler"
COPY --from=0 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=0 /go/bin/go-1.x /opt/bootstrap
COPY --from=0 /go/bin/ /opt
ENTRYPOINT ["/opt/aws-custom-runtime"]
