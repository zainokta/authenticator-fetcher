const express = require('express')
const router = express.Router()

router.get('/verify', async(req, res, next) => {
    return res.status(200).json({
        message: req.User
    })
})

module.exports = router