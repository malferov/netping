FROM golang AS build
ARG sha
ARG ver
ARG date
ARG CGO_ENABLED=0
WORKDIR /go/src
COPY dns.go .
RUN go mod init dns
RUN go mod tidy
RUN go build -o dns -ldflags "-X main.sha=$sha -X main.version=$ver -X 'main.date=$date'"

FROM centos:7
RUN sed -i s/mirror.centos.org/vault.centos.org/g /etc/yum.repos.d/CentOS-*.repo
RUN sed -i s/^#.*baseurl=http/baseurl=http/g /etc/yum.repos.d/CentOS-*.repo
RUN sed -i s/^mirrorlist=http/#mirrorlist=http/g /etc/yum.repos.d/CentOS-*.repo
RUN yum install -y bind-utils
EXPOSE 5000
COPY --from=build /go/src/dns .
COPY dig.sh .
ENTRYPOINT ["./dns"]
CMD ["-stderrthreshold=ERROR"]
