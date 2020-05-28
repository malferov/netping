FROM golang AS build
ARG ver
ARG date
WORKDIR /go/src
COPY ping.go .
RUN go get github.com/gin-gonic/gin
RUN go get github.com/golang/glog
RUN go get github.com/gin-contrib/cors
RUN go build -o ping -ldflags "-X main.version=$ver -X 'main.date=$date'"

FROM centos
EXPOSE 5000
COPY --from=build /go/src/ping .
ENTRYPOINT ["./ping"]
CMD ["-stderrthreshold=ERROR"]
