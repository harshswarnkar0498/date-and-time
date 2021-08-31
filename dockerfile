FROM golang:latest

RUN git clone https://github.com/harshswarnkar0498/date-and-time.git

WORKDIR date-and-time

CMD go run web.go

EXPOSE 6001
