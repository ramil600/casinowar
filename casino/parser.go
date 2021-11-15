package casino

import (
	"encoding/json"
)

//ParseCardsDealed will take wrapped TCPData json and Unmarshal byte part of it
//to CardsDealed json struct
func ParseCardsDealed(tcpdata TCPData) (CardsDealed, error) {
    cardsdealed := CardsDealed{}
	err := json.Unmarshal(tcpdata.Data, &cardsdealed)
	if err != nil {
		return CardsDealed{}, err
	}
	return cardsdealed, nil
}
