FROM ubuntu:18.04
RUN  apt-get update

RUN apt-get install -y openssh-server

RUN apt-get clean && \
    rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

RUN echo 'root:root' | chpasswd
RUN useradd --create-home admin

RUN mkdir /var/run/sshd

COPY ssh/sshd_config /etc/ssh/sshd_config
COPY ssh/users_ca.pub /etc/ssh/users_ca.pub

RUN mkdir /root/.ssh

EXPOSE 22

CMD ["/usr/sbin/sshd", "-D"]