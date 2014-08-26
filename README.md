Simple example of how to conditional run integration tests.

Normally, a_test.go will not be run as part of the tests because of the build tag.

To cause it to run, add  `-tag=integration` to the commmand line:
`go test -tags=integration .`

We may or may not want to include the test data for integration testing in the source control. Optionally, we may include a link of where to download the test data. It will depend on how large the test data is, and how likely we are to want to reproduce the tests of an old build.
