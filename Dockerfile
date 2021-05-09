# Start from the latest golang base image
FROM golang:latest

# Add Maintainer Info 
LABEL maintainer="Austin Spencer <abspencer2097@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source from the current directory to the Working Directory inside the container
COPY ./bin/main .

# Expose port 8558 to the outside world
EXPOSE 8558

# Command to run the executable
CMD ["./main"]