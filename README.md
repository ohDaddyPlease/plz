# plz
Declarative helper

Toml|Yaml draft:

name: "command name"
description: "command description"
type: "command type"
args:
  - first arg
  - second arg
  - N arg
command: "command to exec"
dependency:
  - dependency type (run after|before)
  - dependency command name
