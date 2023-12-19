# jwt/basic

Basic usage of JWT in Go.

## How to Run

The server can be started simply by:

```bash
# Install dependencies
make deps

# Run the backend
make
```

To test the endpoints use the provided shell script `test`. For example:

```bash
# Create a new admin
./test signup-admin

# Sign-in the admin
./test singin-admin

# Request an admin-only route
./test admin
```

Available commands for the test script can be viewed by:

```bash
./test help
```

