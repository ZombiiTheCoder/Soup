## Soup

##### Version 0.0.1 - [Golang](https://github.com/ZombiiTheCoder/SoupLang) Alpha
##### Version 0.0.6 - [TypeScript](https://github.com/ZombiiTheCoder/Soup) Alpha

<!-- <link rel="stylesheet" href="https://raw.githubusercontent.com/ZombiiTheCoder/SoupLang/master/docs/prism.css">
<script src="https://raw.githubusercontent.com/ZombiiTheCoder/SoupLang/master/docs/prism.js"></script> -->

![feature X](./Soup2.png)

This is a programming language formerly written in TypeScript that can be use for many things but the limitation is that it is currently interpreted and may become compiled in the future.

The Current Lexer Is On Version 2 hece the name [Lex2](./src/lex2/) no longer requiring [Lexer 1's](./src/lexer/) objects and methods

### Building

To build Soup you need to have the prerequisite [Golang](https://go.dev/dl/) installed

Next after go is installed go to this folder in your terminal and type `go build .` to build a binary for Soup.

### Installing

Soup will have an installer soon so you can use soup from anywhere.
<pre class="language-js">

    console.log("e")
    console.log("q")

</pre>

<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/prism/1.23.0/themes/prism.min.css" integrity="sha512-tN7Ec6zAFaVSG3TpNAKtk4DOHNpSwKHxxrsiw4GHKESGPs5njn/0sMCUMl2svV4wo4BK/rCP7juYz+zx+l6oeQ==" crossorigin="anonymous" />
<img src="https://img.shields.io/tokei/lines/github/ZombiiTheCoder/SoupLang?style=plastic"/>

<script>
window.Prism = window.Prism || {};
window.Prism.manual = true;
</script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/prism/1.23.0/prism.min.js"></script>

<script>
    let pre = document.querySelectorAll("pre");
    for (let i = 0; i < pre.length; i++){
        let code = pre[i].innerText;
        pre[i].innerHTML = Prism.highlight(code, Prism.languages.javascript, 'javascript');

  }
</script>
