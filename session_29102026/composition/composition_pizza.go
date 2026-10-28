
package main

import "fmt"

type Pizza struct {
    Name        string
    Size        string
    Toppings    []string
    IsDelicious bool
}

func pizzaStyle(p Pizza) string {
    return p.Name + " pizza is a " + p.Size +
        " pizza with toppings of " + fmt.Sprint(p.Toppings)
}

func main() {
    myPizza := Pizza{
        Name:        "Margherita",
        Size:        "medium",
        Toppings:    []string{"tomatoes", "mozzarella", "basil"},
        IsDelicious: true,
    }

    fmt.Println(pizzaStyle(myPizza))
}
