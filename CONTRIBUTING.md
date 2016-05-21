# We love pull requests. Here's a quick guide:

1. Fork the repo.

2. Run the tests. We only take pull requests with passing tests, and it's great
to know that you have a clean slate.

	$ make test

Tests include ``go vet`` and ``gofmt`` checks.

3. Add a test for your change. Only refactoring and documentation changes
require no new tests. If you are adding functionality or fixing a bug, we need
a test!

4. Make the test pass.

5. Push to your fork and submit a pull request.

6. Verify your tests passed in travis

At this point you're waiting on us. We like to at least comment on, if not
accept, pull requests within three business days. We may suggest some changes
or improvements or alternatives.
