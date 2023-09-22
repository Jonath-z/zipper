# zipper

Zipper is a minimalistic and fast JavaScript bundler written in Go.

## Priciple

To understand how JavaScript is bundled, I read [Build Your Own JS Code Bundler](https://blog.bitsrc.io/building-your-first-bundler-99e4fdf502b2?gi=50efca31fac7) article.

In this bundler as in the mentioned article, here are main operations that are implemented to bundle the JavaScript code code:

- Walk through all the files in `/js` directory
- Read all the files with `.js` extension
- Calculate the dependencies graph of the project
- Build the depedency graph

Note: This bundler does not bundle recursively, however any contribution is welcome

## Dependency Graph structure

```javascript
[
  {
    path: "js/circle.js",
    content:
      "const PI = 3.141;function area(radius) {    return PI * radius * radius;}",
    dependencies: null,
  },
  {
    path: "js/index.js",
    content:
      "console.log('Area of square: ', squareArea(5));console.log('Area of circle', circleArea(5));",
    dependencies: [
      {
        line: 1,
        dependencyPath: "'./square.js';",
        expression: "import squareArea from './square.js';",
      },
      {
        line: 2,
        dependencyPath: "'./circle.js';",
        expression: "import circleArea from './circle.js';",
      },
    ],
  },
  {
    path: "js/square.js",
    content: "function area(side) {    return side * side;}",
    dependencies: null,
  },
];
```

## How to run locally

1. Clone the repository

```shell
git clone git@github.com:Jonath-z/zipper.git
```

2. Run the code with the root permission

```shell
sudo go run main.go
```

This command will bundle the JavaScript code in `/js` directory, the bundled output is located
in `/js/dist/index.js`

3. Run the bundled code

To run the bundle code you need to install Node.js in your local computer or you can it find [here](https://nodejs.org/en/download).

```shell
node js/dist/index.js
```

## Author

Jonathan Z.

- [X](https://twitter.com/JonathanZihind4)
- [LinkedIn](https://www.linkedin.com/in/jonathan-z-0a40ab209/)
