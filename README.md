Nature Remo command line interface
========

[![release badge](https://img.shields.io/github/v/release/chroju/nature-remo-cli.svg)](https://github.com/chroju/nature-remo-cli/releases)
[![Coverage Status](https://coveralls.io/repos/github/chroju/nature-remo-cli/badge.svg?branch=main)](https://coveralls.io/github/chroju/nature-remo-cli?branch=main)


`remo` is an unofficial command line interface for [Nature Remo](https://nature.global/).

Prerequisites
----

You will need your Nature Remo OAuth2 access token. Please read Nature Remo official document [here](https://developer.nature.global/) and generate your own access token.

Usage
----

### Initialize

At first, you must execute `remo init` command to initialize. You will be asked to enter your access token.

```
$ remo init
Nature Remo OAuth Token:
<Input your token>
Initializing ...
Successfully initialized.
```

Now, you can use `remo` !

**Note:** `remo init` command creates a configuration file in `~/.config/remo` by default. The location of the configuration file can be specified with `REMOCONFIG` environment variable.

### Signal

`remo signal list` command shows your available appliance and signal names.

```
$ remo signal list
light on
light off
light brighten
light darken
TV power
TV 1
...
```

`remo signal send` will send actual Nature Remo signal.

```
$ remo signal send light brighten
Succeeded.
```

### Aircon settings

`remo aircon list` will show you current aircon setting.

```
$ remo aircon list
NAME      POWER  TEMP  MODE  VOLUME  DIRECTION
Bed Room  ON     25    warm  2
Dining    OFF    22    auto  auto
```

You can update aircon setting with `remo aircon send`.

```
$ remo aircon send --name 'Bed Room' -t 23
Updated Aircon 'Bed Room' (TEMP: 25 -> 23)

$ remo aircon send --name Dining --on
Updated Aircon 'Dining' (POWER: OFF -> ON)
```

Install
----

### Homebrew

```bash
$ brew install chroju/tap/nature-remo-cli
```

### Download binary

Download the latest binary from [here](https://github.com/chroju/nature-remo-cli/releases) and  place it in the some directory specified by `$PATH`.

### go get

If you have set up Golang environment in your computer, you can also install with `go get` command.

```
$ go get -u github.com/chroju/nature-remo-cli
```

Advanced Usage
----

After you have executed `remo init` command, `remo` will create config yaml file at `$HOME/.config/remo`. This file contains your API token and your Nature Remo information like below.

```yaml
credential:
  token: <YOUR TOKEN>
appliances:
- name: light
  id: abcdefgh-1234-ijkl-5678-mnopqrstuvwx
  signals:
  - id: abcdefgh-1234-ijkl-5678-mnopqrstuvwx
    name: brighten
    image: ico_foo
  - id: abcdefgh-1234-ijkl-5678-mnopqrstuvwx
    name: darken
    image: ico_bar
```

`remo signal` commands load the names of your Nature Remo appliances and signals from here, so you can execute `remo send light brighten` and `remo send light darken` with above file. If you would like to execute commands with the signal names you like, you can rewrite your YAML.

TODO
----

* [x] Write tests.
* [ ] Implement commands for some sensors.
* [ ] Implement commands for TV.
* [ ] Implement `--direction` option to "aircon send" command

LICENSE
----

[MIT](https://github.com/chroju/nature-remo-cli/LICENSE)
