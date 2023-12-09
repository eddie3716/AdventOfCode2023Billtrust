package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func contains(slice []string, element string) bool {
	for _, item := range slice {
		if item == element {
			return true
		}
	}
	return false
}

const (
	HighCard = iota
	OnePair
	TwoPairs
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

type Hand struct {
	Cards string
	Bet   int
}

func (h Hand) HandType() int {
	cards := strings.Split(h.Cards, "")
	uniqueCards := []string{}
	for _, card := range cards {
		if !contains(uniqueCards, card) {
			uniqueCards = append(uniqueCards, card)
		}
	}

	numOfJokers := strings.Count(h.Cards, "J")

	switch len(uniqueCards) {
	case 1:
		return FiveOfAKind
	case 2: // JJJJ7 7777J or JJ777 77JJJ
		occurances := strings.Count(h.Cards, uniqueCards[0])
		if numOfJokers > 0 {
			return FiveOfAKind
		}
		if occurances == 4 || occurances == 1 {
			return FourOfAKind
		} else {
			return FullHouse
		}
	case 3: // JJJ76 777J6 JJ776 7766J
		for _, uniqueCard := range uniqueCards {
			occurances := strings.Count(h.Cards, uniqueCard)
			if occurances == 2 {
				if numOfJokers == 1 {
					return FullHouse
				}
				if numOfJokers == 2 {
					return FourOfAKind
				}
				return TwoPairs
			} else if occurances == 3 {
				if numOfJokers == 1 || numOfJokers == 3 {
					return FourOfAKind
				}
				return ThreeOfAKind
			}
		}
	case 4: // 45JJ7 4566J
		if numOfJokers > 0 {
			return ThreeOfAKind
		}
		return OnePair
	case 5:
		if numOfJokers == 1 {
			return OnePair
		}
		return HighCard
	}
	panic("Found hand type")
	return -1
}

func main() {

	cardRanks := []string{"A", "K", "Q", "T", "9", "8", "7", "6", "5", "4", "3", "2", "J"}
	//cardRanks := []string{"A", "K", "Q", "J", "T", "9", "8", "7", "6", "5", "4", "3", "2"}

	file, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	hands := []Hand{}

	for scanner.Scan() {
		line := scanner.Text()
		var cards string
		var bet int
		for index, token := range strings.Split(line, " ") {
			if index == 0 {
				cards = token
			} else {
				bet, _ = strconv.Atoi(token)
			}
		}
		hands = append(hands, Hand{Cards: cards, Bet: bet})
	}

	slices.SortStableFunc(hands, func(firstHand, secondHand Hand) int {
		firstHandType := firstHand.HandType()
		secondHandType := secondHand.HandType()

		if firstHandType == secondHandType {
			for index := 0; index < len(firstHand.Cards); index++ {
				for _, rank := range cardRanks {
					if string(firstHand.Cards[index]) == rank {
						if string(secondHand.Cards[index]) != rank {
							return 1
						}
					}
					if string(secondHand.Cards[index]) == rank {
						if string(firstHand.Cards[index]) != rank {
							return -1
						}
					}
				}
			}
			return 0
		}

		if firstHandType < secondHandType {
			return -1
		}

		return 1
	})

	totalWinnings := 0
	rank := 0
	for _, hand := range hands {
		rank++

		totalWinnings += rank * hand.Bet
		//fmt.Println(hand, hand.HandType(), rank, hand.Bet, rank*hand.Bet)
	}

	fmt.Println(hands)
	fmt.Println(totalWinnings)
}
