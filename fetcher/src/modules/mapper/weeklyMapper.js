const redisClient = require('../../infrastructure/redis')
const axios = require('axios')
const _ = require('lodash')
const moment = require('moment')

async function filterNull(responseData) {
    return responseData.filter((item) => item.area_provinsi !== null)
}

function getMedianValue(array) {
    array = array.sort();
    if (array.length % 2 === 0) {
        return array[array.length / 2];
    } else {
        return array[(array.length - 1) / 2];
    }
}

exports.weeklyMapper = async (req, res) => {
    let result = await redisClient.getAsync('stein')
    let filteredResults
    if (result) {
        filteredResults = await filterNull(JSON.parse(result))
    } else {
        const response = await axios.get(process.env.BASE_STEIN_URL)
        redisClient.setex('stein', 3600, JSON.stringify(response.data))
        filteredResults = await filterNull(response.data)
    }

    const data = _.chain(filteredResults)
        .groupBy('area_provinsi')
        .mapValues((item) => {
            return _.groupBy(item, (element) => moment(element.tgl_parsed).week())
        })
        .mapValues((weeklyItem) => {
            return _.mapValues(weeklyItem, (item) => {
                return {
                    max_price: parseInt(_.maxBy(item, function (element) {
                        return parseInt(element.price)
                    }).price),
                    min_price: parseInt(_.minBy(item, function (element) {
                        return parseInt(element.price)
                    }).price),
                    avg_price: parseInt(_.meanBy(item, function (element) {
                        return parseInt(element.price)
                    })),
                    max_size: parseInt(_.maxBy(item, function (element) {
                        return parseInt(element.size)
                    }).size),
                    min_size: parseInt(_.minBy(item, function (element) {
                        return parseInt(element.size)
                    }).size),
                    avg_size: parseInt(_.meanBy(item, function (element) {
                        return parseInt(element.size)
                    })),
                    median: getMedianValue(item),
                }
            })
        })

    return res.json(data)
}