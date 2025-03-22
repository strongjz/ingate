FROM python:3-alpine

ARG USER=1001

WORKDIR /docs
COPY docs /docs
COPY tools/docs/requirements.txt /docs  

RUN adduser -h /docs -D -u $USER mkdocs \
&& apk add bash \
&& apk add git 

RUN ls -la 
RUN ls -la /docs

RUN pip install --upgrade pip
RUN pip install -r /docs/requirements.txt