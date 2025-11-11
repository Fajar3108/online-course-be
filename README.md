# Production-Ready Go Boilerplate (Fiber + GORM)

## 🛠️ Tech Stack

  * **Language:** Go (Golang)
  * **Framework:** [Fiber v2](https://gofiber.io/)
  * **Database:** [MySQL](https://www.mysql.com/)
  * **ORM:** [GORM](https://gorm.io/)
  * **Config:** [Viper](https://github.com/spf13/viper) (reading from `.env`)
  * **Authentication:** JWT (`golang-jwt/v5`) with Refresh Token Rotation
  * **Validation:** [Go Playground Validator v10](https://github.com/go-playground/validator)
  * **Email:** [Gomail v2](https://github.com/go-gomail/gomail)

## 📂 Project Structure

```
.
├── cmd/
│   ├── api/main.go         # API application entry point
│   └── migration/main.go   # GORM AutoMigrate script
├── config/                 # Viper setup & .env constants
├── database/               # Thread-safe GORM singleton initializer
├── internal/
│   ├── action/             # Atomic, reusable business logic (e.g., CreateNewUser)
│   ├── controller/         # HTTP (Fiber) layer. Only part that knows about HTTP.
│   ├── request/            # Structs for request body validation.
│   ├── resource/           # Structs for response transformation (DTOs).
│   ├── router/             # Fiber route definitions (Auth, Category, etc).
│   └── service/            # Business logic orchestrators (e.g., AuthService).
├── pkg/
│   ├── error-handler/      # Custom error definitions & global error handler.
│   ├── file-storage/       # Framework-agnostic file upload logic.
│   ├── helpers/            # Utility functions (UUID, slug, response).
│   ├── mail/               # Email sending helper.
│   ├── middleware/         # Custom middleware (e.g., JWT).
│   ├── model/              # GORM data entities (User, Category, etc).
│   ├── token/              # JWT generation & parsing.
│   └── validation/         # Thread-safe Validator singleton.
├── storage/                # (gitignored) Uploaded files are stored here.
├── .env.example            # Configuration file template.
├── go.mod                  # Go dependencies.
└── makefile                # Helper commands (serve, migrate).
```

## 🏁 Getting Started

**Prerequisites:**

  * Go 1.21+
  * MySQL

**Steps:**

1.  **Clone the Repository:**

    ```bash
    git clone https://github.com/Fajar3108/golang-boilerplate.git
    cd golang-boilerplate
    ```

2.  **Create Config File:**
    Copy the example `.env.example` to `.env`.

    ```bash
    cp .env.example .env
    ```

3.  **Set Environment Variables:**
    Open the `.env` file and fill in all required variables, especially:

      * `DB_HOST`, `DB_USER`, `DB_PASSWORD`, `DB_NAME`
      * `JWT_SECRET_KEY` (generate a strong random string)

4.  **Install Dependencies:**

    ```bash
    go mod tidy
    ```

5.  **Run Database Migrations:**
    This command will create the `users`, `categories`, and `user_sessions` tables in your database.

    ```bash
    make migrate
    ```

6.  **Create Storage Symlink:**
    (Optional, if you want to serve files from `storage/public`)

    ```bash
    make storage-link
    ```

7.  **Run the Server:**

    ```bash
    make serve
    ```

    The server will be running at `http://localhost:8080` (or the port you defined in `.env`).

## ⚙️ Environment Variables

The application is configured using the following variables in `.env`:

```ini
# Secret key for signing JWTs
JWT_SECRET_KEY=your_jwt_secret_key
# Access token expiration (in hours)
JWT_EXPIRATION_HOURS=24
# Refresh token expiration (in days)
JWT_REFRESH_EXPIRATION_DAYS=14

# Database Settings
DB_HOST=localhost
DB_PORT=3306
DB_NAME=your_database_name
DB_USER=your_database_user
DB_PASSWORD=your_database_password

# Server Port
APP_PORT=8080

# Key for encrypting cookies
COOKIE_SECRET_KEY="k1Gt+G9LHSdLsPK7SAnpF8kZdFf3pzPzlAdFlvziz0s="

# Email (SMTP) Server Settings
MAIL_HOST="sandbox.smtp.mailtrap.io"
MAIL_PORT=587
MAIL_SENDER="service@mail.com"
MAIL_USERNAME="95f5e242a0c421s"
MAIL_PASSWORD="ec3b9bfd27e252s"
```

## 🗺️ API Endpoints (Included)

The following endpoints are provided out-of-the-box:

### Authentication (`/api/auth`)

| Method | Endpoint | Description | Auth? |
| --- | --- | --- | --- |
| `POST` | `/register` | Register a new user (handles `avatar` via form-data). | No |
| `POST` | `/login` | Log in a user and receive JWT tokens. | No |
| `PUT` | `/refresh-token` | Get a new token pair using a refresh token (with rotation). | No |
| `DELETE` | `/logout` | Revoke the current session. | **Yes (JWT)** |

### Categories (`/api/categories`)

| Method | Endpoint | Description | Auth? |
| --- | --- | --- | --- |
| `GET` | `/` | Get a paginated list of all categories (`?page=1&limit=10`). | **Yes (JWT)** |
| `POST` | `/` | Create a new category. | **Yes (JWT)** |
| `GET` | `/:slug` | Get a single category by its slug. | **Yes (JWT)** |
| `PATCH` | `/:slug` | Update a category by its slug. | **Yes (JWT)** |
| `DELETE` | `/:slug` | Delete a category by its slug. | **Yes (JWT)** |