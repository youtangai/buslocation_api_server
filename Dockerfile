FROM alpine:latest
RUN mkdir /code
COPY redis.json /code
COPY buslocation_api_server /code
WORKDIR /code
EXPOSE 8080
ENTRYPOINT [ "./main" ]