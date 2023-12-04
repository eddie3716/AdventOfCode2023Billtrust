namespace AdventOfCodeParser;

public struct Vertex
{
	public int XPosition { get; private set; }

    public int YPosition { get; private set; }

	public Vertex(int xPosition, int yPosition)
	{
		XPosition = xPosition;
		YPosition = yPosition;
	}

	public override string ToString()
	{
		return $"({XPosition}, {YPosition})";
	}

	public override bool Equals(object obj)
	{
		bool result = false;

		if (obj is Vertex vertex)
		{
			result = XPosition.Equals(vertex.XPosition) && YPosition.Equals(vertex.YPosition);
		}

		return result;
	}

	public override int GetHashCode() => HashCode.Combine(XPosition, YPosition);

	public static bool operator ==(Vertex left, Vertex right)
	{
		return left.Equals(right);
	}

	public static bool operator !=(Vertex left, Vertex right)
	{
		return !(left == right);
	}

	public static Vertex operator +(Vertex left, Vertex right)
	{
		return new Vertex(left.XPosition + right.XPosition, left.YPosition + right.YPosition);
	}

	public static Vertex operator -(Vertex left, Vertex right)
	{
		return new Vertex(left.XPosition - right.XPosition, left.YPosition - right.YPosition);
	}

	public static Vertex operator *(Vertex left, Vertex right)
	{
		return new Vertex(left.XPosition * right.XPosition, left.YPosition * right.YPosition);
	}

	public Vertex DistanceFrom(Vertex other)
	{
		return new Vertex(Math.Abs(XPosition - other.XPosition), Math.Abs(YPosition - other.YPosition));
	}

	public Vertex Magnitude()
	{
		return new Vertex(Math.Abs(XPosition), Math.Abs(YPosition));
	}

	public Vertex Normalize()
	{
		return new Vertex(Math.Sign(XPosition), Math.Sign(YPosition));
	}

	public Vertex RotateLeft()
	{
		return new Vertex(-YPosition, XPosition);
	}

	public Vertex RotateRight()
	{
		return new Vertex(YPosition, -XPosition);
	}

	public Vertex Rotate180()
	{
		return new Vertex(-XPosition, -YPosition);
	}

	public Vertex RotateLeft(int times)
	{
		Vertex result = this;

		for (int i = 0; i < times; i++)
		{
			result = result.RotateLeft();
		}

		return result;
	}

	public Vertex GetVertexInDirection(Direction direction)
	{
		return direction switch
		{
			Direction.Up => new Vertex(XPosition, YPosition - 1),
			Direction.Down => new Vertex(XPosition, YPosition + 1),
			Direction.Left => new Vertex(XPosition - 1, YPosition),
			Direction.Right => new Vertex(XPosition + 1, YPosition),
			Direction.UpLeft => new Vertex(XPosition - 1, YPosition - 1),
			Direction.UpRight => new Vertex(XPosition + 1, YPosition - 1),
			Direction.DownLeft => new Vertex(XPosition - 1, YPosition + 1),
			Direction.DownRight => new Vertex(XPosition + 1, YPosition + 1),
			_ => throw new ArgumentException($"Invalid direction {direction}", nameof(direction))
		};
	}

	public Vertex GetVertexInDirection(Direction direction, int distance)
	{
		Vertex result = this;

		for (int i = 0; i < distance; i++)
		{
			result = result.GetVertexInDirection(direction);
		}

		return result;
	}
}