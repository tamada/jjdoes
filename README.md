[![License](https://img.shields.io/badge/License-WTFPL-blue.svg)](https://github.com/tamada/jjdoe/blob/master/LICENSE)
[![Version](https://img.shields.io/badge/Version-1.0.0-yellowgreen.svg)](https://github.com/tamada/jjdoe/releases/tag/v1.0.0)

# jjdoes

`jjdoes` anonymize given programs for programming courses and their s score for grades.

## Install

### Homebrew

```sh
$ brew tap tamada/brew
$ brew install jjdoes
```

### Go lang

```sh
$ go get github.com/tamada/jjdoes
```

### Build from source codes

```sh
$ git clone https://github.com/tamada/jjdoes.git
$ cd jjdoes
$ make
```

## Usage

### CLI

```sh
$ jjdoes --help
jjdoes [OPTIONS] <ROOT_DIR> <SCORES...>
OPTIONS
    -d, --dest <DIR>           specifies destination of anonymized programs.
                               if this option was not specified, output to 'dest' directory.
    -m, --mapping <MAPPING>    specifies id mapping file. default is `mapping.csv`
    -s, --seed <SEED>          specifies seed for random values.
    -h, --help                 print this message and exit.
    -v, --version              print version and exit.
ROOT_DIR
    the directory contains the programs.  The layout of the directory is arbitrary.
    The user arbitrary defines the names of sub-directories and files.
SCORES...
    show score file, the first row is the header, and following rows
    represent each student, and must be formatted as follows.  The
    first column is id, the second column shows the name, the third
    column is the final score, and the following columns represent the
    scores of assignments.
```

### Docker

```sh
$ docker run --rm -v $PWD:/home/jjdoes tamada/jjdoes:1.0.0 rootdir scores.csv...
```

Above command should run on directory which has `scores.csv` and `rootdir`.

The meaning of the options above command are as follows.

* `--rm`
    * remove container after running Docker.
* `-v $PWD:/home/jjdoes`
    * share volumen `$PWD` in the host OS to `/home/jjdoes` in the container OS.
    * Note that `$PWD` must be the absolute path.

## About

### Developers

* Haruaki Tamada

### `jjdoes`

`jjdoes` means to convert programs into John Doe and Jane Doe.
