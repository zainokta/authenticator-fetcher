const redis = require('redis')
const bluebird = require('bluebird')

bluebird.promisifyAll(redis)
const redisClient = redis.createClient(process.env.REDIS_HOST)

module.exports = redisClient