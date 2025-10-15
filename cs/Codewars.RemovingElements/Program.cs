Kata.RemoveEveryOther(new object[]{"Hello", "Goodbye", "Hello Again"});
Kata.RemoveEveryOther(new object[] { new object[] { 1, 2 } }); 
Kata.RemoveEveryOther(new object[] { 1, 2, 3, 4, 5, 6, 7, 8, 9, 10 });

Console.WriteLine("Hello, World!");

public static class Kata
{
    public static object[] RemoveEveryOther(object[] arr)
    {
        var list = new List<object>();
        for (int i = 0; i < arr.Length; i += 2)
        {
            list.Add(arr[i]);	
        }
        return list.ToArray();
    }
}