## Description

Analysis and process 10000+ rows of data with goroutine

## Tech Stack

1. Golang
2. Postgres
3. Gin Framework
4. Gorm DATABASE ORM

## Getting Started

1. Clone the repository:

    ```bash
    git clone https://github.com/egrizq/product-analysis-with-concurrency.git
    ```

2. Install all package:

    ```bash
    go mod tidy
    ```

3. Copy environment for database:

    ```bash
    cp .env.example .env
    ```

4. Run the application:

    ```bash
    make dev
    ```
    or alternatively 
    ```bash
    go run main.go
    ```

5. The program will run on [http://localhost:8000](http://localhost:8000).

## API ENDPOINT

| Method | Endpoint | Description |
|----------|----------|----------|
| POST | localhost:8000/process/insert | Import JSON & CSV file |
| GET | localhost:8000/process/reports | Process & analysis |
