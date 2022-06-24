<script setup lang="ts">
import { useEtcdStore } from '~/store/etcd'

const etcdStore = useEtcdStore()

const etcdList = asyncComputed(async () => {
  return await etcdStore.getKeyList()
})

const newKey = $ref('')
const newValue = $ref('')
</script>

<template>
  <button
    type="button" class="px-6
py-2.5
bg-blue-600
text-white
font-medium
text-xs
leading-tight
uppercase
rounded
shadow-md
hover:bg-blue-700 hover:shadow-lg
focus:bg-blue-700 focus:shadow-lg focus:outline-none focus:ring-0
active:bg-blue-800 active:shadow-lg
transition
duration-150
ease-in-out" data-bs-toggle="modal" data-bs-target="#exampleModal"
  >
    添加新元素
  </button>

  <!-- Modal -->
  <div
    id="exampleModal"
    class="modal fade fixed top-0 left-0 hidden w-full h-full outline-none overflow-x-hidden overflow-y-auto" tabindex="-1" aria-labelledby="exampleModalLabel" aria-hidden="true"
  >
    <div class="modal-dialog relative w-auto pointer-events-none">
      <div
        class="modal-content border-none shadow-lg relative flex flex-col w-full pointer-events-auto bg-white bg-clip-padding rounded-md outline-none text-current"
      >
        <div
          class="modal-header flex flex-shrink-0 items-center justify-between p-4 border-b border-gray-200 rounded-t-md"
        >
          <h5 id="exampleModalLabel" class="text-xl font-medium leading-normal text-gray-800">
            添加元素
          </h5>
          <button
            type="button"
            class="btn-close box-content w-4 h-4 p-1 text-black border-none rounded-none opacity-50 focus:shadow-none focus:outline-none focus:opacity-100 hover:text-black hover:opacity-75 hover:no-underline"
            data-bs-dismiss="modal" aria-label="Close"
          />
        </div>
        <div class="modal-body relative p-4">
          <div class="block p-6 rounded-lg shadow-lg bg-white max-w-md mx-auto">
            <div class="form-group mb-6">
              <input
                id="exampleInput125"
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
              focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none"
                placeholder="键"
              >
            </div>
            <div class="form-group mb-6">
              <input
                id="exampleInput126" v-model="newValue" type="text" class="form-control block
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
              focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none"
                placeholder="值"
              >
            </div>
          </div>
        </div>
        <div
          class="modal-footer flex flex-shrink-0 flex-wrap items-center justify-end p-4 border-t border-gray-200 rounded-b-md"
        >
          <button
            type="button" class="px-6
    py-2.5
    bg-purple-600
    text-white
    font-medium
    text-xs
    leading-tight
    uppercase
    rounded
    shadow-md
    hover:bg-purple-700 hover:shadow-lg
    focus:bg-purple-700 focus:shadow-lg focus:outline-none focus:ring-0
    active:bg-purple-800 active:shadow-lg
    transition
    duration-150
    ease-in-out" data-bs-dismiss="modal"
          >
            Close
          </button>
          <button
            type="button" class="px-6
py-2.5
bg-blue-600
text-white
font-medium
text-xs
leading-tight
uppercase
rounded
shadow-md
hover:bg-blue-700 hover:shadow-lg
focus:bg-blue-700 focus:shadow-lg focus:outline-none focus:ring-0
active:bg-blue-800 active:shadow-lg
transition
duration-150
ease-in-out
ml-1"
            data-bs-dismiss="modal"
            @click.stop="etcdStore.updateKey(newKey, newValue)"
          >
            确认
          </button>
        </div>
      </div>
    </div>
  </div>
  <div class="flex flex-col">
    <div class="overflow-x-auto sm:-mx-6 lg:-mx-8">
      <div class="py-2 inline-block min-w-full sm:px-6 lg:px-8">
        <div class="overflow-hidden text-left">
          <table class="min-w-full">
            <thead class="border-b">
              <tr>
                <th scope="col" class="text-sm font-medium text-gray-900 px-6 py-4">
                  #
                </th>
                <th scope="col" class="text-sm font-medium text-gray-900 px-6 py-4">
                  Key
                </th>
                <th scope="col" class="text-sm font-medium text-gray-900 px-6 py-4">
                  Value
                </th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="item, i in etcdList" :key="`item-${i}`" class="border-b">
                <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                  {{ i }}
                </td>
                <td class="text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap">
                  {{ item.name }}
                </td>
                <td class="text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap">
                  {{ item.value }}
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
</template>

<style></style>

<route lang="yaml">
meta:
  layout: home
</route>
