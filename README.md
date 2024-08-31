# Gratitude Backend

Gratitude Backend is a full-stack application designed to handle post management for a social media platform. It features an API Gateway with rate limiting, IP blacklisting, and Firebase authentication, as well as a gRPC-based post service. The project includes Docker configurations for local development and deployment.

## Features

- **API Gateway**:
  - **Rate Limiting**: Controls the rate of incoming requests to prevent abuse.
  - **IP Blacklisting**: Blocks requests from blacklisted IP addresses.
  - **CORS Middleware**: Manages Cross-Origin Resource Sharing (CORS) policies.
  - **Firebase Authentication**: Validates user authentication tokens.

- **Post Service**:
  - **gRPC Service**: Manages posts using gRPC for efficient communication.
  - **DynamoDB Integration**: Stores and retrieves post data using AWS DynamoDB.

## gRPC Services

The Post Service includes the following gRPC services:

- **PostService**:
  - **CreatePost**: Creates a new post.
  - **GetPost**: Retrieves a post by its ID.
  - **UpdatePost**: Updates an existing post.
  - **DeletePost**: Deletes a post by its ID.
  - **GetPosts**: Retrieves a list of posts based on pagination and filters.

  ## Prerequisites

- Docker and Docker Compose
- Go 1.18 or later
- Firebase Admin SDK credentials

## Getting Started

### Clone the Repository

```bash
git clone https://github.com/your-username/gratitude-backend-go.git
cd gratitude-backend-go
