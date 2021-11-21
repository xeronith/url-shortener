# Url-Shortener

This sample demonstrates a loosely-coupled extensible architecture with a working implementation of a shortener service. This implementation supports in-memory and SQLite storages and features caching to boost the performance. Everything has been implemented from scratch with no dependency to external libraries except for handling SQLite. The main focus was to create a robust architecture which can be scaled and enhanced upon indefinitely, hence many minor details, extensive logs, numerous tests and improvements were omitted for the sake of brevity and faster delivery. The sample can be run directly or via Docker.

To run the sample from terminal: `go run main.go -p 8080 -b http://localhost:8080`

To build the Docker image you can use: `docker build -t xeronith/url-shortener:v1.0 .`

To run the container you can use: `docker run --rm -e BASE_URL='http://localhost:8080' -e PORT=8080 -p 8080:8080 xeronith/url-shortener:v1.0`

Note that if you want the container to be persistent throughout multiple sessions, you should drop the `--rm` switch.

To create a shortened URL simply send a POST request to `/api/create` with the desired URL in the body of the request as plain text. For example: `curl -X POST -d "http://somedomain.com/a/long/path?with=parameters&and=more_parameters" http://localhost:8080/api/create`