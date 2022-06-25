<script setup lang="ts">
import { useEtcdStore } from '~/store/etcd'

const etcdStore = useEtcdStore()
const router = useRouter()

const etcdStatus = asyncComputed(async () => {
  const status = await etcdStore.getStatus()
  if (status.code !== 0)
    router.push('/etcd/config')
})

const etcdList = asyncComputed(async () => {
  return await etcdStore.getKeyList()
})

const newKey = $ref('')
const newValue = $ref('')
const addRecord = async () => {
  return await etcdStore.updateKey(newKey, newValue)
}
</script>

<template>
  <HButton bs-toggle="modal" bs-target="#exampleModal" value="添加新元素" />
  <Modal modal="exampleModal" :ok-cb="addRecord">
    <template #body>
      <div class="block p-6 rounded-lg shadow-lg bg-white max-w-md mx-auto space-y-2">
        <input
          v-model="newKey" type="text" class="form-control block
              w-full
              px-3
              py-1.5
              text-base
              font-normal
              text-gray-700
              bg-white bg-clip-padding
              border border-solid border-gray-300
              rounded
              transition
              ease-in-out
              m-0
              focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none" placeholder="键"
        >
        <input
          v-model="newValue" type="text" class="form-control block
              w-full
              px-3
              py-1.5
              text-base
              font-normal
              text-gray-700
              bg-white bg-clip-padding
              border border-solid border-gray-300
              rounded
              transition
              ease-in-out
              m-0
              focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none" placeholder="值"
        >
      </div>
    </template>
  </Modal>

  <HTable :tables="etcdList" />
</template>

<style>
</style>

<route lang="yaml">
meta:
  layout: home
</route>
