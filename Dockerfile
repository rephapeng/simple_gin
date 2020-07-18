FROM centos:latest
MAINTAINER Refany Anhar

#read the argument from docker build command
ARG STATE_ENV
ENV STATE=$STATE_ENV

#create user
ARG user=simple_gin
ARG group=simple_gin
ARG uid=1000
ARG gid=1000

#create user-group
RUN groupadd -g ${gid} ${group} \
    && useradd -d "/home/${user}" -u ${uid} -g ${gid} -m -s /bin/bash ${user}
RUN usermod -a -G root ${user}

USER ${user}

RUN mkdir /home/${user}/bin

#add go binary 
ADD simple_gin /home/${user}/bin/

#expose port
EXPOSE 8080

#run the application
CMD ./home/simple_gin/bin/simple_gin
