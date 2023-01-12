**When RPC**

here are a few classic cases where RPC is used over REST:

- When the client and server are both written in the same programming language: In this case, it may be easier to use an RPC framework to invoke remote procedures directly, rather than going through the overhead of making HTTP requests and parsing the responses.
- When the client and server are tightly coupled: If the client and server are part of the same system or are otherwise closely tied together, it may be more efficient to use RPC to communicate between them.
- When the client needs to invoke multiple procedures in a single request: With REST, each request is typically limited to a single operation, but with RPC, it is possible to invoke multiple procedures in a single request.
- When the client needs to pass complex data structures as arguments to procedures: REST typically uses simple data types (e.g., strings, integers, booleans) as arguments, while RPC frameworks may support the ability to pass more complex data structures as arguments.

----

**WHEN REST**

Here are a few classic cases where REST is used over RPC:

- When the client and server are decoupled: REST is a good choice for building APIs that are consumed by a variety of different clients, as it does not require that the client and server be written in the same programming language or be tightly coupled.
- When the client needs to communicate with multiple servers: With REST, it is easy to switch between servers by changing the base URL of the API, while with RPC the client would need to be explicitly configured to communicate with each individual server.
- When the client needs to cache responses: REST APIs often use HTTP cache headers to allow clients to cache responses and improve performance. This is not possible with RPC, as the response is typically tied to a specific procedure invocation.
- When the client needs to interact with the server through a firewall: REST APIs can be accessed using standard HTTP protocols, which are often allowed through firewalls. This is not always the case for RPC, which may use proprietary protocols that are blocked by firewalls.