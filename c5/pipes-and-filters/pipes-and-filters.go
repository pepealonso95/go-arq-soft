package main

import (
	"encoding/json"
	"fmt"
	"time"

	"gopkg.in/yaml.v2"
)

type Filter func(any) (any, error)

type Pipeline struct {
	filters []Filter
}

func (p *Pipeline) Use(f ...Filter) {
	p.filters = append(p.filters, f...)
}

func (p Pipeline) RunChan(input any, done chan struct{}) (err error) {
	defer func() { done <- struct{}{} }()
	return p.Run(input)
}

func (p Pipeline) Run(input any) (err error) {
	for _, f := range p.filters {
		input, err = f(input)
		if err != nil {
			return err
		}
	}
	return err
}

func FilterJSON(input any) (any, error) {
	// Primero chequeo que el filtro reciba un input de tipo apropiado
	// En este caso se usa []byte para que sea compatible con json unmarshal
	if data, ok := input.([]byte); ok {
		var myData map[string]any
		err := json.Unmarshal(data, &myData)
		return myData, err
	} else {
		return nil, fmt.Errorf("Not a valid type, expected []byte")
	}
}

func FilterYAML(input any) (any, error) {
	if data, ok := input.([]byte); ok {
		var myData map[string]any
		err := yaml.Unmarshal(data, &myData)
		return myData, err
	} else {
		return nil, fmt.Errorf("Not a valid type, expected []byte")
	}
}

func FilterMultiplyTwo(input interface{}) (any, error) {
	switch input.(type) {
	case int:
		return input.(int) * 2, nil
	case float64:
		return input.(float64) * 2, nil
	default:
		return nil, fmt.Errorf("Not a valid type, expected int or float64")
	}
}

func FilterAddOne(input interface{}) (any, error) {
	switch input.(type) {
	case int:
		return input.(int) + 1, nil
	case float64:
		return input.(float64) + 1, nil
	default:
		return nil, fmt.Errorf("Not a valid type, expected int or float64")
	}
}

func main() {
	p := Pipeline{}
	for i := 0; i < 1000; i++ {
		p.Use(FilterMultiplyTwo, FilterAddOne)
	}
	now := time.Now()
	for i := 0; i < 100000; i++ {
		p.Run(i)
	}
	fmt.Println("Sequential")
	fmt.Println(time.Now().Sub(now))

	now = time.Now()
	done := make(chan struct{}, 100)
	for i := 0; i < 100000; i++ {
		go p.RunChan(i, done)
	}
	for i := 0; i < 100000; i++ {
		<-done
	}
	fmt.Println("With concurrency")
	fmt.Println(time.Now().Sub(now))
}
