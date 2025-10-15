SmallestK.SmallestKInList<long>([1, 2, 3], 0);
SmallestK.SmallestKInList<long>([1, -2, 3], 1);
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
            return new T[]{};

        if (k >= arr.Length)
            return arr;

        int left = 0;
        int right = arr.Length - 1;
        while (true)
        {
            int pivotIndex = Partition(arr, left, right);
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

        return arr.Take(k).ToArray();
    }

    private static int Partition<T>(T[] arr, int left, int right)
    {
        int pivotIndex = rand.Next(left, right + 1);
        var pivot = arr[pivotIndex];
        (arr[pivotIndex], arr[right]) = (arr[right], arr[pivotIndex]);

        int storeIndex = left;

        for (int i = left; i < right; i++)
        {
            var a = long.Parse(arr[i].ToString());
            var b = long.Parse(pivot.ToString());
            if (a < b)
            {
                (arr[storeIndex], arr[i]) = (arr[i], arr[storeIndex]);
                storeIndex++;
            }
        }
        (arr[storeIndex], arr[right]) = (arr[right], arr[storeIndex]);
        return storeIndex;
    }
}

public class Kata
{
    public static T[] SmallestKElements<T>(T[] arr, int k) where T : IComparable<T>
    {
        if (k <= 0)
        {
            return new T[]{};
        }
        
        var maxHeap = new PriorityQueue<T, long>();

        foreach (var num in arr)
        {
            var tempNum = long.Parse(num.ToString());
            if (maxHeap.Count < k)
            {
                maxHeap.Enqueue(num, -tempNum); 
            }
            else
            {
                long largestOfK = long.Parse(maxHeap.Peek().ToString()); 
                
                if (tempNum < largestOfK)
                {
                    maxHeap.Dequeue();
                    
                    maxHeap.Enqueue(num, -tempNum);
                }
            }
        }
        return maxHeap.UnorderedItems.Select(item => item.Element).ToArray();
    }
}
