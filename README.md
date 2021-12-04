# Simple chat

This is a very simple chat that I did following some examples and I am trying to use what I already know about channels in go.

So the code follows basically 3 steps:
- The main function starts a goroutine `Server.start`, which is responsible to know if has something to print to the users.
- The main function runs an infinite loop to keep the chat running and verify if has new connections, it has one it calls `Server.newClient`
- The `newClient` function store the user in a map and runs an infinite loop to check if the user wrote a message, if has messages they will be pushed on the `conversation` channel.


To test this code run the main.go and then open 2 sessions on the terminal and type this: `telnet localhost 3000`
