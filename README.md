# netping.org

![CI](https://github.com/malferov/netping/workflows/CI/badge.svg)

- Ping - shows how long it takes for packets to reach host
- Traceroute - traces the route of packets to destination host from our service
- DNS lookup - look up DNS records
- Port check - tests if TCP port is opened on specified IP
- Reverse lookup – gets hostname by IP address
- Proxy checker – detects a proxy server
- Country by IP – detects country by IP or hostname
- Network calculator – calculates subnet range by network mask
- WHOIS – lists contact info for an IP or domain
- Bandwidth meter – detects your download speed from our server

## local
```
ver=$(git rev-parse --short HEAD) docker-compose up --build -d
```
## inin
```
kubectl create namespace netping
```
