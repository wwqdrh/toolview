<script setup lang="ts">
import type { EtcdConfig } from '~/api/types'
import { useEtcdStore } from '~/store/etcd'
const etcdStore = useEtcdStore()

const curretcdConf = $ref<EtcdConfig>({
  endpoints: '',
  username: '',
  password: '',
})

onBeforeMount(async () => {
  const conf = await etcdStore.getConf()
  curretcdConf.endpoints = conf.endpoints
  curretcdConf.username = conf.username
  curretcdConf.password = conf.password
})

const etcdConfUpdate = async () => {
  etcdStore.updateConf(curretcdConf)
}
</script>

<template>
  <HButton bs-toggle="modal" bs-target="#configModal" value="查看配置" />
  <Modal modal="configModal" :ok-cb="etcdConfUpdate">
    <template #body>
      <div class="block p-6 rounded-lg shadow-lg bg-white max-w-md mx-auto space-y-2">
        endpoints: <input v-model="curretcdConf.endpoints">
      </div>
      <div class="block p-6 rounded-lg shadow-lg bg-white max-w-md mx-auto space-y-2">
        username: <input v-model="curretcdConf.username">
      </div>
      <div class="block p-6 rounded-lg shadow-lg bg-white max-w-md mx-auto space-y-2">
        password: <input v-model="curretcdConf.password">
      </div>
    </template>
  </Modal>
</template>

<style>
</style>

<route lang="yaml">
meta:
  layout: home
</route>
