## Base project for Golang
### About this branch:
- This `simplified` branch is only for Go beginner to have something concrete to build upon and not for production build.
- Still enforces that the direction of dependency must always be from outer layers to inner layers.
- The structure is much more direct than the `master` branch: there's no layer-boundary interfaces, and inner layers just directly import services from outer layers.
### Structure:
* IO/Infra layer
  * adapter/
    * Infra/IO layer to interact with external systems
    * DB, repositories and the like live here
  * gateway/
    * Infra/IO layer for external systems to interact with this system
    * Typically the place for http/grpc/graphQL... server and their controllers
* Application logic layer
  * application/
    * Main application logic. Consists of multiple independent interactors, each interactor groups a set of related usecase flows
  * services/
    * Services providing reusable activities/steps for usecase flows
    * Each service may wrap multiple adapters
* Core domain layer
  * domain/
    * Core domain model and logic
* common/
  * Commonly used stuffs used at multiple layers, e.g. logger, utilities and so on
* migrations/
  * Migration scripts
