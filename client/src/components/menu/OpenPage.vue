<template>
  <thumbnail
    :name="'' + pageNumber"
    :image="thumnailImage"
    :is-show-date="false"
    @click="openPage"
  />
</template>

<script lang="ts" setup>
import { useRouter } from 'vue-router'
import Thumbnail from './Thumbnail.vue'
const props = withDefaults(
  defineProps<{
    pageNumber: number
    url: string
    thumnailImage: string
  }>(),
  {
    pageNumber: 1,
    url: '/',
    thumnailImage: '/vite.svg'
  }
)
const router = useRouter()

const openPage = function () {
  router.push(props.url)
}
</script>

<style module></style>
