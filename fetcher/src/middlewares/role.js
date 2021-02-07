const checkRole = async (req, res, next) => {
    if (req.User.Role !== "admin") {
        return res.status(403).json({
            message: "Not Authenticated"
        })
    }

    next()
}

module.exports = checkRole