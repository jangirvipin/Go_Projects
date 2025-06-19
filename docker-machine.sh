#!/bin/bash

# CONFIGURATIONS
MACHINE_NAME="aws-golang-machine"
AWS_REGION="us-west-2"
INSTANCE_TYPE="t2.micro"
SSH_KEY_PATH="$HOME/.ssh/id_rsa"
IMAGE_NAME="golang-server"
PORT=8080

echo "=== Creating Docker Machine on AWS ==="
docker-machine create \
  --driver amazonec2 \
  --amazonec2-region "$AWS_REGION" \
  --amazonec2-instance-type "$INSTANCE_TYPE" \
  --amazonec2-ssh-keypath "$SSH_KEY_PATH" \
  "$MACHINE_NAME"

if [ $? -ne 0 ]; then
  echo "Failed to create Docker machine. Exiting."
  exit 1
fi

echo "=== Switching environment to use the remote Docker host ==="
eval "$(docker-machine env "$MACHINE_NAME")"

echo "=== Building Docker image [$IMAGE_NAME] ==="
docker build -t "$IMAGE_NAME" .

echo "=== Running container on AWS EC2 instance ==="
docker run -d -p $PORT:$PORT -it "$IMAGE_NAME"

echo "=== App is now running on AWS! ==="
PUBLIC_IP=$(docker-machine ip "$MACHINE_NAME")
echo "Access it at: http://$PUBLIC_IP:$PORT"
