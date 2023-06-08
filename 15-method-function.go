package main

import "fmt"

type Profile struct {
	People []struct {
		Name string
		Id   int64
	}
}

func (p *Profile) ImportUser(name string, id int64) {
	p.People = append(p.People, struct {
		Name string
		Id   int64
	}{Name: name, Id: id})
}

func ImportUserFunc(p *Profile, name string, id int64) {
	p.People = append(p.People, struct {
		Name string
		Id   int64
	}{Name: name, Id: id})
}

func main() {
	p := Profile{People: []struct {
		Name string
		Id   int64
	}{{Name: "jack", Id: 22}}}

	p.ImportUser("Milo", 30)
	ImportUserFunc(&p, "Mina", 80)
	fmt.Printf("%+v\n", p)
}
