---
title: ":fork_and_knife: Usage"
---

## :paperclip: CLI

```sh
$ tjdoe --help
tjdoe [OPTIONS] <ROOT_DIR> <SCORES...>
OPTIONS
    -d, --dest <DIR>       specifies destination of anonymized programs.
                           if this option was not specified, output to 'dest' directory.
    -s, --score <SCORE>    specifies id mapping file. default is 'anonymized_score.csv'
    -S, --seed <SEED>      specifies seed for random values.
    -h, --help             print this message and exit.
    -v, --version          print version and exit.
ROOT_DIR
    the directory contains the programs.  The layout of the directory is arbitrary.
    The user arbitrary defines the names of sub-directories and files.
SCORES...
    show score file, the first row is the header, and the following rows
    represent each student and must be formatted as follows.
    The first column is the id, the second column shows the name by dividing
    the surname the given name with space, the third column is the final score,
    and the following columns represent the scores of assignments.
```

## :whale: Docker

```sh
$ docker run --rm -v $PWD:/home/tjdoe tamada/tjdoe:1.0.0 rootdir scores.csv...
```

Above command should run on directory which has `scores.csv` and `rootdir`.

The meaning of the options above command are as follows.

* `--rm`
    * remove container after running Docker.
* `-v $PWD:/home/tjdoe`
    * share volumen `$PWD` in the host OS to `/home/tjdoe` in the container OS.
    * Note that `$PWD` must be the absolute path.

[![Docker](https://img.shields.io/badge/docker-tamada%2Ftjdoe%3Alatest-blue?logo=docker&style=social)](https://hub.docker.com/u/tamada/tjdoe)

## Others

### Anonymity rules

* Change directory names representing year (four-digit numbers) to `0000`.
* Change file names, and directory names contains id, name to anonymized id.
    * id and name are obtained from `scores.csv`.
* Replace ids and names appearing in the content of the files, to anonymized id.
    * if there are students have the same surnames, every anonymized ids will be listed.

### Example of `score.csv`

```csv
id,name,final score,a01,a02,a03,a04,a05,a06,a07,a08,a09,a10
123456,Tamada Haruaki,87,5,6,7,4,3,1,8,9,8,6
234567,Yamamoto Taro,53,4,3,3,2,3,,,,,4
345678,山本 次郎,95,10,10,10,10,10,10,,10,10,
```
