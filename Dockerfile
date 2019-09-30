FROM ubuntu:18.04
RUN  apt-get update

RUN apt-get install -y openssh-server
RUN apt-get install -y rsyslog

RUN apt-get clean && \
    rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

RUN echo 'root:root' | chpasswd
RUN useradd --create-home admin -s /bin/bash
RUN useradd --create-home abhishek -s /bin/bash

RUN mkdir -p /var/run/sshd /var/tmp /var/authz/

COPY ssh/run.sh      /var/authz/run.sh
COPY ssh/sshd        /etc/pam.d/sshd
COPY ssh/sshd_config /etc/ssh/sshd_config
COPY users.json /etc/authz/users.json
COPY ssh/users_ca.pub /etc/ssh/users_ca.pub

RUN mkdir /root/.ssh

EXPOSE 22

CMD ["sh", "/var/authz/run.sh"]
