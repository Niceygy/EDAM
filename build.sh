cd src
CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o edam .
cd ..
mv src/edam .

# Build the Docker image
docker build -t niceygy/edam .

# Tag the Docker image
docker tag niceygy/edam ghcr.io/niceygy/edam:latest

# Push the Docker image to GH registry
docker push ghcr.io/niceygy/edam:latest

#Update local container

cd /opt/stacks/elite_apps

docker compose pull

docker compose down

docker compose up -d

docker logs edam -f