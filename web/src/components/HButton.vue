<script setup lang="ts">
const props = defineProps<{
  bsToggle?: string
  bsTarget?: string
  value: string
}>()
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
ease-in-out" :data-bs-toggle="props.bsToggle" :data-bs-target="props.bsTarget"
  >
    {{ props.value }}
  </button>
</template>

<style></style>
