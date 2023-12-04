// See https://aka.ms/new-console-template for more information
var cards = new Dictionary<int, int>();

var tryAddCard = (int card) => {
	if (cards.ContainsKey(card))
	{
		cards[card] = cards[card]+1;
	}
	else
	{
		cards.Add(card, 1);
	}
};

var tryAddCards = (int card, int number) => {
	if (cards.ContainsKey(card))
	{
		cards[card] = cards[card]+number;
	}
	else
	{
		cards.Add(card, number);
	}
};

var matrix = new List<string>();

var currentCard = 0;
foreach (var line in File.ReadLines("Input.txt"))
{
	tryAddCard(currentCard);
	matrix.Add(line);
	currentCard++;	
}

currentCard = 0;
foreach (var line in matrix)
{
	var numberOfCards = cards[currentCard];

	var indexOfColon = line.IndexOf(":");
	var roundsStart = indexOfColon+2;

	var scoreSections =line[roundsStart..].Split(" | ");

	var winningNumbers = scoreSections[0].Split(" ",  StringSplitOptions.RemoveEmptyEntries).Select(x => Convert.ToInt32(x)).ToList();
	var myNumbers = scoreSections[1].Split(" ",  StringSplitOptions.RemoveEmptyEntries).Select(x => Convert.ToInt32(x)).ToList();

	var myScore = 0;
	foreach(var number in myNumbers)
	{
		if (winningNumbers.Contains(number))
		{
			myScore++;
		}
	}
	for (int i = currentCard+1; i < Math.Min(matrix.Count, currentCard+myScore+1); i++)
	{
		tryAddCards(i, numberOfCards);
	}

	currentCard++;
}

var answerPart1 = cards.Values.Sum();
Console.WriteLine($"Sum of scores: {answerPart1}");