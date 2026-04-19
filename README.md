# ltail

`ltail` is a fast, multipurpose tool for AWS CloudWatch Logs.

Based on [TylerBrock/saw](https://github.com/TylerBrock/saw), ltail uses AWS SDK v2 and adds new features.

## Features

- Colorized output that can be formatted in various ways
    - `--expand` Explode JSON objects using indenting
    - `--rawString` Print JSON strings instead of escaping ("\n", ...)
    - `--invert` Invert white colors to black for light color schemes
    - `--raw`, or `--pretty`, for `watch` and `get` commands respectively, toggles display of the timestamp and stream name prefix.
    - `--shorten` or `-s` Truncate lines exceeding 512 characters and append "..."

- Automatic log level colorization
    - `INFO` (all caps) appears with green background and black text
    - `ERROR` (all caps) appears with red background and white text

- Dual split-pane view with Bubble Tea TUI
    - `ltail dual log-group-1 log-group-2` Watch two log groups in fullscreen split view
    - Interactive keyboard navigation with vim-style controls
    - Independent filtering and scrolling for each pane

- Filter logs using CloudWatch patterns
    - `--filter foo` Filter logs for the text "foo"

- Watch aggregated interleaved streams across a log group
    - `ltail watch production` Stream logs from production log group
    - `ltail watch production --prefix api` Stream logs from production log group with prefix "api"

## Usage

- Basic
    ```sh
    # Get list of log groups
    ltail groups

    # Get list of streams for production log group
    ltail streams production
    ```

- Watch
    ```sh
    # Watch production log group
    ltail watch production

    # Watch production log group streams for api
    ltail watch production --prefix api

    # Watch production log group streams for api and filter for "error"
    ltail watch production --prefix api --filter error

    # Watch with shortened lines (truncate long lines at 512 chars)
    ltail watch production --shorten
    ```

- Get
    ```sh
    # Get production log group for the last 2 hours
    ltail get production --start -2h

    # Get production log group for the last 2 hours and filter for "error"
    ltail get production --start -2h --filter error

    # Get production log group for api between 26th June 2018 and 28th June 2018
    ltail get production --prefix api --start 2018-06-26 --stop 2018-06-28

    # Get logs with shortened lines (useful for logs with large payloads)
    ltail get production -s --pretty
    ```

- Dual (Split View)
    ```sh
    # Watch two log groups in split-pane view
    ltail dual production-api production-worker

    # Compare errors between two services
    ltail dual service-a service-b --filter1 ERROR --filter2 ERROR

    # Monitor deployment with old and new versions
    ltail dual production production --prefix1 v1.0 --prefix2 v2.0

    # Different filters per pane
    ltail dual api-logs worker-logs --filter1 "POST /api" --filter2 "job completed"

    # Keyboard controls: q=quit, tab=switch pane, ↑↓/jk=scroll, g/G=top/bottom
    ```

### Profile and Region Support

By default ltail uses the region and credentials in your default profile. You can override these to your liking using the command line flags:

```sh
# Use personal profile
ltail groups --profile personal

# Use us-west-1 region
ltail groups --region us-west-1
```

## Installation

### Mac OS X

```sh
brew tap megaproaktiv/ltail
brew install ltail
```

### Manual Install/Update

- [Install go](https://golang.org/doc/install)
- Configure your `GOPATH` and add `$GOPATH/bin` to your path
- Run `go install github.com/megaproaktiv/ltail@latest`

Alternatively you can hard code these in your shell's init scripts (bashrc, zshrc, etc...):

```sh
# Export profile and region that override the default
export AWS_PROFILE='work_profile'
export AWS_REGION='us-west-1'
```

## Run Tests
From root of repository: `go test -v ./...`

# Development

## Release

### 1. Check current version
task version

### 2. Update VERSION file to new version (manually)
echo "v0.3.1" > VERSION

### 3. Test the release build locally
task release-dry-run

### 4. Commit your changes
git add .
git commit -m "Release v0.3.1"

### 5. Create release (automatically tags and releases)
task release
