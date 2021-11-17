package casino

import (
	"encoding/json"
)

//ParseCardsDealt will take wrapped TCPData json and Unmarshal byte part of it
//to CardsDealed json struct
func ParseCardsDealt(tcpdata TCPData) (CardsDealed, error) {
    cardsdealt := CardsDealed{}
	err := json.Unmarshal(tcpdata.Data, &cardsdealt)
	if err != nil {
		return CardsDealed{}, err
	}
	return cardsdealt, nil
}
