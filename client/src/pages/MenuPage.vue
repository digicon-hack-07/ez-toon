<template>
    <div :class="$style.menu">
        <create-new-project/>
        <open-project v-for="n in 5" :key="n" :project-name="n.toString" />
    </div>
</template>

<script setup lang="ts">
import {  ref } from 'vue';
import CreateNewProject from '../components/menu/CreateNewProject.vue';
import OpenProject from '../components/menu/OpenProject.vue';

const getProjectList = function(){
  const projectList = ref<JSON>();
  return projectList
}
const projectList = getProjectList()


</script>

<style module>
    .menu{
        text-align: left;
    }
</style>

