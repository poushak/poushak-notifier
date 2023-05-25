#! /bin/bash

echo 'Removing existing images'
docker rmi notifier:1.0.0

echo 'Creating new image for poushak-notifier'
docker build --tag notifier:1.0.0 .