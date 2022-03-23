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
plz start local server   on port      8080
    [  command name  ] [first arg] [second arg]  
```

# Parsing
There may be same commands like:
"start server" and "start server local"

And args like: "... on port 8080" and "... on port default"

Need to parse more suitable **command** and **args**.
