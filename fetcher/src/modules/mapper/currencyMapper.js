const redisClient = require('../../infrastructure/redis')
const axios = require('axios')

const CURRENCY_API_URL = 'https://free.currconv.com/api/v7/convert?q=USD_IDR&compact=ultra&apiKey=547807587094635f1a50'

async function getCurrency() {
    let result = await redisClient.getAsync('currency')
    if (result) {
        return result
    }

    try {
        const response = await axios.get(CURRENCY_API_URL)
        redisClient.setex('currency', 3600, JSON.stringify(response.data.USD_IDR))
        result = response.data.USD_IDR
    } catch (err) {
        throw err
    }

    return result
}

async function mapCurrency(responseData) {
    const currency = await getCurrency()
    let mappedData = []
    for (const item of responseData) {
        if (item.price !== null) {
            let obj = {
                ...item,
                price_usd: item.price / currency
            }
            mappedData.push(obj)
        }
    }

    return mappedData
}

module.exports = mapCurrency