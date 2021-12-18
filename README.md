# go_generate_example

stringer as generator:

```
go install golang.org/x/tools/cmd/stringer@latest
go generate ./...
```
This will generate event_string.go under the same working directory as events.go
