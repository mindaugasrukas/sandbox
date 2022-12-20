# Using circular dependencies with Uber.fx

## Dependencies

Declarative dependencies:

    % goda graph ./... | dot -Tsvg -o graph.svg && open graph.svg

![img](graph.svg)

Implicit dependencies:

    % dot -Tsvg -o graph_implicit.svg graph_implicit.dot && open graph_implicit.svg

![img](graph_implicit.svg)

## Run example

    % go run ./cmd
    [Fx] PROVIDE    *service1.Service1 <= sandbox/di/service1.NewService1()
    [Fx] PROVIDE    *service2.Service2 <= sandbox/di/service2.NewService2()
    [Fx] PROVIDE    service3.ServiceA <= fx.Annotate(sandbox/di/service2.NewService2(), fx.As([[service3.ServiceA]])
    [Fx] PROVIDE    *service3.Service3 <= sandbox/di/service3.NewService3()
    [Fx] PROVIDE    fx.Lifecycle <= go.uber.org/fx.New.func1()
    [Fx] PROVIDE    fx.Shutdowner <= go.uber.org/fx.(*App).shutdowner-fm()
    [Fx] PROVIDE    fx.DotGraph <= go.uber.org/fx.(*App).dotGraph-fm()
    [Fx] INVOKE             sandbox/di/service1.ServerLifetimeHooks()
    [Fx] HOOK OnStart               sandbox/di/service1.ServerLifetimeHooks.func1() executing (caller: sandbox/di/service1.ServerLifetimeHooks)
    Service1.Start()
    Service2.DoSomething2()
    Service3.DoSomething()
    Service3.DoSometingWithServiceA()
    Service2.DoSomething()
    [Fx] HOOK OnStart               sandbox/di/service1.ServerLifetimeHooks.func1() called by sandbox/di/service1.ServerLifetimeHooks ran successfully in 5.041µs
    [Fx] RUNNING
    ^C[Fx] INTERRUPT
    [Fx] HOOK OnStop                sandbox/di/service1.ServerLifetimeHooks.func2() executing (caller: sandbox/di/service1.ServerLifetimeHooks)
    Service1.Stop()
    [Fx] HOOK OnStop                sandbox/di/service1.ServerLifetimeHooks.func2() called by sandbox/di/service1.ServerLifetimeHooks ran successfully in 3.75µs
