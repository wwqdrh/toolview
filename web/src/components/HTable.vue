<script setup lang="ts">
interface EtcdList {
  name: string
  value: string
}

const props = defineProps<{
  tables: EtcdList[]
}>()
</script>

<template>
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
              <tr v-for="item, i in props.tables" :key="`item-${i}`" class="border-b">
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
