<template>
  <header :class="$style.header">
    <router-link to="/">Menu</router-link>
  </header>
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
    <transition-group :move-class="$style.listMove" tag="div">
      <open-page
        v-for="page in pageList"
        :key="page.id"
        :page-number="page.index + 1"
        :url="'/page/' + page.id"
        :thumnail-image="'/noImage.svg'"
        @left="decrementIndex"
        @right="incrementIndex"
      />
    </transition-group>
  </div>
</template>

<script lang="ts" setup>
import axios from 'axios'
import { nextTick, onMounted, ref } from 'vue'
import { useRoute } from 'vue-router'
import OpenPage from '../components/menu/OpenPage.vue'
import type { Page } from '../lib/page'
import type { Project } from '../lib/project'
const route = useRoute()
const name = ref()
const nameInput = ref<HTMLInputElement>()

const pages = route.query.pages
const id = ref<number>(5)
const isNameSelect = ref(false)
const pageList = ref<Page[]>([])
const projectData = ref<Project>()

const unSelect = async () => {
  isNameSelect.value = false
  const res = await axios.patch('/api/projects/' + route.params.id, {
    name: name.value
  })
  if (res.status / 100 < 2 && res.status / 100 >= 3) {
    const getRes = await axios.get('/api/projects/' + route.params.id)
    if (res.status / 100 >= 2 && res.status / 100 < 3) {
      name.value = getRes.data.name
    } else {
      name.value = ''
    }
  }
}

const select = () => {
  isNameSelect.value = true
  nextTick(() => {
    nameInput.value?.focus()
  })
}

const decrementIndex = (index: number) => {
  if (
    pageList.value.some(v => v.index === index - 1) &&
    pageList.value.some(v => v.index === index)
  ) {
    //ここにAPI操作を書く
    const pageListValue = pageList.value
    //axios.patch('/api/pages/' + pageListValue[index - 1].id,{operation: 'inc'})
    axios
      .patch('/api/pages/' + pageListValue[index].id + '/index', {
        operation: 'dec'
      })
      .then(responce => {
        if (responce.status !== 200) return
      })

    for (let i = 0; i < pageListValue.length; i++) {
      if (pageListValue[i].index === index - 1) {
        pageList.value[i].index = index
      } else if (pageListValue[i].index === index) {
        pageList.value[i].index = index - 1
      }
    }
    pageList.value.sort((a, b) => {
      return a.index < b.index ? -1 : 1
    })
    //console.log(JSON.stringify(pageList.value))
  }
}

const incrementIndex = (index: number) => {
  //console.log(pageList.value.some(v => v.index === index))
  if (
    pageList.value.some(v => v.index === index) &&
    pageList.value.some(v => v.index === index + 1)
  ) {
    //ここにAPI操作を書く
    //console.log(JSON.stringify(pageList.value))
    const pageListValue = pageList.value
    axios
      .patch('/api/pages/' + pageListValue[index].id + '/index', {
        operation: 'inc'
      })
      .then(responce => {
        if (responce.status !== 200) return
      })
    //axios.patch('/api/pages/' + pageListValue[index + 1].id,{operation: 'dec'})

    for (let i = 0; i < pageListValue.length; i++) {
      if (pageListValue[i].index === index) {
        pageList.value[i].index = index + 1
      } else if (pageListValue[i].index === index + 1) {
        pageList.value[i].index = index
      }
    }
    pageList.value.sort((a, b) => {
      return a.index < b.index ? -1 : 1
    })
    //console.log(JSON.stringify(pageList.value))
  }
}

onMounted(async () => {
  const res = await axios.get('/api/projects/' + route.params.id)
  pageList.value = res.data.pages
  name.value = res.data.name
})
</script>

<style module>
.menu {
  text-align: left;
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

.listMove {
  transition: all 0.3s ease;
}

.header {
  height: 2rem;
  margin-bottom: 1rem;
  display: flex;
  justify-content: center;
  margin-top: 1.5rem;
  font-size: 1.5rem;
}
</style>
