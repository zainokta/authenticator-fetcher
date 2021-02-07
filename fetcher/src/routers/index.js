const express = require('express')
const { auth, checkRole } = require('../middlewares/')
const jwtRouter = require('./jwt')

const router = express.Router()

router.use('/admin', auth, checkRole, async(req, res) => {
    res.json({
        message: "hello from admin api"
    })
})

router.use('/jwt', auth, jwtRouter)

module.exports = router