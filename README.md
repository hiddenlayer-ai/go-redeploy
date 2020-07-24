# go-redeploy

go-redeploy allows you to redeploy go projects that use git. This is great for projects that might be working a server and are not extremely convienent to access. Whenever a change to `origin/master` branch, go-redeploy pulls and run the new executable.

### Configuration

Note that there is a `config.toml` file. There are three parameters that you have to set. `sleep` sets the time in seconds between each check to `origin/master` for an update. `repo-path` is the relative path from `go-deploy` to the go project that you are managing. The `entry-point` is the file that will be run on each update like so: `go run {entry-point}.go`

### Run

```
$ go run redeploy
```

or 

```
$ go build
$ ./go-redeploy
```
