// See https://aka.ms/new-console-template for more information
using System.IO;

var matrix = new List<List<char>>();
var adjacentNumbers = new List<int>();

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
			Console.WriteLine($"Found digit {matrix[row][col]} at {row},{col}");
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
						adjacentNumbers.Add(number);
						Console.WriteLine($"Found adjacent number {number} with symblol {matrix[scanRow][scanCol]} at {scanRow},{scanCol}");
						goto foundSymbol;
					}
				}
			}
			foundSymbol:
			col +=digits.Count-1;
		}
	}
}
// foreach(var number in adjacentNumbers)
// {
// 	Console.WriteLine(number);
// }

var answerPart1 = adjacentNumbers.Sum();
Console.WriteLine($"Answer Day 3 Part1: {answerPart1}");