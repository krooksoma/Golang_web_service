package main

type person struct{
	firstName string
	lastName string
	contact contactInfo
}

type contactInfo struct{
	phone string
	zipCode string
}