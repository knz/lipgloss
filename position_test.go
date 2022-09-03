package lipgloss

import (
	"fmt"
	"strings"
)

func revealNL(o string) string {
	// Make newlines visible.
	o = strings.ReplaceAll(o, "\n", "␤\n")
	o = strings.ReplaceAll(o, " ", "_")
	o = strings.ReplaceAll(o, "\t", ">-->")
	// Add a "no newline at end" marker if there was no newline at the end.
	if len(o) == 0 || o[len(o)-1] != '\n' {
		o += "🛇"
	}
	return o
}

func out(o string) {
	fmt.Println(revealNL(o))
	fmt.Println()
}

func Example_position() {
	fmt.Println("horizontal:")
	out(PlaceHorizontal(10, Left, "hello"))
	out(PlaceHorizontal(10, Right, "hello"))
	out(PlaceHorizontal(10, Center, "hello"))
	fmt.Println("vertical:")
	out(PlaceVertical(3, Top, "hello"))
	out(PlaceVertical(3, Bottom, "hello"))
	out(PlaceVertical(3, Center, "hello"))
	fmt.Println("both:")
	out(Place(10, 3, Left, Top, "hello"))
	out(Place(10, 3, Center, Top, "hello"))
	out(Place(10, 3, Right, Top, "hello"))
	out(Place(10, 3, Left, Center, "hello"))
	out(Place(10, 3, Center, Center, "hello"))
	out(Place(10, 3, Right, Center, "hello"))
	out(Place(10, 3, Left, Bottom, "hello"))
	out(Place(10, 3, Center, Bottom, "hello"))
	out(Place(10, 3, Right, Bottom, "hello"))

	// Output:
	// horizontal:
	// hello_____🛇
	//
	// _____hello🛇
	//
	// __hello___🛇
	//
	// vertical:
	// hello␤
	// _____␤
	// _____🛇
	//
	// _____␤
	// _____␤
	// hello🛇
	//
	// _____␤
	// hello␤
	// _____🛇
	//
	// both:
	// hello_____␤
	// __________␤
	// __________🛇
	//
	// __hello___␤
	// __________␤
	// __________🛇
	//
	// _____hello␤
	// __________␤
	// __________🛇
	//
	// __________␤
	// hello_____␤
	// __________🛇
	//
	// __________␤
	// __hello___␤
	// __________🛇
	//
	// __________␤
	// _____hello␤
	// __________🛇
	//
	// __________␤
	// __________␤
	// hello_____🛇
	//
	// __________␤
	// __________␤
	// __hello___🛇
	//
	// __________␤
	// __________␤
	// _____hello🛇
}

func Example_ws_chars() {
	out(Place(10, 3, Center, Center, "hello",
		WithWhitespaceChars("."),
	))
	out(Place(10, 3, Center, Center, "hello",
		WithWhitespaceChars("☃"),
	))
	out(Place(10, 3, Center, Center, "hello",
		WithWhitespaceChars("12"),
	))

	// Output:
	// ..........␤
	// ..hello...␤
	// ..........🛇
	//
	//☃☃☃☃☃☃☃☃☃☃␤
	//☃☃hello☃☃☃␤
	//☃☃☃☃☃☃☃☃☃☃🛇
	//
	// 1212121212␤
	// 12hello121␤
	// 1212121212🛇
}
