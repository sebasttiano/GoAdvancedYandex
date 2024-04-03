package main

import "fmt"

// Pizza2 — сложная система с простым интерфейсом.
type Pizza2 struct {
	site     Site
	kitchen  Kitchen
	delivery Delivery
}

// Site отвечает за получение заказа на сайте.
type Site struct {
}

func (s *Site) GetOrder() {
	fmt.Println("Оформляем заказ.")
}

func (s *Site) RedirectOrder() {
	fmt.Println("Отправляем заказ на кухню.")
}

func (s *Site) CloseOrder() {
	fmt.Println("Заказ выполнен.")
}

// Kitchen готовит пиццу.
type Kitchen struct {
}

func (k *Kitchen) CookPizza() {
	fmt.Println("Готовим пиццу.")
}

// Delivery занимается доставкой пиццы.
type Delivery struct {
}

func (d *Delivery) SearchCourier() {
	fmt.Println("Ищем курьера.")
}

func (d *Delivery) DeliverPizza() {
	fmt.Println("Доставляем пиццу.")
}

func (d *Delivery) GetPayment() {
	fmt.Println("Получаем оплату.")
}

// StartOrder — метод для заказа пиццы.
func (p *Pizza2) StartOrder() {
	p.site.GetOrder()
	p.site.RedirectOrder()
	p.kitchen.CookPizza()
	p.delivery.SearchCourier()
	p.delivery.DeliverPizza()
}

// FinishOrder — выполнение заказа.
func (p *Pizza2) FinishOrder() {
	p.delivery.GetPayment()
	p.site.CloseOrder()
}

func main() {
	pizza := &Pizza2{}

	pizza.StartOrder()
	pizza.FinishOrder()
}
