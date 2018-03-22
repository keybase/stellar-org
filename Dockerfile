FROM golang:1.9
RUN bash -c "curl https://glide.sh/get | sh"
WORKDIR /go/src/github.com/keybase/stellar-org

COPY glide.lock /go/src/github.com/keybase/stellar-org
COPY glide.yaml /go/src/github.com/keybase/stellar-org
RUN glide install

COPY . .
RUN go install github.com/keybase/stellar-org/tools/...
RUN go install github.com/keybase/stellar-org/services/...
