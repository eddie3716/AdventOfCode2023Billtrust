namespace AdventOfCodeParser;

public interface ILexer
{
    bool Next();

    Token GetToken();
}