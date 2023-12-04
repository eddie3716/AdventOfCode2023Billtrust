// See https://aka.ms/new-console-template for more information
List<int> ints = new List<int>();

var file = File.OpenText("PuzzleInput.txt");

string? currentLine;
currentLine = file.ReadLine();
var count = 0;
while (currentLine != null)
{
    var tensDigit = currentLine.First(c => Char.IsNumber(c));
    var onesDigit = currentLine.Last(c => Char.IsNumber(c));
    int theDigit = (tensDigit - '0')*10 + onesDigit - '0';
    ints.Add(theDigit);
    Console.WriteLine($"{currentLine}...{theDigit}");
    currentLine = file.ReadLine();
    count++;
    //if (count > 4)
    //    break;
}

Console.WriteLine(ints.Sum());