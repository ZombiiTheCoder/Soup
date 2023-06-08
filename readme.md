## Soup
###### Ver 0.1.1

![feature X](./imgs/Soup.png)

This is a programming language written in Go as being the third rewrite that is currently interpreted with a similar syntax to javascript.

### Building

Building this language's core is simple. You need to have [Golang](https://go.dev/dl/) installed and from the main "Soup" folder you need to run "scripts/build" and it will build all the binaries related to soup and add them to the bin folder.

### Code Example

test.soup
```soup

    use `@std`

    func hello_world(): str{

        val text : str = `Hello World`
        return text;

    }

    std.io.println(hello_world())

```