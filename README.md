# mark
```shell
$ nodebrew ls | sed '$d'
v14.5.0
v14.6.0
v15.2.1
```

```shell
$ nodebrew ls | sed '$d' | ./mark -mark "*" "v15.2.1" "v14.6.0"
  v14.5.0
* v14.6.0
* v15.2.1
```