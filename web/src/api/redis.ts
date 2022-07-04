import type { RedisConfig } from './types'
import request from './index'

export function GetRedisConf() {
  return request({
    url: '/api/redis/conf/status',
    method: 'get',
  })
}

export function UpdateRedisConf(conf: RedisConfig) {
  return request({
    url: '/api/redis/conf/update',
    method: 'post',
    headers: {
      ContentType: 'application/json',
    },
    data: {
      endpoint: conf.endpoint,
      password: conf.password,
    },
  })
}

export function GetRedisConfStatus() {
  return request({
    url: '/api/redis/conf/verify',
    method: 'get',
  })
}

export function GetRedisKey(data: { key: string; type: number }) {
  return request({
    url: '/api/redis/key/get',
    method: 'get',
    headers: {
      ContentType: 'application/json',
    },
    data: {
      key: data.key,
      type: data.type,
    },
  })
}
