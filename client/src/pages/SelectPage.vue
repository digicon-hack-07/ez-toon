<template>
  <div :class="$style.menu">
    <open-page
      v-for="n in id"
      :key="n"
      :page-number="n"
      :url="'/Page/' + id + '/' + n"
    />
  </div>
</template>

<script lang="ts" setup>
import { useRoute } from 'vue-router'
import OpenPage from '../components/menu/OpenPage.vue'
const route = useRoute()

const id = parseInt(route.params.id as string, 10)
</script>

<style module>
.menu {
  text-align: left;
  width: 90vw;
  min-height: 80vh;
}
</style>
