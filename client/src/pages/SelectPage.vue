<template>
  <div :class="$style.menu">
    <div :class="$style.name">
      <h1 :style="{ margin: 0 }">
        <input
          v-if="isNameSelect"
          ref="nameInput"
          v-model="name"
          :class="$style.nameInput"
          @keyup.enter="unSelect"
          @blur="unSelect"
        />
        <span v-else @click="select">
          {{ name }}
          <img src="/pencil.svg" :class="$style.pencilImage" />
        </span>
      </h1>
    </div>
    <open-page
      v-for="page in pageList"
      :key="page.id"
      :page-number="page.index"
      :url="'/Page/' + page.projectID + '/' + page.id"
      :thumnail-image="'/vite.svg'"
    />
  </div>
</template>

<script lang="ts" setup>
import { nextTick, ref } from 'vue'
import { useRoute } from 'vue-router'
import OpenPage from '../components/menu/OpenPage.vue'
import type { Page } from '../lib/page'
const route = useRoute()
const name = ref('test')
const nameInput = ref<HTMLInputElement>()

const pages = route.query.pages
const id = ref<number>(5)
const isNameSelect = ref(false)
const pageList = ref<Page[]>([
  {
    id: '1',
    projectID: 'a',
    index: 1
  },
  {
    id: '2',
    projectID: 'a',
    index: 2
  },
  {
    id: '3',
    projectID: 'a',
    index: 3
  }
])

const unSelect = () => {
  isNameSelect.value = false
}

const select = () => {
  isNameSelect.value = true
  nextTick(() => {
    nameInput.value?.focus()
  })
}

// if (pages) {
//   if (!Array.isArray(pages)) {
//     let pageCount = parseInt(pages.toString(), 10)
//     pageCount = pageCount > 0 ? pageCount : 1
//     id.value = pageCount
//   } else {
//     router.push('/notFound')
//   }
// } else {
//   router.push('/notFound')
// }
</script>

<style module>
.menu {
  text-align: left;
  width: 90vw;
  min-height: 80vh;
}

.name {
  text-align: center;
  margin: 0;
}

.nameInput {
  font-size: 80%;
  text-align: center;
}

.pencilImage {
  height: 1rem;
}
</style>
