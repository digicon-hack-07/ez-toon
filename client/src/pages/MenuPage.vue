<template>
  <div :class="$style.menu">
    <create-new-project />
    <open-project
      v-for="project in projectList"
      :key="project.id"
      :project-name="project.name"
      :url="'/project/' + project.id"
      :pages="project.pages"
      :thumnail-image="project.thumbnail"
      :created-at="new Date(project.created_at)"
      :updated-at="new Date(project.updated_at)"
    />
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import axios from 'axios'
import CreateNewProject from '../components/menu/CreateNewProject.vue'
import OpenProject from '../components/menu/OpenProject.vue'
import type { Project } from '../lib/project'

const projectList = ref<Project[]>()

onMounted(async () => {
  const res = await axios.get('/api/projects')
  projectList.value = res.data
})
</script>

<style module>
.menu {
  text-align: left;
}
</style>
