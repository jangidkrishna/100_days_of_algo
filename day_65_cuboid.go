package main

import "fmt"

func cuboid_draw(drawX, drawY, drawZ int) {
	fmt.Printf("Cuboid %d %d %d:\n", drawX, drawY, drawZ)
	cube_line(drawY+1, drawX, 0, "+-")
	for i := 1; i <= drawY; i++ {
		cube_line(drawY-i+1, drawX, i-1, "/ |")
	}
	cube_line(0, drawX, drawY, "+-|")
	for i := 4*drawZ - drawY - 2; i > 0; i-- {
		cube_line(0, drawX, drawY, "| |")
	}
	cube_line(0, drawX, drawY, "| +")
	for i := 1; i <= drawY; i++ {
		cube_line(0, drawX, drawY-i, "| /")
	}
	cube_line(0, drawX, 0, "+-\n")
}

func cube_line(n, drawX, drawY int, cubeDraw string) {
	fmt.Printf("%*s", n+1, cubeDraw[:1])
	for d := 9*drawX - 1; d > 0; d-- {
		fmt.Print(cubeDraw[1:2])
	}
	fmt.Print(cubeDraw[:1])
	fmt.Printf("%*s\n", drawY+1, cubeDraw[2:])
}

func main() {
	cuboid_draw(3, 7, 4)
}
