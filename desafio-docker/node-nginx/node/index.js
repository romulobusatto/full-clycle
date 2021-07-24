const express = require('express')
var random_name = require('node-random-name')
const app = express()
const port = 3000
const config = {
    host: 'db',
    user: 'root',
    password: 'root',
    database: 'nodedb'
};
const mysql = require('mysql')

const connection = mysql.createConnection(config)

const sql = `create table IF NOT EXISTS people (id int not null auto_increment, name varchar(255) not null, primary key (id))`
connection.query(sql)
connection.end()

app.get('/', (req, res) => {
    var html = '<h1>Full Cycle Rocks!</h1>'

    const connection = mysql.createConnection(config)
    const sql = `INSERT INTO people(name) values('` + random_name() + `')`
    connection.query(sql)

    html += '<ul>'
    connection.query(
        'SELECT * FROM people;',
        null,
        function (err, peoples) {
            if (err) throw err;
            peoples.forEach(function (people) {
                html += '<li>' + people.name + '</li>'

            });
            html += '</ul>'
            res.send(html)
        }
    );
    connection.end()
})

app.listen(port, () => {
    console.log('Rodando na porta ' + port)
})