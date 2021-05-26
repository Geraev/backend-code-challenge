# Backend code challenge

Write a TCP service in Go.

The service should have the following features:
1. Start listening for TCP/UDP connections.
2. Be able to accept connections.
3. Read a json payload `{"user_id": 1, "friends": [2, 3, 4]}`
3. After establishing a successful connection - "store" the payload in memory as you see fit.
4. When another connection is established with the `user_id` from the list of any other
   user's `friends` section, they should be notified about it with message `{"online": true}`
5. When the user goes offline, their `friends` (if they have some and any of them are online)
   should receive a message `{"online": false}`
6. Service should be deployable.

# Questions:
1. What changes if we switch from TCP to UDP?
2. How would you detect the user (connection) is still online?
3. What issues do you foresee if the user can have a lot of friends? How would you change design of your application in case of 10? 100? 1000? 10000 friends? 