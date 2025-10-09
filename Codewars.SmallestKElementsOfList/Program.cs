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
    private static Random rand = new();
    public static T[] SmallestKInList<T>(T[] arr, int k) where T : IComparable<T>
    {
        if (k <= 0)
        {
            return new T[]{};
        }

        if (k >= arr.Length)
        {
            return arr;
        }

        int left = 0;
        int right = arr.Length - 1;
        while (true)
        {
            int pivotIndex = Partition();
            if (pivotIndex == k)
            {
                break;
            }
            else if (pivotIndex < k)
            {
                left = pivotIndex + 1;
            }
            else
            {
                right = pivotIndex - 1;
            }
        }
        Array.Sort(arr);
        T[] result = new T[k];
        Array.Copy(arr, result, k);
        return result;
    }

    private static int Partition<T>(T[] arr, int left, int right)
    {
        int pivotIndex = rand.Next(left, right + 1);
        var pivot = arr[pivotIndex];
        (arr[pivotIndex], arr[right]) = (arr[right], arr[pivotIndex]);

        int storeIndex = left;

        for (int i = left; i < right; i++)
        {
            if (arr[i] < pivot)
            {
                (arr[storeIndex], arr[i]) = (arr[i], arr[storeIndex]);
                storeIndex++;
            }
        }
        (arr[storeIndex], arr[right]) = (arr[right], arr[storeIndex]);
        return storeIndex;
    }
}
