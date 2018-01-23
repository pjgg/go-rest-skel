FROM frolvlad/alpine-glibc

ADD bin/app-linux-amd64 /app
CMD /app
