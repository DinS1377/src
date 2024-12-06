package goroutines

import (
	"fmt"
	"go_project/packages/interfaces"
	"sync"

	"github.com/gen2brain/beeep"
)

// Информация о животных
type AnimalInfo struct {
	Name        string
	HighSpeed   string
	Size        string
	UniqueTrait string
}

func Animals_g() {
	animalTypes := []string{"lion", "cat", "dog", "karas", "bird"}

	var wg sync.WaitGroup
	infoChannel := make(chan AnimalInfo)

	factory := &interfaces.ConcreteAnimalFactory{}

	// горутины
	for _, animalType := range animalTypes {
		wg.Add(1)
		go func(aType string) {
			defer wg.Done()
			animal := factory.CreateAnimal(aType)
			if animal == nil {
				fmt.Printf("Unknown animal type: %s\n", aType)
				return
			}
			processAnimal(animal, aType, infoChannel)
		}(animalType)
	}

	// закртытие канала
	go func() {
		wg.Wait()
		close(infoChannel)
	}()

	// Получаем информацию из канала
	for info := range infoChannel {
		fmt.Printf("Animal: %s\nHighSpeed: %s\nSize: %s\nUniqueTrait: %s\n\n",
			info.Name, info.HighSpeed, info.Size, info.UniqueTrait)
		// Отправка уведомлений
		go func(info AnimalInfo) {
			err := beeep.Notify(
				fmt.Sprintf("Animal Processed: %s", info.Name),
				fmt.Sprintf("HighSpeed: %s, Size: %s, UniqueTrait: %s", info.HighSpeed, info.Size, info.UniqueTrait),
				"")
			if err != nil {
				fmt.Printf("Failed to send notification for %s: %v\n", info.Name, err)
			}
		}(info)
	}
}

// Обработка животного
func processAnimal(animal interfaces.Animal, animalType string, ch chan AnimalInfo) {
	animal.MakeSound()
	animal.Move()
	animal.Feed()
	animal.Sleep()

	// Загружаем информации
	var info AnimalInfo
	switch animalType {
	case "lion":
		info = AnimalInfo{"Lion", "80 km/h", "Large", "King of the jungle"}
	case "cat":
		info = AnimalInfo{"Cat", "48 km/h", "Small", "Can climb trees"}
	case "dog":
		info = AnimalInfo{"Dog", "45 km/h", "Medium", "Recognizes human emotions"}
	case "karas":
		info = AnimalInfo{"Karas", "5 km/h", "Small", "Can survive in cold water"}
	case "bird":
		info = AnimalInfo{"Bird", "50 km/h", "Small", "Can fly high"}
	}
	ch <- info
}
