Nature Remo command line interface
========

Overview
----

This CLI enable you to manipulate your [Nature Remo](https://nature.global/) in terminal.


Prerequisites
----

You will need your Nature Remo OAuth2 access token. Please read Nature Remo official document [here](https://developer.nature.global/).

Example
----

Usage
----

```
$ remo <command> <sub-command> [parameters]
```

* `init` - Initialize CLI with your access token.
* `aircon` - Change the aircon settings and display current ones.
* `list` - Display the appliances or signals.
* `send` - send a signal.

### init

```
$ remo init
```

You need to execute this command at first. The entered token will be written in `~/.config/remo`.


### aircon

```
$ remo aircon <on|off|change|show> [parameters]
```



```
$ remo aircon on

$ remo aircon change -t 28 -p auto

$ remo aircon show

$ remo aircon off
```

### list

```
$ remo list <appliances|signals> [parameters]
```

```
$ remo list appliances
$ remo list signals 1
```

### send

```
$ remo send 1 1
```

#### signal aliases

Select signal with appliance ID and signal ID is , 

```
$ remo send bright
```