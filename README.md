## Build the docker image
docker build -t golang-api .

## Then Run the image with command:
docker run -p 8080:8000 golang-api