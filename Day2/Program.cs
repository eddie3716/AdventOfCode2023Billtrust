// See https://aka.ms/new-console-template for more information
Console.WriteLine("Hello, World!");
//var matchedGames = new List<string>();
var gamePower = new Dictionary<string, int>();
// var MaxReds = 12;
// var MaxBlues = 14;
// var MaxGreens = 13;


//var count = 0;
foreach (var line in File.ReadLines("CubesInput.txt"))
{
	// var isMatch=true;

	var indexOfColon = line.IndexOf(":");
	var roundsStart = indexOfColon+2;

	var biggestRed = 0;
		var biggestBlue = 0;
		var biggestGreen = 0;

	foreach (var round in line[roundsStart..].Split("; "))
	{
		foreach (var cube in round.Split(", "))
		{
			var tokens = cube.Split(" ");
			var number = Convert.ToInt32(tokens[0]);
			var color = tokens[1];
			if (color == "red")
				biggestRed = System.Math.Max(biggestRed, number);
			if (color == "blue")
				biggestBlue = System.Math.Max(biggestBlue, number);
			if (color == "green")
				biggestGreen = System.Math.Max(biggestGreen, number);
			// if ((color == "red" && number > MaxReds)
			// 	|| (color == "blue" && number > MaxBlues)
			// 	|| (color == "green" && number > MaxGreens))
			// 	{
			// 		isMatch = false;
			// 		break;
			// 	}
		}

		// if (!isMatch)
		// {
		// 	break;
		// }
	}

	// if (isMatch)
	// {
	// 	var game = line[0..indexOfColon].ToString();
	// 	matchedGames.Add(game);
	// 	Console.WriteLine($"{game} matched!");
	// }
	var game = line[0..indexOfColon].ToString();
	var power = biggestBlue*biggestGreen*biggestRed;
	gamePower[game] = power;
	Console.WriteLine(line);
}

// Console.WriteLine($"Total games matched: {matchedGames.Count}");

// var sumOfIds = matchedGames.Sum(x => Convert.ToInt32(x[5..]));

// Console.WriteLine($"Sum of ids: {sumOfIds}");

var sumOfPowerSets = gamePower.Values.Sum();
Console.WriteLine($"Sum of power sets: {sumOfPowerSets}");