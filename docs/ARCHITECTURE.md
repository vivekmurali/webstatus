A worker runs in parallel that runs an SQL query every 5 minutes.

It checks for URLs that have their check remaining. You can find how checks work in the [README](README.md#Process).

Each query that returns a URL, a new goroutine is spawned which then sends a request to the given URL with specified parameters. The response code, the headers and the body are then stored.

