FROM golang:1.15 AS build
ARG sha
ARG ver
ARG date
WORKDIR /go/src
COPY dns.go .
RUN go mod init dns
RUN go mod tidy
RUN go build -o dns -ldflags "-X main.sha=$sha -X main.version=$ver -X 'main.date=$date'"

FROM centos
RUN yum install -y bind-utils
EXPOSE 5000
COPY --from=build /go/src/dns .
ENTRYPOINT ["./dns"]
CMD ["-stderrthreshold=ERROR"]
