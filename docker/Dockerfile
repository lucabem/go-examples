FROM public.ecr.aws/lambda/go:1

RUN yum install -y wget

# Install Go
# Download and install Go 1.21
RUN wget https://go.dev/dl/go1.21.0.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf go1.21.0.linux-amd64.tar.gz \
    && rm go1.21.0.linux-amd64.tar.gz

ENV PATH="/usr/local/go/bin:${PATH}"

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to the container
COPY go.mod go.sum ./

# Download and cache Go modules
RUN go mod download

# Copy the rest of the application code to the container
COPY main.go ./
COPY pkg/ ./pkg
COPY src/ ./src
COPY scripts/ ./scripts

# Build the application
RUN go build -o ${LAMBDA_TASK_ROOT}

# Set the CMD to your handler (could also be done as a parameter override outside of the Dockerfile)
CMD [ "main" ]
