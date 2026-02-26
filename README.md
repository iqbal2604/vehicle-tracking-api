![CI](https://github.com/iqbal2604/vehicle-tracking-api/actions/workflows/ci.yml/badge.svg)
# Vehicle Tracking API

A robust backend REST API for a vehicle tracking system built with Golang and the Fiber framework. This API enables real-time vehicle monitoring, user authentication, and GPS data management.

## Features

- **User Authentication**: Register and login with JWT-based authentication.
- **User Profiles**: Retrieve user profile information.
- **Vehicle Management**: Create, read, update, and delete vehicles associated with users.
- **GPS Tracking**: Record GPS locations, retrieve history, last location, and stream real-time data.
- **WebSocket Support**: Real-time communication via WebSockets.
- **Secure Endpoints**: JWT middleware for protected routes.
- **Database Integration**: Uses GORM with MySQL for data persistence.

## Tech Stack

- **Language**: Go 1.25.5
- **Framework**: Fiber v2.52.10
- **Database**: MySQL (via GORM) , Redis
- **Authentication**: JWT (golang-jwt/jwt/v5)
- **WebSockets**: Fiber WebSocket v2.2.1
- **Dependency Injection**: Google Wire
- **Testing**: Testify

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/vehicle-tracking-api.git
   cd vehicle-tracking-api
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Set up environment variables:
   Create a `.env` file or set environment variables for:
   - Database connection (e.g., `DB_HOST`, `DB_USER`, `DB_PASSWORD`, `DB_NAME`)
   - JWT secret: `JWT_SECRET`
   - Redis connection (e.g., `REDIS_HOST`, `REDIS_PORT`, `REDIS_PASSWORD`, `REDIS_DB`)

4. Run the application:
   ```bash
   go run ./cmd
   ```

   The server will start on `http://localhost:3000`.

## Usage

### API Endpoints

The API provides the following endpoints:

- **Authentication**:
  - `POST /api/register` - Register a new user
  - `POST /api/login` - Login user

- **User**:
  - `GET /api/profile` - Get user profile (requires JWT)

- **Vehicles**:
  - `POST /api/vehicles` - Create a vehicle (requires JWT)
  - `GET /api/vehicles` - List user's vehicles (requires JWT)
  - `GET /api/vehicles/{id}` - Get vehicle by ID (requires JWT)
  - `PUT /api/vehicles/{id}` - Update vehicle (requires JWT)
  - `DELETE /api/vehicles/{id}` - Delete vehicle (requires JWT)

- **GPS**:
  - `POST /api/gps` - Create GPS location (requires JWT)
  - `GET /api/gps/history/{vehicle_id}` - Get GPS history (requires JWT)
  - `GET /api/gps/last/{vehicle_id}` - Get last GPS location (requires JWT)
  - `GET /api/gps/stream/{vehicle_id}` - Stream GPS locations (requires JWT)

- **WebSocket**:
  - `GET /ws` - WebSocket connection

### Testing

Use the provided `test.http` file with a REST client like VS Code's REST Client extension or Postman to test the endpoints.

## API Documentation

The full API specification is available in `openapi.yaml`. You can view it using tools like Swagger UI.

## Project Structure

```
.
├── cmd/
│   └── main.go                 # Application entry point
├── config/
│   └── database.go             # Database configuration
├── handlers/                   # HTTP handlers
├── helpers/                    # Utility functions
├── middlewares/                # Middleware functions
├── models/                     # Database models
├── repositories/               # Data access layer
├── requests/                   # Request structs
├── routes/                     # Route definitions
├── services/                   # Business logic
├── dtos/                       # Data transfer objects
├── websocket/                  # WebSocket handlers
├── test.http                   # API test file
├── openapi.yaml                # API specification
└── README.md
```

## Contributing

1. Fork the repository.
2. Create a feature branch: `git checkout -b feature-name`
3. Commit your changes: `git commit -am 'Add some feature'`
4. Push to the branch: `git push origin feature-name`
5. Submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
