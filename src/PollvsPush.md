**When Polling**

Here are a few classic cases where polling is used over Pushing:

- When the client only needs to receive updates periodically: If the client only needs to receive updates at regular intervals (e.g., every minute or every hour), polling may be a good choice, as it allows the client to control the frequency of the updates.
- When the client has limited resources: Polling can be less resource-intensive for the client, as it does not need to constantly maintain a connection to the server and be ready to receive data at any moment.
- When the client needs to prioritize certain updates: With Polling, the client can choose to prioritize certain updates by increasing the frequency of the requests for those updates.
- When the client needs to conserve network bandwidth: Polling can be more efficient in terms of network bandwidth, as it only sends a request when the client needs an update, rather than continuously sending data regardless of whether the client is ready to receive it.


**When Pushing**

Here are a few classic cases where Pushing is used over Polling:

- When the client needs to receive updates in real-time: If the client needs to receive updates as soon as they become available (e.g., for a stock ticker or a live sports feed), Pushing may be a good choice, as it allows the client to receive the updates as soon as they are available.
- When the client needs to receive a high volume of updates: If the client needs to receive a large number of updates in a short period of time, Pushing may be more efficient, as it allows the server to send all of the updates at once, rather than requiring the client to send multiple requests to retrieve them.
- When the client needs to minimize latency: Pushing can minimize latency, as the client does not need to wait for a request to be sent and a response to be received before receiving new data.
- When the client and server are on the same network: If the client and server are on the same network (e.g., in a LAN or WAN), Pushing may be more efficient, as it can take advantage of the lower latency and higher bandwidth of the network.