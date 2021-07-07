# Delfin

![Stars](https://img.shields.io/github/stars/5elenay/delfin)
![Commits](https://img.shields.io/github/commit-activity/w/5elenay/delfin)

🐬 Command line tool for compress & decompress a folder.

## Note

Delfin is in alpha-release so it may have some bugs. If you found a bug, please fork the project and send a pull request!

## Installation

Check [release page](https://github.com/5elenay/delfin/releases/latest) for binary files. Just download the file for your operating system and check with:

- **Windows**: `.\delfin.exe version`
- **Linux/macOs**: `./delfin version`

If it works, you can add Delfin to the path now. If not go to the **Compiling** title.

## Compiling

If you can't find a version for your system / architecture, You can compile Delfin yourself easily.

### Pre-Requests

Make sure you have Go downloaded. For check run `go version`. If it works you already have Go in your computer. If not, download from [official website](https://golang.org/dl/).

### Windows

Download the source code with git (`git clone https://github.com/5elenay/delfin.git`) or download zip and extract it. Now run these commands:

```ps1
cd delfin
.\compile\win.bat
```

and it will compile Delfin now. When its finished it will run `.\delfin.exe help` for make sure it works.

### Linux / MacOs

Its almost same with windows. Download the source code with git or download zip and extract it, Then run these commands:

```bash
cd delfin
./compile/other.sh
```

and it will compile Delfin now. When its finished it will run `./delfin help` for make sure it works.

### Other

Goto the `src` folder and run:

```sh
go build
```

After this, it will give you Delfin's binary file.

## Commands

### help

Shows all the commands. If you add an extra parameter it will show to extra informations about the command.

#### Usage

- `delfin help`
- `delfin help <command>`

#### Example(s)

```bash
delfin help
```

```bash
delfin help compress
```

### version

Shows the version and github url.

#### Usage

- `delfin version`

#### Example

```bash
delfin version
```

### check

Check Delfin for updates.

#### Usage

- `delfin check`

#### Example

```bash
delfin check
```

### compress

Compress a folder to `.delfin` format.

#### Usage

- `delfin compress <folder location> <output location>`

#### Example(s)

```bash
delfin compress ./example .
```

```bash
delfin compress ./example ./files
```

### decompress

Decompress a `.delfin` file.

#### Usage

- `delfin decompress <file location> <output location>`

#### Example(s)

```bash
delfin decompress example.delfin .
```

```bash
delfin decompress example.delfin ./folders
```
