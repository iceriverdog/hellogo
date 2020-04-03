package main
import "fmt"
func sortNums(arr []int, L, R int) {

    if L >= R{
        return
    }
   left, right := L, R
   pivot := arr[left]

   for {
      for left < right && arr[right] >= pivot{
         right--
      }
      if left < right{
         arr[left] = arr[right]
      }
      for left < right && arr[left] <= pivot{
         left++
      }
      if left < right{
         arr[right] = arr[left]
      }
      if left >= right{
         arr[left] = pivot
         break
      }
   }
   sortNums(arr, L, left-1)
   sortNums(arr, left+1, R)
}
func main()  {
   arr := []int{1,-3,5,-7,9}
   sortNums(arr, 0, len(arr) - 1)
   fmt.Println(arr)

}