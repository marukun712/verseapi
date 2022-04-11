import fetch from 'node-fetch';
import dotenv from 'dotenv'
dotenv.config();
import fs from 'fs'

var member = {
    'uno': {
        id: "UCLfAsY3iMUAF2vvDxvIjymQ"
    },
    'itsuki': {
        id: "UCi3aT9cC6YyC0BnOpm2XBEw"
    },
    'otoha': {
        id: "UCYcnLc0n1ryBDZeGWQTVJ_g"
    },
    'sera': {
        id: "UCvXsXmpMKthJuX8XHbsnOjQ"
    },
    'maru': {
        id: "UCOg01LJmZF9UnwFbly73CVw"
    }
}

var items = []

function req() {
    items.length = 0;
    fetch("https://www.googleapis.com/youtube/v3/search?part=snippet&channelId=" + member.itsuki.id + "&key=" + process.env.key + "&eventType=upcoming&type=video")
        .then(res => res.json())
        .then(json => {
            var num = json.items.length;
            for (let i = 0; i < num; i++) {
                var ID = json.items[i].id.videoId
                var title = json.items[i].snippet.title
                var image = json.items[i].snippet.thumbnails.medium.url;
                items.push({ title: title, image: image, id: ID })
            }
        });

    fetch("https://www.googleapis.com/youtube/v3/search?part=snippet&channelId=" + member.uno.id + "&key=" + process.env.key + "&eventType=upcoming&type=video")
        .then(res => res.json())
        .then(json => {
            var num = json.items.length;
            for (let i = 0; i < num; i++) {
                var ID = json.items[i].id.videoId
                var title = json.items[i].snippet.title
                var image = json.items[i].snippet.thumbnails.medium.url;
                items.push({ title: title, image: image, id: ID })
            }
        });

    fetch("https://www.googleapis.com/youtube/v3/search?part=snippet&channelId=" + member.maru.id + "&key=" + process.env.key + "&eventType=upcoming&type=video")
        .then(res => res.json())
        .then(json => {
            var num = json.items.length;
            for (let i = 0; i < num; i++) {
                var ID = json.items[i].id.videoId
                var title = json.items[i].snippet.title
                var image = json.items[i].snippet.thumbnails.medium.url;
                items.push({ title: title, image: image, id: ID })
            }
        });

    fetch("https://www.googleapis.com/youtube/v3/search?part=snippet&channelId=" + member.sera.id + "&key=" + process.env.key + "&eventType=upcoming&type=video")
        .then(res => res.json())
        .then(json => {
            var num = json.items.length;
            for (let i = 0; i < num; i++) {
                var ID = json.items[i].id.videoId
                var title = json.items[i].snippet.title
                var image = json.items[i].snippet.thumbnails.medium.url;
                items.push({ title: title, image: image, id: ID })
            }
        });

    fetch("https://www.googleapis.com/youtube/v3/search?part=snippet&channelId=" + member.otoha.id + "&key=" + process.env.key + "&eventType=upcoming&type=video")
        .then(res => res.json())
        .then(json => {
            var num = json.items.length;
            for (let i = 0; i < num; i++) {
                var ID = json.items[i].id.videoId
                var title = json.items[i].snippet.title
                var image = json.items[i].snippet.thumbnails.medium.url;
                items.push({ title: title, image: image, id: ID })
            }
        });

    function write() {
        fs.writeFile('res.json', JSON.stringify(items, null, '    '), (err) => {
            if (err) console.log(`error!::${err}`);
        });

    }
    setTimeout(write, 1000);
}

setInterval(() => {
    req();
    console.log('情報を更新しました。 ')
}, 5.76e+6);