# TrackPixel

## About The Project

A GoLang REST API server for creating and tracking email pixel tracker images. 

### Built With
* [Go](https://go.dev/)
* [Gin](https://github.com/gin-gonic/gin)
* [QuestDB](https://questdb.io/)

## Getting Started

### Prerequisites
* Go
* QuestDB
* PostgreSQL (To connect to QuestDB using PSQL wire protocol)
* Git

### Installation

1. Clone the repo
```sh
git clone https://github.com/colt005/TrackPixel.git
```
2. Download and install [Quest DB](https://questdb.io/get-questdb/), then run it locally or deploy it on a server. 

3. Set up environment variables
```sh
export PORT=<Port you want to run the server on>
export BASE_URL=<Base url of the server use `http://localhost$PORT` if running locally>
export CONN_STRING=<Connection url of the PSQL wire protocol which questDB exposes. If the QuestDB is running locally it will be `postgresql://admin:quest@localhost:8812/qdb` >
```

4. Build binary
```go
go build -o TrackPixel
```
5. Use the binary to run local server
```sh
./TrackPixel
```

## API Endpoints


### GET /v1/api/status
Check if the API is running.

**Response**
```json
{
  "data": "Pixel tracker API running smoothly",
  "status": "success"
}
```

### GET /v1/api/gen-image
Generate an image for the user to incorporate in an email and track email statistics.

**Response**
```json
{
  "image_url": "http://localhost:8000/image/3DjwvFjAjG8uZOPN.png",
  "unique_code": "3DjwvFjAjG8uZOPN"
}
```
This produced `image url` can be pasted into an email message. 

### GET /v1/api/info/:uniqueCode
Get statistics on the generated image and how it's being used in emails and other places. 

**Path Parameters**
* `:uniqueCode` - The generated unique code


**Response**
```json
{
  "uri": "KgQR3dvAeImUpfkz",
  "total_opens": 4,
  "created_at": "2022-06-16T23:00:21.453445Z",
  "timeseries_data": [
    {
      "uri": "KgQR3dvAeImUpfkz",
      "time_stamp": "2022-06-16T23:04:01.456028Z",
      "ip_address": "127.0.0.1"
    },
    {
      "uri": "KgQR3dvAeImUpfkz",
      "time_stamp": "2022-06-16T23:04:07.951854Z",
      "ip_address": "127.0.0.1"
    },
    {
      "uri": "KgQR3dvAeImUpfkz",
      "time_stamp": "2022-06-16T23:04:09.220944Z",
      "ip_address": "127.0.0.1"
    },
    {
      "uri": "KgQR3dvAeImUpfkz",
      "time_stamp": "2022-06-16T23:04:10.544438Z",
      "ip_address": "127.0.0.1"
    }
  ]
}
```

## Contributing

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## Acknowledgements

This project uses [QuestDB](https://questdb.io/) which is one of the fastest open source time series database. 

## Go Packages used
 * [uniuri](https://github.com/dchest/uniuri)
 * [Gin](https://github.com/gin-gonic/gin)
 * [Logrus](https://github.com/sirupsen/logrus)
 * [pgx](https://github.com/jackc/pgx)
