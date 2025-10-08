SmallestK.SmallestKInList<long>([1, 2, 3], 0);
SmallestK.SmallestKInList<long>([1, 2, 3], 1);
SmallestK.SmallestKInList<long>([1, 2, 3], 2);
SmallestK.SmallestKInList<long>([1, 2, 3], 3);

SmallestK.SmallestKInList<long>([3,2,1], 1);
SmallestK.SmallestKInList<long>([3,2,1], 2);
SmallestK.SmallestKInList<long>([3,2,1], 3);

Console.WriteLine("Hello, World!");

class SmallestK
{
    public static T[] SmallestKInList<T>(T[] arr, int k) where T : IComparable<T>
    {
        Array.Sort(arr);
        T[] result = new T[k];
        Array.Copy(arr, result, k);
        return result;
    }
}
