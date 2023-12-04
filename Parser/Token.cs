using System;

namespace AdventOfCodeParser;
public class Token
{
    public string Name { get; private set; }

    public int TokenLength => Name.Length;

    public Vertex[] Vertices { get; private set; }

    public Token(string name, params Vertex[] vertices)
    {
        Name = name;
        Vertices = vertices;
    }

    public override string ToString()
    {
        return (string.IsNullOrEmpty(Name) ? string.Empty : Name);
    }

    public override bool Equals(object obj)
    {
        bool result = false;

        if (obj is Token token)
        {

            result = Name.Equals(token.Name) && Vertices.Equals(StructuralComparisons.StructuralEqualityComparer.Equals(Vertices));
        }


        return result;
    }

    public override int GetHashCode() => HashCode.Combine(Name, StructuralComparisons.StructuralEqualityComparer.GetHashCode(Vertices));


    public static bool operator ==(Token left, Token right)
    {
        return left.Equals(right);
    }

    public static bool operator !=(Token left, Token right)
    {
        return !(left == right);
    }
}