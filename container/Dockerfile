FROM opensuse/leap:15.6

RUN zypper refresh && \
    zypper in go -y && \
    zypper clean

WORKDIR /app

COPY server.go .

RUN go build -o server server.go

EXPOSE 9977

CMD ["./server"]