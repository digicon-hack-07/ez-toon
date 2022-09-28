<!-- eslint-disable vue/multi-word-component-names -->
<template>
  <thumbnail
    name="create new page"
    :is-show-date="false"
    :image="'/plus.svg'"
    @click="createAndOpenNewPage"
  />
</template>

<script lang="ts" setup>
import axios from 'axios'
import { useRouter } from 'vue-router'
import type { Page } from '../../lib/page'
import Thumbnail from './Thumbnail.vue'
const router = useRouter()
const props = defineProps<{
    projectId: string
    index: number
  }>()
const createAndOpenNewPage = async () => {
  const now = new Date(Date.now()).toISOString()
  const pageData: Page = {
    project_id: props.projectId,
    id: '00000000000000000000000000',
    index: props.index,
    height: 1000,
    width: 707
  }

  const pageRes = await axios.post('/api/pages', pageData)
  if (pageRes.status === 201) {
    router.push('/page/' + pageRes.data.id)
  }
}
</script>

<style module></style>
