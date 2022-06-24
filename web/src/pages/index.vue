<script setup lang="ts">
const router = useRouter()
</script>

<template>
  <div class="container my-12 px-6 mx-auto">
    <!-- Section: Design Block -->
    <section class="mb-32 text-gray-800">
      <h2 class="text-3xl font-bold mb-12 text-center">
        可视化工具
      </h2>

      <div class="grid lg:grid-cols-3 gap-6">
        <div
          class="relative overflow-hidden bg-no-repeat bg-cover shadow-lg rounded-lg"
          style="background-position: 50%;" data-mdb-ripple="true" data-mdb-ripple-color="light"
        >
          <img src="https://mdbootstrap.com/img/new/standard/nature/051.jpg" class="w-full">
          <div class="cursor-pointer" @click.stop="router.push('/etcd')">
            <div
              class="absolute top-0 right-0 bottom-0 left-0 w-full h-full overflow-hidden bg-fixed"
              style="background-color: rgba(0, 0, 0, 0.4)"
            >
              <div class="flex justify-start items-end h-full">
                <div class="text-white m-6">
                  <h5 class="font-bold text-lg mb-3">
                    Etcd
                  </h5>
                  <p>
                    <small>Updated <u>23.06.2022</u> by Ahui</small>
                  </p>
                </div>
              </div>
            </div>
          </div>
        </div>

        <div
          class="relative overflow-hidden bg-no-repeat bg-cover shadow-lg rounded-lg" data-mdb-ripple="true"
          data-mdb-ripple-color="light"
        >
          <img src="https://mdbootstrap.com/img/new/standard/nature/044.jpg" class="w-full">
          <div class="cursor-pointer" @click.stop="router.push('/redis')">
            <div
              class="absolute top-0 right-0 bottom-0 left-0 w-full h-full overflow-hidden bg-fixed"
              style="background-color: rgba(0, 0, 0, 0.4)"
            >
              <div class="flex justify-start items-end h-full">
                <div class="text-white m-6">
                  <h5 class="font-bold text-lg mb-3">
                    Redis
                  </h5>
                  <p>
                    <small>Updated <u>23.06.2022</u> by Ahui</small>
                  </p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>
    <!-- Section: Design Block -->
  </div>
</template>

<route lang="yaml">
meta:
  layout: home
</route>
