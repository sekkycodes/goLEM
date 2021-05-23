## TODOs

* Implement scaffolding for E2E testing:
 * Spinning up docker containers / golems
 * Configuring those golems
 * Providing test data
 * Evaluating test results
 * Storing test results and reports
 * Tearing down docker containers / golems

## Test scenarios to be implemented
* Create a golem and configure it appropriately
* Spin up multiple golems and create a simple provide(from file) -> process -> consume(to file) scenario

## Notes
* This is not the place to write unit tests for other modules! The test module is intended for E2E testing.