import type { EtcdConfig } from './types'
import request from './index'

/**
 *
 * @param {Object} data time,name
 * @returns
 */
export function GetEtcdList() {
  return request({
    url: '/api/etcd/key/list',
    method: 'get',
  })
}

export function UpdateEtcdKey(key: string, value: string) {
  return request({
    url: '/api/etcd/key/put',
    method: 'post',
    headers: {
      ContentType: 'application/json',
    },
    data: {
      key,
      value,
    },
  })
}

export function GetEtcdConf() {
  return request({
    url: '/api/etcd/conf/status',
    method: 'get',
  })
}

export function UpdateEtcdConf(conf: EtcdConfig) {
  return request({
    url: '/api/etcd/conf/update',
    method: 'post',
    headers: {
      ContentType: 'application/json',
    },
    data: {
      endpoints: conf.endpoints,
      username: conf.username,
      password: conf.password,
    },
  })
}

export function GetEtcdConfStatus() {
  return request({
    url: '/api/etcd/conf/verify',
    method: 'get',
  })
}
