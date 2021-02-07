const express = require('express')
const { auth, checkRole } = require('../middlewares/')
const modules = require('../modules/')
const jwtRouter = require('./jwt')

const router = express.Router()

router.use('/admin/fetcher', auth, checkRole, (req, res) => {
    return res.json("not implemented yet")
})

router.use('/fetcher', auth, modules.fetcher)

router.use('/jwt', auth, jwtRouter)

module.exports = router