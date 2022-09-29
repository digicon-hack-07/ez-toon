<template>
  <thumbnail
    :name="'' + pageNumber"
    :image="thumnailImage"
    :is-show-date="false"
    :is-show-arrow="true"
    @click="openPage"
    @left="decrementIndex"
    @right="incrementIndex"
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
    thumnailImage: '/noImage.svg'
  }
)

interface Emits {
  (e: "left", value: number): void
  (e: "right", value: number): void
}

const emit = defineEmits<Emits>();
const router = useRouter()

const openPage =  () => {
  router.push(props.url)
}
const incrementIndex = () => {
  emit('right', props.pageNumber)
}
const decrementIndex = () => {
  emit('left', props.pageNumber)
}
</script>

<style module></style>
