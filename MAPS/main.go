package main

import (
	"fmt"
)

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"black": "#000000",
		"white": "#ffffff",
	}
	// delete(colors, "red")

	// var colors map[string]string
	// colors := make(map[string]string)
	// colors["white"] = "#fffffff"
	// fmt.Println(colors)

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, hex)
	}
}

// What is the different between Map and Struct ?
// Feature      	struct                        map
// Defined at	    Compile time	              Runtime
// Key/Field names	Fixed (field names)	          Dynamic (map keys)
// Type safety	    ✅ Yes	                    ❌ No (especially with interface{})
// Access syntax	obj.Field	                  map["key"]
// Best for     	Structured data / entities	  Dynamic/unpredictable data
// Performance	    Faster (compiler optimized)	  Slower (more overhead