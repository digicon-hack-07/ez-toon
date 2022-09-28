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
import { useRouter } from 'vue-router'
import type { Page } from '../../lib/page';
import type { Project } from '../../lib/project';
import Thumbnail from './Thumbnail.vue'
const router = useRouter()

const createAndOpenNewProject = async () => {
  const now = new Date(Date.now()).toISOString()
  const projectData : Project = {
    id: "00000000000000000000000000",
    name: "untitled",
    thumbnail: "/vite.svg",
    pages: 1,
    created_at: now,
    updated_at: now
  }

  const pageData : Page = {
    project_id: "",
    id: "00000000000000000000000000",
    index: 1,
    height: 600,
    width: 400
  }

  const res = await axios.post('/api/projects', projectData)
  if(res.status === 201){
    pageData.project_id = res.data.id
    const pageRes = await axios.post('/api/pages', pageData)
    if(pageRes.status === 201){
      router.push("/page/" + pageRes.data.id)
    }
  }
}
</script>

<style module></style>
