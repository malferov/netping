FROM golang AS build
ARG sha
ARG ver
ARG date
WORKDIR /go/src
COPY dns.go .
RUN go get github.com/gin-gonic/gin
RUN go get github.com/golang/glog
RUN go get github.com/gin-contrib/cors
RUN go build -o dns -ldflags "-X main.sha=$sha -X main.version=$ver -X 'main.date=$date'"

FROM centos
RUN yum install -y bind-utils
EXPOSE 5000
COPY --from=build /go/src/dns .
ENTRYPOINT ["./dns"]
CMD ["-stderrthreshold=ERROR"]