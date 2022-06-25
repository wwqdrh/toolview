<script setup lang="ts">
const props = defineProps<{
  modal: string
  okCb?: Function
  closeCb?: Function
}>()
</script>

<template>
  <!-- Modal -->
  <div class="container flex justify-center ">
    <div
      :id="props.modal"
      class="modal fade fixed hidden w-1/2 outline-none overflow-x-hidden overflow-y-auto border-stone-600 border-1"
      tabindex="-1" aria-labelledby="exampleModalLabel" aria-hidden="true"
    >
      <div class="modal-dialog relative w-auto pointer-events-none">
        <div
          class="modal-content border-none shadow-lg relative flex flex-col w-full pointer-events-auto bg-clip-padding rounded-md outline-none text-current"
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
            <slot name="body" />
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
              @click.stop="props.closeCb?.()"
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
ml-1" data-bs-dismiss="modal"
              @click.stop="props.okCb?.()"
            >
              确认
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style>
</style>
