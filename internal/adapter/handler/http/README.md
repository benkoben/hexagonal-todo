# Http server

Http server is a driver actor adapter. This is because the HTTP server is an initiator of communication with the core. Its responsibilities are listed below:

**Roles**

1. Perform validation of requests (gin can do this i think)
2. Logic to prepare a response (handler functions)
3. Map requests into logic that prepares a response (done with gin.Engine)
4. Log requests (a logger configured on our http server)
