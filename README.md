# strongify-passgen-go-echo
Here's a README file template for your Strongify Password Generator project. This template includes sections commonly found in README files, such as a project description, setup instructions, usage, and more. You can customize it further based on your specific project details.

---

# Strongify Password Generator

Strongify Password Generator is a secure password generation service that allows users to save custom phrases (collections of words) and generate passwords based on these phrases. The passwords can include user-defined lengths, minimum numbers of numbers, symbols, and a secret from an environment variable. This project is developed using Go, Echo, MySQL, and JWT for authentication.

## Table of Contents

- [Features](#features)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Configuration](#configuration)
- [Usage](#usage)
  - [API Endpoints](#api-endpoints)
- [Contributing](#contributing)
- [License](#license)

## Features

- User authentication with JWT
- Save custom phrases (collections of words)
- Generate secure passwords based on saved phrases
- Customizable password length, minimum number of numbers, and symbols
- Secure storage of secrets using environment variables

## Getting Started

### Prerequisites

- Go (version 1.15 or later)
- MySQL
- Git

### Installation

1. **Clone the repository:**

   ```sh
   git clone https://github.com/your-username/strongify-password-gen.git
   cd strongify-password-gen
   ```

2. **Install dependencies:**

   ```sh
   go mod tidy
   ```

3. **Set up the database:**

   Create a MySQL database and update the connection details in the `.env` file.

   ```sql
   CREATE DATABASE strongify;
   ```

4. **Run database migrations:**

   Use a migration tool or the provided SQL scripts to set up the necessary tables.

5. **Build and run the application:**

   ```sh
   go build
   ./strongify-password-gen
   ```

### Configuration

Create a `.env` file in the project root directory and add the following environment variables:

```env
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_HOST=localhost
DB_PORT=3306
DB_NAME=strongify
JWT_SECRET=your_jwt_secret
APP_SECRET=your_app_secret
```

## Usage

### API Endpoints

#### Authentication

- **Login**
  - `POST /login`
  - Request Body: `{ "username": "user", "password": "pass" }`
  - Response: `{ "token": "jwt_token" }`

- **Register**
  - `POST /register`
  - Request Body: `{ "username": "user", "password": "pass" }`
  - Response: `{ "message": "User registered successfully" }`

#### Phrases

- **Add Phrase**
  - `POST /phrases`
  - Headers: `{ "Authorization": "Bearer jwt_token" }`
  - Request Body: `{ "phrase": "your custom phrase" }`
  - Response: `{ "message": "Phrase added successfully" }`

- **Get Phrases**
  - `GET /phrases`
  - Headers: `{ "Authorization": "Bearer jwt_token" }`
  - Response: `{ "phrases": ["phrase1", "phrase2"] }`

#### Password Generation

- **Generate Password**
  - `POST /generate`
  - Headers: `{ "Authorization": "Bearer jwt_token" }`
  - Request Body:
    ```json
    {
      "phrase_id": 1,
      "length": 16,
      "min_numbers": 2,
      "min_symbols": 2
    }
    ```
  - Response: `{ "password": "generated_password" }`

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

Replace the placeholder URLs, environment variable names, and any specific details as necessary. This README file should provide a clear overview of your project and how to get it up and running.