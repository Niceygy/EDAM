FROM alpine:latest
# Set working directory in the container
WORKDIR /root/

COPY . .

LABEL org.opencontainers.image.description="EDDataCollector"
LABEL org.opencontainers.image.authors="Niceygy (Ava Whale)"

RUN chmod 777 ./edam

# Expose the application port
EXPOSE 3696
# Run the application
CMD ["./edam"]
