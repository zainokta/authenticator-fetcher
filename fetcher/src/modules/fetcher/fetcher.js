const axios = require('axios')
const redisClient = require('../../infrastructure/redis')
const mapCurrency = require('../mapper/currencyMapper')

const BASE_URL = 'https://stein.efishery.com/v1/storages/5e1edf521073e315924ceab4/list'

exports.fetcher = async(req, res) => {
    return redisClient.get('stein', async(err, result) => {
        if (result) {
            const mappedData = await mapCurrency(JSON.parse(result))
            return res.json({
                data: mappedData
            })
        }

        if(err !== null){
            return res.status(500).json(err.message)
        }

        try {
            const response = await axios.get(BASE_URL)
            redisClient.setex('stein', 3600, JSON.stringify(response.data))
            const mappedData = await mapCurrency(response.data)
            res.status(200).json(mappedData)
        }catch(err){
            res.status(500).json(err.message)
        }
    })
}
