const redis = require('redis')
const bluebird = require('bluebird')

bluebird.promisifyAll(redis)
const redisClient = redis.createClient()

module.exports = redisClient