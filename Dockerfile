FROM centos:6.5
MAINTAINER Cean Cheng <cean.ch@gmail.com>

RUN yum groupinstall "Development Tools"
RUN yum -y install pcre* zlib* openssl*



EXPOSE 3000

CMD [ "application"]