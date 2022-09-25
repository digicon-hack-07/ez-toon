<template>
  <thumbnail
    :name="projectName"
    :image="thumnailImage"
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
    createdAt?: Date
    updatedAt?: Date
  }>(),
  {
    projectName: 'noname',
    url: '/',
    thumnailImage: '/vite.svg'
  }
)
const router = useRouter()

const openProject = function () {
  router.push(props.url)
}
</script>

<style module></style>
