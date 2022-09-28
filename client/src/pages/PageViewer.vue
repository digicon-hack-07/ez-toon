<template>
  <div :class="$style.editor">
    <header :class="$style.header">
      <!-- <router-link to="/">Menu</router-link>| -->
      <router-link :to="'/project/' + projectID">Project</router-link>
    </header>
    <div :class="$style.editor">
      <page-editor :page-id="pageID"></page-editor>
    </div>
  </div>
</template>

<script lang="ts" setup>
import axios from 'axios'
import { onMounted, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import PageEditor from '../components/PageEditor/PageEditor.vue'

const route = useRoute()

const projectID = ref<string>()
const pageID = ref<string>('')
const router = useRouter()
if (Array.isArray(route.params.id) || !route.params.id) {
  router.push('/')
} else {
  pageID.value = route.params.id
}

onMounted(async () => {
  const res = await axios.get('/api/pages/' + route.params.id)
  projectID.value = res.data.project_id
})
</script>

<style module>
.header {
  height: 2rem;
  margin-bottom: 1rem;
  display: flex;
  justify-content: center;
  margin-top: 1.5rem;
  font-size: 1.5rem;
  position: fixed;
  z-index: 1;
  left: 2rem;
}

.editor {
  position: fixed;
  top: 0;
  bottom: 0;
  left: 0;
  right: 0;
  max-width: 100%;
  max-height: 100%;
  margin: auto;
}
</style>
