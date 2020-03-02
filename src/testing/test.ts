import acn from 'autocannon'

const book = {
  Isbn: "9984059",
  Title: "Book",
  Author: {
    FirstName: "James",
    LastName: "Tan"
  }

}
const instance = acn({
  url: 'http://localhost:8081/api/books/',
  connections: 10,
  duration: 10,
  pipelining: 2,
  method: "POST",
  body: JSON.stringify(book)
}, console.log)


// just render results
acn.track(instance, { renderProgressBar: true })