# Using circular dependencies with Uber.fx

## Circular dependencies issue

A circular or cyclic dependency occurs when two or more modules or components depend on each other to function correctly, creating a cycle of dependency.

Circular dependencies are generally considered a bad practice in software development, as they can make a system more complex and harder to maintain. However, there may be some cases where creating a circular dependency is necessary or beneficial.

One example of when a circular dependency might be useful is when building a data model that represents a many-to-many relationship between two entities. For example, if you have a "users" and a "groups", and you want to represent the fact that a user can belong to multiple groups and a group can have multiple users.

The dependency graph would be:
```
groups -(depends_on)-> users.
users -(depends_on)-> groups.
```

Circular dependencies can cause problems with initialization, as one module may depend on another module being initialized before it can be initialized itself. This can lead to initialization failures or other errors.

## Solutions

There are several ways to solve circular dependency issues:

1. Refactor the code to remove the dependency cycle: This is generally the best approach, as it involves restructuring the code to eliminate the circular dependency. This may include breaking the cycle by introducing a new component or module that acts as a mediator between the two dependent components or refactoring the code to make one component depend on a third component instead of the other. To help find a breaking point, review SRP (Single Responsibility Principle). I.e., a package should not have too many responsibilities.
2. Use a dependency injection: A dependency injector is a software design pattern that allows you to specify a component's dependencies at runtime rather than compile time. This can help to decouple the components and make it easier to manage the dependencies between them.
3. Use lazy loading: Lazy loading is a technique where a component is only loaded or initialized when needed. This can help to break the dependency cycle, as one component can be initialized before the other, allowing the system to start up and function properly.
4. Use forward declarations: Forward declarations are used to tell the compiler about the existence of a symbol (such as a function or variable) without actually providing its definition. This can allow you to break the dependency cycle by allowing one component to reference another component without actually depending on it.
5. Use interfaces: Interfaces specify a set of related functions that a class or component must implement. By using interfaces, you can decouple the components and break the dependency cycle by allowing one component to depend on an interface rather than a concrete implementation.

## Example

In this PoC example, we will use DI (Uber.fx) and interfaces to solve circular dependency issues.

The original dependency graph:
Service1 depends on service2 and service3;
Service2 depends on service3.

```
service1 -> (service2, service3).
service2 -> service3.
```

We want to add another dependency, `service3 -> service2`.

```
service1 -> (service2, service3).
service2 -> service3.
service3 -> service2.
```

However, this cycle is not allowed.

To solve this issue, we want to keep the original declarative dependency graph but have implicit access to circular dependencies.

To achieve this, we can define the provider interface on the client side:

service3/service3.go:
```
ServiceA interface {
    DoSomething() error
}
```

And annotate the DI container entity that implements mentioned interface. That is possible using the `fx.Annotate` feature:

cmd/main.go
```
fx.Annotate(
    service2.NewService2,
    fx.As(new(service3.ServiceA)),
),
```

### Dependency graphs

Declarative dependencies:

    % goda graph ./... | dot -Tsvg -o graph.svg && open graph.svg

![img](graph.svg)

Implicit dependencies:

    % dot -Tsvg -o graph_implicit.svg graph_implicit.dot && open graph_implicit.svg

![img](graph_implicit.svg)

### Runtime example

    % go run ./...
    [Fx] PROVIDE    *service1.Service1 <= sandbox/di/service1.NewService1()
    [Fx] PROVIDE    *service2.Service2 <= sandbox/di/service2.NewService2()
    [Fx] PROVIDE    service2.Iface <= fx.Annotate(sandbox/di/service2.glob..func1(), fx.As([[service2.Iface]])
    [Fx] PROVIDE    *service3.Service3 <= sandbox/di/service3.NewService3()
    [Fx] PROVIDE    service3.ServiceA <= fx.Annotate(main.main.func1(), fx.As([[service3.ServiceA]])
    [Fx] PROVIDE    fx.Lifecycle <= go.uber.org/fx.New.func1()
    [Fx] PROVIDE    fx.Shutdowner <= go.uber.org/fx.(*App).shutdowner-fm()
    [Fx] PROVIDE    fx.DotGraph <= go.uber.org/fx.(*App).dotGraph-fm()
    [Fx] INVOKE             sandbox/di/service1.ServiceLifecycleHooks()
    Service2.NewService2()
    [Fx] HOOK OnStart               sandbox/di/service1.ServiceLifecycleHooks.func1() executing (caller: sandbox/di/service1.ServiceLifecycleHooks)
    Service1.Start() 0x14000072640
    Service2.DoSomething2() 0x14000018d30, 5577006791947779410
    Service3.DoSomething() 0x14000013c60
    Service3.DoSometingWithServiceA() 0x14000013c60
    Service2.DoSomething() 0x14000018d30, 5577006791947779410
    [Fx] HOOK OnStart               sandbox/di/service1.ServiceLifecycleHooks.func1() called by sandbox/di/service1.ServiceLifecycleHooks ran successfully in 8.25µs
    [Fx] RUNNING
    ^C[Fx] INTERRUPT
    [Fx] HOOK OnStop                sandbox/di/service1.ServiceLifecycleHooks.func2() executing (caller: sandbox/di/service1.ServiceLifecycleHooks)
    Service1.Stop() 0x14000072640
    [Fx] HOOK OnStop                sandbox/di/service1.ServiceLifecycleHooks.func2() called by sandbox/di/service1.ServiceLifecycleHooks ran successfully in 7.167µs
