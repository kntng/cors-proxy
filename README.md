# cors-proxy

A small, fast, concurrent CORS proxy server written in Go.

Available as a binary or a Docker image (both ~2MB).

## Building From Source

If you have Make, use the Makefile and run

```sh
make
```

Or alternatively run

```sh
go build proxy.go
```

to build for your platform.

### Size Optimizations

1. Compiling the program with no flags results in a 8.2 MB binary.

2. Compiling the binary with `-ldflags "-s -w"` reduces the binary size to 5.6 MB.

3. Compressing with UPX reduces the binary size to 2.3 MB.

4. Compressing with UPX with the `--brute` flag reduces the binary size to 1.8 MB.
