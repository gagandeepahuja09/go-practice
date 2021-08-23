package main

type department interface {
	setNext(department)
	execute(*patient)
}
