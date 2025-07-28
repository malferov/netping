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

dig @$SOA +noall +answer $1 ANY
