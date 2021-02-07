const jwt = require('jsonwebtoken')
const { JWT_SECRET } = process.env

const authenticate = async (req, res, next) => {
    const authHeader = req.get('Authorization')
    let token
    if (authHeader && authHeader.startsWith('Bearer')) {
        token = authHeader.split(' ')[1]
    }
    
    if (!token) {
        return res.status(400).json({
            message: "Bad Request"
        })
    }
    
    let claims
    try {
        claims = jwt.verify(token, JWT_SECRET)
    }catch(error){
        return res.status(400).json({
            message: "Invalid Token"
        })
    }
    
    req.User = claims

    next()
}

module.exports = authenticate