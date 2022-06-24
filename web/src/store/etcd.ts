import { acceptHMRUpdate, defineStore } from 'pinia'

import { GetEtcdList, UpdateEtcdKey } from '~/api/etcd'
import type { EtcdList } from '~/api/types'

export const useEtcdStore = defineStore('etcd', () => {
  const version = ref(1)

  async function getKeyList(): Promise<EtcdList[]> {
    const _ = version.value // 触发更新
    const res = await GetEtcdList()
    return res.data
  }

  async function updateKey(key: string, value: string) {
    await UpdateEtcdKey(key, value)
    version.value++
  }

  return {
    getKeyList,
    updateKey,
  }
})

if (import.meta.hot)
  import.meta.hot.accept(acceptHMRUpdate(useEtcdStore, import.meta.hot))
