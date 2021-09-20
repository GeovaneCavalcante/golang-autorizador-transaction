FROM golang:alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

RUN apk add -U make
RUN apk add -U build-base
# Download all the dependencies
RUN make dependencies
RUN make build


## Run the executable
ENTRYPOINT ["/app/autorizador"]