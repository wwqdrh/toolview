import { acceptHMRUpdate, defineStore } from 'pinia'

import { GetEtcdConf, GetEtcdConfStatus, GetEtcdList, UpdateEtcdConf, UpdateEtcdKey } from '~/api/etcd'
import type { APIResponse, EtcdConfig, EtcdList } from '~/api/types'

export const useEtcdStore = defineStore('etcd', () => {
  const version = ref(1)
  const _ = ref<any>() // 只是为了避免lint报错

  async function getStatus(): Promise<APIResponse> {
    _.value = version.value
    return await (await GetEtcdConfStatus()).data
  }

  async function getKeyList(): Promise<EtcdList[]> {
    _.value = version.value // 引用version字段 能够触发更新
    const res = await GetEtcdList()
    return res.data
  }

  async function updateKey(key: string, value: string) {
    await UpdateEtcdKey(key, value)
    version.value++
  }

  async function getConf(): Promise<EtcdConfig> {
    const res = await GetEtcdConf()
    return (res.data as APIResponse).data as EtcdConfig
  }

  async function updateConf(conf: EtcdConfig) {
    return await UpdateEtcdConf(conf)
  }

  return {
    getKeyList,
    getStatus,
    updateKey,
    getConf,
    updateConf,
  }
})

if (import.meta.hot)
  import.meta.hot.accept(acceptHMRUpdate(useEtcdStore, import.meta.hot))
