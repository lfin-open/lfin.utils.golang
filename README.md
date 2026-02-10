# lfin.utils.golang

[![DeepSource](https://deepsource.io/gh/lfin-open/lfin.utils.golang.svg/?label=active+issues&show_trend=true&token=X_N6yHamPQT9n3u955Ey4KPY)](https://deepsource.io/gh/lfin-open/lfin.utils.golang/?ref=repository-badge)
[![Go](https://github.com/lfin-open/lfin.utils.golang/actions/workflows/go_test_report_deepsource.yml/badge.svg)](https://github.com/lfin-open/lfin.utils.golang/actions/workflows/go_test_report_deepsource.yml)
[![Go Version](https://img.shields.io/badge/Go-1.25+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](LICENSE)

A comprehensive collection of utility packages for Go development, providing commonly-used functions for cryptography, date/time handling, string manipulation, and more.

## 📦 Installation

```bash
go get github.com/lfin-open/lfin.utils.golang
```

## 🚀 Quick Start

```go
import (
    "github.com/lfin-open/lfin.utils.golang/uuid"
    "github.com/lfin-open/lfin.utils.golang/date"
    "github.com/lfin-open/lfin.utils.golang/strings"
)

func main() {
    // Generate a time-ordered UUID v7
    id := uuid.GenerateUUIDv7()
    
    // Get current timestamp
    timestamp := date.GetCurrDateTimeYMD24HMMSS()
    
    // Convert to snake_case
    snakeCase := strings.ToSnakeCase("HelloWorld")
}
```

## 📚 Packages

### 🔐 Crypto

Cryptographic utilities for encryption, hashing, and secure random generation.

#### **AES**
- AES-256 encryption and decryption
- Secure key generation and management

#### **AWS KMS**
- Integration with AWS Key Management Service
- Encrypt/decrypt data using AWS KMS keys
- Secure key rotation support

#### **Hash**
- SHA-256, SHA-512 hashing
- HMAC generation and verification
- Secure password hashing

#### **Password**
- Bcrypt password hashing
- Password strength validation
- Secure password comparison

#### **Random**
- Cryptographically secure random number generation
- Random string generation
- UUID generation

**Example:**
```go
import "github.com/lfin-open/lfin.utils.golang/crypto/password"

// Hash a password
hashedPassword := password.HashPassword("mySecurePassword")

// Verify password
isValid := password.VerifyPassword("mySecurePassword", hashedPassword)
```

---

### 📅 Date

Date and time utilities with multiple format support.

**Functions:**
- `GetCurrDateTimeYMD24HMMSSdot6Z()` - Current time with microseconds (e.g., `2022-01-07 14:56:44.545722 KST`)
- `GetCurrDateTimeYMD24HMMSS()` - Current time (e.g., `2022-01-07 14:56:44`)
- `GetCurrDateTimeYMD24HMMSSNoDiv()` - Compact format (e.g., `20220107145644`)
- `GetCurrentUnixTimestampSec()` - Unix timestamp in seconds
- `GetCurrentUnixTimestampMill()` - Unix timestamp in milliseconds
- `GetCurrentUnixTimestampNano()` - Unix timestamp in nanoseconds
- `ConvertTimeYMD24HMMSSdot6Z(time.Time)` - Convert time.Time to formatted string

**Example:**
```go
import "github.com/lfin-open/lfin.utils.golang/date"

// Get current timestamp
now := date.GetCurrDateTimeYMD24HMMSS()
// Output: "2022-01-07 14:56:44"

// Get Unix timestamp in milliseconds
timestamp := date.GetCurrentUnixTimestampMill()
```

---

### 🔤 Encoding

Base64 encoding and decoding utilities.

**Functions:**
- Base64 encoding with standard and URL-safe alphabets
- Automatic padding handling
- Binary data encoding/decoding

**Example:**
```go
import "github.com/lfin-open/lfin.utils.golang/encoding"

encoded := encoding.Base64Encode([]byte("Hello, World!"))
decoded := encoding.Base64Decode(encoded)
```

---

### 🌍 Environment

Environment variable management with type-safe getters.

**Functions:**
- Get environment variables with default values
- Type conversion helpers
- Environment validation

**Example:**
```go
import "github.com/lfin-open/lfin.utils.golang/env"

// Get environment variable with default
port := env.GetEnvWithDefault("PORT", "8080")
```

---

### 📁 File

File system utilities for common operations.

**Functions:**
- Check file/directory existence
- Create files and directories
- Read and write operations
- Path manipulation

**Example:**
```go
import "github.com/lfin-open/lfin.utils.golang/file"

// Check if file exists
exists := file.FileExists("/path/to/file.txt")

// Create file if not exists
file.CreateFileIfNotExists("/path/to/file.txt")
```

---

### 📍 Location

Geographic coordinate utilities and validation.

**Functions:**
- Validate latitude/longitude coordinates
- Convert between coordinate formats (DDM ↔ DM)
- Distance calculations
- Coordinate normalization

**Example:**
```go
import "github.com/lfin-open/lfin.utils.golang/location"

// Validate coordinates
isValid := location.ValidateCoordinates(37.5665, 126.9780)

// Convert coordinate formats
dm := location.ConvertDDMtoDM(37.5665)
```

---

### 🎭 Masking

Data masking utilities for privacy and security.

**Functions:**
- Email masking (e.g., `user@example.com` → `u***@example.com`)
- Phone number masking
- Credit card masking
- Custom pattern masking

**Example:**
```go
import "github.com/lfin-open/lfin.utils.golang/masking"

// Mask email address
masked := masking.MaskEmail("user@example.com")
// Output: "u***@example.com"

// Mask phone number
maskedPhone := masking.MaskPhone("010-1234-5678")
```

---

### 🌐 Net

Network utilities for hostname and MAC address operations.

**Functions:**
- Get hostname
- MAC address retrieval and validation
- Network interface information
- IP address utilities

**Example:**
```go
import "github.com/lfin-open/lfin.utils.golang/net"

// Get hostname
hostname := net.GetHostname()

// Get MAC address
macAddr := net.GetMACAddress()
```

---

### 📄 Paging

Pagination calculation utilities for APIs and databases.

**Functions:**
- Calculate offset and limit
- Total pages calculation
- Page number validation
- Pagination metadata generation

**Example:**
```go
import "github.com/lfin-open/lfin.utils.golang/paging"

// Calculate pagination
pagination := paging.Calc(pageNo, pageSize, totalRecords)
// Returns: Pagination{Offset, CurrentPage, TotalPages, Size, Total}
```

---

### 🔍 Reflection

Go reflection utilities for runtime type inspection.

**Functions:**
- Type inspection
- Struct field iteration
- Dynamic method invocation
- Tag parsing

---

### 📝 Strings

String manipulation and validation utilities.

**Functions:**
- `IsEmptyString(s string)` - Check if string is empty or whitespace
- `I64ToS(i64 int64)` - Convert int64 to string
- `ToSnakeCase(str string)` - Convert to snake_case (e.g., `HelloWorld` → `hello_world`)

**Example:**
```go
import "github.com/lfin-open/lfin.utils.golang/strings"

// Check empty string
isEmpty := strings.IsEmptyString("  ")  // true

// Convert to snake_case
snake := strings.ToSnakeCase("HelloWorld")
// Output: "hello_world"

// Convert int64 to string
str := strings.I64ToS(12345)
```

---

### 🆔 UUID

UUID generation utilities with multiple formats.

**Functions:**
- `GenerateUUID()` - Standard UUID v4 with hyphens (36 characters)
  - Format: `xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx`
- `GenerateUUIDv7()` - Time-ordered UUID v7 without hyphens (32 characters)
  - Format: `xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx`
  - Sequential ordering guaranteed
  - Optimized for database indexing
- `GenerateShortUID8()` - Short 8-character unique identifier
  - Format: `xxxxxxxx`
  - Compact unique IDs for URLs or short references

**Example:**
```go
import "github.com/lfin-open/lfin.utils.golang/uuid"

// Standard UUID v4
id := uuid.GenerateUUID()
// Output: "550e8400-e29b-41d4-a716-446655440000"

// Time-ordered UUID v7 (sequential)
v7id := uuid.GenerateUUIDv7()
// Output: "019c455dfe9f726db72b5ec2a91a0ebc"

// Short 8-character UID
shortId := uuid.GenerateShortUID8()
// Output: "2e31c3df"
```

## 🧪 Testing

Run all tests:
```bash
go test ./...
```

Run tests with coverage:
```bash
go test -cover ./...
```

Run tests for a specific package:
```bash
go test ./uuid -v
```

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'feat: add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

### Commit Convention

We follow [Conventional Commits](https://www.conventionalcommits.org/):

- `feat:` - New features
- `fix:` - Bug fixes
- `docs:` - Documentation changes
- `test:` - Test additions or modifications
- `refactor:` - Code refactoring
- `chore:` - Maintenance tasks

## 📄 License

Copyright (c) 2022 LFin and others. All rights reserved.

## 👥 Contributors

- Ted, KIM

## 🔗 Links

- [GitHub Repository](https://github.com/lfin-open/lfin.utils.golang)
- [Issue Tracker](https://github.com/lfin-open/lfin.utils.golang/issues)
- [DeepSource Analysis](https://deepsource.io/gh/lfin-open/lfin.utils.golang)

## 📊 Project Status

This project is actively maintained and used in production environments. We welcome feedback and contributions from the community.