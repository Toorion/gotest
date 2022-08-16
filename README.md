Test https microservice parallel url loader

Run: go run .

Make http request:

curl -X POST --data-binary "@./links.txt" localhost:8888

Sime links from linx.txt file can return 503 instead of 200


