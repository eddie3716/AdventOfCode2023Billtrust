using System.Collections.Generic;

namespace AdventOfCodeParser;
public interface IParser
{
    HashSet<Token> GetTokens(List<List<char>> rawMap);
}