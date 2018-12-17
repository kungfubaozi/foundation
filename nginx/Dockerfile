FROM nginx:latest

RUN mkdir /etc/nginx/crt
RUN mkdir /etc/nginx/conf

WORKDIR /etc/nginx

ADD crt/server.pem /etc/nginx/crt/server.pem
ADD crt/server.key /etc/nginx/crt/server.key

ADD nginx.conf /etc/nginx/nginx.conf
ADD api.conf /etc/nginx/vhost/api.conf

