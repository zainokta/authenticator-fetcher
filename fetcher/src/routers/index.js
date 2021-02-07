const express = require('express')
const { auth, checkRole } = require('../middlewares/')
const { fetcher, weeklyMapper } = require('../modules/')
const jwtRouter = require('./jwt')

const router = express.Router()

router.use('/admin/fetcher', auth, checkRole, weeklyMapper)

router.use('/fetcher', auth, fetcher)

router.use('/jwt', auth, jwtRouter)

module.exports = router