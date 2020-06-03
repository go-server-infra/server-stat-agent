FROM golang:alpine as gobuild
WORKDIR /app
ADD ./ .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/agent


FROM scratch
WORKDIR /app
COPY --from=gobuild /app/build/agent ./

CMD [ "/app/agent" ]
