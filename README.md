# train-guidance
Where to switch trains.
https://train-guidance.appspot.com/

# debugging train-guidance

Example of how to debug this kind of issue using `log.Infof` to print the status
of the BFS function as it traverses the graph.

You can start up the app directly via `dev_appserver.py` and issue requests via
web browser and watch the real requests. but the "real" train network data is
complex, so it's easier to run smaller tests, and that's what `app_test.go`
does. You can run it via

    go test *.go

You'll only see logs if one of the tests fails (and currently one of them will). But if you'd like to see logs even when the tests are passing, you can run:

    go test -v *.go
