// See https://aka.ms/new-console-template for more information
var cardScores = new Dictionary<string, int>();

var count = 1;
foreach (var line in File.ReadLines("Input.txt"))
{
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
			if (myScore == 0)
				myScore = 1;
			else
				myScore *= 2;
		}
	}
	cardScores.Add($"Game {count}", myScore);

	count++;
}

var answerPart1 = cardScores.Values.Sum();
Console.WriteLine($"Sum of scores: {answerPart1}");