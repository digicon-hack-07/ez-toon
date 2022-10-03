<template>
  <div :class="$style.menu">
    <header :class="$style.header">
      <button :class="$style.back" @click="push">
        <img src="/leftArrow.svg" />
      </button>
    </header>
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
      <button :class="$style.remove" @click="remove">削除</button>
    </div>
    <transition-group :move-class="$style.listMove" tag="span">
      <open-page
        v-for="page in pageList"
        :key="page.id"
        :page-number="page.index"
        :url="'/page/' + page.id"
        :thumnail-image="'/noImage.svg'"
        @left="decrementIndex"
        @right="incrementIndex"
      />
    </transition-group>
    <create-page
      :index="pageList.length"
      :project-id="(route.params.id as string)"
    ></create-page>
  </div>
</template>

<script lang="ts" setup>
import axios from 'axios'
import { nextTick, onMounted, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import OpenPage from '../components/menu/OpenPage.vue'
import type { Page } from '../lib/page'
import CreatePage from '../components/menu/CreatePage.vue'
const route = useRoute()
const name = ref()
const nameInput = ref<HTMLInputElement>()

const pages = route.query.pages
const id = ref<number>(5)
const isNameSelect = ref(false)
const pageList = ref<Page[]>([])
const router = useRouter()

if (Array.isArray(route.params.id) || !route.params.id) {
  router.push('/')
}

const remove = async () => {
  const res = await axios.delete('/api/projects/' + route.params.id)
  if (res.status / 100 >= 2 && res.status / 100 < 3) {
    router.push('/')
  }
}
const unSelect = async () => {
  isNameSelect.value = false
  const res = await axios.patch('/api/projects/' + route.params.id, {
    name: name.value
  })
  if (res.status / 100 >= 2 && res.status / 100 < 3) {
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

const decrementIndex = async (index: number) => {
  if (
    pageList.value.some(v => v.index === index - 1) &&
    pageList.value.some(v => v.index === index)
  ) {
    //ここにAPI操作を書く
    const pageListValue = pageList.value
    // alert(JSON.stringify(pageList.value))
    //axios.patch('/api/pages/' + pageListValue[index - 1].id,{operation: 'inc'})
    const res = await axios.patch(
      '/api/pages/' + pageListValue[index - 1].id + '/index',
      {
        operation: 'dec'
      }
    )
    if (res.status === 200) {
      pageList.value = res.data
    }
    // await axios.patch(
    //   '/api/pages/' + pageListValue[index - 2].id + '/index',
    //   {
    //     operation: 'inc'
    //   }
    // )
    // for (let i = 0; i < pageListValue.length; i++) {
    //   if (pageListValue[i].index === index - 1) {
    //     pageList.value[i].index = index
    //   } else if (pageListValue[i].index === index) {
    //     pageList.value[i].index = index - 1
    //   }
    // }
    // pageList.value.sort((a, b) => {
    //   return a.index < b.index ? -1 : 1
    // })
    //console.log(JSON.stringify(pageList.value))
  }
}

const incrementIndex = async (index: number) => {
  //console.log(pageList.value.some(v => v.index === index))
  if (
    pageList.value.some(v => v.index === index) &&
    pageList.value.some(v => v.index === index + 1)
  ) {
    //ここにAPI操作を書く
    //console.log(JSON.stringify(pageList.value))
    const pageListValue = pageList.value
    //axios.patch('/api/pages/' + pageListValue[index + 1].id,{operation: 'dec'})
    const res = await axios.patch('/api/pages/' + pageListValue[index - 1].id + '/index', {
      operation: 'inc'
    })
    if (res.status === 200) {
      pageList.value = res.data
    }
    // await axios.patch(
    //   '/api/pages/' + pageListValue[index].id + '/index',
    //   {
    //     operation: 'dec'
    //   }
    // )

    // for (let i = 0; i < pageListValue.length; i++) {
    //   if (pageListValue[i].index === index) {
    //     pageList.value[i].index = index + 1
    //   } else if (pageListValue[i].index === index + 1) {
    //     pageList.value[i].index = index
    //   }
    // }
    // pageList.value.sort((a, b) => {
    //   return a.index < b.index ? -1 : 1
    // })
    //axios.get('/api/projects/' + route.params.id)
  }
}

onMounted(async () => {
  const res = await axios.get('/api/projects/' + route.params.id)
  pageList.value = res.data.pages
  name.value = res.data.name
})

const push = () => {
  router.push('/')
}
</script>

<style module>
.menu {
  text-align: left;
}

.name {
  text-align: center;
  margin-top: 1rem;
  margin-bottom: 1rem;
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

.remove {
  position: absolute;
  top: 2rem;
  right: 3rem;
  color: crimson;
}
.remove:hover {
  border-color: crimson;
}

.header {
  margin-bottom: 3rem;
  justify-content: center;
  margin-top: 1rem;
  font-size: 1.5rem;
  position: absolute;
  z-index: 1;
  left: 2rem;
}

.back {
  position: absolute;
  width: 4rem;
  height: 4rem;
  padding: 0.5rem;
  vertical-align: middle;
  border-radius: 0.3rem;
}
</style>
