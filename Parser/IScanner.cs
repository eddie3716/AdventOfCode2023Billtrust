namespace AdventOfCodeParser;
public interface IScanner<T, S>
{
    bool Next();

    bool Previous();

    T Get();

    int CurrentIndex { get; }

    S Source { get; }
}