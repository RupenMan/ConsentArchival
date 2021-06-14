FROM public.ecr.aws/bitnami/golang:1.16 as build-image

# WORKDIR /go/src
WORKDIR $HOME/GoLang/ConsentArchival

COPY go.mod go.sum main.go ./

# RUN go build -o $HOME/go/bin

RUN go build -o ../bin

RUN go mod vendor

FROM public.ecr.aws/lambda/go:1

COPY --from=build-image /go/bin/ /var/task/

# Command can be overwritten by providing a different command in the template directly.
CMD ["consent-archival"]
