# Keep Secrets

A secure service for one-time secret sharing with client-side encryption.

## Architecture

- **Client**: Nuxt 3 + Vue 3 + Nuxt UI 2 + Tailwind + TypeScript + TanStack Vue Query + Axios 
- **Server**: Go 1.21 + SQLite + Gorilla Mux

## Features

- **End-to-end encryption** (AES-256-GCM)
- Encryption key is passed in URL fragment (never sent to server)
- **One-time use** - secrets are automatically deleted after viewing
- Modern UI with dark mode support
- Responsive design
- SEO optimized

## Installation & Setup

### Docker Compose (Recommended)

The easiest way to run the entire project with a single command:

```bash
docker compose up --build
```

This will build and run both services:

- **Server** will be available at `http://localhost:8080`
- **Client** will be available at `http://localhost:3000`

To stop:

```bash
docker compose down
```

To run in background:

```bash
docker compose up -d --build
```

### Local Development

#### Server (Go)

```bash
cd server
go mod download
go run main.go
```

The server will start on port 8080 (or the port specified in the `PORT` environment variable).

#### Client (Nuxt)

```bash
cd client
pnpm install
pnpm dev
```

The client will start on port 3000 (or another port if 3000 is busy).

## Usage

1. Open the client in your browser
2. Enter your secret in the text field
3. Click "Create secret link"
4. Copy the link and share it
5. When the link is opened, the secret will be displayed once and then deleted

## Security

- **Client-side encryption** using Web Crypto API (AES-256-GCM)
- **Encryption key never sent to server** - it's passed in the URL fragment (`#key=...`)
- **One-time use** - secrets are automatically deleted after viewing
- **No server-side decryption** - server only stores encrypted data
