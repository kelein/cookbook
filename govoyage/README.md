# GRPC Gateway

```bash
          ┌────────────────────────┐           ┌────────────────────────┐
 Request  │          HTTP          │           │          gRPC          │
──────────►         Server         │           │         Server         │
          │                        │           │                        │
          │                        │           │                        │
          │           ┌──────────┐ │           │                        │
          │           │   gRPC   ◄─┼───────────►   ┌────────────────┐   │
 Response │           │  Client  │ │           │   │ ExampleService │   │
◄─────────┤           └──────────┘ │           │   └────────────────┘   │
          │                        │           │                        │
          └────────────────────────┘           └────────────────────────┘
```
