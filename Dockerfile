# Use an official Golang runtime as a parent image
FROM golang:alpine

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Build the Go app
RUN go build -o main .

# Inform Docker that port 2384 is to be exposed
EXPOSE 2384

# Command to run the executable
CMD ["./main"]