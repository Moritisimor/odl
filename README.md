# Object Definition Language
A Domain-Specific-Language for generating Classes/Structs in other languages

## What is this project about?
ODL is a tiny, declarative Language for defining an object and transpiling it to many different targets, such as Java, C#, Python, Rust, Go, SQL etc.

## What does it look like?
As of now, the language is not very well defined, but current designs look like this:
```odl
class person
    string name
    int age
    bool is_unemployed
    float monthly_salary
end
```
