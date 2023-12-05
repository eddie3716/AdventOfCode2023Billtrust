// See https://aka.ms/new-console-template for more information
using ConsoleApp1;
using System.Linq;
using System.Security.Cryptography.X509Certificates;

List<int> ints = new();
var translation = new Dictionary<string, int>
{
    { "1", 1 },
    { "2", 2 },
    { "3", 3 },
    { "4", 4 },
    { "5", 5 },
    { "6", 6 },
    { "7", 7 },
    { "8", 8 },
    { "9", 9 },
    { "one", 1 },
    { "two", 2 },
    { "three", 3 },
    { "four", 4 },
    { "five", 5 },
    { "six", 6 },
    { "seven", 7 },
    { "eight", 8 },
    { "nine", 9 }
};
var tokens = new []{"1", "2", "3", "4", "5", "6", "7", "8", "9", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"};
var reversedTokens = tokens.Select(x => x.BetterReverse());
var file = File.OpenText("PuzzleInput.txt");

var tokenFinder = (string theString, IEnumerable<string> theTokens) =>
{
    for (var i = 0; i < theString.Length; i++)
    {
        foreach(var token in theTokens)
        {
            if (theString[i..].StartsWith(token))
            {
                return token;
            }
        }
    }

    throw new Exception("what kind of crapped up logic is this?");
};

string? currentLine;
currentLine = file.ReadLine();
var count = 0;
while (currentLine != null)
{
    var tensDigit = translation[tokenFinder(currentLine, tokens)];
    var onesDigit = translation[tokenFinder(currentLine.BetterReverse(), reversedTokens).BetterReverse()];
    int theDigit = tensDigit*10 + onesDigit;
    ints.Add(theDigit);
    Console.WriteLine($"{currentLine}...{theDigit}");
    currentLine = file.ReadLine();
    count++;
    //if (count > 4)
    //    break;
}

Console.WriteLine(ints.Sum());