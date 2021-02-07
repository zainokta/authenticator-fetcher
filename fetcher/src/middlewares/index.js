const authenticate = require('./auth')
const checkRole = require('./role')

module.exports = {
    auth: authenticate,
    checkRole: checkRole
}