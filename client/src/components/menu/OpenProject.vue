<template>
  <thumbnail
    :name="projectName"
    :image="thumnailImage"
    :is-show-date="true"
    :created-at="createdAt"
    :undated-at="updatedAt"
    @click="openProject"
  />
</template>

<script lang="ts" setup>
import { useRouter } from 'vue-router'
import Thumbnail from './Thumbnail.vue'
const props = withDefaults(
  defineProps<{
    projectName: string
    url: string
    thumnailImage: string
    pages: number
    createdAt?: Date
    updatedAt?: Date
  }>(),
  {
    projectName: 'noname',
    url: '/',
    thumnailImage: '/vite.svg',
    pages: 0
  }
)
const router = useRouter()

const openProject = function () {
  router.push({ path: props.url})
}
</script>

<style module></style>
