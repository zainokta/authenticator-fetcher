
const express = require('express');
const cors = require('cors')
const router = require('./routers')
const app = express();

app.use(cors())
app.use(express.json({
    limit: '50mb'
}));
app.use(express.urlencoded({
    extended: false,
    limit: '50mb'
}));

app.use('/api/v1', router)

app.get('*', (req, res, next) => {
    res.status(200).json({
        message: "Welcome to the beginning of nothingness.",
        status: 200
    })
})

module.exports = app