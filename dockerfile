FROM golang:1.24.3

LABEL Project="dockerize" \ 
      Member="iallaoui - melghama"


COPY . .

RUN go build -o server

EXPOSE 8080

CMD [ "./server" ]


