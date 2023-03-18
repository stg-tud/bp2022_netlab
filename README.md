```
            _   _       _
 _ __   ___| |_| | __ _| |__
| '_ \ / _ \ __| |/ _' | '_ \
| | | |  __/ |_| | (_| | |_) |
|_| |_|\___|\__|_|\__,_|_.__/
```

# bp2022_netlab

netlab helps you to quickly generate configuration files for network experiment softwares such as [CORE](http://coreemu.github.io/core/), [coreemu-lab](https://github.com/gh0st42/coreemu-lab) and [The ONE](https://github.com/akeranen/the-one) using a single TOML file. Its aim is to make your work easier, thus it handles annoying tasks such as generating IP addresses, movement patterns (using [BonnMotion](https://sys.cs.uos.de/bonnmotion/)) and multiple parameterized runs.

netlab was developed by a group of five students for the [Software Technology Group](https://www.stg.tu-darmstadt.de/main_stg/index.en.jsp) as part of a bachelor internship of the [department of computer science](https://www.informatik.tu-darmstadt.de/fb20/index.en.jsp) at [TU Darmstadt](https://www.tu-darmstadt.de/index.en.jsp) in summer term 2023.

## Requirements

In order to generate movement patterns, [BonnMotion](https://sys.cs.uos.de/bonnmotion/) version 3.0.1 or higher must be installed. It must either be available under the name "bm" (which is default) in your `$PATH`/`%Path%` or you must specify the path of the BonnMotion executable using the `-b <path-to-bonnmotion>` flag of netlab.

If you do not want to use our pre-built [Releases](https://github.com/stg-tud/bp2022_netlab/releases) you will also need [go](https://go.dev/) version 1.19 or higher installed.

## Installation

To install netlab, you can either use the pre-built binaries provided [here](https://github.com/stg-tud/bp2022_netlab/releases) or build one yourself. To build from source, run these commands:

```bash
# HTTPS:
git clone https://github.com/stg-tud/bp2022_netlab.git netlab
# SSH:
git clone git@github.com:stg-tud/bp2022_netlab.git netlab

cd netlab

# Linux, macOS:
go build -o netlab main.go
# Windows:
go build -o netlab.exe main.go
```

## Usage

netlab is a simple command line application. Thus you have to run it from a command line. Either install it to a path in `$PATH` (`%Path%` for Windows respectively) or run it inside the current folder, like so:

```bash
# installed in $PATH:
netlab

# in current folder:
./netlab
```

The usage itself should be pretty self-explainatory. You can always run `netlab help` or `netlab help <command>` for help on that.

Currently, these **commands** are supported:

- `netlab version`: Shows the current version number of netlab.
- `netlab test [filename]`: Tests the given TOML file for errors but does not generate any output files.
- `netlab generate [filename]` (or `netlab gen [filename]` for short): Loads the given TOML file and generates output files for all targets.

There are also some **flags** which might interest you:

- `-h --help`: shows help (same as `netlab help`)
- `-v --version`: shows the version (same as `netlab version`)
- `-d --debug`: enabled debug output/logging
- `-o --overwrite`: overwrites existing files _(`generate` only)_
- `-f <name> --folder <name>`: outputs to the specified folder name instead of the default `output` _(`generate` only)_
- `-t <target> --target <target>`: outputs for the specified target format. If specified, the targets definition from the TOML file will be ignored. Can be specified multiple times for multiple targets. Possible values are for example `core`, `coreemu-lab` and `the-one`. _(`generate` only)_
- `-b <path-to-bonnmotion> --bonnmotion <path-to-bonnmotion>`: path of the BonnMotion executable to run, if not installed as `bm` in `$PATH`/`%Path%`. _(`generate` only)_

## Configuration files

netlab uses TOML files for experiment configurations. For more information on how to create those, have a look at the [Wiki](https://github.com/stg-tud/bp2022_netlab/wiki/). You can also find a few example files in the [examples](https://github.com/stg-tud/bp2022_netlab/tree/documentation/examples) folder.
