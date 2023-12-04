// See https://aka.ms/new-console-template for more information
using System.IO;

var matrix = new List<List<char>>();
var potentialGears = new List<Tuple<int, int, int>>();

foreach (var line in File.ReadLines("Input.txt"))
{
	matrix.Add(line.ToList());	
}

for (int row = 0; row < matrix.Count; row++)
{
	for (int col = 0; col < matrix[row].Count; col++)
	{
		if (char.IsDigit(matrix[row][col]))
		{
			//Console.WriteLine($"Found digit {matrix[row][col]} at {row},{col}");
			var digits = new List<char>();

			for (int digitCol=col; digitCol < matrix[row].Count && char.IsDigit(matrix[row][digitCol]);digitCol++)
			{
				digits.Add(matrix[row][digitCol]);
			}

			var number = int.Parse(new string(digits.ToArray()));
			for (int scanRow = Math.Max(0, row-1); scanRow < Math.Min(matrix.Count, row+2); scanRow++)
			{
				for (int scanCol = Math.Max(0, col-1); scanCol < Math.Min(matrix[scanRow].Count, col + digits.Count+1); scanCol++)
				{
					if (matrix[scanRow][scanCol] != '.' && !char.IsDigit(matrix[scanRow][scanCol]))
					{
						potentialGears.Add(new Tuple<int, int, int>(number, scanRow, scanCol));
						//Console.WriteLine($"Found adjacent number {number} with symblol {matrix[scanRow][scanCol]} at {scanRow},{scanCol}");
						goto foundSymbol;
					}
				}
			}
			foundSymbol:
			col +=digits.Count-1;
		}
	}
}

var gearPowers = new List<int>();
for (int i = 0; i < potentialGears.Count; i++)
{
	var potentialGear = potentialGears[i];
	if (i != potentialGears.Count-1)
	{
		int checkafter = i+1;
		if (potentialGears.Skip(checkafter).Any(x => x.Item2 == potentialGear.Item2 && x.Item3 == potentialGear.Item3))
		{
			Console.WriteLine($"Found potential gear {potentialGear.Item1} at {potentialGear.Item2},{potentialGear.Item3}");
			var gearBuddy = potentialGears.Skip(checkafter).First(x => x.Item2 == potentialGear.Item2 && x.Item3 == potentialGear.Item3);
			Console.WriteLine($"Found gear buddy {gearBuddy.Item1} at {gearBuddy.Item2},{gearBuddy.Item3}");
			
			gearPowers.Add(potentialGear.Item1 * gearBuddy.Item1);
		}
	}
}

var answerPart2 = gearPowers.Sum();
Console.WriteLine($"Answer Day 3 Part2: {answerPart2}");