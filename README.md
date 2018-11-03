Nature Remo command line interface
========

`remo` is the command line interface for [Nature Remo](https://nature.global/).

Prerequisites
----

You will need your Nature Remo OAuth2 access token. Please read Nature Remo official document [here](https://developer.nature.global/).

Usage
----

At first, you must execute `remo init` command to initialize. You will be prompted for your OAuth2 token.

```
$ remo init
Nature Remo OAuth Token:
<Input your token>
Initializing ...
Successfully initialized.
```

`remo list` command shows your available appliance and signal names.

```
$ remo list
light on
light off
light brighten
light darken
TV power
TV 1
...
```

`remo send` will send actual Nature Remo signal.

```
$ remo send light brighten
Success.
```

Install
----

### Download binary

Download binary from [here](https://github.com/chroju/nature-remo-cli/releases) and put it in your `$PATH` directory.

### go get

If you have installed Golang environment in your PC, you also install with `go get` command.

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

`remo list` and `remo send` commands load your Nature Remo config from here, so you can execute `remo send light brighten` and `remo send light darken` with above file. If you would like to execute commands with the signal names you like, you can rewrite your YAML.

TODO
----

* [ ] Write tests.
* [ ] Implement commands for aircon settings.
* [ ] Implement commands for some sensors.

Author
----

[chroju](https://github.com/chroju)

LICENSE
----

[MIT](https://github.com/chroju/nature-remo-cli/LICENSE)
