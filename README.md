# netping.org

![CI](https://github.com/malferov/netping/workflows/CI/badge.svg)

- Ping - used to test the reachability of a host and measures the round-trip time
- DNS lookup - used to obtain domain name or IP address mapping, or other DNS records
- Port check - used to determine whether specific port is up and running
- Traceroute - indicates the route and measures transit delays
- Reverse lookup – determines the domain name associated with an IP address
- Proxy checker – detects a proxy server
- Country by IP – detects country by IP address or hostname
- Network calculator – used to enumerate addresses belong to subnet
- WHOIS – shows the registered users or assignees of a domain name
- Bandwidth meter – measures the maximum rate of data transfer

## local
```
ver=$(git rev-parse --short HEAD) docker-compose up --build -d
```
## kube init
```
kubectl create namespace netping
docker login docker.pkg.github.com
kubectl create secret generic regcred \
  --from-file=.dockerconfigjson=<path/to/.docker/config.json> \
  --type=kubernetes.io/dockerconfigjson \
  --namespace netping
```
## terraform init
```
remote backend -> netping workspace -> general settings -> local mode
```
