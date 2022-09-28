<template>
  <header :class="$style.header">
        <router-link to="/">Menu</router-link>|
        <router-link :to="'/project/' + projectID">Project</router-link>
  </header>
  <div>ここに編集画面が入るよ</div>
</template>

<script lang="ts" setup>
import axios from 'axios';
import { onMounted, ref } from 'vue';
import { useRoute } from 'vue-router';


const route = useRoute()

const projectID = ref<string>()

onMounted(async () => {
  const res = await axios.get('/api/pages/' + route.params.id)
  projectID.value = res.data.project_id
})
</script>

<style module>

.header {
  height:2rem;
  margin-bottom: 1rem;
  display: flex;
  justify-content: center;
  margin-top: 1.5rem;
  font-size: 1.5rem;
}
</style>
