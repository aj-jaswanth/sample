FROM alpine:latest

COPY ./app /app
RUN chmod +x ./app

ENTRYPOINT ./app
