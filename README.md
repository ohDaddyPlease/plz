# plz
Declarative helper

# Toml|Yaml draft:

```
name: "command name"
description: "command description"
type: "command type"
args:
  - 1
    - name: "first arg name"
    - value: "first arg value"
  - 2
    - name: "second arg name"
    - value: "second arg value"
command: "command to exec"
dependency:
  - dependency type (run after|before)
  - dependency command name
```

# Example:
```
plz start_local_server   on_port      8080
    [  command name  ] [first arg] [second arg]  
```

# Reserved command

- plz validate cfg - validate syntax ofthe main config file
- plz init - initialize cfg
- plz where cfg - show full path to cfg
