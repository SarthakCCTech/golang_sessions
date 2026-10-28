package main
import "fmt"


type DataProcessor interface {
    Process(data []byte) error
}

type DataValidator interface {
    Validate(data []byte) error
}

type DataHandler interface {
    DataProcessor
    DataValidator
}

type JSONHandler struct{}

func (j JSONHandler) Process(data []byte) error {
    fmt.Println("Processing JSON data...")
    return nil
}

func (j JSONHandler) Validate(data []byte) error {
    fmt.Println("Validating JSON format...")
    return nil
}