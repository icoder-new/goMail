# goMail

goMail is a simple REST API built in Golang that allows you to send SMS via email. It utilizes a combination of technologies including Golang, Gin, Lumberjack v2, SMTP, GORM, GORM Postgres driver, PostgreSQL, godotenv, and viper.

## Table of Contents
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Configuration](#configuration)
- [Usage](#usage)
- [Endpoints](#endpoints)
- [Example](#example)
- [Contributing](#contributing)
- [License](#license)

## Prerequisites

Before you begin, ensure you have met the following requirements:

- Go (Golang) installed. You can download it from the [official website](https://golang.org/dl/).
- PostgreSQL database with necessary configurations.
- SMTP server details to send emails.
- Basic understanding of the stack used in this project.

## Installation

1. Clone this repository:

   ```bash
   git clone github.com/icoder-new/goMail
   cd goMail
   ```

2. Install the required dependencies:

   ```bash
   go mod tidy
   ```

## Configuration

1. Create a `.env` file in the project root directory and configure the following environment variables:

   ```
   GIN_MODE=debug

   DB_PASSWORD=database_password
   SMTP_PASSWORD=smtp_password
   ```

2. Modify the SMTP, database, and other settings as per your requirements.

## Usage

To start the goMail REST API server, run the following command:

```bash
go run ./cmd
```

The API server will start on the specified port (default is 3333).

## Endpoints

The goMail REST API provides the following endpoints:

- `GET /` - A simple endpoint to check if the server is running.
- `GET /api/v1` - A test endpoint for the POST method to send SMS via email.
- `POST /api/v1/send` - Send an SMS via email. You need to provide a JSON request body with the following fields:
  - `firstname` (string): First name of the recipient.
  - `lastname` (string): Last name of the recipient.
  - `email` (string): Email address of the recipient.
  - `subject` (string): Subject of the email.
  - `message` (string): Message to send via email.

## Example

Here's an example of how to send an SMS via email using the API:

```bash
curl -X POST http://localhost:3333/api/v1/send -H "Content-Type: application/json" -d '{
    "firstname": "John",
    "lastname": "Doe",
    "email": "johndoe@example.com",
    "subject": "Hello from goMail",
    "message": "This is a test message sent via goMail."
}'
```

## Contributing

Contributions are welcome! If you'd like to contribute to this project, please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes and submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.