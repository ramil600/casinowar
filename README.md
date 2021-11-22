This is the implementation of a game Casino War.

Implemented over tcp connection. Server handles multiple players. 
* One player per table. 
* If card has a higher rank player wins 
* If ranks match player has an option to increase bet and go to war.

## CLI Usage

**To run server:**

`
Usage: .\server.exe [options]

OPTIONS

--port/$SERVER_PORT  <string>  (default: 8081)

--help/-h

display this help message
`

**To run client:**

To run from container execute /bin/ash:
> docker exec -it [containerId] /bin/ash

Run:
> .\client

.\client [options]

OPTIONS

--host/$SERVER_HOST  <string>  (default: 0.0.0.0)

--port/$SERVER_PORT  <string>  (default: 8081)

--help/-h


### THE GAME LOGIC

The game logic is stored in State struct in game.go file. It holds player's and dealer's 
cards. The deck that the game  is played with includes 6 * 52 cards. TopCard is pointing to
the top of the deck. And the player information that participates in the game.
###THE SERVER
For each concurrent player HandleConnection function initiates a new deck, shuffles the cards
in the deck.
It enters the game loop by accepting the player's bet. If bet is 0, it means that the player
decided to quit. 
The server also checks if the bet is a war bet and applies the required logic.
The server then calls State#DealCards to create struct with cards to be sent to the player, 
and SendCards to send the json obeject to the client.

If dealt cards are a draw then server expects Y/N answer from the client to handle the war game
scenario.
If player agrees server also calls State#DealCards, this function has a logic to 
update player's win and loss according to War Game scenario. Then server calls SendCards
to send json object to the client.

If player opts out server calls State#ProcessWarOut function that handles partially refunding
player's bet.

###THE CLIENT
Every client is given 10000$ in the bank to start the game and offered to make original and
side bets. Inputs are taken through os.Stdin and marshalled json objects are sent to server.
