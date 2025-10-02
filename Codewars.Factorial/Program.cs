Kata.Factorial(0);
Kata.Factorial(1);
Kata.Factorial(2);
Kata.Factorial(3);

Console.WriteLine("Hello, World!");

public static class Kata
{
    public static int Factorial(int n)
    {
        if (n < 0 || n > 12)
        {
            throw new ArgumentOutOfRangeException();
        }

        int res = 1;
        void Recursive(int val)
        {
            if (val > 1)
            {
                res *= val;
                Recursive(val-1);
            }
        }
        Recursive(n);
        return res;
    }
    
    
    public static int Factorial2(int n)
    {
        if (n < 0 || n > 12)
        {
            throw new ArgumentOutOfRangeException();
        }
        int Recursive(int val)
        {
            if (val == 0)
            {
                return 1;
            }
            return val * Recursive(val - 1);
        }
        return Recursive(n);
    }
    public static int Factorial3(int n)
    {
        return (n < 0 || n > 12) ? throw new ArgumentOutOfRangeException() : n > 1 ? n * Factorial3(n-1) : 1;
    }
}