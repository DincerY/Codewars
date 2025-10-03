using Solution;
Kata.IsValidIp("12.34.56 .1");
Kata.IsValidIp("12.034.056.0");
Kata.IsValidIp("0.0.0.0");
Kata.IsValidIp("12.255.56.1");
Kata.IsValidIp("137.255.156.100");
Kata.IsValidIp("");
Kata.IsValidIp("abc.def.ghi.jkl");
Kata.IsValidIp("123.456.789.0");
Kata.IsValidIp("12.34.56");

Console.WriteLine("Hello, World!");

namespace Solution
{
    class Kata
    {
        public static bool IsValidIp(string ipAddres)
        {
            var partition = ipAddres.Split('.');
            if (partition.Length != 4)
            {
                return false;
            }

            foreach (var p in partition)
            {
                if (p.Contains(' '))
                {
                    return false;
                }

                if (!int.TryParse(p, out var a))
                {
                    return false;
                }
                if (a <= 255 && a >= 0)
                {
                    if (p.Length > 1 && p[0] == '0')
                    {
                        return false;
                    }
                }
                else
                {
                    return false;
                }
            }
            return true;
        }
    }
}