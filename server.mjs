import express from "express";
var app = express();
var port = 2100
import dotenv from 'dotenv'
dotenv.config();
import fs from 'fs'

app.get('/api', function(req, res) {
    const json = JSON.parse(fs.readFileSync('./res.json', 'utf8'));
    res.json(json)
})

app.get('/', function(req, res) {
    res.send('welcome to verseapi')
})


app.listen(port, function() {
    console.log("http://localhost:" + port + "で起動")
})