package main

import "fmt"

type Creature struct {
	Name string
	Attack, Defense int
}

func NewCreature(name string, attack int defense int) *Creature{
	return &Creature{Name: name, Attack: attack, Defense: defense}
}

func (c *Creature) String() string{
	return fmt.Sprintf("%s (%d/%d)", c.Name, c.Attack, c.Defense)
}

type Modifier interface {
	Add(m Modifier)
	Handle()
}

type CreatureModifier struct {
	creature *Creature
	next Modifier
}

func ()  {
	
}

func main()  {
	
}