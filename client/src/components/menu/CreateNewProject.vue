<!-- eslint-disable vue/multi-word-component-names -->
<template>
  <thumbnail
    name="create new project"
    :created-at="createdAt"
    :undated-at="updatedAt"
    :image="image"
    @click="createAndOpenNewProject"
  />
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import Thumbnail from './Thumbnail.vue'
withDefaults(
  defineProps<{
    name: string
    image: string
    createdAt?: Date
    updatedAt?: Date
  }>(),
  {
    name: 'create new project',
    image: '/plus.svg'
  }
)
const router = useRouter()

const createAndOpenNewProject = function () {
  const url = ref<string>('/')
  /*以下にサーバーに通信して新しいプロジェクトを作らせ、そこにアクセスするためのデータを入手するプログラムを書く */
  router.push(url.value)
}
</script>

<style module></style>
