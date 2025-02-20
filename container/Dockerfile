FROM opensuse/leap:15.6


RUN zypper --non-interactive refresh && \
    zypper --non-interactive install go && \
    zypper --non-interactive clean

WORKDIR /app

COPY server.go .

RUN go build -o server server.go

EXPOSE 9977

CMD ["./server"]