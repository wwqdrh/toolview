<script setup lang="ts">
import type { RedisConfig } from '~/api/types'

import { useRedisStore } from '~/store/redis'

const redisStore = useRedisStore()

const curretcdConf = $ref<RedisConfig>({
  endpoint: '',
  password: '',
})

const searchKey = $ref('')
const searchResult = ref('')

const search = async () => {
  const res = await redisStore.GetKey({ key: searchKey, type: 0 })
  searchResult.value = (res.data as any).data
}

const status = asyncComputed(async () => {
  return await redisStore.getStatus()
})

onBeforeMount(async () => {
  const conf = await redisStore.getConf()
  curretcdConf.endpoint = conf.endpoint
  curretcdConf.password = conf.password
})
</script>

<template>
  <div>
    {{ status }}
    <div>
      <div class="block p-6 rounded-lg shadow-lg bg-white max-w-md mx-auto space-y-2">
        endpoints: <input v-model="curretcdConf.endpoint">
      </div>
      <div class="block p-6 rounded-lg shadow-lg bg-white max-w-md mx-auto space-y-2">
        password: <input v-model="curretcdConf.password">
      </div>
    </div>
  </div>
  <div>
    <input v-model="searchKey" placeholder="输入搜索key" @keydown.enter="search">
    <input v-model="searchResult">
  </div>
</template>

<style>
</style>

<route lang="yaml">
meta:
  layout: home
</route>
