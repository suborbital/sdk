# Build and test
docker build . -t se2-sdk-bindings
docker run -t se2-sdk-bindings ls -a /tmp

# Extract the built files
id=$(docker create se2-sdk-bindings:latest)
docker cp $id:/tmp ./bindings
docker rm -v $id
