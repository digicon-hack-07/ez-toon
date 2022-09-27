<!-- eslint-disable vue/multi-word-component-names -->
<template>
  <thumbnail
    name="create new project"
    :is-show-date="false"
    :image="'/plus.svg'"
    @click="createAndOpenNewProject"
  />
</template>

<script lang="ts" setup>
import axios from 'axios';
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import type { Project } from '../../lib/project';
import Thumbnail from './Thumbnail.vue'
const router = useRouter()

const createAndOpenNewProject = async () => {
  const now = new Date(Date.now()).toISOString()
  const data : Project = {
    id: "00000000000000000000000000",
    name: "untitled",
    thumbnail: "/vite.svg",
    pages: 1,
    created_at: now,
    updated_at: now
  }

  const res = await axios.post('/api/projects', data)
  if(res.status === 201){
    router.push("/project/" + res.data.id)
  }
}
</script>

<style module></style>
