require('dotenv').config()
const http = require('http')
const app = require('./src/app')

const port = process.env.PORT || '8001'
const server = http.createServer(app)

server.listen(port, async () => {
    console.log(`Server running at port ${port}`)
})