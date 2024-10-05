# Build stage for Terraform, kubectl, and AWS CLI
FROM --platform=${BUILDPLATFORM:-linux/amd64} alpine AS builder

# Set environment variables for Terraform version, AWS CLI version, and kubectl version
ENV TERRAFORM_VERSION=1.9.5
ENV KUBECTL_VERSION=v1.29.2

# Install dependencies for building and installing
RUN apk add --no-cache \
    curl \
    unzip \
    jq

# Detect the architecture (amd64 or arm64)
ARG TARGETPLATFORM
RUN echo "Building for $TARGETPLATFORM"

# Install Terraform
RUN if [ "$TARGETPLATFORM" = "linux/arm64" ]; then \
        curl -LO https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_linux_arm64.zip; \
    else \
        curl -LO https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_linux_amd64.zip; \
    fi \
    && unzip terraform_${TERRAFORM_VERSION}_linux_*.zip \
    && mv terraform /usr/local/bin/ \
    && rm terraform_${TERRAFORM_VERSION}_linux_*.zip

# Install kubectl
RUN if [ "$TARGETPLATFORM" = "linux/arm64" ]; then \
        curl -LO "https://dl.k8s.io/release/${KUBECTL_VERSION}/bin/linux/arm64/kubectl"; \
    else \
        curl -LO "https://dl.k8s.io/release/${KUBECTL_VERSION}/bin/linux/amd64/kubectl"; \
    fi \
    && chmod +x kubectl \
    && mv kubectl /usr/local/bin/

# Final stage - minimal image
FROM --platform=${TARGETPLATFORM:-linux/amd64} alpine

RUN apk add --no-cache aws-cli

# Copy binaries from builder stage
COPY --from=builder /usr/local/bin/terraform /usr/local/bin/terraform
COPY --from=builder /usr/local/bin/kubectl /usr/local/bin/kubectl

# Set the entrypoint to bash, ensuring bash is executable
CMD ["/bin/sh"]

