FROM    vikings/alpine:latest
LABEL   maintainer=ztao8607@gmail.com
COPY    bin/dns-antman /dns-antman
EXPOSE  8000
ENTRYPOINT ["/dns-antman"]