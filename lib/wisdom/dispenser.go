package wisdom

import (
	"encoding/json"
	"github.eagleview.com/engineering/pa-gosdk/errors"
	"io/ioutil"
	"math/rand"
)

//-------------------------------------------------------------------------------------------------

type Dispenser interface {
	Count() int
	Get(n int) Quote
	Random() Quote
}

func NewDispenser(quotes []Quote) Dispenser {
	return &memoryDispenser{quotes: quotes}
}

//-------------------------------------------------------------------------------------------------

type memoryDispenser struct {
	quotes []Quote
}

func (md *memoryDispenser) Count() int {
	return len(md.quotes)
}

func (md *memoryDispenser) Get(n int) Quote {
	return md.quotes[n]
}

func (md *memoryDispenser) Random() Quote {
	return md.quotes[rand.Intn(len(md.quotes))]
}

//-------------------------------------------------------------------------------------------------

func FromFile(filename string) (Dispenser, error) {

	quotes := []Quote{}

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, errors.Wrap(err, "ioutil.ReadFile failed")
	}

	err = json.Unmarshal(bytes, &quotes)
	if err != nil {
		return nil, errors.Wrap(err, "json.Unmarshal failed")
	}

	return NewDispenser(quotes), nil

}

//-------------------------------------------------------------------------------------------------
