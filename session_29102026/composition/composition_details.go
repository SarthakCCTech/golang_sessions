package main
import "fmt"

type details struct{
genre string
genreRating string
reviews string
}

type game struct{
name string
price string 

details
}

func main(){

g := game{}
g.name = "COC"
g.genre = "first world"
g.genreRating = "5 Star"
g.reviews = "Overall Good"
g.price = "$1000"

fmt.Printf("Name: %s\n", g.name)

}
