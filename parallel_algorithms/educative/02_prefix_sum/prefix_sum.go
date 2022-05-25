package _2_prefix_sum

func PrefixSum(myArray, myOutput []int, parent chan int) {
	if len(myArray) < 2 {
		parent <- myArray[0]
		myOutput[0] = myArray[0] + <-parent
	} else if len(myArray) < 1 {
		parent <- 0
		<-parent
	} else {
		mid := len(myArray) / 2
		left, right := make(chan int), make(chan int)

		go PrefixSum(myArray[:mid], myOutput[:mid], left)
		go PrefixSum(myArray[mid:], myOutput[mid:], right)

		leftSum := <-left
		parent <- leftSum + <-right
		fromLeft := <-parent
		left <- fromLeft
		right <- fromLeft + leftSum
		<-left
		<-right
	}

	parent <- 0
}
