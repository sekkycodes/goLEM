## TODOs

* Implement functionality for starting up new golems
 * On startup a golem discovers all other nodes in the network using the core/discover package and creates a local store with that information using the core/registry package
 * Creation can be done using the CLI (later ther might be a webapp or a config file based startup mechanism as well)
 * Design a way to hand information to golems when creating them:
  * What type of golem is this node? (provider, processor, or consumer?)
  * What exactly does it do?
  * For providers and processors: Define the next node in the process flow
  * For consumers: Define endpoint
* Containerize the golem in Docker containers
 * Registry (and possibly other local storage) should be persistent - i.e. a docker volume