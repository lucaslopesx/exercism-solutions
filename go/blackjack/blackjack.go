package blackjack

const stand = "S"
const hit = "H"
const split = "P"
const autoWin = "W"

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
		case "ace":
			return 11
		case "two":
			return 2
		case "three":
			return 3
		case "four":
			return 4
		case "five":
			return 5
		case "six":
			return 6
		case "seven":
			return 7
		case "eight":
			return 8
		case "nine":
			return 9
		case "ten", "jack", "queen", "king":
			return 10
		default:
			return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	var valueCard1 = ParseCard(card1);
	var valueCard2 = ParseCard(card2);
	var valueHand = valueCard1 + valueCard2
	var valueDealerCard = ParseCard(dealerCard);

	if(card1 == "ace" && card2 == "ace"){
		return split;
	}

	if(valueHand <= 11){
		return hit;
	}

	if(valueHand <= 16 && valueDealerCard < 7){
		return stand
	}

	if(valueHand <= 16 && valueDealerCard >= 7){
		return hit
	}

	if(valueHand <= 20){
		return stand
	}

	if(valueHand == 21 && dealerCard != "ace" && valueDealerCard != 10){
		return autoWin
	}
	
	return stand;
}
