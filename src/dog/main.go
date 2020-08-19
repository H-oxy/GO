//  Package dog Allows the conversion from dog years to human years and the other way around.
package dog


//  Dyears converts from dog years to human years.
func Dyears(d int) int {
	return d * 7
}


// Hyear converts from human years to dog years.
func Hyear(i int) int{
	return i / 7


}

