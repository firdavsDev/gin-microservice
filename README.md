# Gin microservice for PDF generation

## Description
This is a simple microservice that generates a PDF file using the gofpdf library. The microservice uses the Gin framework to handle the HTTP requests and the GORM library to interact with the PostgreSQL database.

## Installation

```bash
git clone github.com/firdavsdev/gin-microservice
go mod tidy
go mod init
```

## Running the microservice
To run the microservice, you need to have a PostgreSQL database running. You can use the following command to start a PostgreSQL database using Docker:

```bash
go run main.go
// or
go build -o gin-microservice
./gin-microservice
// or
docker-compose up --build --remove-orphans
```

## API Endpoints
The microservice has the following API endpoints:

- `POST /api/v1/generate-pdf?name=save || download`: This endpoint generates a PDF file with the specified file name and saves it in the `pdfs` directory. If the `name` query parameter is set to `download`, the PDF file is downloaded by the client.
- `GET /api/v1/get-pdfs`: This endpoint returns a list of all the PDF files in the `pdfs` directory.
- `GET /api/v1/get-pdf?file=file_name`: This endpoint generates a PDF file with the specified file name. The file name should be a string without spaces. The generated PDF file is saved in the `pdfs` directory.

