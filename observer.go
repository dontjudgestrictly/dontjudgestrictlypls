package main
import "fmt"

type Deistvie interface {
	Podpisatsa(Potrebitel)
	Otpisatsa(Potrebitel)
	Uvedomlyat()
}


type Potrebitel interface {
	Update(pubName string)
	GetImya() string
}



type Magazine struct {
	Imya        string
	Potrebiteli map[string]Potrebitel
}

func (pub *Magazine) Podpisatsa(potrebitel Potrebitel) {
	pub.Potrebiteli[potrebitel.GetImya()] = potrebitel
}

func (pub *Magazine) Otpisatsa(consumer Potrebitel) {
	delete(pub.Potrebiteli, consumer.GetImya())
}

func (pub *Magazine) Uvedomlyat() {
	for _, consumer := range pub.Potrebiteli {
		consumer.Update(pub.Imya)
	}
}



type Subscriber struct {
	Imya string
}

func (cons *Subscriber) Update(pubName string) {
	fmt.Printf("Sending to subscriber %s from publisher %s\n", cons.GetImya(), pubName)
}

func (cons *Subscriber) GetImya() string {
	return cons.Imya
}



func main() {
	sub1 := &Subscriber{Imya: "Timon"}
	sub2 := &Subscriber{Imya: "Asel"}

	channel := Magazine{
		Imya:        "Dress Magazine",
		Potrebiteli: map[string]Potrebitel{},
	}
	channel.Podpisatsa(sub1)
	channel.Podpisatsa(sub2)
	channel.Uvedomlyat()
}
