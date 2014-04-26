package main

import color "../"
import "fmt"

func main() {
	fmt.Print("RGB2HSL: ")
	fmt.Println(color.RGB2HSL(121, 167, 22))
	fmt.Print("RGB2HSL: ")
	fmt.Println(color.RGB2HSL(69, 209, 237))
	fmt.Print("RGB2HSL: ")
	fmt.Println(color.RGB2HSL(254, 207, 37))
	fmt.Print("RGB2HSL: ")
	fmt.Println(color.RGB2HSL(122, 167, 255))
	fmt.Print("RGB2HSL: ")
	fmt.Println(color.RGB2HSL(255, 255, 255))
	fmt.Print("HSV2RGB: ")
	fmt.Println(color.HSV2RGB(184, 71, 81))
	fmt.Print("HEX2RGB: ")
	fmt.Println(color.HEX2RGB("#C3C3FF"))
	fmt.Print("RGB2HEX(HEX2RGB): ")
	fmt.Println(color.RGB2HEX(color.HEX2RGB("c3")))
	fmt.Print("HSL2RGB(RGB2HSL): ")
	fmt.Println(color.HSL2RGB(color.RGB2HSL(200, 100, 50)))
	fmt.Print("RGB2HSL(HSV2RGB): ")
	fmt.Println(color.RGB2HSL(color.HSV2RGB(184, 71, 81)))
	fmt.Print("RGB2HSL(CMYK2RGB): ")
	fmt.Println(color.RGB2HSL(color.CMYK2RGB(215, 55, 0, 20)))
}
