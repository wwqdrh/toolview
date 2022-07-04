import { acceptHMRUpdate, defineStore } from 'pinia'

import { GetRedisConf, GetRedisConfStatus, GetRedisKey, UpdateRedisConf } from '~/api/redis'
import type { APIResponse, RedisConfig } from '~/api/types'

export const useRedisStore = defineStore('redis', () => {
  const version = ref(1)
  const _ = ref<any>() // 只是为了避免lint报错

  async function getStatus(): Promise<APIResponse> {
    _.value = version.value
    return await (await GetRedisConfStatus()).data
  }

  async function getConf(): Promise<RedisConfig> {
    const res = await GetRedisConf()
    return (res.data as APIResponse).data as RedisConfig
  }

  async function updateConf(conf: RedisConfig) {
    return await UpdateRedisConf(conf)
  }

  async function GetKey(data: { key: string; type: number }): Promise<APIResponse> {
    return await (await GetRedisKey(data)).data
  }

  return {
    getStatus,
    getConf,
    updateConf,
    GetKey,
  }
})

if (import.meta.hot)
  import.meta.hot.accept(acceptHMRUpdate(useRedisStore, import.meta.hot))
