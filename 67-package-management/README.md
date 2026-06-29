
- go mod is the root path
- a project can be a binary project or a library project or both together
- every  package must have a directory, the name of the package must be(there can be alias as well) the name of the immediate directory

- what kind of project this is --> This is a binary and library application

- go does not have public private kind of access specifiers or modifier
- The whole encapsulation is done only by using a concept called restricted or unrestricted
- or exported or unexported

- any type, method, field or function in a package, if it stats with UpperCase it is exported/unrestricted --> public
- any type, method, field or function in a package, if it stats with lowerCase it is unexported/restricted --> private