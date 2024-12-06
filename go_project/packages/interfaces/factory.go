package interfaces

// AnimalFactory - интерфейс фабрики животных
type AnimalFactory interface {
	CreateAnimal(animalType string) Animal
}

// ConcreteAnimalFactory - конкретная реализация фабрики
type ConcreteAnimalFactory struct{}

// CreateAnimal - создает животное по типу
func (f *ConcreteAnimalFactory) CreateAnimal(animalType string) Animal {
	switch animalType {
	case "lion":
		return &Lion{}
	case "cat":
		return &Cat{}
	case "dog":
		return &Dog{}
	case "karas":
		return &Karas{}
	case "bird":
		return &Bird{}
	default:
		return nil
	}
}
