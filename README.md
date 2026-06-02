# Object Definition Language
A Domain-Specific-Language for generating Classes/Structs in other languages

## What is this project about?
ODL is a tiny, declarative Language for defining an object and transpiling it to many different targets, such as Java, C#, Python, Rust, Go, SQL etc.

## What does the language itself look like?
The syntax is very very minimal and easy to learn. Take this example:
```odl
class person
    string name
    int age
    bool is_unemployed
    float monthly_salary
end
```

A class simply consists of a name and its fields. 

Please note that EVERYTHING must be `snake_case`. This is not just a convention, the transpiler expects `snake_case` so it can construct proper, idiomatic names for the target language.

## Supported targets
- Java
- Python
- Rust

More on the way.

## Examples
Take this small example:
```odl
class person
    string name
    int age
end
```

With Java as your target, the generated code would look like this:
```java
public class Person {
	private String name;
	private int age;

	public String getName() {
		return this.name;
	}

	public int getAge() {
		return this.age;
	}

	public void setName(String name) {
		this.name = name;
	}

	public void setAge(int age) {
		this.age = age;
	}

	public Person(
		String name,
		int age
	) {
		this.name = name;
		this.age = age;
	}
}
```

Getters, Setters, and an All-Args-Constructor are generated for you!

## Cloning & Compilation
```bash
git clone https://github.com/Moritisimor/odl
cd odl
go build -ldflags="-s -w" -o odl cmd/odl/main.go
```

## Usage
```bash
odl <input_file> <options...>
```

For example, if you had a file named `employee.odl` and you want to generate Python code, saving the file in `employee.py`:
```bash
odl employee.odl -t python -o employee
```

ODL will append the appropriate suffix for you, although you can write out the suffix if you want to.

If you had a file named `person.odl` and you want to generate Java code:
```bash
odl person.odl -t java
```

If you're transpiling for Java, then for each class in your file, the transpiler generates a seperate file where the class is stored.

This is also why you MUST omit the output flag when transpiling for Java.
