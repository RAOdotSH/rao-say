# rao-says

`rao-says` is a command-line tool that generates random ASCII bots, inspired by `cowsay`.

## Installation

### Using Precompiled Binaries

Download the latest release from the [GitHub Releases](https://github.com/your-username/rao-says/releases) page. Extract the archive and move the binary to a directory in your `PATH`.

#### Linux

1. Download the `tar.gz` or `zip` file for your architecture.
2. Extract the archive.
    ```sh
    tar -xzvf rao-says_v1.0.0_Linux_x86_64.tar.gz
    # or
    unzip rao-says_v1.0.0_Linux_x86_64.zip
    ```
3. Move the binary to `/usr/local/bin` or another directory in your `PATH`.
    ```sh
    sudo mv rao-says /usr/local/bin/
    ```

#### macOS

1. Download the `tar.gz` or `zip` file for your architecture.
2. Extract the archive.
    ```sh
    tar -xzvf rao-says_v1.0.0_Darwin_x86_64.tar.gz
    # or
    unzip rao-says_v1.0.0_Darwin_x86_64.zip
    ```
3. Move the binary to `/usr/local/bin` or another directory in your `PATH`.
    ```sh
    sudo mv rao-says /usr/local/bin/
    ```

#### Windows

1. Download the `zip` file for your architecture.
2. Extract the zip file using your preferred tool.
3. Move `rao-says.exe` to a directory in your `PATH`.

### Using `go install`

If you have Go installed, you can install the latest version of `rao-says` directly from the source:

```sh
go install github.com/your-username/rao-says/home@latest
