FROM golang:latest

RUN git clone https://github.com/deadbedpk/web_app_ymdt.git

WORKDIR web_app_ymdt

CMD go run app.go

EXPOSE 8000
