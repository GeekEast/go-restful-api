import express, { json } from 'express';
import _ from 'lodash';
const app = express();

const books = [{
  ID: "1",
  Isbn: "44873",
  Title: "Book One",
  Author: {
    FirstName: "James",
    LastName: "Tan"
  }
}]

app.use(json())

app.get("/api/books", (req, res) => {
  res.json(books);
})

app.post("/api/books", (req, res) => {
  console.log('create new book in express')
  const book = req.body;
  const id = String(_.random(0, 10000000))
  const newBook = { id, ...book };
  books.push(newBook);
  res.json(newBook)
})


app.listen(8081, () => console.log("app listening on port 8081!"))