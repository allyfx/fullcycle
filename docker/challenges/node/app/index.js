const express = require('express')
const app = express()
const port = 3000
const config = {
  host: 'db',
  user: 'root',
  password: 'root',
  database: 'nodedb'
}
const mysql = require('mysql')
const connection = mysql.createConnection(config)

// Drop table for testing purposes
const cleanDatabase = `DROP TABLE people`
connection.query(cleanDatabase)

const createTableSql = `
  CREATE TABLE people(
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255)
  )
`
const createPersonSql = `INSERT INTO people (name) VALUES ('Ally')`
const getPeopleSql = `SELECT * from people`

connection.query(createTableSql)
connection.query(createPersonSql)
let people
connection.query(getPeopleSql, (err, result, fields) => {
  people = result
})

connection.end()

app.get('/', (req, res) => {
  res.send(`
    <h1>Full Cycle Rocks!</h1>
    <ol>
      ${people.map(person => `<li>${person.name}</li>`)}
    </ol>
  `)
})

app.listen(port, () => {
  console.log('Rodando na porta ' + port)
})