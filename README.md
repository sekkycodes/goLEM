# golem

Note: this project is still in its conception and POC phase and might be subject to heavy changes (including core concepts!).

## The Idea

What if you could run a set of functions similar to Linux pipes, but scalable as Docker containers?
What if further you could extend that functionality with small pieces of input, processing, or output programs?
The idea is to provide an environment where writing complex processes easily.
Picture a set of conveyor belts, where raw input material is processed into fully finished products.

Yes, tools for doing all this already exist - and to some extend they influence the functionality of golem.
This project is a fresh look on the capabilities of such tools with heavy focus on simple extendability and built-in scalability.

## Basic Functionality

The components of the software (let's call them "golems") come in three flavors:
* Providers - those are the "entry points" into the process, taking a raw input and handing it over to processors 
* Processors - make up the main bulk of the golems and may come in many different types again, some of which are Transformers, Filters, or Delegators
* Consumers - the "exit" points of the process, typically sending the final result to a service (i.e. a REST endpoint), or storing it as a file or database entry

A typical scenario might be:

A JSON file is discovered on the file system by a JsonFileProvider golem, who hands it over to a JsonOrderTransformer golem. The JsonOrderTransformer maps the raw JSON into a usable canonical Order item. This item can then be further processed by a Filter golem to determine what to do with the Order and hand it to the correct Consumer. For example, invalid Orders might be sent to a REST endpoint, while valid orders are stored on the file system for later processing, being grabbed at any point by another OrderProvider golem...

Every golem is hosted as an individual container and therefore fully scalable. The underlying framework makes sure that no items are dropped, handled twice, or lost somewhere along the way. A set of golems is provided out-of-the-box to steer and monitor the individual processes (i.e. loggers, error queues, etc.).
The framework provides and handles retries, timeouts and erroneous or invalid content by default with the properties of those mechanisms (# of retries, ms to timeout etc) being customizable.

The plan is further to make the framework runnable in the cloud, so resources can be provisioned to scale processes with high load up during peak times as needed. However, if somebody wants to run simple processes installed locally, this should also be possible - maybe with docker-compose or by deploying the golems outside containers.

### Testability

For many such processing frameworks Testability is a huge challenge. The concept here is to have golems be mocked and stubbed out to run processes (or partial processes) in isolation. Use of interfaces for implementation simplifies mocking on a unit testing level when working on a golem.

## Development

### Build and Test

To create te bazel build files:

    bazel run //:gazelle

To build the project run

    // build all
    bazel build //...

    // build specific
    bazel bulid //core/mock

    // test all
    bazel test //...

    // test specific
    bazel test //core/mock

Useful go commands:

    go test