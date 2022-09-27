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
import { onMounted, ref } from 'vue';
import axios from "axios";
import CreateNewProject from '../components/menu/CreateNewProject.vue'
import OpenProject from '../components/menu/OpenProject.vue'
import type { Project } from '../lib/project'

const projectList = ref<Project[]>()

onMounted(async () => {
  const res = await axios.get("/api/projects");
  projectList.value = res.data;
  //alert(JSON.stringify(projectList.value[0].created_at))
})

// const getProjectList = function () {
//   const projectList: Project[] = [
//     {
//       id: '1',
//       name: '1',
//       pages: 7,
//       thumbnail: '/vite.svg',
//       createdAt: new Date(0),
//       updatedAt: new Date(0)
//     },
//     {
//       id: '2',
//       name: '2',
//       pages: 10,
//       thumbnail: '/vite.svg',
//       createdAt: new Date(10000000000),
//       updatedAt: new Date(10000000000)
//     },
//     {
//       id: '3',
//       name: '3',
//       pages: 10,
//       thumbnail: '/vite.svg',
//       createdAt: new Date(20000000000),
//       updatedAt: new Date(20000000000)
//     },
//     {
//       id: '10',
//       name: '10',
//       pages: 10,
//       thumbnail: '/vite.svg',
//       createdAt: new Date(30000000000),
//       updatedAt: new Date(30000000000)
//     }
//   ]
//   return projectList
// }


// const projectList = getProjectList()
</script>

<style module>
.menu {
  text-align: left;
  width: 90vw;
  min-height: 80vh;
}
</style>
