FROM golang:alpine

WORKDIR /go/src/app

# Get Reflex for live reload in development
ENV GO111MODULE=on
RUN go get github.com/cespare/reflex

# Copy mod files
COPY go.mod .
COPY go.sum .

# Install dependencies
RUN go mod download

# Copy the app to the working directory
COPY . .

# Copy ENV file
COPY .env .

# Build and run the app
RUN go build -o app
CMD ["./app"]