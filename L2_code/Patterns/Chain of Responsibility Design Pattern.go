package main

import (
	"fmt"
)

type departement interface {
	execute(*patient)
	setNext(departement)
}

type reception struct {
	next departement
}

func (r *reception) execute(p *patient) {
	if p.registrationDone {
		fmt.Println("Pacient have registration, done")
		r.next.execute(p)
		return
	}

	fmt.Println("Create registration pacient")
	p.registrationDone = true
	r.next.execute(p)
}

func (r *reception) setNext(next departement) {
	r.next = next
}

type doctor struct {
	next departement
}

func (d *doctor) execute(p *patient) {
	if p.doctorCheckUpDone {
		fmt.Println("Doctor checkoup already done")
		d.next.execute(p)
		return
	}

	fmt.Println("Doctor checking pacient")
	p.doctorCheckUpDone = true
	d.next.execute(p)
}

func (d *doctor) setNext(next departement) {
	d.next = next
}

type medecine struct {
	next departement
}

func (m *medecine) execute(p *patient) {
	if p.medicineDone {
		fmt.Println("Medicine done")
		m.next.execute(p)
		return
	}

	fmt.Println("Given medicine to pacient")
	p.medicineDone = true
	m.next.execute(p)
}

func (m *medecine) setNext(next departement) {
	m.next = next
}

type cashier struct {
	next departement
}

func (c *cashier) execute(p *patient) {
	if p.paymentDone {
		fmt.Println("Payment done")
	} else {
		p.paymentDone = true
		fmt.Println("Cashier getting money from pacient")
	}
}

func (c *cashier) setNext(next departement) {
	c.next = next
}

type patient struct {
	name              string
	registrationDone  bool
	doctorCheckUpDone bool
	medicineDone      bool
	paymentDone       bool
}

func main() {
	cashier := &cashier{}
	medical := &medecine{}

	medical.setNext(cashier)
	
	doctor := &doctor{}
	doctor.setNext(medical)
	
	reception := &reception{}
	reception.setNext(doctor)
	patient := &patient{
		name: "Anatoliy",
	}

	reception.execute(patient)
	fmt.Println()
	reception.execute(patient)
}

// Шаблон проектирования цепочки ответственности — это поведенческий шаблон проектирования.
// Он позволяет создать цепочку обработчиков запросов. Для каждого входящего
// запроса он проходит черезцепочку и каждый из обработчиков:
// Обрабатывает запрос или пропускает обработку.
// Решает, передать ли запрос следующему обработчику в цепочке или нет.
