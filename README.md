# MyFTP Server

A lightweight, secure FTP server implementation written in Go that supports standard FTP protocol commands with built-in security features.

## Features

- **Full FTP Protocol Support** - Implements core FTP commands for file transfer and navigation
- **Dual Transfer Modes** - Supports both Active (PORT) and Passive (PASV) data transfer modes
- **User Authentication** - Secure login system with username and password verification
- **Directory Isolation** - Sandboxed file access prevents directory traversal attacks
- **Concurrent Connections** - Multi-client support with goroutine-based connection handling
- **Standards Compliant** - RFC 959 compatible responses and behavior

## Supported Commands

| Command | Description |
|---------|-------------|
| `USER` | Specify username for authentication |
| `PASS` | Provide password for authentication |
| `SYST` | Display system type |
| `FEAT` | List server features |
| `PWD` | Print working directory |
| `CWD` | Change working directory |
| `CDUP` | Change to parent directory |
| `LIST` | List files in current directory |
| `RETR` | Download a file from server |
| `DELE` | Delete a file on server |
| `PASV` | Enter passive mode |
| `PORT` | Specify client data port (active mode) |
| `HELP` | Display available commands |
| `QUIT` | Close connection |

## Requirements

- Go 1.22.2 or higher

## Installation

```bash
# Clone the repository
git clone <repository-url>
cd FTP

# Build the server
go build -o myftp
```

## Usage

```bash
# Run the server
./myftp <port> <path>

# Example
./myftp 2121 /home/user/ftp-files
```

**Parameters:**
- `port` - Port number for the FTP server to listen on
- `path` - Root directory for FTP file access (anonymous are sandboxed to this directory)

## Connecting to the Server

You can connect using any standard FTP client:

```bash
# Using command-line FTP client
ftp localhost 2121

# Using FileZilla, WinSCP, or other GUI clients
# Host: localhost
# Port: 2121
# Protocol: FTP
```