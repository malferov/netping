#!/bin/sh

name=${1%.}

for i in `seq 6`; do
    SOA=`dig +noall +answer $name SOA | grep SOA | tr -s '[:space:]' ' ' | cut -d' ' -f5`
    if [ ${#SOA} -gt 0 ]; then break; fi
    name=${name#*.}
    if [ $name = ${name#*.} ]; then break; fi
done

if [ ${#SOA} -lt 4 ];
    then echo Not found; exit
fi

dig @$SOA +noall +answer +multiline $1 SOA $1 A $1 AAAA $1 TXT $1 MX $1 NS $1 PTR $1 CNAME $1 DNAME $1 CAA $1 CERT $1 HTTPS $1 DNSKEY $1 DS
