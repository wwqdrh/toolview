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
