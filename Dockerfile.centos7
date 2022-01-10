FROM openshift/origin-release:golang-1.14 AS build
COPY . /go/src/github.com/mintel/eventrouter
WORKDIR /go/src/github.com/mintel/eventrouter
RUN go build .
FROM centos:7
COPY --from=build /go/src/github.com/mintel/eventrouter/eventrouter /bin/eventrouter
CMD ["/bin/eventrouter", "-v", "3", "-logtostderr"]
LABEL version=v0.3
