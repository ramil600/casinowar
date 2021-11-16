This is the implementation of a game Casino War.

Implemented over tcp connection. Server handles multiple players. 
* One player per table. 
* If card has a higher rank player wins 
* If ranks match player has an option to increase bet and go to war.

To run server:
> /casinowar 8081

To run client:
> client/client 127.0.0.1:8081
