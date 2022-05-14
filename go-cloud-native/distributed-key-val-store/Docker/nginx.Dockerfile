# The parent image. At build time, this image will be pulled and subsequent
# instructions run against it.
FROM ubuntu:20.04

# Update apt cache and install nginx without any prompt approval.
RUN apt-get update && apt-get install --yes nginx

# Tell docker container this image's container will use port 80
EXPOSE 80

# Run nginx in the foreground. This is important: without a foreground process
# the container will automatically stop
CMD ["nginx", "-g", "daemon off;"]