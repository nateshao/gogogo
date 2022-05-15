package main

func main() {

}

//func BenchmarkInline(b *testing.B) {
//	x := genInteger()
//	y := genInteger()
//	for i := 0; i < b.N; i++ {
//		addInLine(x, y)
//	}
//}
//func addInline(a, b int) int {
//	return a + b
//}
//func BenchmarkInlineDisabLed(b *testing.B) {
//	x := genInteger()
//	y := genInteger()
//	for i := 0; i < b.N; i++ {
//		addNoIntine(x, y)
//	}
//}
//
////go :noinline
//func addNoInLine(a, b int) int {
//	return a + b
//}
