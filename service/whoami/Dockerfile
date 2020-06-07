FROM golang AS build
ARG sha
ARG ver
ARG date
WORKDIR /go/src
COPY whoami.go .
RUN go get github.com/gin-gonic/gin
RUN go get github.com/golang/glog
RUN go get github.com/gin-contrib/cors
RUN go build -o whoami -ldflags "-X main.sha=$sha -X main.version=$ver -X 'main.date=$date'"

FROM centos
EXPOSE 5000
COPY --from=build /go/src/whoami .
ENTRYPOINT ["./whoami"]
CMD ["-stderrthreshold=ERROR"]
