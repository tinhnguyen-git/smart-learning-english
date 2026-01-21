# API Reference
Base URL: `http://localhost:8080`

## Authentication

### Register
Create a new user account.

- **URL**: `/auth/register`
- **Method**: `POST`
- **Auth Required**: No
- **Body**:
  ```json
  {
    "email": "user@example.com",
    "password": "securepassword",
    "full_name": "John Doe"
  }
  ```
- **Response** (`201 Created`):
  ```json
  {
    "token": "jwt_token_here",
    "user": {
      "id": "uuid",
      "email": "user@example.com",
      "full_name": "John Doe"
    }
  }
  ```

### Login
Authenticate an existing user.

- **URL**: `/auth/login`
- **Method**: `POST`
- **Auth Required**: No
- **Body**:
  ```json
  {
    "email": "user@example.com",
    "password": "securepassword"
  }
  ```
- **Response** (`200 OK`):
  ```json
  {
    "token": "jwt_token_here",
    "user": {
      "id": "uuid",
      "email": "user@example.com",
      "full_name": "John Doe"
    }
  }
  ```

---

## Stories

### Scrape Story
Trigger the scraper to fetch a story from a URL and save it to the database.

- **URL**: `/stories/scrape`
- **Method**: `POST`
- **Auth Required**: No (Ideally Admin only, but currently public)
- **Body**:
  ```json
  {
    "url": "https://example.com/story-page"
  }
  ```
- **Response** (`201 Created`):
  ```json
  {
    "id": "uuid",
    "title": "Extracted Title",
    "content": "Full extracted content...",
    "source_url": "https://example.com/story-page",
    "created_at": "timestamp"
  }
  ```

---

## Users

### Upgrade to Premium
Upgrade the current user to premium status.

- **URL**: `/users/upgrade`
- **Method**: `POST`
- **Auth Required**: Yes (Bearer Token)
- **Headers**:
  - `Authorization`: `Bearer <jwt_token>`
- **Body**: Empty
- **Response** (`200 OK`):
  ```json
  {
    "message": "User upgraded to premium successfully"
  }
  ```
