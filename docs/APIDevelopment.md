<div id="top-header" style="with:100%;height:auto;text-align:right;">
    <img src="./images/pr-banner-long.png">
</div>

# WORKTIME CONTROLLER - GOLANG 1.25

- [/README.md](../README.md)
<br><br>

# API Development

If you are using the https://github.com/pabloripoll/platform-docker-nginx-go-1.25 repository, you may notice that `./platform/nginx-go-1.25/docker/config/supervisor/conf.d-sample` you have two different Supervisor service configurations. Use the following one for development:
```sh
[program:go]
command=air -c .air.toml # or: go run cmd/api/main.go
stdout_logfile=/dev/stdout
stdout_logfile_maxbytes=0
stderr_logfile=/dev/stderr
stderr_logfile_maxbytes=0
autorestart=true
startretries=3
```
<br>

## API start up

Once the platform is installed and the REST API container is running, you can execute the initialization commands to start the application

1. Access into container
```bash
# Local machine
$ make apirest-ssh

# Docker container
/var/www $
```

2. Initialize module *(if not done before)*
```bash
/var/www $ go mod init
```

3. Install the application Tidy dependencies *(on every container built)*
```bash
/var/www $ go mod tidy
```

4. Run the app or as supervisor is already running, you can exit the container and restart it
```bash
/var/www $ go run cmd/api/main.go
```
<br><br>

## API Improvement

If API is keeping live with AIR,

On the other hand, you may prefer to run `main.go` without building step
```bash
# Clean Go cache
/var/www $ go clean -cache

# Verify module
/var/www $ go mod tidy

# Run
/var/www $ go run cmd/api/main.go
```

Deeply cache clearing
```bash
# Remove all build artifacts
/var/www $ go clean -cache -modcache -i -r

# Verify go.mod
/var/www $ go mod tidy

# Run
/var/www $ go run cmd/api/main.go
```

#### Container re-build

On every container re-build `go mod tidy` command must be executed.
<br>

#### Database Migrations

There is a function using GORM's AutoMigrate() but commented because is great for prototyping, is terrible for production as it doesn't drop columns, can't easily handle data transformations, and leaves no history.

To get framework-style migrations in Go, the industry standard is golang-migrate/migrate.

Step 1: Install the CLI tool (on your local machine)

The github.com/golang-migrate/migrate/v4/cmd/migrate@latest is a standalone program for your terminal. You install it globally on your machine, not into your project. Run this once on your machine/container:
```bash
# On Mac/Linux
/var/www $ go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

Step 2: Create your first migration files Run this command in your project root:
```bash
# On Mac/Linux
/var/www $ migrate create -ext sql -dir ./migrations create_example_table
```

This generates two files in your migrations/ folder:
```bash
20240211221748_create_example_table.up.sql # What runs when you migrate
20240211221748_create_example_table.down.sql # What runs when you rollback
```

For a better production-grade approach, there is a command to run migrations
```bash
# To run migrations forward
/var/www $ go run ./cmd/db-migrate/main.go -up

# To undo the last migration
/var/www $ go run ./cmd/db-migrate/main.go -down
```
<br>

## Build and run the API

Try building
```sh
/var/www $ go build ./cmd/api
```
If output is as follows
```sh
error obtaining VCS status: exit status 128
	Use -buildvcs=false to disable VCS stamping.
```
This is a Git-related warning during build. You can safely ignore it or disable it. The error is just about Git version control metadata - it won't affect your application. The build should still succeed. Here are the options:

Option 1: Disable VCS stamping (Quick fix)
```sh
/var/www $ go build -buildvcs=false -o bin/api ./cmd/api
```

Option 2: Fix Git issues (if any) add and commit changes. Then build normally
```sh
/var/www $ go build ./cmd/api
```

Option 3: Set environment variable permanently
```sh
# Add to your ~/.bashrc or ~/.zshrc
export GOFLAGS="-buildvcs=false"

# Or set it for current session
export GOFLAGS="-buildvcs=false"
```

Run the API
```sh
/var/www $ ./bin/api
```
<br>

## API Deployment

Update Dockerfile for Production/Staging You don't need to install air in your staging Dockerfile. Instead, compile the app during the Docker build process.

Build the GO Application
```sh
/var/www $ RUN go build -buildvcs=false -ldflags="-w -s" -o /var/www/bin/api ./cmd/api
```

*(Note: -ldflags="-w -s" strips debugging info, making your binary much smaller and faster).*

<!-- FOOTER -->
<br>

---

<br>

- [GO TOP ⮙](#top-header)

<div style="with:100%;height:auto;text-align:right;">
    <img src="./images/pr-banner-long.png">
</div>