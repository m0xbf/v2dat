# v2dat

A cli tool that can unpack v2ray data packages (also known as `geoip.dat` and `geosite.dat`) to text files.

## Usage

```shell
v2dat unpack geoip [-d output_dir] [-f tag]... geoip_file
v2dat unpack geosite [-d output_dir] [-f tag[@attr]...]... geosite_file
```

- If `-d` was omitted, the current working dir `.` will be used.
- If no filter `-f` was given. All tags will be unpacked.
- If multiple `@attr` were given. Entries that don't contain any of given attrs will be ignored.
- Unpacked text files will be named as `<geo_filename>_<filter>.txt`.

### Example

Here is an example of how to use the `v2dat` tool to unpack `geoip` and `geosite` files:

```sh
v2dat unpack geoip -o ./output_dir -f cn ./geoip.dat
v2dat unpack geosite -o ./output_dir -f cn -f 'geolocation-!cn' ./geosite.dat
```

## Unpacked IP Data

Unpacked IP text files contain a list of CIDRs.

```text
2.16.33.76/32
2.19.128.0/20
2.20.32.0/22
```

## Unpacked Domain Data

`geosite` contains four types of domain rule expression: `domain`, `keyword`, `regexp`, `full`. Each expression can have several attributes `@attr`. More info about `geosite` can be found in [here](https://github.com/v2fly/domain-list-community).

`v2dat` will split type and expression with a `:`. But omits the `domain` prefix and attributes.

```text
google.com
keyword:google
regexp:www\.google\.com$
full:www.google.com
```

## Build

### Linux AMD64

To compile the project for the Linux AMD64 platform, use the following command:

```shell
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o v2dat
```

### ImmortalWrt ARM64

**Download ImmortalWrt SDK**

First, go to [ImmortalWrt SDK](https://downloads.immortalwrt.org/releases/23.05.2/targets/rockchip/armv8/immortalwrt-sdk-23.05.2-rockchip-armv8_gcc-12.3.0_musl.Linux-x86_64.tar.xz) to download the `.tar.xz` compressed package and extract it to the `~/ImmortalWrt/Sdk/` directory.

**Build the program**

Assume `~/Workspaces/Golang/Projects/Desktop/v2dat` is your project path.

```shell
export STAGING_DIR=~/ImmortalWrt/Sdk/immortalwrt-sdk-23.05.2-rockchip-armv8_gcc-12.3.0_musl.Linux-x86_64/staging_dir
export TOOLCHAIN_DIR=$STAGING_DIR/toolchain-aarch64_generic_gcc-12.3.0_musl
export PATH=$TOOLCHAIN_DIR/bin:$PATH

cd ~/Workspaces/Golang/Projects/Desktop/v2dat
export GOOS=linux
export GOARCH=arm64
export CC=$TOOLCHAIN_DIR/bin/aarch64-openwrt-linux-gcc
export CXX=$TOOLCHAIN_DIR/bin/aarch64-openwrt-linux-g++

go build -ldflags="-s -w" -o v2dat
```

### Description

- `GOOS=linux` specifies the target operating system.

- `GOARCH=amd64` or `GOARCH=arm64` specifies the target architecture.

- `-ldflags="-s -w"` reduces the binary size by stripping the symbol table and debug information.

- `-o v2dat` specifies the output file name.
