Simple log formatter.

```
# echo "foo" 2>&1 | ./logformat -t bar >> /path/to/log

# cat /path/to/log
[Thu Aug 2 16:31:30 -0700 2018] [bar] foo
```

