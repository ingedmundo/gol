package main

import "testing"

var c = Cell{ status: true }

func TestIsAlive(t *testing.T){
  if c.IsAlive() == false{
    t.Error("Cell must be alive when initialized")
  }
}

func TestDie(t *testing.T){
  c.Die()

  if c.IsAlive() == true {
    t.Error("Cell must be dead by now")
  }
}

func TestLive(t *testing.T){
  c.Die()
  c.Live()

  if c.IsAlive() == false {
    t.Error("Cell must be alive by now")
  }
}

func TestUniverse(t *testing.T){
  u := new(Universe)

  if u.Statistics().Alive > 0 {
    t.Error("Universe houdl not have any alive cell")
  }

}

func TestBigBang(t *testing.T){
  u := new (Universe)

  u.BigBang(10)

  if u.Statistics().Alive == 0 {
    t.Error("Universe should have some alive cells by now")
  }

}

