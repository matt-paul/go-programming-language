### Leaning Go

This repo is mainly here to write as much go code as possible to become used to the language.
Almost everything here is copied from various sources on the web.

#### Exciting things

Go has the concept of a 'naked' return, however it can harm readability so only use in short functions if at all.

```func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
```

If an initializer for a var declaration is included, the type can be ommited.

& in front of variable name is used to retrieve the address of where this variable’s value is stored. That address is what the pointer is going to store.

- in front of a type name, means that the declared variable will store an address of another variable of that type (not a value of that type).

- in front of a variable of pointer type is used to retrieve a value stored at given address. In Go speak this is called dereferencing.
