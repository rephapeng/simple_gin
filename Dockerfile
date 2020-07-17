FROM centos:latest
MAINTAINER Refany Anhar

ARG STATE_ENV
ENV STATE=$STATE_ENV

#create user
ARG user=simple_gin
ARG group=simple_gin
ARG uid=1000
ARG gid=1000

RUN groupadd -g ${gid} ${group} \
    && useradd -d "/home/${user}" -u ${uid} -g ${gid} -m -s /bin/bash ${user}
RUN usermod -a -G root ${user}

USER ${user}

RUN mkdir /home/${user}/bin

#add binary 
ADD simple_gin /home/${user}/bin/

#expose port
EXPOSE 8080

CMD ./home/simple_gin/bin/simple_gin
