package main

type SpecialArray []interface{}

// Tip: Each element of `array` can either be cast
// to `SpecialArray` or `int`.

func ProductSum(array []interface{}) int {
	return helper(array, 1)
}

func helper(array SpecialArray, multiplier int) int {
	sum := 0
	for _, el := range array {
		if cast, ok := el.(SpecialArray); ok {
			sum += helper(cast, multiplier+1)
		} else if cast, ok := el.(int); ok {
			sum += cast
		}
	}
	return sum * multiplier
}

/*
func ProductSum(array []interface{}) int {
	indx :=0
	depth :=1
	return ProductSumHelper(array,indx,depth)
}

func ProductSumHelper(array []interface{}, indx int,depth int) int {
	if indx == len(array)-1{
		i,isSpecial := array[indx].(SpecialArray)
		if isSpecial{
			temp := i
			return (depth + 1) *  ProductSumHelper(temp,0,depth + 1)
		}else{
			return array[indx].(int)
		}
	}
	j, isInt := array[indx].(int)
	if isInt{
		return j + ProductSumHelper(array,indx + 1,depth)
	}else{
		temp := array[indx].(SpecialArray)
		return (depth + 1) *  ProductSumHelper(temp,0,depth + 1) + ProductSumHelper(array,indx + 1,depth)
	}
} */
