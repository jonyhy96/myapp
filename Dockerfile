FROM centos:latest
COPY . /
MAINTAINER haoyun
EXPOSE 8080 
ENV OEM myapp
ENV VER 2.0