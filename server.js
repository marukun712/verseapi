import express from "express";
var app = express();
var port = 2100
import dotenv from 'dotenv'
dotenv.config();
import fs from 'fs'
const allowCrossDomain = function(req, res, next) {
    res.header('Access-Control-Allow-Origin', '*')
    res.header('Access-Control-Allow-Methods', 'GET,PUT,POST,DELETE')
    res.header(
        'Access-Control-Allow-Headers',
        'Content-Type, Authorization, access_token'
    )

    // intercept OPTIONS method
    if ('OPTIONS' === req.method) {
        res.send(200)
    } else {
        next()
    }
}
app.use(allowCrossDomain)

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