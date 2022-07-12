# Awesome Project
---

The purpose of this repository is mostly to understand how the implicit implementation of interfaces in Go works, and how this can be used for unit testing in which dependencies are mocked.

It also shows how simple a Go program can look like, especially for engineers that come from a C#/Java or other object oriented languages background.

Although Go has OO capabilities as well (there's no concept of `class` but objects, composition and encapsulation are present), Go is defined by many people smarter than me as a `Data Oriented language`, and I completely agree with that.

### How is this solution structured?

Two main folders are present: `cmd`, that contains our main Go file that will start the program, and `internal`, that contains two packages we don't want to be exported outside of the application (hence they residing inside of the `internal` folder). The first package is called `repo` and is, well, a repo of sorts; the only thing it does now is return a message. The other package, `user`, is responsible to orchestrating the user operations in our program. This package also defines a `userStorer` interface, that `repo` knows nothing about, but that it also satisfies implicitly. This way we can use `userStorer` inside the `user` package to add a layer of abstraction between `service` and `repo`. From a program perspective of using concrete types (check `main.go` inside the `cmd` folder) nothing changes, but if you open `service_test.go` in the `user` package, you can see how we can leverage this implicitly satisfied interface to unit-test our service (using table tests, my favorite kind!)

### Plans for the future

I don't know if I'm ever gonna extend this repository. This is mostly used to remind me of some patterns as well, so I might add those patterns to it in the future; in its current state, though, I believe it already accomplishes the mission of making engineers from other backgrounds understand how powerful and flexible the interface model of Go is.